package errors

import (
	"errors"
	"fmt"
)

var (
	InvalidAddressError     = errors.New("invalid address access")
	InvalidMnemonicError    = errors.New("invalid mnemonic")
	InvalidOperandError     = errors.New("invalid operand")
	InvalidInstructionError = errors.New("invalid instruction")
	NotImplementedError     = errors.New("not implemented")
)

func NewInvalidOperandError(operand any) error {
	return fmt.Errorf("%w: %s (%T)", InvalidOperandError, operand, operand)
}

func NewOperandSymbolError(got, want any) error {
	return fmt.Errorf("%w: invalid symbol: got: %T, want: %T", InvalidOperandError, got, want)
}

func NewAccessError(addr uint16, resource string) error {
	return fmt.Errorf("%w: unable to access %s at 0x%04x", InvalidAddressError, resource, addr)
}

func NewReadError(addr uint16, resource string) error {
	return fmt.Errorf("%w: unable to read %s at 0x%04x", InvalidAddressError, resource, addr)
}

func NewWriteError(addr uint16, resource string) error {
	return fmt.Errorf("%w: unable to write %s at 0x%04x", InvalidAddressError, resource, addr)
}

func NewInvalidMnemonicError(mnemonic string) error {
	return fmt.Errorf("%w: %q", InvalidMnemonicError, mnemonic)
}

func NewUnknownOPCodeError(opcode byte) error {
	return fmt.Errorf("%w: unknown opcode 0x%02x", InvalidInstructionError, opcode)
}
