package switching

import (
	"errors"
	"fmt"
)

// HeavyCase is a single case for `MultipleHeavyCase`.
// You are only allowed to initiate this type by using `NewHeavyCase`.
type HeavyCase struct {
	evaluator heavyEvaluatorFunc
	command   heavyCommandFunc
}

// NewHeavyCase it generate a single case for `MultipleHeavyCase`.
func NewHeavyCase(evaluator heavyEvaluatorFunc, command heavyCommandFunc) HeavyCase {
	return HeavyCase{
		evaluator: evaluator,
		command:   command,
	}
}

type heavyEvaluatorFunc func() (bool, error)
type heavyCommandFunc func() error

// MultipleHeavyCase it let you evaluate multiple case given an evaluator and run each command accordingly.
// Example:
//  value := 2
//  err := switching.MultipleHeavyCase(
//			switching.NewHeavyCase(
//				func() (bool, error) {
//					return value < 4, nil
//				},
//				func() error {
//					fmt.Println("case 1")
// 					return nil
//				},
//			),
//			switching.NewHeavyCase(
//				func() (bool, error) {
//					return value == 2, nil
//				},
//				func() error {
//					fmt.Println("case 2")
// 					return nil
//				},
//			),
//		)
// Given some cases those need to be evaluated, it will evaluate be evaluated in for each loop.
// For each case, the `evaluator()` would be call.
// f the result is `true`, it will call `command()` right away.
func MultipleHeavyCase(cases ...HeavyCase) (err error) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			err = errors.New(fmt.Sprint(recovered))
		}
	}()

	for _, kasus := range cases {
		ok, err := kasus.evaluator()
		if err != nil {
			return err
		} else if ok {
			err = kasus.command()
			if err != nil {
				return err
			}
		}
	}

	return
}
