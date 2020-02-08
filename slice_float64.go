package ameda

import (
	"sort"
)

// Float64Slice float64 slice object
type Float64Slice []float64

// NewFloat64Slice creates an Float64Slice object.
func NewFloat64Slice(a []float64) *Float64Slice {
	f := Float64Slice(a)
	return &f
}

// Float64sCopy creates a copy of the float64 slice.
func Float64sCopy(f []float64) []float64 {
	b := make([]float64, len(f))
	copy(b, f)
	return b
}

// Copy creates a copy of the float64 slice.
func (f Float64Slice) Copy() []float64 {
	return Float64sCopy(f)
}

// Float64sToInterfaces converts float64 slice to interface slice.
func Float64sToInterfaces(f []float64) []interface{} {
	r := make([]interface{}, len(f))
	for k, v := range f {
		r[k] = Float64ToInterface(v)
	}
	return r
}

// Interfaces converts float64 slice to interface slice.
func (f Float64Slice) Interfaces() []interface{} {
	return Float64sToInterfaces(f)
}

// Float64sToStrings converts float64 slice to string slice.
func Float64sToStrings(f []float64) []string {
	r := make([]string, len(f))
	for k, v := range f {
		r[k] = Float64ToString(v)
	}
	return r
}

// Strings converts float64 slice to string slice.
func (f Float64Slice) Strings() []string {
	return Float64sToStrings(f)
}

// Float64sToBools converts float64 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func Float64sToBools(f []float64) []bool {
	r := make([]bool, len(f))
	for k, v := range f {
		r[k] = Float64ToBool(v)
	}
	return r
}

// Bools converts float64 slice to bool slice.
// NOTE:
//  0 is false, everything else is true
func (f Float64Slice) Bools() []bool {
	return Float64sToBools(f)
}

