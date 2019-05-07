package campaign

import "github.com/Ananto30/go-grpc/model"

// Repository ...
type Repository interface {
	GetByID(id string) (*model.Campaign, error)
	Store(model.Campaign) (*model.Campaign, error)
}
