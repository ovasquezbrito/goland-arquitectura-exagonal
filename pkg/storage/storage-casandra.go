package storage

import (
	"log"
	"github.com/gocql/gocql"
)

// Storage . base de datos
type Storage struct {
	db *gocql.Session
}

// SetupStorage .
func SetupStorage() (*Storage, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "candy_shop_db"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{db: session}, nil
}

// GetAllCandyNames .
func (s *Storage) GetAllCandyNames() ([]string, error) {
	var candy string
	var candies []string
	iter := s.db.Query(`SELECT name FROM candies`).Iter()
	for iter.Scan(&candy) {
		candies = append(candies, candy)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return candies, nil
}