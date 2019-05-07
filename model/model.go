package model

import (
	mgo "gopkg.in/mgo.v2"
)

// Model is an interface for database model
type Model interface {
	Collection(sess *mgo.Session) *mgo.Collection
	Indexes() []mgo.Index
}

// Campaigns returns the collection of campaigns
func Campaigns(sess *mgo.Session) *mgo.Collection {
	return sess.DB("").C("campaigns")
}
