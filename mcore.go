package mcore

import (
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/bbolt"
)

// DataUnit represents a stored object (обьем хранения)
type DataUnit struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"` // secret, link, note, image_path
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
}

// Mcore - is the main core structure (оснавная структура ядра)
type Mcore struct {
	db *bbolt.DB
}

// NewEngine initializes the storage engine (инициализирует движок хранения)
func NewEngine(dbPath string) (*Mcore, error) {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	// Create storage bucket if it does not exist (создаем хранилище, если его нет)
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Vault"))
		return err
	})

	return &Mcore{db: db}, err
}

// Save stores a data unit in the database (сохранение обьекта в базе)
func (e *Mcore) Save(unit DataUnit) error {
	return e.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Vault"))
		unit.CreatedAt = time.Now()

		// Encode struct to JSON (преобразуем структуру в JSON)
		encoded, err := json.Marshal(unit)
		if err != nil {
			return err
		}

		return b.Put([]byte(unit.ID), encoded)
	})
}

// Get retrieves a data unit by ID (Функция поиска, ядра)
func (e *Mcore) Get(id string) (DataUnit, error) {
	var unit DataUnit
	err := e.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("Vault"))
		v := b.Get([]byte(id))
		if v == nil {
			return fmt.Errorf("error")
		}
		return json.Unmarshal(v, &unit)
	})
	return unit, err
}
