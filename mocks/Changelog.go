// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	entry "github.com/chelnak/gh-changelog/pkg/entry"
	mock "github.com/stretchr/testify/mock"
)

// Changelog is an autogenerated mock type for the Changelog type
type Changelog struct {
	mock.Mock
}

// AddUnreleased provides a mock function with given fields: _a0
func (_m *Changelog) AddUnreleased(_a0 []string) {
	_m.Called(_a0)
}

// GetEntries provides a mock function with given fields:
func (_m *Changelog) GetEntries() []*entry.Entry {
	ret := _m.Called()

	var r0 []*entry.Entry
	if rf, ok := ret.Get(0).(func() []*entry.Entry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entry.Entry)
		}
	}

	return r0
}

// GetRepoName provides a mock function with given fields:
func (_m *Changelog) GetRepoName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRepoOwner provides a mock function with given fields:
func (_m *Changelog) GetRepoOwner() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUnreleased provides a mock function with given fields:
func (_m *Changelog) GetUnreleased() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Head provides a mock function with given fields:
func (_m *Changelog) Head() *entry.Entry {
	ret := _m.Called()

	var r0 *entry.Entry
	if rf, ok := ret.Get(0).(func() *entry.Entry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entry.Entry)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: _a0
func (_m *Changelog) Insert(_a0 entry.Entry) {
	_m.Called(_a0)
}

// Tail provides a mock function with given fields:
func (_m *Changelog) Tail() *entry.Entry {
	ret := _m.Called()

	var r0 *entry.Entry
	if rf, ok := ret.Get(0).(func() *entry.Entry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entry.Entry)
		}
	}

	return r0
}

type mockConstructorTestingTNewChangelog interface {
	mock.TestingT
	Cleanup(func())
}

// NewChangelog creates a new instance of Changelog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChangelog(t mockConstructorTestingTNewChangelog) *Changelog {
	mock := &Changelog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
