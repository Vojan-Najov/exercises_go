/*
 * 1. Cache patterns
 *
 * Создать хранилище сессий пользователей. Время работы каждой сессии (TTL)
 * составляет 2 минуты. Описание структур и функций.
 *
 * type Session struct {
 *   ID        string
 *   UserID    string
 *   ExpiresAt time.Time
 * }

 * type SessionManager struct {
 *   sessions map[string]Session
 *   mutex    sync.RWMutex
 * }

 * func NewSessionManager() *SessionManager
 * func (sm *SessionManager) StartSession(userID string) string
 * func (sm *SessionManager) GetSession(sessionID string) (Session, bool)
 *
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var sessionCount int

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
}

type SessionManager struct {
	sessions map[string]Session
	mutex    sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]Session),
	}
}

func (sm *SessionManager) StartSession(userID string) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sessionCount++
	id := fmt.Sprint(sessionCount)
	sm.sessions[id] = Session{
		ID:        id,
		UserID:    userID,
		ExpiresAt: time.Now().Add(2 * time.Minute),
	}
	return id
}

func (sm *SessionManager) GetSession(sessionID string) (Session, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	session, found := sm.sessions[sessionID]
	if !found || time.Now().After(session.ExpiresAt) {
		return Session{}, false
	}
	return session, true
}
