package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/Ananto30/go-grpc/campaign/mocks"
	pb2 "github.com/Ananto30/go-grpc/proto/common"
	pb "github.com/Ananto30/go-grpc/proto/fortress"
	"google.golang.org/grpc"
)

func TestGetCampaign(t *testing.T) {
	// md := data.NewMockData()
	// srvr := NewServer(md)
	// req := pb.ReqGetCampaign{
	// 	Id: "PATHAORIDES",
	// }

	// res, err := srvr.GetCampaign(context.Background(), &req)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if res.GetId() != req.GetId() {
	// 	t.Error("Campaign ID should be:", res.GetId())
	// }
	// if res.GetName() != "Pathao Rides Campaigns" {
	// 	t.Error("Campaign Name should be: 'Pathao Rides Campaign' not", res.GetName())
	// }

	// t.Log("Campaign-ID:", res.GetId())
	// t.Log("Campaign Name:", res.GetName())

}

func TestToAllowedPaymentStr(t *testing.T) {
	r := toAllowedPaymentStr([]pb.PaymentType{pb.PaymentType_CASH, pb.PaymentType_DIGITAL_PAYMENT})
	if !((r[0] == "cash" && r[1] == "digital") || (r[1] == "cash" && r[0] == "digital")) {
		t.Fatal("toAllowedPaymentStr() should map appropriately")
	}

	t.Log(r)
}

func TestCreateCampaign(t *testing.T) {
	// repo := new(mocks.Repository)
	uc := new(mocks.Usecase)
	s := grpc.NewServer()
	srvr := NewCampaignServer(s, uc)

	// Validation test
	req := pb.ReqCreateCampaign{
		Title:       "",
		Code:        "TEST",
		Partner:     "Partner-ID",
		BurnSharing: 0,
		Budget:      10000000,
		Promos: map[string]*pb.Document{
			"food": &pb.Document{
				List: []string{"ABCD", "ACBD", "BACD"},
			},
		},
		StartDate: "2019-01-01:12:00:01",
	}
	res, err := srvr.CreateCampaign(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	if res.GetError() == nil {
		t.Fatal("Campaign error response to CreateCampaign is <nil>")
	}
	details := []pb2.Error_Detail{
		pb2.Error_Detail{Field: "title", Messages: []string{"is required"}},
		pb2.Error_Detail{Field: "burn_sharing", Messages: []string{"cannot be zero"}},
		pb2.Error_Detail{Field: "start_date", Messages: []string{"format should be like - 2006-01-02"}},
		pb2.Error_Detail{Field: "end_date", Messages: []string{"is required", "format should be like - 2006-01-02"}},
	}
	if res.GetError().Details == nil {
		t.Fatal("Validation check for error details are missing")
	}
	for i := range details {
		found := false
		for j := range res.GetError().Details {
			if details[i].Field == res.GetError().Details[j].Field {
				found = true
			}
		}
		if found == false {
			t.Fatal("Validation field not found ", details[i].Field, " ", res.GetError().Details[i].Field)
		}
	}

	// Creation test
	req = pb.ReqCreateCampaign{
		Title:       "Test Campaign",
		Code:        "TEST",
		Partner:     "Partner-ID",
		BurnSharing: 15.5,
		Budget:      10000000,
		Promos: map[string]*pb.Document{
			"food": &pb.Document{
				List: []string{"ABCD", "ACBD", "BACD"},
			},
		},
		StartDate: "2019-01-01",
		EndDate:   "2019-01-31",
	}

	res, err = srvr.CreateCampaign(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	if res.GetCampaign() == nil {
		t.Fatal("Campaign response to CreateCampaign is <nil>")
	}
	if res.GetCampaign().Code != req.GetCode() {
		t.Error("Campaign code didn't match. Expected: '", req.GetCode(), "', Got '", res.GetCampaign().GetCode(), "'")
	}
	// t.Log(res)
	// spew.Dump(res)
}

func TestTime(t *testing.T) {
	tme, err := time.Parse("2006-01-02", "2019-01-31")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tme)
}
