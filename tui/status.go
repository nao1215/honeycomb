package tui

// status is the status of the honeycomb operation.
type status uint

const (
	// statusNone is the status when the honeycomb operation is not executed.
	statusNone status = iota
	// statusPrivateKeyInput is the status when the private key input view is displayed.
	statusPrivateKeyInput
	// statusPrivateKeyValidateErr is the status when the private key validation error occurs.
	statusPrivateKeyValidateErr
	// statusPrivateKeySaveErr is the status when the private key save error occurs.
	statusPrivateKeySaveErr
)
