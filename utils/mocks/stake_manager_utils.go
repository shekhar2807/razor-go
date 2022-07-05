// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "razor/core/types"
)

// StakeManagerUtils is an autogenerated mock type for the StakeManagerUtils type
type StakeManagerUtils struct {
	mock.Mock
}

// EpochLimitForUpdateCommission provides a mock function with given fields: client
func (_m *StakeManagerUtils) EpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	ret := _m.Called(client)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint16); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumStakers provides a mock function with given fields: client
func (_m *StakeManagerUtils) GetNumStakers(client *ethclient.Client) (uint32, error) {
	ret := _m.Called(client)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStaker provides a mock function with given fields: client, stakerId
func (_m *StakeManagerUtils) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(client, stakerId)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsStaker); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerId provides a mock function with given fields: client, address
func (_m *StakeManagerUtils) GetStakerId(client *ethclient.Client, address common.Address) (uint32, error) {
	ret := _m.Called(client, address)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) uint32); ok {
		r0 = rf(client, address)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, common.Address) error); ok {
		r1 = rf(client, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Locks provides a mock function with given fields: client, address, address1, lockType
func (_m *StakeManagerUtils) Locks(client *ethclient.Client, address common.Address, address1 common.Address, lockType uint8) (types.Locks, error) {
	ret := _m.Called(client, address, address1, lockType)

	var r0 types.Locks
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address, common.Address, uint8) types.Locks); ok {
		r0 = rf(client, address, address1, lockType)
	} else {
		r0 = ret.Get(0).(types.Locks)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, common.Address, common.Address, uint8) error); ok {
		r1 = rf(client, address, address1, lockType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxCommission provides a mock function with given fields: client
func (_m *StakeManagerUtils) MaxCommission(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinSafeRazor provides a mock function with given fields: client
func (_m *StakeManagerUtils) MinSafeRazor(client *ethclient.Client) (*big.Int, error) {
	ret := _m.Called(client)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *big.Int); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawInitiationPeriod provides a mock function with given fields: client
func (_m *StakeManagerUtils) WithdrawInitiationPeriod(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithdrawLockPeriod provides a mock function with given fields: client
func (_m *StakeManagerUtils) WithdrawLockPeriod(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStakeManagerUtils interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakeManagerUtils creates a new instance of StakeManagerUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakeManagerUtils(t mockConstructorTestingTNewStakeManagerUtils) *StakeManagerUtils {
	mock := &StakeManagerUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
