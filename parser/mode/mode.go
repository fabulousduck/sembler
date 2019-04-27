package mode

/*
Mode is a struct holding info about the operation mode
*/
type Mode struct {
	Name, Variable string
}

/*
NewMode returns a new mode pointer
*/
func NewMode() *Mode {
	return new(Mode)
}
