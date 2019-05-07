package campaign

import (
	"time"

	"github.com/Ananto30/go-grpc/model"
)

// Campaign represents a Promotional Campaign
type Campaign struct {
	ID                 string
	Title              string
	Code               string
	AllowedPaymentType []string
	Partner            string
	BurnSharing        float64
	Budget             int64
	StartDate          time.Time
	EndDate            time.Time
	Promos             map[string][]string
	CashBacks          map[string][]string
	Discounts          map[string][]string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// ToModel ...
func (cmp *Campaign) ToModel() *model.Campaign {
	return &model.Campaign{
		Title:              cmp.Title,
		Code:               cmp.Code,
		AllowedPaymentType: cmp.AllowedPaymentType,
		Partner:            cmp.Partner,
		BurnSharing:        cmp.BurnSharing,
		Budget:             cmp.Budget,
		StartDate:          cmp.StartDate,
		EndDate:            cmp.EndDate,
		Promos:             cmp.Promos,
		CashBacks:          cmp.CashBacks,
		// Discounts:          cmp.Discounts,
	}
}

func FormatCampaign(m *model.Campaign) *Campaign {
	return &Campaign{
		ID:                 m.ID.Hex(),
		Title:              m.Title,
		Code:               m.Code,
		AllowedPaymentType: m.AllowedPaymentType,
		Partner:            m.Partner,
		BurnSharing:        m.BurnSharing,
		Budget:             m.Budget,
		StartDate:          m.StartDate,
		EndDate:            m.EndDate,
		Promos:             m.Promos,
		CashBacks:          m.CashBacks,
		// Discounts:          m.Discounts,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
