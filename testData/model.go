package testdata

// Food is just an helper to test different types.
type Food int

const (
	// Ice is ... cold.
	Ice Food = iota
	// Burger consists of a bun, a patty and some toppings.
	Burger Food = iota
)

// Person represents a person and is used to test the code generation.
// go:generate partialstructupdater -type=Person
type Person struct {
	Firstname    string
	Lastname     string
	FavoriteFood Food
}
