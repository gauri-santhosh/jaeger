// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	storage_v1 "github.com/jaegertracing/jaeger/proto-gen/storage_v1"
)

// SpanReaderPlugin_GetTraceClient is an autogenerated mock type for the SpanReaderPlugin_GetTraceClient type
type SpanReaderPlugin_GetTraceClient struct {
	mock.Mock
}

// CloseSend provides a mock function with given fields:
func (_m *SpanReaderPlugin_GetTraceClient) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *SpanReaderPlugin_GetTraceClient) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Header provides a mock function with given fields:
func (_m *SpanReaderPlugin_GetTraceClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Recv provides a mock function with given fields:
func (_m *SpanReaderPlugin_GetTraceClient) Recv() (*storage_v1.SpansResponseChunk, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *storage_v1.SpansResponseChunk
	var r1 error
	if rf, ok := ret.Get(0).(func() (*storage_v1.SpansResponseChunk, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *storage_v1.SpansResponseChunk); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage_v1.SpansResponseChunk)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecvMsg provides a mock function with given fields: m
func (_m *SpanReaderPlugin_GetTraceClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *SpanReaderPlugin_GetTraceClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trailer provides a mock function with given fields:
func (_m *SpanReaderPlugin_GetTraceClient) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// NewSpanReaderPlugin_GetTraceClient creates a new instance of SpanReaderPlugin_GetTraceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpanReaderPlugin_GetTraceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpanReaderPlugin_GetTraceClient {
	mock := &SpanReaderPlugin_GetTraceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
