package utils

import (
    "fmt"
)

func IterateOverNumbers(elements ...int) {

    fmt.Println(elements)

    for _, i := range elements {
	fmt.Println(i)
    }
}

type Person struct {
    Name string
    Surname string
    Age int
}
