package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Manager handles active WebSocket connections.
type Manager struct {
	connections map[*websocket.Conn]struct{}
	mu          sync.Mutex
}

// NewManager returns a new connection manager.
func NewManager() *Manager {
	return &Manager{connections: make(map[*websocket.Conn]struct{})}
}

// Add registers a new connection.
func (m *Manager) Add(conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.connections[conn] = struct{}{}
}

// Remove unregisters a connection.
func (m *Manager) Remove(conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.connections, conn)
}
