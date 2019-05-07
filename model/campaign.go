package model

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Campaign represents a Campaign Model
type Campaign struct {
	ID                 bson.ObjectId       `bson:"_id"`
	Title              string              `bson:"title"`
	Code               string              `bson:"code"`
	AllowedPaymentType []string            `bson:"allowedPaymentType"`
	Partner            string              `bson:"partner"`
	BurnSharing        float64             `bson:"burnSharing"`
	Budget             int64               `bson:"budget"`
	StartDate          time.Time           `bson:"startDate"`
	EndDate            time.Time           `bson:"endDate"`
	Promos             map[string][]string `bson:"promo"`
	CashBacks          map[string][]string `bson:"cashback"`
	Discounts          []string            `bson:"discount"`
	CreatedAt          time.Time           `bson:"createdAt"`
	UpdatedAt          time.Time           `bson:"updatedAt"`
}

// Collection represents the database collection of Campaign
func (c Campaign) Collection(sess *mgo.Session) *mgo.Collection {
	return Campaigns(sess)
}

// Indexes represents the database indexes of Campaign
func (c Campaign) Indexes() []mgo.Index {
	return []mgo.Index{
		{
			Key:    []string{"code"},
			Unique: true,
		},
		{
			Key: []string{"allowedPaymentType"},
		},
		{
			Key: []string{"partner"},
		},
		{
			Key: []string{"startDate", "endDate"},
		},
	}
}

func (c *Campaign) pre() error {
	c.UpdatedAt = time.Now()
	if c.ID == "" {
		c.ID = bson.NewObjectId()
		c.CreatedAt = c.UpdatedAt
	}

	return nil
}

func (c *Campaign) Save(sess *mgo.Session) error {
	if err := c.pre(); err != nil {
		return err
	}
	_, err := c.Collection(sess).UpsertId(c.ID, c)
	return err
}
