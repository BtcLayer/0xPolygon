// Code generated by mockery. DO NOT EDIT.

package synchronizer

import (
	context "context"

	etherman "github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"
)

// storageMock is an autogenerated mock type for the storageInterface type
type storageMock struct {
	mock.Mock
}

type storageMock_Expecter struct {
	mock *mock.Mock
}

func (_m *storageMock) EXPECT() *storageMock_Expecter {
	return &storageMock_Expecter{mock: &_m.Mock}
}

// AddBlock provides a mock function with given fields: ctx, block, dbTx
func (_m *storageMock) AddBlock(ctx context.Context, block *etherman.Block, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, block, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddBlock")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.Block, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, block, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.Block, pgx.Tx) uint64); ok {
		r0 = rf(ctx, block, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *etherman.Block, pgx.Tx) error); ok {
		r1 = rf(ctx, block, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_AddBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBlock'
type storageMock_AddBlock_Call struct {
	*mock.Call
}

// AddBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - block *etherman.Block
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) AddBlock(ctx interface{}, block interface{}, dbTx interface{}) *storageMock_AddBlock_Call {
	return &storageMock_AddBlock_Call{Call: _e.mock.On("AddBlock", ctx, block, dbTx)}
}

func (_c *storageMock_AddBlock_Call) Run(run func(ctx context.Context, block *etherman.Block, dbTx pgx.Tx)) *storageMock_AddBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*etherman.Block), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_AddBlock_Call) Return(_a0 uint64, _a1 error) *storageMock_AddBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_AddBlock_Call) RunAndReturn(run func(context.Context, *etherman.Block, pgx.Tx) (uint64, error)) *storageMock_AddBlock_Call {
	_c.Call.Return(run)
	return _c
}

// AddClaim provides a mock function with given fields: ctx, claim, dbTx
func (_m *storageMock) AddClaim(ctx context.Context, claim *etherman.Claim, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, claim, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddClaim")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.Claim, pgx.Tx) error); ok {
		r0 = rf(ctx, claim, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageMock_AddClaim_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddClaim'
type storageMock_AddClaim_Call struct {
	*mock.Call
}

// AddClaim is a helper method to define mock.On call
//   - ctx context.Context
//   - claim *etherman.Claim
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) AddClaim(ctx interface{}, claim interface{}, dbTx interface{}) *storageMock_AddClaim_Call {
	return &storageMock_AddClaim_Call{Call: _e.mock.On("AddClaim", ctx, claim, dbTx)}
}

func (_c *storageMock_AddClaim_Call) Run(run func(ctx context.Context, claim *etherman.Claim, dbTx pgx.Tx)) *storageMock_AddClaim_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*etherman.Claim), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_AddClaim_Call) Return(_a0 error) *storageMock_AddClaim_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageMock_AddClaim_Call) RunAndReturn(run func(context.Context, *etherman.Claim, pgx.Tx) error) *storageMock_AddClaim_Call {
	_c.Call.Return(run)
	return _c
}

// AddDeposit provides a mock function with given fields: ctx, deposit, dbTx
func (_m *storageMock) AddDeposit(ctx context.Context, deposit *etherman.Deposit, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, deposit, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddDeposit")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.Deposit, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, deposit, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.Deposit, pgx.Tx) uint64); ok {
		r0 = rf(ctx, deposit, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *etherman.Deposit, pgx.Tx) error); ok {
		r1 = rf(ctx, deposit, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_AddDeposit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddDeposit'
type storageMock_AddDeposit_Call struct {
	*mock.Call
}

// AddDeposit is a helper method to define mock.On call
//   - ctx context.Context
//   - deposit *etherman.Deposit
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) AddDeposit(ctx interface{}, deposit interface{}, dbTx interface{}) *storageMock_AddDeposit_Call {
	return &storageMock_AddDeposit_Call{Call: _e.mock.On("AddDeposit", ctx, deposit, dbTx)}
}

func (_c *storageMock_AddDeposit_Call) Run(run func(ctx context.Context, deposit *etherman.Deposit, dbTx pgx.Tx)) *storageMock_AddDeposit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*etherman.Deposit), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_AddDeposit_Call) Return(_a0 uint64, _a1 error) *storageMock_AddDeposit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_AddDeposit_Call) RunAndReturn(run func(context.Context, *etherman.Deposit, pgx.Tx) (uint64, error)) *storageMock_AddDeposit_Call {
	_c.Call.Return(run)
	return _c
}

