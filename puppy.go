package puppy

import (
	"fmt"

	"github.com/kaliaskarn/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From12() {
	fmt.Println("this is for version 1.0.0")
}
