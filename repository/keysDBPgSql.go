package repository

import (
	"Atlan_Collect_KGS/entity"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type KeyStoreSql struct {
	pool *pgxpool.Pool
}

func NewKeyStoreSql(pool *pgxpool.Pool) *KeyStoreSql {
	return &KeyStoreSql{
		pool: pool,
	}
}

func (r *KeyStoreSql) Create(e []*entity.Key, last int) error {
	log.Printf("Inserting Keys....")
	var rows = [][]interface{}{}
	for _, r := range e {
		temp := make([]interface{}, 0)
		temp = append(temp, r.KeyValue, last)
		rows = append(rows, temp)
	}
	copyCount, err := r.pool.CopyFrom(context.Background(), pgx.Identifier{"key_store"}, []string{"keyval", "lastval"}, pgx.CopyFromRows(rows))
	if err != nil {
		fmt.Printf("error : %s", err)
		return err
	}
	if copyCount >= 1 {
		log.Printf("Insertion Complete Rows affected : ", copyCount)
	}
	return nil
}

func (r *KeyStoreSql) Get() (int, error) {
	var last int
	log.Printf("Fetching Last key....")
	row, err := r.pool.Query(context.Background(), "SELECT lastval FROM key_store ORDER BY keyid DESC LIMIT 1")

	if err != nil {
		log.Fatalf("Error while fetching last key...")
		return -1, err
	}
	for row.Next() {
		er := row.Scan(&last)
		if er != nil {
			log.Fatalf("Error while iterrating")
		}
	}
	log.Printf("Lastkey = %d", last)
	return last, nil
}
