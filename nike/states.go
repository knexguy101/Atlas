package nike

/*
	States

	Used to identify which step the task is on. Names should be relevant to what's happening, hopefully
 */

const (
	FINISHED = 0
	INIT = 1
	LOAD = 2
	LOGIN = 3
	ATC = 4
	CLEAR = 5
	ADDRESS = 6
	PAYMENT = 7
	SUBMIT = 8
	PENDING = 9
)
