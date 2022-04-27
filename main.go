package main

import (
	"Atlan_Collect_KGS/repository"
	"Atlan_Collect_KGS/usecase/key"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func main() {
	DATABASE_URL := os.Getenv("KEY_DATABASE_URL")
	ctx := context.Background()
	cofig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, cofig)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	defer pool.Close()

	repo := repository.NewKeyStoreSql(pool)
	serv := key.NewService(repo)

	lastKey, err := serv.GetLastKey()
	if err != nil {
		return
	}
	err, lastKey = serv.AddKeys(lastKey, 10000)
	if err != nil {
		return
	}

}
