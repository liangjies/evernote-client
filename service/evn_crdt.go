package service

import (
	"crypto/rand"
	"encoding/hex"
	"evernote-client/global"
	"evernote-client/model"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// CRDT服务
type CRDTService struct {
	mu sync.RWMutex
	// 笔记ID -> 连接映射
	connections map[uint]map[string]*websocket.Conn
	// 用户时钟
	clocks map[uint]int64
}

func NewCRDTService() *CRDTService {
	return &CRDTService{
		connections: make(map[uint]map[string]*websocket.Conn),
		clocks:      make(map[uint]int64),
	}
}

// 生成唯一操作ID
func generateOperationID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// 获取用户时钟
func (s *CRDTService) getClock(userID uint) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	if clock, exists := s.clocks[userID]; exists {
		s.clocks[userID] = clock + 1
		return s.clocks[userID]
	}
	s.clocks[userID] = 1
	return 1
}

func (s *CRDTService) JoinCollaboration(noteID uint, userID uint, username string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sessionID := fmt.Sprintf("%d_%d", userID, time.Now().UnixNano())

	session := model.CollaborationSession{
		ID:       sessionID,
		NoteID:   noteID,
		UserID:   userID,
		Username: username,
		JoinTime: time.Now(),
		LastSeen: time.Now(),
		IsActive: true,
	}

	global.DB.Create(&session)

	if s.connections[noteID] == nil {
		s.connections[noteID] = make(map[string]*websocket.Conn)
	}
	s.connections[noteID][sessionID] = conn

	s.broadcastToNote(noteID, map[string]interface{}{
		"type":      "user_joined",
		"sessionId": sessionID,
		"userId":    userID,
		"username":  username,
	}, sessionID)
}

func (s *CRDTService) LeaveCollaboration(noteID uint, sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	global.DB.Model(&model.CollaborationSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"is_active": false,
			"last_seen": time.Now(),
		})

	if connections, exists := s.connections[noteID]; exists {
		delete(connections, sessionID)
		if len(connections) == 0 {
			delete(s.connections, noteID)
		}
	}

	s.broadcastToNote(noteID, map[string]interface{}{
		"type":      "user_left",
		"sessionId": sessionID,
	}, sessionID)
}

// 处理CRDT操作
func (s *CRDTService) HandleOperation(noteID uint, userID uint, opType string, position int, content string) {
	operationID := generateOperationID()
	clock := s.getClock(userID)

	operation := model.CRDTOperation{
		ID:        operationID,
		NoteID:    noteID,
		Type:      opType,
		Position:  position,
		Content:   content,
		UserID:    userID,
		Timestamp: clock,
		CreatedAt: time.Now(),
	}

	global.DB.Create(&operation)

	s.broadcastToNote(noteID, map[string]interface{}{
		"type":      "operation",
		"operation": operation,
	}, "")
}

func (s *CRDTService) GetNoteOperations(noteID uint) ([]model.CRDTOperation, error) {
	var operations []model.CRDTOperation
	err := global.DB.Where("note_id = ?", noteID).
		Order("timestamp ASC, user_id ASC").
		Find(&operations).Error
	return operations, err
}

func (s *CRDTService) ApplyOperationsToContent(content string, operations []model.CRDTOperation) string {
	sort.Slice(operations, func(i, j int) bool {
		if operations[i].Timestamp != operations[j].Timestamp {
			return operations[i].Timestamp < operations[j].Timestamp
		}
		return operations[i].UserID < operations[j].UserID
	})

	result := content
	offset := 0

	for _, op := range operations {
		switch op.Type {
		case model.OP_INSERT:
			pos := op.Position + offset
			if pos >= 0 && pos <= len(result) {
				result = result[:pos] + op.Content + result[pos:]
				offset += len(op.Content)
			}
		case model.OP_DELETE:
			pos := op.Position + offset
			if pos >= 0 && pos < len(result) {
				deleteLen := len(op.Content)
				if pos+deleteLen <= len(result) {
					result = result[:pos] + result[pos+deleteLen:]
					offset -= deleteLen
				}
			}
		}
	}

	return result
}

func (s *CRDTService) broadcastToNote(noteID uint, message map[string]interface{}, excludeSessionID string) {
	if connections, exists := s.connections[noteID]; exists {
		for sessionID, conn := range connections {
			if sessionID != excludeSessionID {
				err := conn.WriteJSON(message)
				if err != nil {
					delete(connections, sessionID)
				}
			}
		}
	}
}

func (s *CRDTService) GetActiveUsers(noteID uint) ([]model.CollaborationSession, error) {
	var sessions []model.CollaborationSession
	err := global.DB.Where("note_id = ? AND is_active = ?", noteID, true).
		Order("join_time ASC").
		Find(&sessions).Error
	return sessions, err
}
