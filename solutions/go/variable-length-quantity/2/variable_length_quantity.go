package variablelengthquantity

import "errors"

func EncodeVarint(input []uint32) []byte {
	encoded := []byte{}
	
	for _, n := range input {
		bytes := []byte{}
		i := 0
		for n > 0 {
			v := n & 0x7f
			if i > 0 {
				v |= 0x80
			}
			bytes = append(bytes, byte(v))
			n >>= 7
			i++
		}
		if len(bytes) > 0 {
			// reverse order
			for i,j := 0,len(bytes)-1; i < j; i, j = i+1, j-1 {
				bytes[i], bytes[j] = bytes[j], bytes[i]
			}
			encoded = append(encoded, bytes...)
		} else {
			encoded = append(encoded, 0x0)
		}
	}

	return encoded
}

func DecodeVarint(input []byte) ([]uint32, error) {
	// validate sequence	
	if input[len(input) - 1] & 0x80 == 0x80 {
		return []uint32{}, errors.New("invalid sequence")
	}

	nums := []uint32{}

	current := uint32(0)
	for _, n := range input {
		v := uint32(n) & 0x7f
		current <<= 7
		current |= v
		if uint32(n) & 0x80 == 0 {
			nums = append(nums, current)
			current = 0
		}
	} 

	return nums, nil
}
