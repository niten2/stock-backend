package db

import (
	"fmt"
	"os"
	"gopkg.in/mgo.v2"
)

var (
	Session *mgo.Session
	Mongo *mgo.DialInfo
  Db *mgo.Database
)

func Connect() {
	uri := os.Getenv("DB_URL")

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to mongo", uri)

	Session = s
	Mongo = mongo
  Db = s.DB(os.Getenv("DB_NAME"))
}
