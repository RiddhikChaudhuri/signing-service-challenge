package main

import (
	"fmt"
	"log"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/api"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	_ "github.com/lib/pq"
)

const (
	ListenAddress = ":8080"
	postgreshost  = "localhost"
	port          = "5432"
	user          = "admin"
	password      = "password"
	dbname        = "signing-service"
	// TODO: add further configuration parameters here ...
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", postgreshost, port, user, password, dbname)
	fmt.Println(psqlInfo)
	//Removing PostgreSQL Database
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	storage := persistence.NewMemoryStore()

	//defer db.Close()
	// check db
	//err = db.Ping()
	//CheckError(err)

	//fmt.Println("Connected!")
	server := api.NewServer(ListenAddress, storage)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
