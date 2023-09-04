package cmdModbus

import (
	"github.com/MickMake/GoUnify/Only"
	"github.com/simonvetter/modbus"
)

const (
	TypeSigned8Bit    = "S8"
	TypeSigned16Bit   = "S16"
	TypeSigned32Bit   = "S32"
	TypeSigned64Bit   = "S64"

	TypeUnsigned8Bit  = "U8"
	TypeUnsigned16Bit = "U16"
	TypeUnsigned32Bit = "U32"
	TypeUnsigned64Bit = "U64"

	TypeFloat32Bit    = "F32"
	TypeFloat64Bit    = "F64"

	TypeString        = "UTF-8"
)


func (m *Modbus) Read(address Address, quantity Quantity, valueType string) string {
	var ret string

	for range Only.Once {
		width := 8

		switch valueType {
			case TypeUnsigned8Bit:
				fallthrough
			case TypeSigned8Bit:
				v := m.ReadByte(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeUnsigned16Bit:
				fallthrough
			case TypeSigned16Bit:
				v := m.Read16Bit(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeUnsigned32Bit:
				fallthrough
			case TypeSigned32Bit:
				v := m.Read32Bit(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeUnsigned64Bit:
				fallthrough
			case TypeSigned64Bit:
				v := m.Read64Bit(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeFloat32Bit:
				v := m.ReadFloat32(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeFloat64Bit:
				v := m.ReadFloat64(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeString:
				v := m.Read16Bit(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			default:
				v := m.Read16Bit(address, quantity, modbus.INPUT_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)
		}
	}

	return ret
}

func (m *Modbus) ReadHolding(address Address, quantity Quantity, valueType string) string {
	var ret string

	for range Only.Once {
		width := 8

		switch valueType {
			case TypeUnsigned8Bit:
				fallthrough
			case TypeSigned8Bit:
				v := m.ReadByte(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeUnsigned16Bit:
				fallthrough
			case TypeSigned16Bit:
				v := m.Read16Bit(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeUnsigned32Bit:
				fallthrough
			case TypeSigned32Bit:
				v := m.Read32Bit(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeUnsigned64Bit:
				fallthrough
			case TypeSigned64Bit:
				v := m.Read64Bit(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeFloat32Bit:
				v := m.ReadFloat32(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeFloat64Bit:
				v := m.ReadFloat64(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			case TypeString:
				v := m.Read16Bit(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)

			default:
				v := m.Read16Bit(address, quantity, modbus.HOLDING_REGISTER)
				if m.err != nil {
					break
				}
				ret = v.Table(width)
		}
	}

	return ret
}
