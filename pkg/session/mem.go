package session

type Memory struct {
	val map[int64]map[string]interface{}
}

func NewMemory() Memory {
	return Memory{val: map[int64]map[string]interface{}{}}
}

func (m Memory) Set(id int64, key string, value interface{}) {
	m.val[id][key] = value
}

func (m Memory) Get(id int64, key string) (value interface{}, err error) {
	return m.val[id][key], nil
}

func (m Memory) GetAll(id int64) (map[string]interface{}, error) {
	return m.val[id], nil
}
