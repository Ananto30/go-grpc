package model

// ValidationError type error is returned from model validation
type ValidationError struct {
	Field, Msg string
}

func (v *ValidationError) Error() string {
	return "model: the " + v.Field + " field " + v.Msg
}
