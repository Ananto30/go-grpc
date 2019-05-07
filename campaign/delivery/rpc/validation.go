package rpc

import (
	"time"

	"magic.pathao.com/nightwing/fortress/campaign"
	"magic.pathao.com/nightwing/fortress/helpers"
	pb "magic.pathao.com/nightwing/proto/fortress"
)

// FormatCreateCampaign validates the campaign protobuff and returns data campaign
func formatCreateCampaign(r *pb.ReqCreateCampaign) (*campaign.Campaign, *helpers.ValidationError) {

	errs := helpers.ValidationError{}
	if r.GetTitle() == "" {
		errs.Add("title", "is required")
	}
	if r.GetPartner() == "" {
		errs.Add("partner", "is required")
	}
	if r.GetBudget() == 0 {
		errs.Add("budget", "cannot be zero")
	}
	if r.GetBurnSharing() == 0 {
		errs.Add("burn_sharing", "cannot be zero")
	}
	if r.GetStartDate() == "" {
		errs.Add("start_date", "is required")
	}
	if r.GetEndDate() == "" {
		errs.Add("end_date", "is required")
	}
	sDt, err := time.Parse("2006-01-02", r.GetStartDate())
	if err != nil {
		errs.Add("start_date", "format should be like - 2006-01-02")
	}
	eDt, err := time.Parse("2006-01-02", r.GetEndDate())
	if err != nil {
		errs.Add("end_date", "format should be like - 2006-01-02")
	}
	if len(errs) > 0 {
		return nil, &errs
	}

	c := &campaign.Campaign{
		Title:       r.GetTitle(),
		Code:        r.GetCode(),
		Partner:     r.GetPartner(),
		BurnSharing: r.GetBurnSharing(),
		Budget:      r.GetBudget(),
		Promos:      toMapStringArray(r.GetPromos()),
		CashBacks:   toMapStringArray(r.GetCashBacks()),
		Discounts:   toMapStringArray(r.GetDiscounts()),
		StartDate:   sDt,
		EndDate:     eDt,
	}

	return c, nil
}
