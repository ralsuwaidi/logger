package main

import (
	"fmt"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var connection = "localhost"
var database = "ProfileService"
var collection = "log"

func connectAtlas() (*mgo.Session, error) {
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{"ds133630.mlab.com:33630"},
		Database: "personal",
		Username: "rashed",
		Password: "password_123",
		Timeout:  60 * time.Second,
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	fmt.Println(session.LiveServers())
	return session, err
}

// Profile - is the memory representation of one user profile
type Profile struct {
	Name        string `json:"username"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	LastUpdated time.Time
}

// GetLogs - Returns all the profile in the Profiles Collection
func GetLogs() []Log {
	session, err := mgo.Dial(connection)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)
	var logs []Log
	err = c.Find(bson.M{}).All(&logs)

	return logs
}

// ShowProfile - Returns the profile in the Profiles Collection with name equal to the id parameter (id == name)
func ShowProfile(id string) Profile {
	session, err := mgo.Dial(connection)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return Profile{}
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)
	profile := Profile{}
	err = c.Find(bson.M{"name": id}).One(&profile)

	return profile
}

// DeleteProfile - Deletes the profile in the Profiles Collection with name equal to the id parameter (id == name)
func DeleteProfile(id string) bool {
	session, err := mgo.Dial(connection)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)
	err = c.RemoveId(id)

	return true
}

// CreateOrUpdateProfile - Creates or Updates (Upsert) the profile in the Profiles Collection with id parameter
func (p *Profile) CreateOrUpdateProfile() bool {
	session, err := mgo.Dial(connection)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)
	_, err = c.UpsertId(p.Name, p)
	if err != nil {
		log.Println("Error creating Profile: ", err.Error())
		return false
	}
	return true
}

// CreateLog - Creates or Updates (Upsert) the profile in the Profiles Collection with id parameter
func (p *Log) CreateLog() bool {
	session, err := connectAtlas()
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("personal").C("logs")
	_, err = c.UpsertId(p.Title, p)
	if err != nil {
		log.Println("Error creating Log: ", err.Error())
		return false
	}
	return true
}
