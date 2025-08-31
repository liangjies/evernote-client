package service

import (
	"evernote-client/global"
	"evernote-client/model"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type LikeService struct {
	ctx context.Context
}

func NewLikeService() *LikeService {
	return &LikeService{
		ctx: context.Background(),
	}
}

func (s *LikeService) LikeNote(noteID uint, userID uint, username string) error {
	likeKey := fmt.Sprintf("note_like:%d", noteID)
	userLikeKey := fmt.Sprintf("user_like:%d:%d", userID, noteID)

	exists, err := global.REDIS.Exists(s.ctx, userLikeKey).Result()
	if err != nil {
		return err
	}

	if exists == 1 {
		return fmt.Errorf("用户已经点赞过该笔记")
	}

	likeRecord := model.LikeRecord{
		NoteID:    noteID,
		UserID:    userID,
		Username:  username,
		CreatedAt: time.Now(),
	}

	if err := global.DB.Create(&likeRecord).Error; err != nil {
		return err
	}

	// 使用Redis ZSET记录点赞时间戳（用于排序）
	score := float64(time.Now().Unix())
	err = global.REDIS.ZAdd(s.ctx, likeKey, &redis.Z{
		Score:  score,
		Member: userID,
	}).Err()
	if err != nil {
		return err
	}

	err = global.REDIS.Set(s.ctx, userLikeKey, "1", 0).Err()
	if err != nil {
		return err
	}

	noteLikeCountKey := fmt.Sprintf("note_like_count:%d", noteID)
	err = global.REDIS.Incr(s.ctx, noteLikeCountKey).Err()
	if err != nil {
		return err
	}

	globalLikeRankKey := "global_like_rank"
	err = global.REDIS.ZIncrBy(s.ctx, globalLikeRankKey, 1, strconv.FormatUint(uint64(noteID), 10)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *LikeService) UnlikeNote(noteID uint, userID uint) error {
	likeKey := fmt.Sprintf("note_like:%d", noteID)
	userLikeKey := fmt.Sprintf("user_like:%d:%d", userID, noteID)

	exists, err := global.REDIS.Exists(s.ctx, userLikeKey).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("用户未点赞该笔记")
	}

	err = global.DB.Where("note_id = ? AND user_id = ?", noteID, userID).
		Delete(&model.LikeRecord{}).Error
	if err != nil {
		return err
	}

	err = global.REDIS.ZRem(s.ctx, likeKey, userID).Err()
	if err != nil {
		return err
	}

	err = global.REDIS.Del(s.ctx, userLikeKey).Err()
	if err != nil {
		return err
	}

	noteLikeCountKey := fmt.Sprintf("note_like_count:%d", noteID)
	err = global.REDIS.Decr(s.ctx, noteLikeCountKey).Err()
	if err != nil {
		return err
	}

	globalLikeRankKey := "global_like_rank"
	err = global.REDIS.ZIncrBy(s.ctx, globalLikeRankKey, -1, strconv.FormatUint(uint64(noteID), 10)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *LikeService) IsUserLiked(noteID uint, userID uint) (bool, error) {
	userLikeKey := fmt.Sprintf("user_like:%d:%d", userID, noteID)
	exists, err := global.REDIS.Exists(s.ctx, userLikeKey).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func (s *LikeService) GetNoteLikeCount(noteID uint) (int64, error) {
	noteLikeCountKey := fmt.Sprintf("note_like_count:%d", noteID)
	count, err := global.REDIS.Get(s.ctx, noteLikeCountKey).Int64()
	if err == redis.Nil {
		// 如果Redis中没有，从数据库计算
		var count int64
		err = global.DB.Model(&model.LikeRecord{}).
			Where("note_id = ?", noteID).
			Count(&count).Error
		if err != nil {
			return 0, err
		}
		// 缓存到Redis
		global.REDIS.Set(s.ctx, noteLikeCountKey, count, 0)
		return count, nil
	}
	return count, err
}

func (s *LikeService) GetNoteLikeUsers(noteID uint, limit int) ([]model.LikeRecord, error) {
	var records []model.LikeRecord
	err := global.DB.Where("note_id = ?", noteID).
		Order("created_at DESC").
		Limit(limit).
		Find(&records).Error
	return records, err
}

func (s *LikeService) GetLikeRanking(limit int) ([]map[string]interface{}, error) {
	globalLikeRankKey := "global_like_rank"

	result, err := global.REDIS.ZRevRangeWithScores(s.ctx, globalLikeRankKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	var ranking []map[string]interface{}
	for _, z := range result {
		noteID, err := strconv.ParseUint(z.Member.(string), 10, 64)
		if err != nil {
			continue
		}

		var note model.EvnNote
		err = global.DB.Where("id = ?", noteID).First(&note).Error
		if err != nil {
			continue
		}

		ranking = append(ranking, map[string]interface{}{
			"noteId":    noteID,
			"title":     note.Title,
			"snippet":   note.Snippet,
			"likeCount": int(z.Score),
		})
	}

	return ranking, nil
}

func (s *LikeService) GetUserLikedNotes(userID uint, page, pageSize int) ([]model.EvnNote, int64, error) {
	var notes []model.EvnNote
	var total int64

	err := global.DB.Model(&model.LikeRecord{}).
		Joins("JOIN evn_notes ON evn_like_records.note_id = evn_notes.id").
		Where("evn_like_records.user_id = ?", userID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = global.DB.Model(&model.LikeRecord{}).
		Select("evn_notes.*").
		Joins("JOIN evn_notes ON evn_like_records.note_id = evn_notes.id").
		Where("evn_like_records.user_id = ?", userID).
		Order("evn_like_records.created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&notes).Error

	return notes, total, err
}

func (s *LikeService) GetHotNotes(limit int) ([]model.EvnNote, error) {
	var notes []model.EvnNote

	err := global.DB.Model(&model.EvnNote{}).
		Select("evn_notes.*, COUNT(evn_like_records.id) as like_count").
		Joins("LEFT JOIN evn_like_records ON evn_notes.id = evn_like_records.note_id").
		Where("evn_notes.del_flag = ?", false).
		Group("evn_notes.id").
		Order("like_count DESC").
		Limit(limit).
		Find(&notes).Error

	return notes, err
}

// 同步Redis和数据库的点赞数据
func (s *LikeService) SyncLikeData() error {
	// 从数据库获取所有点赞记录，同步到Redis
	var records []model.LikeRecord
	err := global.DB.Find(&records).Error
	if err != nil {
		return err
	}

	for _, record := range records {
		likeKey := fmt.Sprintf("note_like:%d", record.NoteID)
		userLikeKey := fmt.Sprintf("user_like:%d:%d", record.UserID, record.NoteID)
		noteLikeCountKey := fmt.Sprintf("note_like_count:%d", record.NoteID)
		globalLikeRankKey := "global_like_rank"

		global.REDIS.Set(s.ctx, userLikeKey, "1", 0)

		score := float64(record.CreatedAt.Unix())
		global.REDIS.ZAdd(s.ctx, likeKey, &redis.Z{
			Score:  score,
			Member: record.UserID,
		})

		global.REDIS.Incr(s.ctx, noteLikeCountKey)

		global.REDIS.ZIncrBy(s.ctx, globalLikeRankKey, 1, strconv.FormatUint(uint64(record.NoteID), 10))
	}

	return nil
}
