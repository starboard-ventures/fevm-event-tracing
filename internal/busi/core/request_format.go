package core

import "errors"

type RequestHeight struct {
	MinHeight uint64 `form:"from" json:"from"`
	MaxHeight uint64 `form:"to" json:"to"`

	Lotus0 string `form:"-" json:"-"`
}

func (r *RequestHeight) Validate() error {
	if !(r.MinHeight < r.MaxHeight) {
		return errors.New("'from' should less or equal than 'to'")
	}

	return nil
}
