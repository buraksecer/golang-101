package main

import (
	"errors"
	"fmt"
	"log"
)

type errKind int

const (
	_ errKind = iota // skip 0
	invalidTime
	invalidName
)

type Errors struct {
	err  error
	kind errKind
}

func (e Errors) Error() string {
	switch e.kind {
	case invalidName:
		return "please check name"
	case invalidTime:
		return "please check time"
	}
	return "default error"
}

var (
	InvalidTimeError = Errors{kind: invalidTime}

	InvalidNameError = Errors{kind: invalidName}

	DefaultError = Errors{kind: 3131}
)

func do(p int) error {
	if p == 0 {
		return InvalidNameError
	} else if p == 1 {
		return InvalidTimeError
	} else if p == 2 {
		return DefaultError
	}
	return nil
}

func doSomething() error {
	return errors.New("this is an error")
}

func main() {

	if err := do(1); err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}

	if err := do(2); err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}

	if err := do(0); err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}

	err := doSomething()
	if err != nil {
		log.Fatal(fmt.Errorf("we have an error: %w", err)) // %w
	}
}
