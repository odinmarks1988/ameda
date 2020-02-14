package ameda

import (
	"sort"
)

// Uint8Slice uint8 slice object
type Uint8Slice []uint8

// NewUint8Slice creates an Uint8Slice object.
func NewUint8Slice(a []uint8) *Uint8Slice {
	u := Uint8Slice(a)
	return &u
}

// Uint8sCopy creates a copy of the uint8 slice.
func Uint8sCopy(u []uint8) []uint8 {
	b := make([]uint8, len(u))
	copy(b, u)
	return b
}

// Copy creates a copy of the uint8 slice.
func (u Uint8Slice) Copy() []uint8 {
	return Uint8sCopy(u)
}

// Uint8sToInterfaces converts uint8 slice to interface slice.
func Uint8sToInterfaces(u []uint8) []interface{} {
	r := make([]interface{}, len(u))
	for k, v := range u {
		r[k] = Uint8ToInterface(v)
	}
	return r
}

// Interfaces converts uint8 slice to interface slice.
func (u Uint8Slice) Interfaces() []interface{} {
	return Uint8sToInterfaces(u)
}

// Uint8sToStrings converts uint8 slice to string slice.
func Uint8sToStrings(u []uint8) []string {
	r := make([]string, len(u))
	for k, v := range u {
		r[k] = Uint8ToString(v)
	}
	return r
}

// Strings converts uint8 slice to string slice.
func (u Uint8Slice) Strings() []string {
	return Uint8sToStrings(u)
}

// Uint8sToBools converts uint8 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func Uint8sToBools(u []uint8) []bool {
	r := make([]bool, len(u))
	for k, v := range u {
		r[k] = Uint8ToBool(v)
	}
	return r
}

// Bools converts uint8 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func (u Uint8Slice) Bools() []bool {
	return Uint8sToBools(u)
}

// Uint8sToFloat32s converts uint8 slice to float32 slice.
func Uint8sToFloat32s(u []uint8) []float32 {
	r := make([]float32, len(u))
	for k, v := range u {
		r[k] = Uint8ToFloat32(v)
	}
	return r
}

// Float32s converts uint8 slice to float32 slice.
func (u Uint8Slice) Float32s() []float32 {
	return Uint8sToFloat32s(u)
}

// Uint8sToFloat64s converts uint8 slice to float64 slice.
func Uint8sToFloat64s(u []uint8) []float64 {
	r := make([]float64, len(u))
	for k, v := range u {
		r[k] = Uint8ToFloat64(v)
	}
	return r
}

// Float64s converts uint8 slice to float64 slice.
func (u Uint8Slice) Float64s() []float64 {
	return Uint8sToFloat64s(u)
}

// Uint8sToInts converts uint8 slice to int slice.
func Uint8sToInts(u []uint8) []int {
	r := make([]int, len(u))
	for k, v := range u {
		r[k] = Uint8ToInt(v)
	}
	return r
}

// Ints converts uint8 slice to int slice.
func (u Uint8Slice) Ints() []int {
	return Uint8sToInts(u)
}

