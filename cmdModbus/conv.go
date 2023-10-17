//go:build !(freebsd && amd64)

package cmdModbus

import (
	"strconv"

	"github.com/MickMake/GoUnify/Only"
)

type Address uint16

// func (m Address) String() string {
// 	return fmt.Sprintf("%2X ", m)
// }

func StringToAddress(address string) (Address, error) {
	var ret Address
	var err error
	for range Only.Once {
		var i int64
		i, err = strconv.ParseInt(address, 0, 64)
		if err != nil {
			break
		}
		ret = Address(i)
	}
	return ret, err
}

type Value uint16

// func (m Value) String() string {
// 	return fmt.Sprintf("%2X ", m)
// }

func StringToValue(value string) (Value, error) {
	var ret Value
	var err error
	for range Only.Once {
		var i int64
		i, err = strconv.ParseInt(value, 0, 64)
		if err != nil {
			break
		}
		ret = Value(i)
	}
	return ret, err
}

type Quantity uint16

// func (m Quantity) String() string {
// 	return fmt.Sprintf("%2X ", m)
// }

func StringToQuantity(quantity string) (Quantity, error) {
	var ret Quantity
	var err error
	for range Only.Once {
		var i int64
		i, err = strconv.ParseInt(quantity, 0, 64)
		if err != nil {
			break
		}
		ret = Quantity(i)
	}
	return ret, err
}
