package main

import (
	"fmt"
	"github.com/all4dich/golang/buildanalysis/builddata"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	dbUrl := "156.147.69.55:27017"
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		log.Fatal("Can't open a db connection")
		panic(err)
	} else {
		log.Println("SUccess")
	}
	defer session.Close()
	dbName := "git_api_server"
	dbColl := "buildjob_status"
	dbUser := "log_manager"
	dbPassword := "Sanfrancisco"
	jobName := "starfish-gld4tv-verify-lm18a"
	db := session.DB(dbName)
	db.Login(dbUser, dbPassword)
	coll := db.C(dbColl)
	latest_job := builddata.BuildData{}
	err2 := coll.Find(bson.M{"jobname": jobName}).Sort("-buildnumber").One(&latest_job)
	if err2 != nil {
		panic(err2)
	}
	//latest_job := coll.Find(bson.M{"jobname": jobName}).Sort("buildnumber")
	if latest_job.Buildnumber <= 4000 {
		fmt.Println(latest_job.Buildnumber)
	}
	fmt.Println("123")
}
