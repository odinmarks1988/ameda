package ameda

import (
	"fmt"
	"math"
	"sort"
)

// Float32Slice float32 slice object
type Float32Slice []float32

// NewFloat32Slice creates an Float32Slice object.
func NewFloat32Slice(a []float32) *Float32Slice {
	i := Float32Slice(a)
	return &i
}

// Copy creates a copy of the float32 slice.
func (u Float32Slice) Copy() []float32 {
	b := make([]float32, len(u))
	copy(b, u)
	return b
}

// Interfaces converts float32 slice to interface{} slice.
func (u Float32Slice) Interfaces() []interface{} {
	r := make([]interface{}, len(u))
	for k, v := range u {
		r[k] = v
	}
	return r
}

// Strings converts float32 slice to string slice.
func (u Float32Slice) Strings() []string {
	r := make([]string, len(u))
	for k, v := range u {
		r[k] = fmt.Sprintf("%f", v)
	}
	return r
}

// Bools converts float32 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func (u Float32Slice) Bools() []bool {
	r := make([]bool, len(u))
	for k, v := range u {
		r[k] = v != 0
	}
	return r
}

// Float32s converts to []float32.
func (u Float32Slice) Float32s() []float32 {
	return []float32(u)
}

// Float32s converts float32 slice to float64 slice.
func (u Float32Slice) Float64s() []float64 {
	r := make([]float64, len(u))
	for k, v := range u {
		r[k] = float64(v)
	}
	return r
}

// Ints converts float32 slice to int slice.
func (u Float32Slice) Ints() ([]int, error) {
	r := make([]int, len(u))
	if is64BitPlatform {
		for k, v := range u {
			if v > math.MaxInt64 {
				return nil, errOverflowValue
			}
			r[k] = int(v)
		}
	} else {
		for k, v := range u {
			if v > math.MaxInt32 {
				return nil, errOverflowValue
			}
			r[k] = int(v)
		}
	}
	return r, nil
}

// Int8s converts float32 slice to int8 slice.
func (u Float32Slice) Int8s() ([]int8, error) {
	r := make([]int8, len(u))
	for k, v := range u {
		if v > math.MaxInt8 {
			return nil, errOverflowValue
		}
		r[k] = int8(v)
	}
	return r, nil
}

// Int16s converts float32 slice to int16 slice.
func (u Float32Slice) Int16s() ([]int16, error) {
	r := make([]int16, len(u))
	for k, v := range u {
		if v > math.MaxInt16 {
			return nil, errOverflowValue
		}
		r[k] = int16(v)
	}
	return r, nil
}

// Int32s converts float32 slice to int32 slice.
func (u Float32Slice) Int32s() ([]int32, error) {
	r := make([]int32, len(u))
	for k, v := range u {
		if v > math.MaxInt32 {
			return nil, errOverflowValue
		}
		r[k] = int32(v)
	}
	return r, nil
}

// Int64s converts float32 slice to int64 slice.
func (u Float32Slice) Int64s() ([]int64, error) {
	r := make([]int64, len(u))
	for k, v := range u {
		if v > math.MaxInt64 {
			return nil, errOverflowValue
		}
		r[k] = int64(v)
	}
	return r, nil
}

// Uints converts float32 slice to uint slice.
func (u Float32Slice) Uints() ([]uint, error) {
	r := make([]uint, len(u))
	if is64BitPlatform {
		for k, v := range u {
			if v > math.MaxUint64 {
				return nil, errOverflowValue
			}
			r[k] = uint(v)
		}
	} else {
		for k, v := range u {
			if v > math.MaxUint32 {
				return nil, errOverflowValue
			}
			r[k] = uint(v)
		}
	}
	return r, nil
}

// Uint8s converts float32 slice to uint8 slice.
func (u Float32Slice) Uint8s() ([]uint8, error) {
	r := make([]uint8, len(u))
	for k, v := range u {
		if v > math.MaxUint8 {
			return nil, errOverflowValue
		}
		r[k] = uint8(v)
	}
	return r, nil
}

