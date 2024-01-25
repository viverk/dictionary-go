package dictionary

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const dbFile = "dictionary.db"
const bucketName = "entries"

type Dictionary struct {
	db *bolt.DB
}

func New() (*Dictionary, error) {
	// Open the BoltDB database file
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Create a bucket if it doesn't exist
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("error creating bucket: %v", err)
	}

	return &Dictionary{db: db}, nil
}

func (d *Dictionary) Close() {
	d.db.Close()
}

const (
	minWordLength       = 3
	maxWordLength       = 50
	minDefinitionLength = 5
	maxDefinitionLength = 500
)

func (d *Dictionary) Add(word, definition string) error {

	// Validate data
	if len(word) < minWordLength || len(word) > maxWordLength {
		return fmt.Errorf("La taille du mot doit etre entre %d et %d caractere", minWordLength, maxWordLength)
	}

	if len(definition) < minDefinitionLength || len(definition) > maxDefinitionLength {
		return fmt.Errorf("La taille de la definition doit etre entre %d et %d caractere", minDefinitionLength, maxDefinitionLength)
	}

	return d.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		return bucket.Put([]byte(word), []byte(definition))
	})
}

func (d *Dictionary) Get(word string) (string, error) {
	var result string
	err := d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		val := bucket.Get([]byte(word))
		if val == nil {
			return fmt.Errorf("word not found")
		}

		result = string(val)
		return nil
	})
	return result, err
}

func (d *Dictionary) Remove(word string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		return bucket.Delete([]byte(word))
	})
}

func (d *Dictionary) List() (map[string]string, error) {
	result := make(map[string]string)
	err := d.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}

		return bucket.ForEach(func(k, v []byte) error {
			result[string(k)] = string(v)
			return nil
		})
	})
	return result, err
}
