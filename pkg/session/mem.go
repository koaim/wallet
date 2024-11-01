package session

import "sync"

type Memory struct {
	val map[int64]map[string]interface{}
	mu  *sync.Mutex
}

func NewMemory() Memory {
	return Memory{val: map[int64]map[string]interface{}{}, mu: &sync.Mutex{}}
}

func (m Memory) Set(id int64, key string, value interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.val[id]
	if !ok {
		m.val[id] = map[string]interface{}{}
	}

	m.val[id][key] = value

	return nil
}

func (m Memory) Get(id int64, key string) (interface{}, error) {
	return m.val[id][key], nil
}

func (m Memory) Clear(id int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.val[id] = map[string]interface{}{}

	return nil
}

func (m Memory) GetAll(id int64) (map[string]interface{}, error) {
	return m.val[id], nil
}