// Uint8sToInt8s converts uint8 slice to int8 slice.
func Uint8sToInt8s(u []uint8) ([]int8, error) {
	var err error
	r := make([]int8, len(u))
	for k, v := range u {
		r[k], err = Uint8ToInt8(v)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// Int8s converts uint8 slice to int8 slice.
func (u Uint8Slice) Int8s() ([]int8, error) {
	return Uint8sToInt8s(u)
}

// Uint8sToInt16s converts uint8 slice to int16 slice.
func Uint8sToInt16s(u []uint8) []int16 {
	r := make([]int16, len(u))
	for k, v := range u {
		r[k] = Uint8ToInt16(v)
	}
	return r
}

// Int16s converts uint8 slice to int16 slice.
func (u Uint8Slice) Int16s() []int16 {
	return Uint8sToInt16s(u)
}

// Uint8sToInt32s converts uint8 slice to int32 slice.
func Uint8sToInt32s(u []uint8) []int32 {
	r := make([]int32, len(u))
	for k, v := range u {
		r[k] = Uint8ToInt32(v)
	}
	return r
}

// Int32s converts uint8 slice to int32 slice.
func (u Uint8Slice) Int32s() []int32 {
	return Uint8sToInt32s(u)
}

// Uint8sToInt64s converts uint8 slice to int64 slice.
func Uint8sToInt64s(u []uint8) []int64 {
	r := make([]int64, len(u))
	for k, v := range u {
		r[k] = Uint8ToInt64(v)
	}
	return r
}

// Int64s converts uint8 slice to int64 slice.
func (u Uint8Slice) Int64s() []int64 {
	return Uint8sToInt64s(u)
}

// Uint8sToUints converts uint8 slice to uint slice.
func Uint8sToUints(u []uint8) []uint {
	r := make([]uint, len(u))
	for k, v := range u {
		r[k] = Uint8ToUint(v)
	}
	return r
}

// Uints converts uint8 slice to uint slice.
func (u Uint8Slice) Uints() []uint {
	return Uint8sToUints(u)
}

// Uint8s converts to []uint8.
func (u Uint8Slice) Uint8s() []uint8 {
	return []uint8(u)
}

// Uint8sToUint16s converts uint8 slice to uint16 slice.
func Uint8sToUint16s(u []uint8) []uint16 {
	r := make([]uint16, len(u))
	for k, v := range u {
		r[k] = Uint8ToUint16(v)
	}
	return r
}

// Uint16s converts uint8 slice to uint16 slice.
func (u Uint8Slice) Uint16s() []uint16 {
	return Uint8sToUint16s(u)
}

// Uint8sToUint32s converts uint8 slice to uint32 slice.
func Uint8sToUint32s(u []uint8) []uint32 {
	r := make([]uint32, len(u))
	for k, v := range u {
		r[k] = Uint8ToUint32(v)
	}
	return r
}

// Uint32s converts uint8 slice to uint32 slice.
func (u Uint8Slice) Uint32s() []uint32 {
	return Uint8sToUint32s(u)
}

// Uint8sToUint64s converts uint8 slice to uint64 slice.
func Uint8sToUint64s(u []uint8) []uint64 {
	r := make([]uint64, len(u))
	for k, v := range u {
		r[k] = Uint8ToUint64(v)
	}
	return r
}

// Uint64s converts uint8 slice to uint64 slice.
func (u Uint8Slice) Uint64s() []uint64 {
	return Uint8sToUint64s(u)
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (u Uint8Slice) Concat(a ...[]uint8) []uint8 {
	totalLen := len(u)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]uint8, totalLen)
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
func (u Uint8Slice) CopyWithin(target, start int, end ...int) {
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
func (u Uint8Slice) Every(fn func(curr Uint8Slice, k int, v uint8) bool) bool {
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
func (u Uint8Slice) Fill(value uint8, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(u), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		u[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (u Uint8Slice) Filter(fn func(curr Uint8Slice, k int, v uint8) bool) []uint8 {
	ret := make([]uint8, 0)
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
func (u Uint8Slice) Find(fn func(curr Uint8Slice, k int, v uint8) bool) (k int, v uint8) {
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
func (u Uint8Slice) Includes(valueToFind uint8, fromIndex ...int) bool {
	return u.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (u Uint8Slice) IndexOf(searchElement uint8, fromIndex ...int) int {
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
func (u Uint8Slice) LastIndexOf(searchElement uint8, fromIndex ...int) int {
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
func (u Uint8Slice) Map(fn func(curr Uint8Slice, k int, v uint8) uint8) []uint8 {
	ret := make([]uint8, len(u))
	for k, v := range u {
		ret[k] = fn(u, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (u *Uint8Slice) Pop() (uint8, bool) {
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
func (u *Uint8Slice) Push(element ...uint8) int {
	*u = append(*u, element...)
	return len(*u)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (u *Uint8Slice) PushOnce(element ...uint8) int {
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
func (u Uint8Slice) Reduce(
	fn func(curr Uint8Slice, k int, v, accumulator uint8) uint8, initialValue ...uint8,
) uint8 {
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
func (u Uint8Slice) ReduceRight(
	fn func(curr Uint8Slice, k int, v, accumulator uint8) uint8, initialValue ...uint8,
) uint8 {
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
func (u Uint8Slice) Reverse() {
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
func (u *Uint8Slice) Shift() (uint8, bool) {
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
func (u Uint8Slice) Slice(begin int, end ...int) []uint8 {
	fixedStart, fixedEnd, ok := fixRange(len(u), begin, end...)
	if !ok {
		return []uint8{}
	}
	return u[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (u Uint8Slice) Some(fn func(curr Uint8Slice, k int, v uint8) bool) bool {
	for k, v := range u {
		if fn(u, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (u Uint8Slice) Len() int {
	return len(u)
}

// Less reports whether the element with
// index m should sort before the element with index n.
func (u Uint8Slice) Less(m, n int) bool {
	return u[m] < u[n]
}

// Swap swaps the elements with indexes m and n.
func (u Uint8Slice) Swap(m, n int) {
	u[m], u[n] = u[n], u[m]
}

// Sort sorts the elements of an slice in place and returns the sorted slice.
func (u Uint8Slice) Sort() {
	sort.Sort(u)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (u *Uint8Slice) Splice(start, deleteCount int, items ...uint8) {
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
func (u *Uint8Slice) Unshift(element ...uint8) int {
	*u = append(element, *u...)
	return len(*u)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (u *Uint8Slice) UnshiftOnce(element ...uint8) int {
	a := *u
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[uint8]bool, len(element))
	r := make([]uint8, 0, len(a)+len(element))
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
func (u *Uint8Slice) Distinct() int {
	a := (*u)[:0]
	m := make(map[uint8]bool, len(a))
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
func (u *Uint8Slice) RemoveOne(element ...uint8) int {
	a := *u
	m := make(map[uint8]bool, len(element))
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

// RemoveEvery removes all the elements from the slice,
// and returns the new length of the slice.
func (u *Uint8Slice) RemoveEvery(element ...uint8) int {
	a := *u
	m := make(map[uint8]bool, len(element))
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
