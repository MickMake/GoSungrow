//go:build !(freebsd && amd64)
package cmdModbus

import (
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/simonvetter/modbus"
)


type Float32 struct {
	address Address
	values []float32
}

func (m Float32) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%d: [%.8X] %f\n", int(m.address) + i, int(m.address) + i, b, b)
	}
	return ret
}

func (m Float32) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("[%.8X]%f\t", b, b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) ReadFloat32(address Address, quantity Quantity, regType modbus.RegType) Float32 {
	var ret Float32

	for range Only.Once {
		if quantity == 1 {
			var r float32
			r, m.err = m.client.ReadFloat32(uint16(address), regType)
			ret = Float32{address: address, values: []float32{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadFloat32s(uint16(address), uint16(quantity), regType)
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) ReadFloat32Input(address Address, quantity Quantity) Float32 {
	return m.ReadFloat32(address, quantity, modbus.INPUT_REGISTER)
}

func (m *Modbus) ReadFloat32Holding(address Address, quantity Quantity) Float32 {
	return m.ReadFloat32(address, quantity, modbus.HOLDING_REGISTER)
}


type Float64 struct {
	address Address
	values []float64
}

func (m Float64) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%d: [%.16X] %f\n", int(m.address) + i, int(m.address) + i, b, b)
	}
	return ret
}

func (m Float64) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("[%.16X]%f\t", b, b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) ReadFloat64(address Address, quantity Quantity, regType modbus.RegType) Float64 {
	var ret Float64

	for range Only.Once {
		if quantity == 1 {
			var r float64
			r, m.err = m.client.ReadFloat64(uint16(address), regType)
			ret = Float64{address: address, values: []float64{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadFloat64s(uint16(address), uint16(quantity), regType)
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) ReadFloat64Input(address Address, quantity Quantity) Float64 {
	return m.ReadFloat64(address, quantity, modbus.INPUT_REGISTER)
}

func (m *Modbus) ReadFloat64Holding(address Address, quantity Quantity) Float64 {
	return m.ReadFloat64(address, quantity, modbus.HOLDING_REGISTER)
}
