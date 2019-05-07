package campaign

import (
	"github.com/Ananto30/go-grpc/model"
)

// Usecase ...
type Usecase interface {
	GetByID(id string) (*model.Campaign, error)
	Store(Campaign) (*model.Campaign, error)
}
