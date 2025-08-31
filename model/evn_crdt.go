package model

import (
	"time"
)

// CRDT操作类型
const (
	OP_INSERT = "insert"
	OP_DELETE = "delete"
)

// CRDT操作结构
type CRDTOperation struct {
	ID        string    `json:"id" gorm:"primaryKey;column:id;comment:操作唯一标识"`
	NoteID    uint      `json:"noteId" gorm:"column:note_id;comment:笔记ID"`
	Type      string    `json:"type" gorm:"column:type;comment:操作类型"`
	Position  int       `json:"position" gorm:"column:position;comment:操作位置"`
	Content   string    `json:"content" gorm:"column:content;comment:操作内容"`
	UserID    uint      `json:"userId" gorm:"column:user_id;comment:操作用户ID"`
	Timestamp int64     `json:"timestamp" gorm:"column:timestamp;comment:时间戳"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;comment:创建时间"`
}

// CRDT位置标识符
type CRDTPosition struct {
	SiteID    uint   `json:"siteId"`    // 站点ID（用户ID）
	Clock     int64  `json:"clock"`     // 逻辑时钟
	Position  int    `json:"position"`  // 位置
	Character string `json:"character"` // 字符内容
}

// 协同编辑会话
type CollaborationSession struct {
	ID       string    `json:"id" gorm:"primaryKey;column:id;comment:会话ID"`
	NoteID   uint      `json:"noteId" gorm:"column:note_id;comment:笔记ID"`
	UserID   uint      `json:"userId" gorm:"column:user_id;comment:用户ID"`
	Username string    `json:"username" gorm:"column:username;comment:用户名"`
	JoinTime time.Time `json:"joinTime" gorm:"column:join_time;comment:加入时间"`
	LastSeen time.Time `json:"lastSeen" gorm:"column:last_seen;comment:最后活跃时间"`
	IsActive bool      `json:"isActive" gorm:"column:is_active;comment:是否活跃"`
}

// 点赞记录
type LikeRecord struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:id;comment:点赞记录ID"`
	NoteID    uint      `json:"noteId" gorm:"column:note_id;comment:笔记ID"`
	UserID    uint      `json:"userId" gorm:"column:user_id;comment:用户ID"`
	Username  string    `json:"username" gorm:"column:username;comment:用户名"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;comment:点赞时间"`
}

// 笔记点赞统计
type NoteLikeStats struct {
	NoteID    uint `json:"noteId" gorm:"column:note_id;comment:笔记ID"`
	LikeCount int  `json:"likeCount" gorm:"column:like_count;comment:点赞数量"`
}
