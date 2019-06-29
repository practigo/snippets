package snippets

type constErr string

func (e constErr) Error() string {
	return string(e)
}

// Exported errors ...
const (
	DemoConstErr = constErr("demo constant error")
	FungibleErr  = constErr("demo constant error")
)

// GetDemoErr ...
func GetDemoErr() error {
	return DemoConstErr
}
