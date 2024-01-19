package main

import (
	"github.com/nutsdb/nutsdb"
	"log"
)

var db *nutsdb.DB

func initDB() {
	db, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir("./tmp/db"))
	if err != nil {
		log.Fatal(err)
	}
}

func update() {

}
