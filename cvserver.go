package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/remotejob/kaukotyoeu/domains"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var username string
var password string
var addrs []string

func init() {

	if os.Getenv("SECRET_USERNAME") != "" {

		username = os.Getenv("SECRET_USERNAME")
		password = os.Getenv("SECRET_PASSWORD")
	}

	log.Println("pass", username, password)
	addrs = []string{"cv-service"}

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  "admin",
		Username:  username,
		Password:  password,
		Mechanism: "SCRAM-SHA-1",
	}

	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	dbsession.SetMode(mgo.Monotonic, true)
	c := dbsession.DB("cv_employers").C("employers")

	var results []domains.JobOffer
	//	err := c.Find(bson.M{"externallink": bson.M{"$ne": ""}, "location": bson.RegEx{Pattern: "Sweden", Options: "i"}, "applied": false}).All(&results)
	err = c.Find(bson.M{"externallink": bson.M{"$ne": ""}, "applied": false}).All(&results)

	if err != nil {

		log.Fatal(err)
	}

	for _, result := range results {

		log.Println(result.Title)

	}

}

func hello(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