// Float32s converts float32 slice to uint16 slice.
func (u Float32Slice) Uint16s() ([]uint16, error) {
	r := make([]uint16, len(u))
	for k, v := range u {
		if v > math.MaxUint16 {
			return nil, errOverflowValue
		}
		r[k] = uint16(v)
	}
	return r, nil
}

// Float32s converts float32 slice to uint32 slice.
func (u Float32Slice) Uint32s() ([]uint32, error) {
	r := make([]uint32, len(u))
	for k, v := range u {
		if v > math.MaxUint32 {
			return nil, errOverflowValue
		}
		r[k] = uint32(v)
	}
	return r, nil
}

// Float32s converts float32 slice to uint64 slice.
func (u Float32Slice) Uint64s() ([]uint64, error) {
	r := make([]uint64, len(u))
	for k, v := range u {
		if v > math.MaxUint64 {
			return nil, errOverflowValue
		}
		r[k] = uint64(v)
	}
	return r, nil
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (u Float32Slice) Concat(a ...[]float32) []float32 {
	totalLen := len(u)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]float32, totalLen)
	n := copy(ret, u)
	dst := ret[n:]
	for _, v := range a {
		n := copy(dst, v)
		dst = dst[n:]
	}
	return ret
}

// CopyWithin copies part of an slice to another location in the current slice.
// @target
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func (u Float32Slice) CopyWithin(target, start int, end ...int) {
	target = fixIndex(len(u), target, true)
	if target == len(u) {
		return
	}
	sub := u.Slice(start, end...)
	for k, v := range sub {
		u[target+k] = v
	}
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func (u Float32Slice) Every(fn func(curr Float32Slice, k int, v float32) bool) bool {
	for k, v := range u {
		if !fn(u, k, v) {
			return false
		}
	}
	return true
}

// Fill changes all elements in the current slice to a value, from a start index to an end index.
// @value
//  Zero-based index at which to copy the sequence to. If negative, target will be counted from the end.
// @start
//  Zero-based index at which to start copying elements from. If negative, start will be counted from the end.
// @end
//  Zero-based index at which to end copying elements from. CopyWithin copies up to but not including end.
//  If negative, end will be counted from the end.
//  If end is omitted, CopyWithin will copy until the last index (default to len(s)).
func (u Float32Slice) Fill(value float32, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(u), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		u[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (u Float32Slice) Filter(fn func(curr Float32Slice, k int, v float32) bool) []float32 {
	ret := make([]float32, 0)
	for k, v := range u {
		if fn(u, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func (u Float32Slice) Find(fn func(curr Float32Slice, k int, v float32) bool) (k int, v float32) {
	for k, v := range u {
		if fn(u, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// Includes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (u Float32Slice) Includes(valueToFind float32, fromIndex ...int) bool {
	return u.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (u Float32Slice) IndexOf(searchElement float32, fromIndex ...int) int {
	idx := getFromIndex(len(u), fromIndex...)
	for k, v := range u[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (u Float32Slice) LastIndexOf(searchElement float32, fromIndex ...int) int {
	idx := getFromIndex(len(u), fromIndex...)
	for k := len(u) - 1; k >= idx; k-- {
		if searchElement == u[k] {
			return k
		}
	}
	return -1
}

// Map creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func (u Float32Slice) Map(fn func(curr Float32Slice, k int, v float32) float32) []float32 {
	ret := make([]float32, len(u))
	for k, v := range u {
		ret[k] = fn(u, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (u *Float32Slice) Pop() (float32, bool) {
	a := *u
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*u = a[:len(a):len(a)]
	return last, true
}

// Push adds one or more elements to the end of an slice and returns the new length of the slice.
func (u *Float32Slice) Push(element ...float32) int {
	*u = append(*u, element...)
	return len(*u)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (u *Float32Slice) PushOnce(element ...float32) int {
	a := *u
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*u = a
	return len(a)
}

// Reduce executes a reducer function (that you provide) on each element of the slice,
// resulting in a single output value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func (u Float32Slice) Reduce(
	fn func(curr Float32Slice, k int, v float32, accumulator float32) float32, initialValue ...float32,
) float32 {
	if len(u) == 0 {
		return 0
	}
	start := 0
	acc := u[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(u); k++ {
		acc = fn(u, k, u[k], acc)
	}
	return acc
}

// ReduceRight applies a function against an accumulator and each value of the slice (from right-to-left)
// to reduce it to a single value.
// @accumulator
//  The accumulator accumulates callback's return values.
//  It is the accumulated value previously returned in the last invocation of the callback—or initialValue,
//  if it was supplied (see below).
// @initialValue
//  A value to use as the first argument to the first call of the callback.
//  If no initialValue is supplied, the first element in the slice will be used and skipped.
func (u Float32Slice) ReduceRight(
	fn func(curr Float32Slice, k int, v float32, accumulator float32) float32, initialValue ...float32,
) float32 {
	if len(u) == 0 {
		return 0
	}
	end := len(u) - 1
	acc := u[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(u, k, u[k], acc)
	}
	return acc
}

// Reverse reverses an slice in place.
func (u Float32Slice) Reverse() {
	first := 0
	last := len(u) - 1
	for first < last {
		u[first], u[last] = u[last], u[first]
		first++
		last--
	}
}

// Shift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func (u *Float32Slice) Shift() (float32, bool) {
	a := *u
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*u = a[:len(a):len(a)]
	return first, true
}

// Slice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func (u Float32Slice) Slice(begin int, end ...int) []float32 {
	fixedStart, fixedEnd, ok := fixRange(len(u), begin, end...)
	if !ok {
		return []float32{}
	}
	return u[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (u Float32Slice) Some(fn func(curr Float32Slice, k int, v float32) bool) bool {
	for k, v := range u {
		if fn(u, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (u Float32Slice) Len() int {
	return len(u)
}

// Less reports whether the element with
// index m should sort before the element with index n.
func (u Float32Slice) Less(m, n int) bool {
	return u[m] < u[n]
}

// Swap swaps the elements with indexes m and n.
func (u Float32Slice) Swap(m, n int) {
	u[m], u[n] = u[n], u[m]
}

// Sort sorts the elements of an slice in place and returns the sorted slice.
func (u Float32Slice) Sort() {
	sort.Sort(u)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (u *Float32Slice) Splice(start, deleteCount int, items ...float32) {
	a := *u
	if deleteCount < 0 {
		deleteCount = 0
	}
	start, end, _ := fixRange(len(a), start, start+1+deleteCount)
	deleteCount = end - start - 1
	for k := 0; k < len(items); k++ {
		if deleteCount > 0 {
			// replace
			a[start] = items[k]
			deleteCount--
			start++
		} else {
			// insert
			lastSlice := a[start:].Copy()
			items = items[k:]
			a = append(a[:start], items...)
			a = append(a[:start+len(items)], lastSlice...)
			*u = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*u = a[:len(a):len(a)]
}

// Unshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func (u *Float32Slice) Unshift(element ...float32) int {
	*u = append(element, *u...)
	return len(*u)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (u *Float32Slice) UnshiftOnce(element ...float32) int {
	a := *u
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[float32]bool, len(element))
	r := make([]float32, 0, len(a)+len(element))
L:
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		r = append(r, v)
	}
	r = append(r, a...)
	*u = r[:len(r):len(r)]
	return len(r)
}

// Distinct creates an new slice in place set that removes the same elements
// and returns the new length of the slice.
func (u *Float32Slice) Distinct() int {
	a := (*u)[:0]
	m := make(map[float32]bool, len(a))
	for _, v := range *u {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*u = a[:n:n]
	return n
}

// RemoveOne removes the first matched elements from the slice,
// and returns the new length of the slice.
func (u *Float32Slice) RemoveOne(element ...float32) int {
	a := *u
	m := make(map[float32]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
				break
			}
		}
	}
	n := len(a)
	*u = a[:n:n]
	return n
}

// RemoveAll removes all the elements from the slice,
// and returns the new length of the slice.
func (u *Float32Slice) RemoveAll(element ...float32) int {
	a := *u
	m := make(map[float32]bool, len(element))
	for _, v := range element {
		if m[v] {
			continue
		}
		m[v] = true
		for kk, vv := range a {
			if vv == v {
				a = append(a[:kk], a[kk+1:]...)
			}
		}
	}
	n := len(a)
	*u = a[:n:n]
	return n
}
