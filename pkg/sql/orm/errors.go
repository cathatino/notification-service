package orm

import "errors"

var (
	ErrInvalidLengthBetweenColsAndVals = errors.New("model object's columns length is inconsistent with values length")
	ErrModelObjIsNotPtr                = errors.New("model object is not pointer")
	ErrModelObjIsPtr                   = errors.New("model object is pointer")
	ErrNonZeroSliceLength              = errors.New("slice length should be zero")
	ErrPtrIsNotSlice                   = errors.New("Ptr is not slice")
)
