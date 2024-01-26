// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	datacatalog "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/datacatalog"

	mock "github.com/stretchr/testify/mock"
)

// ArtifactManager is an autogenerated mock type for the ArtifactManager type
type ArtifactManager struct {
	mock.Mock
}

type ArtifactManager_CreateArtifact struct {
	*mock.Call
}

func (_m ArtifactManager_CreateArtifact) Return(_a0 *datacatalog.CreateArtifactResponse, _a1 error) *ArtifactManager_CreateArtifact {
	return &ArtifactManager_CreateArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactManager) OnCreateArtifact(ctx context.Context, request *datacatalog.CreateArtifactRequest) *ArtifactManager_CreateArtifact {
	c_call := _m.On("CreateArtifact", ctx, request)
	return &ArtifactManager_CreateArtifact{Call: c_call}
}

func (_m *ArtifactManager) OnCreateArtifactMatch(matchers ...interface{}) *ArtifactManager_CreateArtifact {
	c_call := _m.On("CreateArtifact", matchers...)
	return &ArtifactManager_CreateArtifact{Call: c_call}
}

// CreateArtifact provides a mock function with given fields: ctx, request
func (_m *ArtifactManager) CreateArtifact(ctx context.Context, request *datacatalog.CreateArtifactRequest) (*datacatalog.CreateArtifactResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *datacatalog.CreateArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.CreateArtifactRequest) *datacatalog.CreateArtifactResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.CreateArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.CreateArtifactRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactManager_DeleteArtifact struct {
	*mock.Call
}

func (_m ArtifactManager_DeleteArtifact) Return(_a0 *datacatalog.DeleteArtifactResponse, _a1 error) *ArtifactManager_DeleteArtifact {
	return &ArtifactManager_DeleteArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactManager) OnDeleteArtifact(ctx context.Context, request *datacatalog.DeleteArtifactRequest) *ArtifactManager_DeleteArtifact {
	c_call := _m.On("DeleteArtifact", ctx, request)
	return &ArtifactManager_DeleteArtifact{Call: c_call}
}

func (_m *ArtifactManager) OnDeleteArtifactMatch(matchers ...interface{}) *ArtifactManager_DeleteArtifact {
	c_call := _m.On("DeleteArtifact", matchers...)
	return &ArtifactManager_DeleteArtifact{Call: c_call}
}

// DeleteArtifact provides a mock function with given fields: ctx, request
func (_m *ArtifactManager) DeleteArtifact(ctx context.Context, request *datacatalog.DeleteArtifactRequest) (*datacatalog.DeleteArtifactResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *datacatalog.DeleteArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.DeleteArtifactRequest) *datacatalog.DeleteArtifactResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.DeleteArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.DeleteArtifactRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactManager_GetArtifact struct {
	*mock.Call
}

func (_m ArtifactManager_GetArtifact) Return(_a0 *datacatalog.GetArtifactResponse, _a1 error) *ArtifactManager_GetArtifact {
	return &ArtifactManager_GetArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactManager) OnGetArtifact(ctx context.Context, request *datacatalog.GetArtifactRequest) *ArtifactManager_GetArtifact {
	c_call := _m.On("GetArtifact", ctx, request)
	return &ArtifactManager_GetArtifact{Call: c_call}
}

func (_m *ArtifactManager) OnGetArtifactMatch(matchers ...interface{}) *ArtifactManager_GetArtifact {
	c_call := _m.On("GetArtifact", matchers...)
	return &ArtifactManager_GetArtifact{Call: c_call}
}

// GetArtifact provides a mock function with given fields: ctx, request
func (_m *ArtifactManager) GetArtifact(ctx context.Context, request *datacatalog.GetArtifactRequest) (*datacatalog.GetArtifactResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *datacatalog.GetArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetArtifactRequest) *datacatalog.GetArtifactResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetArtifactRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactManager_ListArtifacts struct {
	*mock.Call
}

func (_m ArtifactManager_ListArtifacts) Return(_a0 *datacatalog.ListArtifactsResponse, _a1 error) *ArtifactManager_ListArtifacts {
	return &ArtifactManager_ListArtifacts{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactManager) OnListArtifacts(ctx context.Context, request *datacatalog.ListArtifactsRequest) *ArtifactManager_ListArtifacts {
	c_call := _m.On("ListArtifacts", ctx, request)
	return &ArtifactManager_ListArtifacts{Call: c_call}
}

func (_m *ArtifactManager) OnListArtifactsMatch(matchers ...interface{}) *ArtifactManager_ListArtifacts {
	c_call := _m.On("ListArtifacts", matchers...)
	return &ArtifactManager_ListArtifacts{Call: c_call}
}

// ListArtifacts provides a mock function with given fields: ctx, request
func (_m *ArtifactManager) ListArtifacts(ctx context.Context, request *datacatalog.ListArtifactsRequest) (*datacatalog.ListArtifactsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *datacatalog.ListArtifactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.ListArtifactsRequest) *datacatalog.ListArtifactsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.ListArtifactsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.ListArtifactsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactManager_UpdateArtifact struct {
	*mock.Call
}

func (_m ArtifactManager_UpdateArtifact) Return(_a0 *datacatalog.UpdateArtifactResponse, _a1 error) *ArtifactManager_UpdateArtifact {
	return &ArtifactManager_UpdateArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactManager) OnUpdateArtifact(ctx context.Context, request *datacatalog.UpdateArtifactRequest) *ArtifactManager_UpdateArtifact {
	c_call := _m.On("UpdateArtifact", ctx, request)
	return &ArtifactManager_UpdateArtifact{Call: c_call}
}

func (_m *ArtifactManager) OnUpdateArtifactMatch(matchers ...interface{}) *ArtifactManager_UpdateArtifact {
	c_call := _m.On("UpdateArtifact", matchers...)
	return &ArtifactManager_UpdateArtifact{Call: c_call}
}

// UpdateArtifact provides a mock function with given fields: ctx, request
func (_m *ArtifactManager) UpdateArtifact(ctx context.Context, request *datacatalog.UpdateArtifactRequest) (*datacatalog.UpdateArtifactResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *datacatalog.UpdateArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.UpdateArtifactRequest) *datacatalog.UpdateArtifactResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.UpdateArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.UpdateArtifactRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