// AddGlobalExitRoot provides a mock function with given fields: ctx, exitRoot, dbTx
func (_m *storageMock) AddGlobalExitRoot(ctx context.Context, exitRoot *etherman.GlobalExitRoot, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, exitRoot, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddGlobalExitRoot")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.GlobalExitRoot, pgx.Tx) error); ok {
		r0 = rf(ctx, exitRoot, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageMock_AddGlobalExitRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGlobalExitRoot'
type storageMock_AddGlobalExitRoot_Call struct {
	*mock.Call
}

// AddGlobalExitRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - exitRoot *etherman.GlobalExitRoot
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) AddGlobalExitRoot(ctx interface{}, exitRoot interface{}, dbTx interface{}) *storageMock_AddGlobalExitRoot_Call {
	return &storageMock_AddGlobalExitRoot_Call{Call: _e.mock.On("AddGlobalExitRoot", ctx, exitRoot, dbTx)}
}

func (_c *storageMock_AddGlobalExitRoot_Call) Run(run func(ctx context.Context, exitRoot *etherman.GlobalExitRoot, dbTx pgx.Tx)) *storageMock_AddGlobalExitRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*etherman.GlobalExitRoot), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_AddGlobalExitRoot_Call) Return(_a0 error) *storageMock_AddGlobalExitRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageMock_AddGlobalExitRoot_Call) RunAndReturn(run func(context.Context, *etherman.GlobalExitRoot, pgx.Tx) error) *storageMock_AddGlobalExitRoot_Call {
	_c.Call.Return(run)
	return _c
}

// AddTokenWrapped provides a mock function with given fields: ctx, tokenWrapped, dbTx
func (_m *storageMock) AddTokenWrapped(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, tokenWrapped, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddTokenWrapped")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.TokenWrapped, pgx.Tx) error); ok {
		r0 = rf(ctx, tokenWrapped, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageMock_AddTokenWrapped_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTokenWrapped'
type storageMock_AddTokenWrapped_Call struct {
	*mock.Call
}

// AddTokenWrapped is a helper method to define mock.On call
//   - ctx context.Context
//   - tokenWrapped *etherman.TokenWrapped
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) AddTokenWrapped(ctx interface{}, tokenWrapped interface{}, dbTx interface{}) *storageMock_AddTokenWrapped_Call {
	return &storageMock_AddTokenWrapped_Call{Call: _e.mock.On("AddTokenWrapped", ctx, tokenWrapped, dbTx)}
}

func (_c *storageMock_AddTokenWrapped_Call) Run(run func(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx pgx.Tx)) *storageMock_AddTokenWrapped_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*etherman.TokenWrapped), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_AddTokenWrapped_Call) Return(_a0 error) *storageMock_AddTokenWrapped_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageMock_AddTokenWrapped_Call) RunAndReturn(run func(context.Context, *etherman.TokenWrapped, pgx.Tx) error) *storageMock_AddTokenWrapped_Call {
	_c.Call.Return(run)
	return _c
}

// AddTrustedGlobalExitRoot provides a mock function with given fields: ctx, trustedExitRoot, dbTx
func (_m *storageMock) AddTrustedGlobalExitRoot(ctx context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, trustedExitRoot, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for AddTrustedGlobalExitRoot")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.GlobalExitRoot, pgx.Tx) (bool, error)); ok {
		return rf(ctx, trustedExitRoot, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *etherman.GlobalExitRoot, pgx.Tx) bool); ok {
		r0 = rf(ctx, trustedExitRoot, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *etherman.GlobalExitRoot, pgx.Tx) error); ok {
		r1 = rf(ctx, trustedExitRoot, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_AddTrustedGlobalExitRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTrustedGlobalExitRoot'
type storageMock_AddTrustedGlobalExitRoot_Call struct {
	*mock.Call
}

// AddTrustedGlobalExitRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - trustedExitRoot *etherman.GlobalExitRoot
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) AddTrustedGlobalExitRoot(ctx interface{}, trustedExitRoot interface{}, dbTx interface{}) *storageMock_AddTrustedGlobalExitRoot_Call {
	return &storageMock_AddTrustedGlobalExitRoot_Call{Call: _e.mock.On("AddTrustedGlobalExitRoot", ctx, trustedExitRoot, dbTx)}
}

func (_c *storageMock_AddTrustedGlobalExitRoot_Call) Run(run func(ctx context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx pgx.Tx)) *storageMock_AddTrustedGlobalExitRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*etherman.GlobalExitRoot), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_AddTrustedGlobalExitRoot_Call) Return(_a0 bool, _a1 error) *storageMock_AddTrustedGlobalExitRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_AddTrustedGlobalExitRoot_Call) RunAndReturn(run func(context.Context, *etherman.GlobalExitRoot, pgx.Tx) (bool, error)) *storageMock_AddTrustedGlobalExitRoot_Call {
	_c.Call.Return(run)
	return _c
}

