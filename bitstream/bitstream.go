package bitstream

import "errors"

var (
	ErrOutOfRange      = errors.New("index out of range")
	ErrInvalidBitValue = errors.New("invalid bit value")
)

// Bitstream represents a sequence of bits storedd in a byte slice.
type Bitstream struct {
	data []byte // byte slice to store bits
	len  int    // number of bits in the bitstream (length of the byte slice * 8)
}

// NewBitstream creates a new Bitstream from the provided byte slice.
func NewBitstream(data []byte) *Bitstream {
	return &Bitstream{data: data, len: len(data) * 8}
}

// Len returns the length of the bitsream in bits.
func (bs *Bitstream) Len() int {
	return bs.len
}

// Bit returns the bit value at the sepcified index.
// It returns an error if the index is out of range.
func (bs *Bitstream) Bit(index int) (byte, error) {
	if index < 0 || index >= bs.len {
		return 0, ErrOutOfRange
	}

	byteIndex, bitIndex := getIndexes(index)
	return (bs.data[byteIndex] >> uint(7-bitIndex)) & 1, nil
}

// SetBit sets the bit value at the specified index.
// It returns an error if the index is out of range.
func (bs *Bitstream) SetBit(index int, bit byte) error {
	if index < 0 || index >= bs.len {
		return ErrOutOfRange
	}

	byteIndex, bitIndex := getIndexes(index)
	mask := byte(1 << uint(7-bitIndex))
	if bit == 1 {
		bs.data[byteIndex] |= mask
	} else {
		bs.data[byteIndex] &= ^mask
	}
	return nil
}

// Append appends a bit to the end of the bitstream.
// It returns an error if the provided bit value is not 0 or 1.
func (bs *Bitstream) Append(bit byte) error {
	if bit != 0 && bit != 1 {
		return ErrInvalidBitValue
	}

	byteIndex, bitIndex := getIndexes(bs.len)
	if bitIndex == 0 {
		bs.data = append(bs.data, 0)
	}

	if bit == 1 {
		bs.data[byteIndex] |= 1 << uint(7-bitIndex)
	}

	bs.len++
	return nil
}

// Bytes returns the underlying byte slice of the bitstream.
func (bs *Bitstream) Bytes() []byte {
	return bs.data
}

// getIndexes is a helper function that calculates the byte and bit index within the byte
// for the given bit position in the bitsream.
func getIndexes(index int) (byteInex, bitIndex int) {
	return index / 8, index % 8
}