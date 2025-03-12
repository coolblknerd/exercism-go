package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amt, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (amt * factor) / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return InvalidCowsError{Cows: cows}.Error("there are no negative cows")
	} else if cows == 0 {
		return InvalidCowsError{Cows: cows}.Error("no cows don't need food")
	}

	return nil
}

type InvalidCowsError struct {
	Cows int
}

func (e InvalidCowsError) Error(msg string) error {
	return fmt.Errorf("%d cows are invalid: %s", e.Cows, msg)
}
