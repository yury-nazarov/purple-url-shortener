/*
Архитектура модуля:
- Load() - Вызываем, что бы при старте приложения загружает файл с данными в память или инициировать in-memeory
- Add() / Validate() - Все запросы через ручки - работают с in-memrory
- Snapshoot() - Сохраняет данные в файл (сейчас: При остановке приложения)

TODO:
- По тикеру регулярный флеш данных в файл
*/

package store

import (
	"encoding/json"
	"os"
	"sync"
)

type Storage struct {
	mu   sync.RWMutex
	data map[string]string
	path string
}

func New(filePath string) *Storage {
	return &Storage{
		path: filePath,
		data: make(map[string]string),
	}
}

// Читает данные с диска и создает временное хранилище в памяти
func (s *Storage) Load() error {
	// Читаем файл
	file, err := os.ReadFile(s.path)
	if err != nil {
		return err
	}
	// Сериализуем RAW JSON в map
	if err := json.Unmarshal(file, &s.data); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Add(hash, email string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[hash] = email
}

func (s *Storage) Validate(hash string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[hash]; ok {
		delete(s.data, hash)
		return true
	}
	return false
}

func (s *Storage) Snapshoot() error {
	// Копируем данные в памяти, для асинхронной записи на диск
	s.mu.RLock()
	copyMap := make(map[string]string, len(s.data))
	for k, v := range s.data {
		copyMap[k] = v
	}
	s.mu.RUnlock()

	// Сохраняем снепшот на диск
	data, err := json.Marshal(copyMap)
	if err != nil {
		return err
	}

	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	// Если все ок, переименовываем
	return os.Rename(tmp, s.path)

}
