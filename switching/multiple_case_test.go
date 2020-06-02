package switching_test

import (
	"errors"
	"testing"

	"github.com/satriahrh/arsy/switching"
	"github.com/stretchr/testify/assert"
)

func TestNewCase(t *testing.T) {
	c := switching.NewCase(
		func() bool {
			return true
		},
		func() {
		},
	)
	assert.IsType(t, switching.Case{}, c)
}

func TestMultipleCase_ErrorHandling(t *testing.T) {
	expectedError := errors.New("some error")
	t.Run("Panic", func(t *testing.T) {
		called := make([]bool, 3)
		assert.PanicsWithError(t, expectedError.Error(), func() {
			switching.MultipleCase(
				switching.NewCase(
					func() bool {
						return true
					},
					func() {
						called[0] = true
					},
				),
				switching.NewCase(
					func() bool {
						return true
					},
					func() {
						panic(expectedError)
					},
				),
				switching.NewCase(
					func() bool {
						return true
					},
					func() {
						called[2] = true
					},
				),
			)
		})
		assert.Equal(t, []bool{true, false, false}, called)
	})
}

func TestMultipleCase_OneCaseMatched(t *testing.T) {
	called := make([]bool, 2)
	assert.NotPanics(t, func() {
		switching.MultipleCase(
			switching.NewCase(
				func() bool {
					return true
				},
				func() {
					called[0] = true
				},
			),
			switching.NewCase(
				func() bool {
					return false
				},
				func() {
					called[1] = true
				},
			),
		)
	})
	assert.Equal(t, []bool{true, false}, called)
}

func TestMultipleCase_AllCaseMatched(t *testing.T) {
	called := make([]bool, 2)
	assert.NotPanics(t, func() {
		switching.MultipleCase(
			switching.NewCase(
				func() bool {
					return true
				},
				func() {
					called[0] = true
				},
			),
			switching.NewCase(
				func() bool {
					return true
				},
				func() {
					called[1] = true
				},
			),
		)
	})
	assert.Equal(t, []bool{true, true}, called)
}

func TestMultipleCase_NoCaseMatched(t *testing.T) {
	called := make([]bool, 2)
	var err error
	assert.NotPanics(t, func() {
		switching.MultipleCase(
			switching.NewCase(
				func() bool {
					return false
				},
				func() {
					called[0] = true
				},
			),
			switching.NewCase(
				func() bool {
					return false
				},
				func() {
					called[1] = true
				},
			),
		)

	})
	assert.NoError(t, err)
	assert.Equal(t, []bool{false, false}, called)
}

func TestMultipleCase_NoCaseGiven(t *testing.T) {
	assert.NotPanics(t, func() {
		switching.MultipleCase()
	})
}
