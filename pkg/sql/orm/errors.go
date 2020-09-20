package orm

import "errors"

var (
	ErrModelObjIsNotPtr                = errors.New("model object is not pointer")
	ErrInvalidLengthBetweenColsAndVals = errors.New("model object's columns length is inconsistent with values length")
	ErrNonZeroSliceLength              = errors.New("slice length should be zero")
)
