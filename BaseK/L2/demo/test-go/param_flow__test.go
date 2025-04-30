package demo

import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	fmt.Println("flow start main success")
	choiceType(CC)
	fmt.Println()
}

type Gender = int

const (
	Male Gender = iota
	FaMale
	CC
)

func choiceType(param any) {
	fmt.Println("choice func ", param)
	switch param {
	case Male, CC:
		fmt.Println("male or cc person")
	case FaMale:
		fmt.Println("faMale person")
	default:
		fmt.Println("default enum")
	}

}
