// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package coin

import "github.com/amherag/skycoin/src/cipher/encoder"

// encodeSizeBlockHeader computes the size of an encoded object of type BlockHeader
func encodeSizeBlockHeader(obj *BlockHeader) uint64 {
	i0 := uint64(0)

	// obj.Version
	i0 += 4

	// obj.Time
	i0 += 8

	// obj.BkSeq
	i0 += 8

	// obj.Fee
	i0 += 8

	// obj.PrevHash
	i0 += 32

	// obj.BodyHash
	i0 += 32

	// obj.UxHash
	i0 += 32

	return i0
}

// encodeBlockHeader encodes an object of type BlockHeader to a buffer allocated to the exact size
// required to encode the object.
func encodeBlockHeader(obj *BlockHeader) ([]byte, error) {
	n := encodeSizeBlockHeader(obj)
	buf := make([]byte, n)

	if err := encodeBlockHeaderToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeBlockHeaderToBuffer encodes an object of type BlockHeader to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeBlockHeaderToBuffer(buf []byte, obj *BlockHeader) error {
	if uint64(len(buf)) < encodeSizeBlockHeader(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Version
	e.Uint32(obj.Version)

	// obj.Time
	e.Uint64(obj.Time)

	// obj.BkSeq
	e.Uint64(obj.BkSeq)

	// obj.Fee
	e.Uint64(obj.Fee)

	// obj.PrevHash
	e.CopyBytes(obj.PrevHash[:])

	// obj.BodyHash
	e.CopyBytes(obj.BodyHash[:])

	// obj.UxHash
	e.CopyBytes(obj.UxHash[:])

	return nil
}

// decodeBlockHeader decodes an object of type BlockHeader from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeBlockHeader(buf []byte, obj *BlockHeader) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Version
		i, err := d.Uint32()
		if err != nil {
			return 0, err
		}
		obj.Version = i
	}

	{
		// obj.Time
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Time = i
	}

	{
		// obj.BkSeq
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.BkSeq = i
	}

	{
		// obj.Fee
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Fee = i
	}

	{
		// obj.PrevHash
		if len(d.Buffer) < len(obj.PrevHash) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.PrevHash[:], d.Buffer[:len(obj.PrevHash)])
		d.Buffer = d.Buffer[len(obj.PrevHash):]
	}

	{
		// obj.BodyHash
		if len(d.Buffer) < len(obj.BodyHash) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.BodyHash[:], d.Buffer[:len(obj.BodyHash)])
		d.Buffer = d.Buffer[len(obj.BodyHash):]
	}

	{
		// obj.UxHash
		if len(d.Buffer) < len(obj.UxHash) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.UxHash[:], d.Buffer[:len(obj.UxHash)])
		d.Buffer = d.Buffer[len(obj.UxHash):]
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeBlockHeaderExact decodes an object of type BlockHeader from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeBlockHeaderExact(buf []byte, obj *BlockHeader) error {
	if n, err := decodeBlockHeader(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
