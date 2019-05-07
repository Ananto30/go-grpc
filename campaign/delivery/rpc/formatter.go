package rpc

import (
	"magic.pathao.com/nightwing/fortress/campaign"
	pb "magic.pathao.com/nightwing/proto/fortress"
)

func formatAllowedPaymentFromStr(apt []string) []pb.PaymentType {
	r := []pb.PaymentType{}

	m := map[string]pb.PaymentType{}
	m["cash"] = pb.PaymentType_CASH
	m["digital"] = pb.PaymentType_DIGITAL_PAYMENT

	for _, c := range apt {
		r = append(r, m[c])
	}

	return r
}

func toAllowedPaymentStr(pt []pb.PaymentType) []string {
	r := []string{}

	m := map[pb.PaymentType]string{}
	m[pb.PaymentType_CASH] = "cash"
	m[pb.PaymentType_DIGITAL_PAYMENT] = "digital"

	for _, c := range pt {
		r = append(r, m[c])
	}

	return r
}

func toMapStringArray(m map[string]*pb.Document) map[string][]string {
	res := map[string][]string{}
	for k, v := range m {
		res[k] = append(res[k], v.List...)
	}
	return res
}

func formatCampaign(cmp *campaign.Campaign) *pb.Campaign {
	return &pb.Campaign{
		Id:          cmp.ID,
		Title:       cmp.Title,
		Code:        cmp.Code,
		Partner:     cmp.Partner,
		BurnSharing: cmp.BurnSharing,
		Budget:      cmp.Budget,
		StartDate:   cmp.StartDate.Format("2006-01-02"),
		EndDate:     cmp.EndDate.Format("2006-01-02"),
		// Promos:    cmp.Promos,
		// CashBacks: cmp.CashBacks,
		// Discounts: cmp.Discounts,
		CreatedAt: cmp.CreatedAt.Unix(),
		UpdatedAt: cmp.UpdatedAt.Unix(),
	}
}
