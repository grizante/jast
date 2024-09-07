package storage

type Storage struct {
	hashmap map[string]any
}

func (s *Storage) Set(key string, value any) {
	s.hashmap[key] = value
}

func (s *Storage) Read() any {
	return s.hashmap
}


func NewStorage() *Storage {
        return &Storage{hashmap: make(map[string]any)}
}