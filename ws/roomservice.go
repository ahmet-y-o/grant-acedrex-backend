package ws

import (
	"acedrex/game"
	"sync"
	"time"
)

type RoomService struct {
	Rooms           map[string]*Room
	mu              sync.RWMutex
	stopChan        chan struct{}
	cleanUpInterval time.Duration
}

func (rs *RoomService) GetRoom(roomId string) (*Room, bool) {
	room, exists := rs.Rooms[roomId]
	if !exists {
		return nil, false
	}

	return room, true
}

func NewRoomService() *RoomService {
	rs := &RoomService{
		Rooms:           make(map[string]*Room),
		stopChan:        make(chan struct{}),
		cleanUpInterval: 30 * time.Second,
	}

	go rs.StartCleanupRoutine()
	return rs
}

func (rs *RoomService) StartCleanupRoutine() {
	ticker := time.NewTicker(rs.cleanUpInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rs.cleanupEmptyRooms()
		case <-rs.stopChan:
			return
		}
	}
}

// cleanupEmptyRooms removes rooms that have been empty for longer than the TTL
func (rs *RoomService) cleanupEmptyRooms() {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	for roomID, room := range rs.Rooms {
		if room.IsEmpty() {
			delete(rs.Rooms, roomID)
		}
	}
}

func (rs *RoomService) DeleteRoom(roomId string) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	delete(rs.Rooms, roomId)
}

// Stop gracefully stops the cleanup routine
func (rs *RoomService) Stop() {
	close(rs.stopChan)
}

func (rs *RoomService) GetRoomsList() []string {
	ids := []string{}
	for id := range rs.Rooms {
		ids = append(ids, id)
	}
	return ids
}

func (rs *RoomService) NewRoom() string {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	uuid := generateHashID()
	room := &Room{
		RoomId:  uuid,
		Game:    game.InitStandardGame(),
		Started: false,
	}
	rs.Rooms[uuid] = room

	return uuid
}