// BeginDBTransaction provides a mock function with given fields: ctx
func (_m *storageMock) BeginDBTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginDBTransaction")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_BeginDBTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginDBTransaction'
type storageMock_BeginDBTransaction_Call struct {
	*mock.Call
}

// BeginDBTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *storageMock_Expecter) BeginDBTransaction(ctx interface{}) *storageMock_BeginDBTransaction_Call {
	return &storageMock_BeginDBTransaction_Call{Call: _e.mock.On("BeginDBTransaction", ctx)}
}

func (_c *storageMock_BeginDBTransaction_Call) Run(run func(ctx context.Context)) *storageMock_BeginDBTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *storageMock_BeginDBTransaction_Call) Return(_a0 pgx.Tx, _a1 error) *storageMock_BeginDBTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_BeginDBTransaction_Call) RunAndReturn(run func(context.Context) (pgx.Tx, error)) *storageMock_BeginDBTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// CheckIfRootExists provides a mock function with given fields: ctx, root, network, dbTx
func (_m *storageMock) CheckIfRootExists(ctx context.Context, root []byte, network uint8, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, root, network, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfRootExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint8, pgx.Tx) (bool, error)); ok {
		return rf(ctx, root, network, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint8, pgx.Tx) bool); ok {
		r0 = rf(ctx, root, network, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, uint8, pgx.Tx) error); ok {
		r1 = rf(ctx, root, network, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_CheckIfRootExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckIfRootExists'
type storageMock_CheckIfRootExists_Call struct {
	*mock.Call
}

// CheckIfRootExists is a helper method to define mock.On call
//   - ctx context.Context
//   - root []byte
//   - network uint8
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) CheckIfRootExists(ctx interface{}, root interface{}, network interface{}, dbTx interface{}) *storageMock_CheckIfRootExists_Call {
	return &storageMock_CheckIfRootExists_Call{Call: _e.mock.On("CheckIfRootExists", ctx, root, network, dbTx)}
}

func (_c *storageMock_CheckIfRootExists_Call) Run(run func(ctx context.Context, root []byte, network uint8, dbTx pgx.Tx)) *storageMock_CheckIfRootExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(uint8), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_CheckIfRootExists_Call) Return(_a0 bool, _a1 error) *storageMock_CheckIfRootExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_CheckIfRootExists_Call) RunAndReturn(run func(context.Context, []byte, uint8, pgx.Tx) (bool, error)) *storageMock_CheckIfRootExists_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields: ctx, dbTx
func (_m *storageMock) Commit(ctx context.Context, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) error); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageMock_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type storageMock_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) Commit(ctx interface{}, dbTx interface{}) *storageMock_Commit_Call {
	return &storageMock_Commit_Call{Call: _e.mock.On("Commit", ctx, dbTx)}
}

