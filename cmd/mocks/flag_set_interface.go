// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	pflag "github.com/spf13/pflag"
	mock "github.com/stretchr/testify/mock"
)

// FlagSetInterface is an autogenerated mock type for the FlagSetInterface type
type FlagSetInterface struct {
	mock.Mock
}

// GetBoolAutoVote provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetBoolAutoVote(flagSet *pflag.FlagSet) (bool, error) {
	ret := _m.Called(flagSet)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) bool); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBoolRogue provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetBoolRogue(flagSet *pflag.FlagSet) (bool, error) {
	ret := _m.Called(flagSet)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) bool); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFloat32GasLimit provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error) {
	ret := _m.Called(flagSet)

	var r0 float32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) float32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFloat32GasMultiplier provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error) {
	ret := _m.Called(flagSet)

	var r0 float32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) float32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt32Buffer provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error) {
	ret := _m.Called(flagSet)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt32GasPrice provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error) {
	ret := _m.Called(flagSet)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt32Wait provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetInt32Wait(flagSet *pflag.FlagSet) (int32, error) {
	ret := _m.Called(flagSet)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt8Power provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	ret := _m.Called(flagSet)

	var r0 int8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootFloat32GasLimit provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootFloat32GasLimit() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootFloat32GasMultiplier provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootFloat32GasMultiplier() (float32, error) {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootInt32Buffer provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootInt32Buffer() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootInt32GasPrice provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootInt32GasPrice() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootInt32Wait provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootInt32Wait() (int32, error) {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootStringLogLevel provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootStringLogLevel() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootStringProvider provides a mock function with given fields:
func (_m *FlagSetInterface) GetRootStringProvider() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringAddress provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringCertFile provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringCertFile(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringCertKey provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringCertKey(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringExposeMetrics provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringExposeMetrics(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringFrom provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringLogLevel provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringLogLevel(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringName provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringPow provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringPow(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringProvider provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringProvider(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSelector provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSliceRogueMode provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringSliceRogueMode(flagSet *pflag.FlagSet) ([]string, error) {
	ret := _m.Called(flagSet)

	var r0 []string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []string); ok {
		r0 = rf(flagSet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringStatus provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringTo provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringUrl provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringValue provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringValue(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16AssetId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint16AssetId(flagSet *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(flagSet)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16CollectionId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint16CollectionId(flagSet *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(flagSet)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16JobId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint16JobId(flagSet *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(flagSet)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32Aggregation provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32BountyId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32StakerId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32Tolerance provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32Tolerance(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8Commission provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(flagSet)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8SelectorType provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(flagSet)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8Weight provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(flagSet)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUintSliceJobIds provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	ret := _m.Called(flagSet)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []uint); ok {
		r0 = rf(flagSet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFlagSetInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewFlagSetInterface creates a new instance of FlagSetInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFlagSetInterface(t mockConstructorTestingTNewFlagSetInterface) *FlagSetInterface {
	mock := &FlagSetInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