// Float64sToFloat32s converts float64 slice to float32 slice.
func Float64sToFloat32s(f []float64) ([]float32, error) {
	var err error
	r := make([]float32, len(f))
	for k, v := range f {
		r[k], err = Float64ToFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Float32s converts float64 slice to float32 slice.
func (f Float64Slice) Float32s() ([]float32, error) {
	return Float64sToFloat32s(f)
}

// Float64s converts to []float64.
func (f Float64Slice) Float64s() []float64 {
	return []float64(f)
}

// Float64sToInts converts float64 slice to int slice.
func Float64sToInts(f []float64) ([]int, error) {
	var err error
	r := make([]int, len(f))
	for k, v := range f {
		r[k], err = Float64ToInt(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Ints converts float64 slice to int slice.
func (f Float64Slice) Ints() ([]int, error) {
	return Float64sToInts(f)
}

// Float64sToInt8s converts float64 slice to int8 slice.
func Float64sToInt8s(f []float64) ([]int8, error) {
	var err error
	r := make([]int8, len(f))
	for k, v := range f {
		r[k], err = Float64ToInt8(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Int8s converts float64 slice to int8 slice.
func (f Float64Slice) Int8s() ([]int8, error) {
	return Float64sToInt8s(f)
}

// Float64sToInt16s converts float64 slice to int16 slice.
func Float64sToInt16s(f []float64) ([]int16, error) {
	var err error
	r := make([]int16, len(f))
	for k, v := range f {
		r[k], err = Float64ToInt16(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Int16s converts float64 slice to int16 slice.
func (f Float64Slice) Int16s() ([]int16, error) {
	return Float64sToInt16s(f)
}

// Float64sToInt32s converts float64 slice to int32 slice.
func Float64sToInt32s(f []float64) ([]int32, error) {
	var err error
	r := make([]int32, len(f))
	for k, v := range f {
		r[k], err = Float64ToInt32(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Int32s converts float64 slice to int32 slice.
func (f Float64Slice) Int32s() ([]int32, error) {
	return Float64sToInt32s(f)
}

// Float64sToInt64s converts float64 slice to int64 slice.
func Float64sToInt64s(f []float64) ([]int64, error) {
	var err error
	r := make([]int64, len(f))
	for k, v := range f {
		r[k], err = Float64ToInt64(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Int64s converts float64 slice to int64 slice.
func (f Float64Slice) Int64s() ([]int64, error) {
	return Float64sToInt64s(f)
}

// Float64sToUints converts float64 slice to uint slice.
func Float64sToUints(f []float64) ([]uint, error) {
	var err error
	r := make([]uint, len(f))
	for k, v := range f {
		r[k], err = Float64ToUint(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uints converts float64 slice to uint slice.
func (f Float64Slice) Uints() ([]uint, error) {
	return Float64sToUints(f)
}

// Float64sToUint8s converts float64 slice to uint8 slice.
func Float64sToUint8s(f []float64) ([]uint8, error) {
	var err error
	r := make([]uint8, len(f))
	for k, v := range f {
		r[k], err = Float64ToUint8(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint8s converts float64 slice to uint8 slice.
func (f Float64Slice) Uint8s() ([]uint8, error) {
	return Float64sToUint8s(f)
}

// Float64sToUint16s converts float64 slice to uint16 slice.
func Float64sToUint16s(f []float64) ([]uint16, error) {
	var err error
	r := make([]uint16, len(f))
	for k, v := range f {
		r[k], err = Float64ToUint16(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint16s converts float64 slice to uint16 slice.
func (f Float64Slice) Uint16s() ([]uint16, error) {
	return Float64sToUint16s(f)
}

// Float64sToUint32s converts float64 slice to uint32 slice.
func Float64sToUint32s(f []float64) ([]uint32, error) {
	var err error
	r := make([]uint32, len(f))
	for k, v := range f {
		r[k], err = Float64ToUint32(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint32s converts float64 slice to uint32 slice.
func (f Float64Slice) Uint32s() ([]uint32, error) {
	return Float64sToUint32s(f)
}

// Float64sToUint64s converts float64 slice to uint64 slice.
func Float64sToUint64s(f []float64) ([]uint64, error) {
	var err error
	r := make([]uint64, len(f))
	for k, v := range f {
		r[k], err = Float64ToUint64(v)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// Uint64s converts float64 slice to uint64 slice.
func (f Float64Slice) Uint64s() ([]uint64, error) {
	return Float64sToUint64s(f)
}

// Concat is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func (f Float64Slice) Concat(a ...[]float64) []float64 {
	totalLen := len(f)
	for _, v := range a {
		totalLen += len(v)
	}
	ret := make([]float64, totalLen)
	n := copy(ret, f)
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
func (f Float64Slice) CopyWithin(target, start int, end ...int) {
	target = fixIndex(len(f), target, true)
	if target == len(f) {
		return
	}
	sub := f.Slice(start, end...)
	for k, v := range sub {
		f[target+k] = v
	}
}

// Every tests whether all elements in the slice pass the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice will return true for any condition!
func (f Float64Slice) Every(fn func(curr Float64Slice, k int, v float64) bool) bool {
	for k, v := range f {
		if !fn(f, k, v) {
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
func (f Float64Slice) Fill(value float64, start int, end ...int) {
	fixedStart, fixedEnd, ok := fixRange(len(f), start, end...)
	if !ok {
		return
	}
	for k := fixedStart; k < fixedEnd; k++ {
		f[k] = value
	}
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (f Float64Slice) Filter(fn func(curr Float64Slice, k int, v float64) bool) []float64 {
	ret := make([]float64, 0)
	for k, v := range f {
		if fn(f, k, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find returns the key-value of the first element in the provided slice that satisfies the provided testing function.
// NOTE:
//  If not found, k = -1
func (f Float64Slice) Find(fn func(curr Float64Slice, k int, v float64) bool) (k int, v float64) {
	for k, v := range f {
		if fn(f, k, v) {
			return k, v
		}
	}
	return -1, 0
}

// Includes determines whether an slice includes a certain value among its entries.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (f Float64Slice) Includes(valueToFind float64, fromIndex ...int) bool {
	return f.IndexOf(valueToFind, fromIndex...) > -1
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (f Float64Slice) IndexOf(searchElement float64, fromIndex ...int) int {
	idx := getFromIndex(len(f), fromIndex...)
	for k, v := range f[idx:] {
		if searchElement == v {
			return k + idx
		}
	}
	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
// @fromIndex
//  The index to start the search at. Defaults to 0.
func (f Float64Slice) LastIndexOf(searchElement float64, fromIndex ...int) int {
	idx := getFromIndex(len(f), fromIndex...)
	for k := len(f) - 1; k >= idx; k-- {
		if searchElement == f[k] {
			return k
		}
	}
	return -1
}

// Map creates a new slice populated with the results of calling a provided function
// on every element in the calling slice.
func (f Float64Slice) Map(fn func(curr Float64Slice, k int, v float64) float64) []float64 {
	ret := make([]float64, len(f))
	for k, v := range f {
		ret[k] = fn(f, k, v)
	}
	return ret
}

// Pop removes the last element from an slice and returns that element.
// This method changes the length of the slice.
func (f *Float64Slice) Pop() (float64, bool) {
	a := *f
	if len(a) == 0 {
		return 0, false
	}
	lastIndex := len(a) - 1
	last := a[lastIndex]
	a = a[:lastIndex]
	*f = a[:len(a):len(a)]
	return last, true
}

// Push adds one or more elements to the end of an slice and returns the new length of the slice.
func (f *Float64Slice) Push(element ...float64) int {
	*f = append(*f, element...)
	return len(*f)
}

// PushOnce adds one or more new elements that do not exist in the current slice at the end
// and returns the new length of the slice.
func (f *Float64Slice) PushOnce(element ...float64) int {
	a := *f
L:
	for _, v := range element {
		for _, vv := range a {
			if vv == v {
				continue L
			}
		}
		a = append(a, v)
	}
	*f = a
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
func (f Float64Slice) Reduce(
	fn func(curr Float64Slice, k int, v float64, accumulator float64) float64, initialValue ...float64,
) float64 {
	if len(f) == 0 {
		return 0
	}
	start := 0
	acc := f[start]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		start += 1
	}
	for k := start; k < len(f); k++ {
		acc = fn(f, k, f[k], acc)
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
func (f Float64Slice) ReduceRight(
	fn func(curr Float64Slice, k int, v float64, accumulator float64) float64, initialValue ...float64,
) float64 {
	if len(f) == 0 {
		return 0
	}
	end := len(f) - 1
	acc := f[end]
	if len(initialValue) > 0 {
		acc = initialValue[0]
	} else {
		end -= 1
	}
	for k := end; k >= 0; k-- {
		acc = fn(f, k, f[k], acc)
	}
	return acc
}

// Reverse reverses an slice in place.
func (f Float64Slice) Reverse() {
	first := 0
	last := len(f) - 1
	for first < last {
		f[first], f[last] = f[last], f[first]
		first++
		last--
	}
}

// Shift removes the first element from an slice and returns that removed element.
// This method changes the length of the slice.
func (f *Float64Slice) Shift() (float64, bool) {
	a := *f
	if len(a) == 0 {
		return 0, false
	}
	first := a[0]
	a = a[1:]
	*f = a[:len(a):len(a)]
	return first, true
}

// Slice returns a copy of a portion of an slice into a new slice object selected
// from begin to end (end not included) where begin and end represent the index of items in that slice.
// The original slice will not be modified.
func (f Float64Slice) Slice(begin int, end ...int) []float64 {
	fixedStart, fixedEnd, ok := fixRange(len(f), begin, end...)
	if !ok {
		return []float64{}
	}
	return f[fixedStart:fixedEnd].Copy()
}

// Some tests whether at least one element in the slice passes the test implemented by the provided function.
// NOTE:
//  Calling this method on an empty slice returns false for any condition!
func (f Float64Slice) Some(fn func(curr Float64Slice, k int, v float64) bool) bool {
	for k, v := range f {
		if fn(f, k, v) {
			return true
		}
	}
	return false
}

// Len is the number of elements in the collection.
func (f Float64Slice) Len() int {
	return len(f)
}

// Less reports whether the element with
// index m should sort before the element with index n.
func (f Float64Slice) Less(m, n int) bool {
	return f[m] < f[n]
}

// Swap swaps the elements with indexes m and n.
func (f Float64Slice) Swap(m, n int) {
	f[m], f[n] = f[n], f[m]
}

// Sort sorts the elements of an slice in place and returns the sorted slice.
func (f Float64Slice) Sort() {
	sort.Sort(f)
}

// Splice changes the contents of an slice by removing or replacing
// existing elements and/or adding new elements in place.
func (f *Float64Slice) Splice(start, deleteCount int, items ...float64) {
	a := *f
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
			*f = a[:len(a):len(a)]
			return
		}
	}
	if deleteCount > 0 {
		a = append(a[:start], a[start+1+deleteCount:]...)
	}
	*f = a[:len(a):len(a)]
}

// Unshift adds one or more elements to the beginning of an slice and returns the new length of the slice.
func (f *Float64Slice) Unshift(element ...float64) int {
	*f = append(element, *f...)
	return len(*f)
}

// UnshiftOnce adds one or more new elements that do not exist in the current slice to the beginning
// and returns the new length of the slice.
func (f *Float64Slice) UnshiftOnce(element ...float64) int {
	a := *f
	if len(element) == 0 {
		return len(a)
	}
	m := make(map[float64]bool, len(element))
	r := make([]float64, 0, len(a)+len(element))
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
	*f = r[:len(r):len(r)]
	return len(r)
}

// Distinct creates an new slice in place set that removes the same elements
// and returns the new length of the slice.
func (f *Float64Slice) Distinct() int {
	a := (*f)[:0]
	m := make(map[float64]bool, len(a))
	for _, v := range *f {
		if m[v] {
			continue
		}
		a = append(a, v)
		m[v] = true
	}
	n := len(m)
	*f = a[:n:n]
	return n
}

// RemoveOne removes the first matched elements from the slice,
// and returns the new length of the slice.
func (f *Float64Slice) RemoveOne(element ...float64) int {
	a := *f
	m := make(map[float64]bool, len(element))
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
	*f = a[:n:n]
	return n
}

// RemoveAll removes all the elements from the slice,
// and returns the new length of the slice.
func (f *Float64Slice) RemoveAll(element ...float64) int {
	a := *f
	m := make(map[float64]bool, len(element))
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
	*f = a[:n:n]
	return n
}
