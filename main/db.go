package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
)

var token_id = 0;

type Repository struct {
}

const SERVER = "mongodb://localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "token"

// COLLECTION is the name of the collection in DB
const COLLECTION = "test"

func (r Repository) AddToken(token Token) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	token_id += 1
	token.ID = token_id
	session.DB(DBNAME).C(COLLECTION).Insert(token)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Addet new Token ", token.ID)
	return true
}

func (r Repository) DeleteToken(id int) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Remove product
	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Token - ", id)
	// Write status
	return "OK"
}