/*
 * Copyright (C) 2024 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */
// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TelemetryInitializer is an autogenerated mock type for the TelemetryInitializer type
type TelemetryInitializer struct {
	mock.Mock
}

// InitOpenTelemetry provides a mock function with given fields: ctx
func (_m *TelemetryInitializer) InitOpenTelemetry(ctx context.Context) (func(context.Context) error, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for InitOpenTelemetry")
	}

	var r0 func(context.Context) error
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (func(context.Context) error, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(context.Context) error)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTelemetryInitializer creates a new instance of TelemetryInitializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTelemetryInitializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *TelemetryInitializer {
	mock := &TelemetryInitializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
