// Code generated by mockery. DO NOT EDIT.

package controller_mock

import mock "github.com/stretchr/testify/mock"

// LocalizationMock is an autogenerated mock type for the Localization type
type LocalizationMock struct {
	mock.Mock
}

type LocalizationMock_Expecter struct {
	mock *mock.Mock
}

func (_m *LocalizationMock) EXPECT() *LocalizationMock_Expecter {
	return &LocalizationMock_Expecter{mock: &_m.Mock}
}

// Localize provides a mock function with given fields: code, locale, vars
func (_m *LocalizationMock) Localize(code string, locale string, vars map[string]string) (string, error) {
	ret := _m.Called(code, locale, vars)

	if len(ret) == 0 {
		panic("no return value specified for Localize")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) (string, error)); ok {
		return rf(code, locale, vars)
	}
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) string); ok {
		r0 = rf(code, locale, vars)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, map[string]string) error); ok {
		r1 = rf(code, locale, vars)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalizationMock_Localize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Localize'
type LocalizationMock_Localize_Call struct {
	*mock.Call
}

// Localize is a helper method to define mock.On call
//   - code string
//   - locale string
//   - vars map[string]string
func (_e *LocalizationMock_Expecter) Localize(code interface{}, locale interface{}, vars interface{}) *LocalizationMock_Localize_Call {
	return &LocalizationMock_Localize_Call{Call: _e.mock.On("Localize", code, locale, vars)}
}

func (_c *LocalizationMock_Localize_Call) Run(run func(code string, locale string, vars map[string]string)) *LocalizationMock_Localize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(map[string]string))
	})
	return _c
}

func (_c *LocalizationMock_Localize_Call) Return(_a0 string, _a1 error) *LocalizationMock_Localize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LocalizationMock_Localize_Call) RunAndReturn(run func(string, string, map[string]string) (string, error)) *LocalizationMock_Localize_Call {
	_c.Call.Return(run)
	return _c
}

// NewLocalizationMock creates a new instance of LocalizationMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLocalizationMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *LocalizationMock {
	mock := &LocalizationMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
