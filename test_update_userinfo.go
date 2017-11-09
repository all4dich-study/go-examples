package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"time"

	"github.com/all4dich/golang/buildanalysis/builddata"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]" + ":mecha"
}

func typeName(v interface{}) string {

	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

type Deadline struct {
	time.Time
}

func main() {
	session, err := mgo.Dial("156.147.69.55:27017")
	var _ = err
	defer session.Close()
	db := session.DB("git_api_server")
	db.Login("log_manager", "Sanfrancisco")
	coll := db.C("z-testone")
	var test = []builddata.BuildData{}
	err = coll.Find(bson.M{"jobname": bson.M{"$regex": bson.RegEx{`starfish.*`, ""}}}).All(&test)
	if err != nil {
		log.Fatal("Error: Can't find")
	}
	for _, a := range test {
		fmt.Println(a.Start)
		fmt.Println(time.Unix(int64(a.Start), 0))
	}
	n := runtime.NumCPU()

	log.Println("Done: ", n)
}
