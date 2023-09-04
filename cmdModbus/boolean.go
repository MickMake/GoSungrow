//go:build !(freebsd && amd64)
package cmdModbus

import (
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)


type Boolean struct {
	address Address
	values []bool
}

func (m Boolean) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%d: %v\n", int(m.address) + i, int(m.address) + i, b)
	}
	return ret
}

func (m Boolean) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("%v\t", b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) ReadBool(address Address, quantity Quantity) Boolean {
	var ret Boolean

	for range Only.Once {
		if quantity == 1 {
			var r bool
			r, m.err = m.client.ReadCoil(uint16(address))
			ret = Boolean{address: address, values: []bool{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadCoils(uint16(address), uint16(quantity))
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) ReadDiscreteInput(address Address, quantity Quantity) Boolean {
	var ret Boolean

	for range Only.Once {
		if quantity == 1 {
			var r bool
			r, m.err = m.client.ReadDiscreteInput(uint16(address))
			ret = Boolean{address: address, values: []bool{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadDiscreteInputs(uint16(address), uint16(quantity))
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}
