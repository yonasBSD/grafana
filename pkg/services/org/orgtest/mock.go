// Code generated by mockery v2.42.1. DO NOT EDIT.

package orgtest

import (
	context "context"

	org "github.com/grafana/grafana/pkg/services/org"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddOrgUser provides a mock function with given fields: _a0, _a1
func (_m *MockService) AddOrgUser(_a0 context.Context, _a1 *org.AddOrgUserCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddOrgUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.AddOrgUserCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWithMember provides a mock function with given fields: _a0, _a1
func (_m *MockService) CreateWithMember(_a0 context.Context, _a1 *org.CreateOrgCommand) (*org.Org, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateWithMember")
	}

	var r0 *org.Org
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.CreateOrgCommand) (*org.Org, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.CreateOrgCommand) *org.Org); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*org.Org)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.CreateOrgCommand) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockService) Delete(_a0 context.Context, _a1 *org.DeleteOrgCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.DeleteOrgCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserFromAll provides a mock function with given fields: _a0, _a1
func (_m *MockService) DeleteUserFromAll(_a0 context.Context, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserFromAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetByID(_a0 context.Context, _a1 *org.GetOrgByIDQuery) (*org.Org, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *org.Org
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetOrgByIDQuery) (*org.Org, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetOrgByIDQuery) *org.Org); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*org.Org)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.GetOrgByIDQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetByName(_a0 context.Context, _a1 *org.GetOrgByNameQuery) (*org.Org, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 *org.Org
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetOrgByNameQuery) (*org.Org, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetOrgByNameQuery) *org.Org); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*org.Org)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.GetOrgByNameQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIDForNewUser provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetIDForNewUser(_a0 context.Context, _a1 org.GetOrgIDForNewUserCommand) (int64, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetIDForNewUser")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, org.GetOrgIDForNewUserCommand) (int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, org.GetOrgIDForNewUserCommand) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, org.GetOrgIDForNewUserCommand) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrCreate provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetOrCreate(_a0 context.Context, _a1 string) (int64, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetOrCreate")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrgUsers provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetOrgUsers(_a0 context.Context, _a1 *org.GetOrgUsersQuery) ([]*org.OrgUserDTO, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetOrgUsers")
	}

	var r0 []*org.OrgUserDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetOrgUsersQuery) ([]*org.OrgUserDTO, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetOrgUsersQuery) []*org.OrgUserDTO); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*org.OrgUserDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.GetOrgUsersQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserOrgList provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetUserOrgList(_a0 context.Context, _a1 *org.GetUserOrgListQuery) ([]*org.UserOrgDTO, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUserOrgList")
	}

	var r0 []*org.UserOrgDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetUserOrgListQuery) ([]*org.UserOrgDTO, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.GetUserOrgListQuery) []*org.UserOrgDTO); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*org.UserOrgDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.GetUserOrgListQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOrgUser provides a mock function with given fields: _a0, _a1
func (_m *MockService) InsertOrgUser(_a0 context.Context, _a1 *org.OrgUser) (int64, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for InsertOrgUser")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.OrgUser) (int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.OrgUser) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.OrgUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDelete provides a mock function with given fields: query
func (_m *MockService) RegisterDelete(query string) {
	_m.Called(query)
}

// RemoveOrgUser provides a mock function with given fields: _a0, _a1
func (_m *MockService) RemoveOrgUser(_a0 context.Context, _a1 *org.RemoveOrgUserCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RemoveOrgUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.RemoveOrgUserCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockService) Search(_a0 context.Context, _a1 *org.SearchOrgsQuery) ([]*org.OrgDTO, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []*org.OrgDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.SearchOrgsQuery) ([]*org.OrgDTO, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.SearchOrgsQuery) []*org.OrgDTO); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*org.OrgDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.SearchOrgsQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchOrgUsers provides a mock function with given fields: _a0, _a1
func (_m *MockService) SearchOrgUsers(_a0 context.Context, _a1 *org.SearchOrgUsersQuery) (*org.SearchOrgUsersQueryResult, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SearchOrgUsers")
	}

	var r0 *org.SearchOrgUsersQueryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.SearchOrgUsersQuery) (*org.SearchOrgUsersQueryResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *org.SearchOrgUsersQuery) *org.SearchOrgUsersQueryResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*org.SearchOrgUsersQueryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *org.SearchOrgUsersQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAddress provides a mock function with given fields: _a0, _a1
func (_m *MockService) UpdateAddress(_a0 context.Context, _a1 *org.UpdateOrgAddressCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAddress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.UpdateOrgAddressCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrg provides a mock function with given fields: _a0, _a1
func (_m *MockService) UpdateOrg(_a0 context.Context, _a1 *org.UpdateOrgCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.UpdateOrgCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrgUser provides a mock function with given fields: _a0, _a1
func (_m *MockService) UpdateOrgUser(_a0 context.Context, _a1 *org.UpdateOrgUserCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrgUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *org.UpdateOrgUserCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}