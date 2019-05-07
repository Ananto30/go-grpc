package helpers

import (
	"github.com/satori/go.uuid"
	pb "magic.pathao.com/nightwing/proto/common"
)

type ValidationError map[string][]string

func (v ValidationError) Error() *pb.Error {
	details := []*pb.Error_Detail{}
	for k, v := range v {
		detail := &pb.Error_Detail{
			Field:    k,
			Messages: v,
		}
		details = append(details, detail)
	}
	return &pb.Error{
		Id:      uuid.Must(uuid.NewV4()).String(),
		Code:    "422001",
		Title:   "Validation error",
		Details: details,
	}
}

func (v ValidationError) Add(key string, val string) {
	v[key] = append(v[key], val)
}

func (v ValidationError) Extend(prefix string, err *ValidationError) {
	if err == nil {
		return
	}
	for k, e := range *err {
		k = prefix + k
		v[k] = append(v[k], e...)
	}
}

func FormatError(err error) *pb.Error {
	return &pb.Error{
		Id:    uuid.Must(uuid.NewV4()).String(),
		Code:  "500001",
		Title: "Internal Server Error",
	}
}
