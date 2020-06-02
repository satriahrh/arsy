package switching

// Case is a single case for `MultipleCase`.
// You are only allowed to initiate this type by using `NewCase`.
type Case struct {
	evaluator evaluatorFunc
	command   commandFunc
}

// NewCase it generate a single case for `MultipleCase`.
func NewCase(evaluator evaluatorFunc, command commandFunc) Case {
	return Case{
		evaluator: evaluator,
		command:   command,
	}
}

type evaluatorFunc func() bool
type commandFunc func()

// MultipleCase it let you evaluate multiple case given an evaluator and run each command accordingly.
//
// Example:
//  value := 2
//  switching.MultipleCase(
//			switching.NewCase(
//				func() bool {
//					return value < 4
//				},
//				func() {
//					fmt.Println("case 1")
//				},
//			),
//			switching.NewCase(
//				func() bool {
//					return value == 2
//				},
//				func() {
//					fmt.Println("case 2")
//				},
//			),
//		)
//
// Given some cases those need to be evaluated, it will evaluate be evaluated in for each loop.
// For each case, the `evaluator()` would be call.
// If the result is `true`, it will call `command()` right away.
func MultipleCase(cases ...Case) {
	for _, kasus := range cases {
		if kasus.evaluator() {
			kasus.command()
		}
	}

	return
}
