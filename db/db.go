package db

import (
  // "fmt"
	"log"
  "os"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func ConnectDb () *mgo.Database {
	session, err := mgo.Dial(os.Getenv("DB_URL"))
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}

	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)
  db := session.DB(os.Getenv("DB_NAME"))

  return db
}

