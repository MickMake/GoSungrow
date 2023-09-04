package cmdModbus

import (
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/simonvetter/modbus"
)


type Byte struct {
	address Address
	values []byte
}

func (m Byte) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%5d: [%.2X] %3d\n", int(m.address) + i, int(m.address) + i, b, b)
	}
	return ret
}

func (m *Byte) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%5d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("[%.2X]%3d\t", b, b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) ReadByte(address Address, quantity Quantity, regType modbus.RegType) Byte {
	var ret Byte

	for range Only.Once {
		ret.address = address
		ret.values, m.err = m.client.ReadBytes(uint16(address), uint16(quantity), regType)
		// ret, m.err = m.client.ReadRawBytes(uint16(address), uint16(quantity), regType)

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) ReadByteInput(address Address, quantity Quantity) Byte {
	return m.ReadByte(address, quantity, modbus.INPUT_REGISTER)
}

func (m *Modbus) ReadByteHolding(address Address, quantity Quantity) Byte {
	return m.ReadByte(address, quantity, modbus.HOLDING_REGISTER)
}


type U16Bit struct {
	address Address
	values []uint16
}

func (m U16Bit) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%5d: [%.4X] %5d\n", int(m.address) + i, int(m.address) + i, b, b)
	}
	return ret
}

func (m U16Bit) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%5d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("[%.4X]%5d\t", b, b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) Read16Bit(address Address, quantity Quantity, regType modbus.RegType) U16Bit {
	var ret U16Bit

	for range Only.Once {
		if quantity == 1 {
			var r uint16
			r, m.err = m.client.ReadRegister(uint16(address), regType)
			ret = U16Bit{address: address, values: []uint16{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadRegisters(uint16(address), uint16(quantity), regType)
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) Read16BitInput(address Address, quantity Quantity) U16Bit {
	return m.Read16Bit(address, quantity, modbus.INPUT_REGISTER)
}

func (m *Modbus) Read16BitHolding(address Address, quantity Quantity) U16Bit {
	return m.Read16Bit(address, quantity, modbus.HOLDING_REGISTER)
}


type U32Bit struct {
	address Address
	values []uint32
}

func (m U32Bit) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%5d: [%.8X] %10d\n", int(m.address) + i, int(m.address) + i, b, b)
	}
	return ret
}

func (m U32Bit) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%5d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("[%.8X]%10d\t", b, b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) Read32Bit(address Address, quantity Quantity, regType modbus.RegType) U32Bit {
	var ret U32Bit

	for range Only.Once {
		if quantity == 1 {
			var r uint32
			r, m.err = m.client.ReadUint32(uint16(address), regType)
			ret = U32Bit{address: address, values: []uint32{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadUint32s(uint16(address), uint16(quantity), regType)
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) Read32BitInput(address Address, quantity Quantity) U32Bit {
	return m.Read32Bit(address, quantity, modbus.INPUT_REGISTER)
}

func (m *Modbus) Read32BitHolding(address Address, quantity Quantity) U32Bit {
	return m.Read32Bit(address, quantity, modbus.HOLDING_REGISTER)
}


type U64Bit struct {
	address Address
	values []uint64
}

func (m U64Bit) String() string {
	var ret string
	for i, b := range m.values {
		ret += fmt.Sprintf("[%.4X]%5d: [%.16X] %20d\n", int(m.address) + i, int(m.address) + i, b, b)
	}
	return ret
}

func (m U64Bit) Table(width int) string {
	var ret string
	for i, b := range m.values {
		if i % width == 0 {
			ret += fmt.Sprintf("\n[%.4X]%5d:\t", int(m.address) + i, int(m.address) + i)
		}
		ret += fmt.Sprintf("[%.16X]%20d\t", b, b)
	}
	ret += "\n"
	return ret
}

func (m *Modbus) Read64Bit(address Address, quantity Quantity, regType modbus.RegType) U64Bit {
	var ret U64Bit

	for range Only.Once {
		if quantity == 1 {
			var r uint64
			r, m.err = m.client.ReadUint64(uint16(address), regType)
			ret = U64Bit{address: address, values: []uint64{r}}
		} else {
			ret.address = address
			ret.values, m.err = m.client.ReadUint64s(uint16(address), uint16(quantity), regType)
		}

		if m.err != nil {
			m.err = errors.New(fmt.Sprintf("failed to read register %v: %v\n", address, m.err))
			break
		}
	}

	return ret
}

func (m *Modbus) Read64BitInput(address Address, quantity Quantity) U64Bit {
	return m.Read64Bit(address, quantity, modbus.INPUT_REGISTER)
}

func (m *Modbus) Read64BitHolding(address Address, quantity Quantity) U64Bit {
	return m.Read64Bit(address, quantity, modbus.HOLDING_REGISTER)
}
