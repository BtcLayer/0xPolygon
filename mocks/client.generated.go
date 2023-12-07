// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygon/cdk-data-availability/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// GetOffChainData provides a mock function with given fields: ctx, hash
func (_m *Client) GetOffChainData(ctx context.Context, hash common.Hash) ([]byte, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for GetOffChainData")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) ([]byte, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []byte); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetOffChainData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOffChainData'
type Client_GetOffChainData_Call struct {
	*mock.Call
}

// GetOffChainData is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *Client_Expecter) GetOffChainData(ctx interface{}, hash interface{}) *Client_GetOffChainData_Call {
	return &Client_GetOffChainData_Call{Call: _e.mock.On("GetOffChainData", ctx, hash)}
}

func (_c *Client_GetOffChainData_Call) Run(run func(ctx context.Context, hash common.Hash)) *Client_GetOffChainData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *Client_GetOffChainData_Call) Return(_a0 []byte, _a1 error) *Client_GetOffChainData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetOffChainData_Call) RunAndReturn(run func(context.Context, common.Hash) ([]byte, error)) *Client_GetOffChainData_Call {
	_c.Call.Return(run)
	return _c
}

// GetStatus provides a mock function with given fields: ctx
func (_m *Client) GetStatus(ctx context.Context) (*types.DACStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetStatus")
	}

	var r0 *types.DACStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.DACStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.DACStatus); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.DACStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStatus'
type Client_GetStatus_Call struct {
	*mock.Call
}

// GetStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) GetStatus(ctx interface{}) *Client_GetStatus_Call {
	return &Client_GetStatus_Call{Call: _e.mock.On("GetStatus", ctx)}
}

func (_c *Client_GetStatus_Call) Run(run func(ctx context.Context)) *Client_GetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_GetStatus_Call) Return(_a0 *types.DACStatus, _a1 error) *Client_GetStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetStatus_Call) RunAndReturn(run func(context.Context) (*types.DACStatus, error)) *Client_GetStatus_Call {
	_c.Call.Return(run)
	return _c
}

// ListOffChainData provides a mock function with given fields: ctx, hashes
func (_m *Client) ListOffChainData(ctx context.Context, hashes []common.Hash) (map[common.Hash][]byte, error) {
	ret := _m.Called(ctx, hashes)

	if len(ret) == 0 {
		panic("no return value specified for ListOffChainData")
	}

	var r0 map[common.Hash][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []common.Hash) (map[common.Hash][]byte, error)); ok {
		return rf(ctx, hashes)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []common.Hash) map[common.Hash][]byte); ok {
		r0 = rf(ctx, hashes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Hash][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []common.Hash) error); ok {
		r1 = rf(ctx, hashes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListOffChainData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOffChainData'
type Client_ListOffChainData_Call struct {
	*mock.Call
}

// ListOffChainData is a helper method to define mock.On call
//   - ctx context.Context
//   - hashes []common.Hash
func (_e *Client_Expecter) ListOffChainData(ctx interface{}, hashes interface{}) *Client_ListOffChainData_Call {
	return &Client_ListOffChainData_Call{Call: _e.mock.On("ListOffChainData", ctx, hashes)}
}

func (_c *Client_ListOffChainData_Call) Run(run func(ctx context.Context, hashes []common.Hash)) *Client_ListOffChainData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]common.Hash))
	})
	return _c
}

func (_c *Client_ListOffChainData_Call) Return(_a0 map[common.Hash][]byte, _a1 error) *Client_ListOffChainData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListOffChainData_Call) RunAndReturn(run func(context.Context, []common.Hash) (map[common.Hash][]byte, error)) *Client_ListOffChainData_Call {
	_c.Call.Return(run)
	return _c
}

// SignSequence provides a mock function with given fields: ctx, signedSequence
func (_m *Client) SignSequence(ctx context.Context, signedSequence types.SignedSequence) ([]byte, error) {
	ret := _m.Called(ctx, signedSequence)

	if len(ret) == 0 {
		panic("no return value specified for SignSequence")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.SignedSequence) ([]byte, error)); ok {
		return rf(ctx, signedSequence)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.SignedSequence) []byte); ok {
		r0 = rf(ctx, signedSequence)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.SignedSequence) error); ok {
		r1 = rf(ctx, signedSequence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_SignSequence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignSequence'
type Client_SignSequence_Call struct {
	*mock.Call
}

// SignSequence is a helper method to define mock.On call
//   - ctx context.Context
//   - signedSequence types.SignedSequence
func (_e *Client_Expecter) SignSequence(ctx interface{}, signedSequence interface{}) *Client_SignSequence_Call {
	return &Client_SignSequence_Call{Call: _e.mock.On("SignSequence", ctx, signedSequence)}
}

func (_c *Client_SignSequence_Call) Run(run func(ctx context.Context, signedSequence types.SignedSequence)) *Client_SignSequence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.SignedSequence))
	})
	return _c
}

func (_c *Client_SignSequence_Call) Return(_a0 []byte, _a1 error) *Client_SignSequence_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_SignSequence_Call) RunAndReturn(run func(context.Context, types.SignedSequence) ([]byte, error)) *Client_SignSequence_Call {
	_c.Call.Return(run)
	return _c
}

// SignSequenceBanana provides a mock function with given fields: ctx, signedSequence
func (_m *Client) SignSequenceBanana(ctx context.Context, signedSequence types.SignedSequenceBanana) ([]byte, error) {
	ret := _m.Called(ctx, signedSequence)

	if len(ret) == 0 {
		panic("no return value specified for SignSequenceBanana")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.SignedSequenceBanana) ([]byte, error)); ok {
		return rf(ctx, signedSequence)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.SignedSequenceBanana) []byte); ok {
		r0 = rf(ctx, signedSequence)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.SignedSequenceBanana) error); ok {
		r1 = rf(ctx, signedSequence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_SignSequenceBanana_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignSequenceBanana'
type Client_SignSequenceBanana_Call struct {
	*mock.Call
}

// SignSequenceBanana is a helper method to define mock.On call
//   - ctx context.Context
//   - signedSequence types.SignedSequenceBanana
func (_e *Client_Expecter) SignSequenceBanana(ctx interface{}, signedSequence interface{}) *Client_SignSequenceBanana_Call {
	return &Client_SignSequenceBanana_Call{Call: _e.mock.On("SignSequenceBanana", ctx, signedSequence)}
}

func (_c *Client_SignSequenceBanana_Call) Run(run func(ctx context.Context, signedSequence types.SignedSequenceBanana)) *Client_SignSequenceBanana_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.SignedSequenceBanana))
	})
	return _c
}

func (_c *Client_SignSequenceBanana_Call) Return(_a0 []byte, _a1 error) *Client_SignSequenceBanana_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_SignSequenceBanana_Call) RunAndReturn(run func(context.Context, types.SignedSequenceBanana) ([]byte, error)) *Client_SignSequenceBanana_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
