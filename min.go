// Package slice provides typesafe functions for common Go slice operations.
package slice

import "errors"

// MinByte returns the minimum value of a byte slice or an error in case of a nil or empty slice
func MinByte(a []byte) (byte, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinFloat32 returns the minimum value of a float32 slice or an error in case of a nil or empty slice
func MinFloat32(a []float32) (float32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinFloat64 returns the minimum value of a float64 slice or an error in case of a nil or empty slice
func MinFloat64(a []float64) (float64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt returns the minimum value of an int slice or an error in case of a nil or empty slice
func MinInt(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt8 returns the minimum value of an int8 slice or an error in case of a nil or empty slice
func MinInt8(a []int8) (int8, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt16 returns the minimum value of an int16 slice or an error in case of a nil or empty slice
func MinInt16(a []int16) (int16, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt32 returns the minimum value of an int32 slice or an error in case of a nil or empty slice
func MinInt32(a []int32) (int32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinInt64 returns the minimum value of an int64 slice or an error in case of a nil or empty slice
func MinInt64(a []int64) (int64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinRune returns the minimum value of a rune slice or an error in case of a nil or empty slice
func MinRune(a []rune) (rune, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint returns the minimum value of a uint slice or an error in case of a nil or empty slice
func MinUint(a []uint) (uint, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint8 returns the minimum value of a uint8 slice or an error in case of a nil or empty slice
func MinUint8(a []uint8) (uint8, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint16 returns the minimum value of a uint16 slice or an error in case of a nil or empty slice
func MinUint16(a []uint16) (uint16, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint32 returns the minimum value of a uint32 slice or an error in case of a nil or empty slice
func MinUint32(a []uint32) (uint32, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUint64 returns the minimum value of a uint64 slice or an error in case of a nil or empty slice
func MinUint64(a []uint64) (uint64, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}

// MinUintptr returns the minimum value of a uintptr slice or an error in case of a nil or empty slice
func MinUintptr(a []uintptr) (uintptr, error) {
	if len(a) == 0 {
		return 0, errors.New("Cannot get the minimum of a nil or empty slice")
	}

	min := a[0]
	for k := 1; k < len(a); k++ {
		if a[k] < min {
			min = a[k]
		}
	}

	return min, nil
}
