package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// https://github.com/jmoiron/sqlx
// https://jmoiron.github.io/sqlx/
// https://github.com/go-jet/jet
// https://github.com/go-pg/pg
// SQL
// drop table transactions;
// create table transactions (
// 	"tx_hash" text,
// 	"from" text,
// 	"to" text,
// 	"amount" BIGINT,
// 	"timestamp" TIMESTAMP
// );

type Transaction struct {
	TxHash    string    `db:"tx_hash" json:"tx_hash"`
	From      string    `db:"from" json:"from"`
	To        string    `db:"to" json:"to"`
	Amount    uint64    `db:"amount" json:"amount"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
}

func main() {
	// Connect to database
	db, err := sqlx.Connect("postgres", "dbname=sample sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Create a transaction
	tx := Transaction{
		TxHash:    "0x0b8106d23c9b8aebbd429b5a5b9c104e77dfc872569bc342d4731965a63f20f5",
		From:      "0xf2f5c73fa04406b1995e397b55c24ab1f3ea726c",
		To:        "0x0b69af521c94e17f51dade83ada119f89056d55a",
		Amount:    797,
		Timestamp: time.Now(),
	}

	// Store a transaction
	_, err = db.NamedExec(`INSERT INTO transactions ("tx_hash","from","to","amount","timestamp")
        VALUES (:tx_hash,:from,:to,:amount,:timestamp)`, tx)
	if err != nil {
		log.Fatalln(err)
	}

	txs := []Transaction{
		{
			TxHash:    "0x9bbed0730a4eaaf1f2e66d585d7cf5f075a065c6c50f341ec9d4d6df633e10f9",
			From:      "0xd24400ae8bfebb18ca49be86258a3c749cf46853",
			To:        "0x5f65f7b609678448494de4c87521cdf6cef1e932",
			Amount:    73,
			Timestamp: time.Now(),
		},
		{
			TxHash:    "0x99ecaa8174ec780c5d12d647ad815c14931eddd07c2a2fc6d85f1bb0465ccc7c",
			From:      "0x150966164761cd49b2730340e4322b6da1a54b68",
			To:        "0xdac17f958d2ee523a2206206994597c13d831ec7",
			Amount:    90,
			Timestamp: time.Now(),
		},
		{
			TxHash:    "0xe0b6dc3af720fadcec1e9bd0b333f2ee83078dd2c6dd1026d1eef2fc11ee8551",
			From:      "0x08d2d7df1ff6a3fc555fc84510a5b04817cbfb36",
			To:        "0x3b794929566e3ba0f25e4263e1987828b5c87161",
			Amount:    105,
			Timestamp: time.Now(),
		},
	}

	_, err = db.NamedExec(`INSERT INTO transactions ("tx_hash","from","to","amount","timestamp")
        VALUES (:tx_hash,:from,:to,:amount,:timestamp)`, txs)
	if err != nil {
		log.Fatalln(err)
	}

	var transactions []Transaction
	err = db.Select(&transactions, `SELECT * FROM transactions;`)
	if err != nil {
		log.Fatalln(err)
	}

	// pp.Println(transactions)
	content, err := json.Marshal(transactions)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", content)
}
