package main

import (
	"fmt"
	"github.com/all4dich/golang/buildanalysis/builddata"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//	"github.com/all4dich/golang/buildanalysis/oebuildjobs"

func main() {
	dbHost := "arion.lge.com"
	dbPort := "27017"
	dbName := "git_api_server"
	dbColl := "buildjob_status"
	dbUser := "log_manager"
	dbPass := "Sanfrancisco"
	dbUrl := fmt.Sprintf("%s:%s", dbHost, dbPort)
	session, err := mgo.Dial(dbUrl)
	log.Println(err)
	/*
		index := mgo.Index{
			Key:        []string{"buildjob", "data"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
	*/
	db := session.DB(dbName)
	db.Login(dbUser, dbPass)
	coll := db.C(dbColl)
	/*
		err := coll.EnsureIndex(index)
			if err != nil {
				panic(err)
			}
	*/
	jobName := "starfish-gld4tv-official-o18"
	//err_check := coll.Find(bson.M{"jobname": jobName})
	var youd []builddata.BuildData
	n := coll.Find(bson.M{"jobname": jobName}).All(&youd)
	log.Println(n)
	log.Println(youd[0].Jobname)
	buildnumbers := make([]int, len(youd))
	for i, v := range youd {
		buildnumbers[i] = v.Duration
	}
	fmt.Println(buildnumbers)
	defer session.Close()
}
