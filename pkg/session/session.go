package session

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/wawayes/bi-chatgpt-golang/models"
	"github.com/wawayes/bi-chatgpt-golang/pkg/logx"
	"time"
)

type Session struct {
	SessionID uuid.UUID
	UserInfo  models.User
}

// Save 将会话存入Session
func (s *Session) Save(ctx context.Context, client *redis.Client) error {
	// json.Marshal 将会话转为json格式
	data, err := json.Marshal(s)
	if err != nil {
		logx.Warning(err.Error())
		return err
	}
	// 将session存入redis
	err = client.Set(ctx, s.SessionID.String(), data, 24*time.Hour).Err()
	if err != nil {
		logx.Warning(err.Error())
		return err
	}
	return nil
}

// GetSession 从Redis中获取会话信息
func GetSession(ctx context.Context, client *redis.Client, sessionID string) (*Session, error) {
	data, err := client.Get(ctx, sessionID).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("session does not exist")
		}
		return nil, err
	}
	var session Session
	err = json.Unmarshal(data, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