func (_c *storageMock_Commit_Call) Run(run func(ctx context.Context, dbTx pgx.Tx)) *storageMock_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_Commit_Call) Return(_a0 error) *storageMock_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageMock_Commit_Call) RunAndReturn(run func(context.Context, pgx.Tx) error) *storageMock_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastBlock provides a mock function with given fields: ctx, networkID, dbTx
func (_m *storageMock) GetLastBlock(ctx context.Context, networkID uint, dbTx pgx.Tx) (*etherman.Block, error) {
	ret := _m.Called(ctx, networkID, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastBlock")
	}

	var r0 *etherman.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, pgx.Tx) (*etherman.Block, error)); ok {
		return rf(ctx, networkID, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, pgx.Tx) *etherman.Block); ok {
		r0 = rf(ctx, networkID, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*etherman.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, pgx.Tx) error); ok {
		r1 = rf(ctx, networkID, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_GetLastBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastBlock'
type storageMock_GetLastBlock_Call struct {
	*mock.Call
}

// GetLastBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - networkID uint
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) GetLastBlock(ctx interface{}, networkID interface{}, dbTx interface{}) *storageMock_GetLastBlock_Call {
	return &storageMock_GetLastBlock_Call{Call: _e.mock.On("GetLastBlock", ctx, networkID, dbTx)}
}

func (_c *storageMock_GetLastBlock_Call) Run(run func(ctx context.Context, networkID uint, dbTx pgx.Tx)) *storageMock_GetLastBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_GetLastBlock_Call) Return(_a0 *etherman.Block, _a1 error) *storageMock_GetLastBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_GetLastBlock_Call) RunAndReturn(run func(context.Context, uint, pgx.Tx) (*etherman.Block, error)) *storageMock_GetLastBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestL1SyncedExitRoot provides a mock function with given fields: ctx, dbTx
func (_m *storageMock) GetLatestL1SyncedExitRoot(ctx context.Context, dbTx pgx.Tx) (*etherman.GlobalExitRoot, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestL1SyncedExitRoot")
	}

	var r0 *etherman.GlobalExitRoot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*etherman.GlobalExitRoot, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *etherman.GlobalExitRoot); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*etherman.GlobalExitRoot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_GetLatestL1SyncedExitRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestL1SyncedExitRoot'
type storageMock_GetLatestL1SyncedExitRoot_Call struct {
	*mock.Call
}

// GetLatestL1SyncedExitRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) GetLatestL1SyncedExitRoot(ctx interface{}, dbTx interface{}) *storageMock_GetLatestL1SyncedExitRoot_Call {
	return &storageMock_GetLatestL1SyncedExitRoot_Call{Call: _e.mock.On("GetLatestL1SyncedExitRoot", ctx, dbTx)}
}

func (_c *storageMock_GetLatestL1SyncedExitRoot_Call) Run(run func(ctx context.Context, dbTx pgx.Tx)) *storageMock_GetLatestL1SyncedExitRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_GetLatestL1SyncedExitRoot_Call) Return(_a0 *etherman.GlobalExitRoot, _a1 error) *storageMock_GetLatestL1SyncedExitRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_GetLatestL1SyncedExitRoot_Call) RunAndReturn(run func(context.Context, pgx.Tx) (*etherman.GlobalExitRoot, error)) *storageMock_GetLatestL1SyncedExitRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumberDeposits provides a mock function with given fields: ctx, origNetworkID, blockNumber, dbTx
func (_m *storageMock) GetNumberDeposits(ctx context.Context, origNetworkID uint, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, origNetworkID, blockNumber, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetNumberDeposits")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint64, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, origNetworkID, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, origNetworkID, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, origNetworkID, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_GetNumberDeposits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumberDeposits'
type storageMock_GetNumberDeposits_Call struct {
	*mock.Call
}

// GetNumberDeposits is a helper method to define mock.On call
//   - ctx context.Context
//   - origNetworkID uint
//   - blockNumber uint64
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) GetNumberDeposits(ctx interface{}, origNetworkID interface{}, blockNumber interface{}, dbTx interface{}) *storageMock_GetNumberDeposits_Call {
	return &storageMock_GetNumberDeposits_Call{Call: _e.mock.On("GetNumberDeposits", ctx, origNetworkID, blockNumber, dbTx)}
}

func (_c *storageMock_GetNumberDeposits_Call) Run(run func(ctx context.Context, origNetworkID uint, blockNumber uint64, dbTx pgx.Tx)) *storageMock_GetNumberDeposits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(uint64), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_GetNumberDeposits_Call) Return(_a0 uint64, _a1 error) *storageMock_GetNumberDeposits_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_GetNumberDeposits_Call) RunAndReturn(run func(context.Context, uint, uint64, pgx.Tx) (uint64, error)) *storageMock_GetNumberDeposits_Call {
	_c.Call.Return(run)
	return _c
}

