package animals

// Animal represents information about all animals.
type Animal struct {
	Name string
	Age  int
}

// Dog represents information about dogs.
type Dog struct {
	Animal
	BarkStrength int
}
