// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	rand "math/rand"

	mock "github.com/stretchr/testify/mock"

	random "github.com/flyteorg/flyte/stdlib/random"
)

// WeightedRandomList is an autogenerated mock type for the WeightedRandomList type
type WeightedRandomList struct {
	mock.Mock
}

type WeightedRandomList_Get struct {
	*mock.Call
}

func (_m WeightedRandomList_Get) Return(_a0 random.Comparable) *WeightedRandomList_Get {
	return &WeightedRandomList_Get{Call: _m.Call.Return(_a0)}
}

func (_m *WeightedRandomList) OnGet() *WeightedRandomList_Get {
	c := _m.On("Get")
	return &WeightedRandomList_Get{Call: c}
}

func (_m *WeightedRandomList) OnGetMatch(matchers ...interface{}) *WeightedRandomList_Get {
	c := _m.On("Get", matchers...)
	return &WeightedRandomList_Get{Call: c}
}

// Get provides a mock function with given fields:
func (_m *WeightedRandomList) Get() random.Comparable {
	ret := _m.Called()

	var r0 random.Comparable
	if rf, ok := ret.Get(0).(func() random.Comparable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(random.Comparable)
		}
	}

	return r0
}

type WeightedRandomList_GetWithSeed struct {
	*mock.Call
}

func (_m WeightedRandomList_GetWithSeed) Return(_a0 random.Comparable, _a1 error) *WeightedRandomList_GetWithSeed {
	return &WeightedRandomList_GetWithSeed{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *WeightedRandomList) OnGetWithSeed(seed rand.Source) *WeightedRandomList_GetWithSeed {
	c := _m.On("GetWithSeed", seed)
	return &WeightedRandomList_GetWithSeed{Call: c}
}

func (_m *WeightedRandomList) OnGetWithSeedMatch(matchers ...interface{}) *WeightedRandomList_GetWithSeed {
	c := _m.On("GetWithSeed", matchers...)
	return &WeightedRandomList_GetWithSeed{Call: c}
}

// GetWithSeed provides a mock function with given fields: seed
func (_m *WeightedRandomList) GetWithSeed(seed rand.Source) (random.Comparable, error) {
	ret := _m.Called(seed)

	var r0 random.Comparable
	if rf, ok := ret.Get(0).(func(rand.Source) random.Comparable); ok {
		r0 = rf(seed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(random.Comparable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(rand.Source) error); ok {
		r1 = rf(seed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type WeightedRandomList_Len struct {
	*mock.Call
}

func (_m WeightedRandomList_Len) Return(_a0 int) *WeightedRandomList_Len {
	return &WeightedRandomList_Len{Call: _m.Call.Return(_a0)}
}

func (_m *WeightedRandomList) OnLen() *WeightedRandomList_Len {
	c := _m.On("Len")
	return &WeightedRandomList_Len{Call: c}
}

func (_m *WeightedRandomList) OnLenMatch(matchers ...interface{}) *WeightedRandomList_Len {
	c := _m.On("Len", matchers...)
	return &WeightedRandomList_Len{Call: c}
}

// Len provides a mock function with given fields:
func (_m *WeightedRandomList) Len() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type WeightedRandomList_List struct {
	*mock.Call
}

func (_m WeightedRandomList_List) Return(_a0 []random.Comparable) *WeightedRandomList_List {
	return &WeightedRandomList_List{Call: _m.Call.Return(_a0)}
}

func (_m *WeightedRandomList) OnList() *WeightedRandomList_List {
	c := _m.On("List")
	return &WeightedRandomList_List{Call: c}
}

func (_m *WeightedRandomList) OnListMatch(matchers ...interface{}) *WeightedRandomList_List {
	c := _m.On("List", matchers...)
	return &WeightedRandomList_List{Call: c}
}

// List provides a mock function with given fields:
func (_m *WeightedRandomList) List() []random.Comparable {
	ret := _m.Called()

	var r0 []random.Comparable
	if rf, ok := ret.Get(0).(func() []random.Comparable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]random.Comparable)
		}
	}

	return r0
}
