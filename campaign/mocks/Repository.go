// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Ananto30/go-grpc/model"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *Repository) GetByID(id string) (*model.Campaign, error) {
	ret := _m.Called(id)

	var r0 *model.Campaign
	if rf, ok := ret.Get(0).(func(string) *model.Campaign); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Campaign)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *Repository) Store(_a0 model.Campaign) (*model.Campaign, error) {
	ret := _m.Called(_a0)

	var r0 *model.Campaign
	if rf, ok := ret.Get(0).(func(model.Campaign) *model.Campaign); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Campaign)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Campaign) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
