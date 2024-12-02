package mocks

import mock "github.com/stretchr/testify/mock"

// MockRowsDB is an autogenerated mock type for the RowsDB type
type MockRowsDB struct {
	mock.Mock
}

type MockRowsDB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRowsDB) EXPECT() *MockRowsDB_Expecter {
	return &MockRowsDB_Expecter{mock: &_m.Mock}
}

// Next provides a mock function with given fields:
func (_m *MockRowsDB) Next() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Next")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRowsDB_Next_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Next'
type MockRowsDB_Next_Call struct {
	*mock.Call
}

// Next is a helper method to define mock.On call
func (_e *MockRowsDB_Expecter) Next() *MockRowsDB_Next_Call {
	return &MockRowsDB_Next_Call{Call: _e.mock.On("Next")}
}

func (_c *MockRowsDB_Next_Call) Run(run func()) *MockRowsDB_Next_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRowsDB_Next_Call) Return(_a0 bool) *MockRowsDB_Next_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRowsDB_Next_Call) RunAndReturn(run func() bool) *MockRowsDB_Next_Call {
	_c.Call.Return(run)
	return _c
}

// Scan provides a mock function with given fields: dest
func (_m *MockRowsDB) Scan(dest ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Scan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dest...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRowsDB_Scan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Scan'
type MockRowsDB_Scan_Call struct {
	*mock.Call
}

// Scan is a helper method to define mock.On call
//   - dest ...interface{}
func (_e *MockRowsDB_Expecter) Scan(dest ...interface{}) *MockRowsDB_Scan_Call {
	return &MockRowsDB_Scan_Call{Call: _e.mock.On("Scan",
		append([]interface{}{}, dest...)...)}
}

func (_c *MockRowsDB_Scan_Call) Run(run func(dest ...interface{})) *MockRowsDB_Scan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockRowsDB_Scan_Call) Return(_a0 error) *MockRowsDB_Scan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRowsDB_Scan_Call) RunAndReturn(run func(...interface{}) error) *MockRowsDB_Scan_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRowsDB creates a new instance of MockRowsDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRowsDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRowsDB {
	mock := &MockRowsDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}