// GetPreviousBlock provides a mock function with given fields: ctx, networkID, offset, dbTx
func (_m *storageMock) GetPreviousBlock(ctx context.Context, networkID uint, offset uint64, dbTx pgx.Tx) (*etherman.Block, error) {
	ret := _m.Called(ctx, networkID, offset, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for GetPreviousBlock")
	}

	var r0 *etherman.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint64, pgx.Tx) (*etherman.Block, error)); ok {
		return rf(ctx, networkID, offset, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint64, pgx.Tx) *etherman.Block); ok {
		r0 = rf(ctx, networkID, offset, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*etherman.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, networkID, offset, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_GetPreviousBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPreviousBlock'
type storageMock_GetPreviousBlock_Call struct {
	*mock.Call
}

// GetPreviousBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - networkID uint
//   - offset uint64
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) GetPreviousBlock(ctx interface{}, networkID interface{}, offset interface{}, dbTx interface{}) *storageMock_GetPreviousBlock_Call {
	return &storageMock_GetPreviousBlock_Call{Call: _e.mock.On("GetPreviousBlock", ctx, networkID, offset, dbTx)}
}

func (_c *storageMock_GetPreviousBlock_Call) Run(run func(ctx context.Context, networkID uint, offset uint64, dbTx pgx.Tx)) *storageMock_GetPreviousBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(uint64), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_GetPreviousBlock_Call) Return(_a0 *etherman.Block, _a1 error) *storageMock_GetPreviousBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_GetPreviousBlock_Call) RunAndReturn(run func(context.Context, uint, uint64, pgx.Tx) (*etherman.Block, error)) *storageMock_GetPreviousBlock_Call {
	_c.Call.Return(run)
	return _c
}

// IsLxLyActivated provides a mock function with given fields: ctx, dbTx
func (_m *storageMock) IsLxLyActivated(ctx context.Context, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for IsLxLyActivated")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (bool, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) bool); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// storageMock_IsLxLyActivated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsLxLyActivated'
type storageMock_IsLxLyActivated_Call struct {
	*mock.Call
}

// IsLxLyActivated is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) IsLxLyActivated(ctx interface{}, dbTx interface{}) *storageMock_IsLxLyActivated_Call {
	return &storageMock_IsLxLyActivated_Call{Call: _e.mock.On("IsLxLyActivated", ctx, dbTx)}
}

func (_c *storageMock_IsLxLyActivated_Call) Run(run func(ctx context.Context, dbTx pgx.Tx)) *storageMock_IsLxLyActivated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_IsLxLyActivated_Call) Return(_a0 bool, _a1 error) *storageMock_IsLxLyActivated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *storageMock_IsLxLyActivated_Call) RunAndReturn(run func(context.Context, pgx.Tx) (bool, error)) *storageMock_IsLxLyActivated_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields: ctx, blockNumber, networkID, dbTx
func (_m *storageMock) Reset(ctx context.Context, blockNumber uint64, networkID uint, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, blockNumber, networkID, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Reset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint, pgx.Tx) error); ok {
		r0 = rf(ctx, blockNumber, networkID, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageMock_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type storageMock_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
//   - networkID uint
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) Reset(ctx interface{}, blockNumber interface{}, networkID interface{}, dbTx interface{}) *storageMock_Reset_Call {
	return &storageMock_Reset_Call{Call: _e.mock.On("Reset", ctx, blockNumber, networkID, dbTx)}
}

func (_c *storageMock_Reset_Call) Run(run func(ctx context.Context, blockNumber uint64, networkID uint, dbTx pgx.Tx)) *storageMock_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(uint), args[3].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_Reset_Call) Return(_a0 error) *storageMock_Reset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageMock_Reset_Call) RunAndReturn(run func(context.Context, uint64, uint, pgx.Tx) error) *storageMock_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields: ctx, dbTx
func (_m *storageMock) Rollback(ctx context.Context, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) error); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// storageMock_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type storageMock_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - ctx context.Context
//   - dbTx pgx.Tx
func (_e *storageMock_Expecter) Rollback(ctx interface{}, dbTx interface{}) *storageMock_Rollback_Call {
	return &storageMock_Rollback_Call{Call: _e.mock.On("Rollback", ctx, dbTx)}
}

func (_c *storageMock_Rollback_Call) Run(run func(ctx context.Context, dbTx pgx.Tx)) *storageMock_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(pgx.Tx))
	})
	return _c
}

func (_c *storageMock_Rollback_Call) Return(_a0 error) *storageMock_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *storageMock_Rollback_Call) RunAndReturn(run func(context.Context, pgx.Tx) error) *storageMock_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// newStorageMock creates a new instance of storageMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStorageMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *storageMock {
	mock := &storageMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
