package repository

import (
	"github.com/Ananto30/go-grpc/campaign"
	"github.com/Ananto30/go-grpc/model"
	mgo "gopkg.in/mgo.v2"
)

type mgoRepository struct {
	sess *mgo.Session
}

// NewMGORepository ...
func NewMGORepository(sess *mgo.Session) (campaign.Repository, error) {
	ms := &mgoRepository{
		sess,
	}
	return ms, nil

}

// GetByID ...
func (m *mgoRepository) GetByID(id string) (*model.Campaign, error) {
	sess := m.sess.Copy()
	defer sess.Close()

	cmp := model.Campaign{}
	err := model.Campaigns(sess).FindId(id).One(&cmp)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &cmp, nil
}

// Store ...
func (m *mgoRepository) Store(cmp model.Campaign) (*model.Campaign, error) {
	sess := m.sess.Copy()
	defer sess.Close()

	err := cmp.Save(sess)
	if err != nil {
		return nil, err
	}
	return &cmp, nil
}
