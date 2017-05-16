package testcases

// Voider does something and returns nothing.
type Voider interface {
	Void() 
}

// VoiderFunc is an adapter to allow ordinary functions to be used as Voider implementations.
type VoiderFunc func() 

// Void calls f.
func (f VoiderFunc) Void()  {
	return f()
}

// Compile time assertion that VoiderFunc implements Voider.
var _ Voider = new(VoiderFunc)
