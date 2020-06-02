package switching_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/satriahrh/arsy/switching"
)

func TestNewHeavyCase(t *testing.T) {
	c := switching.NewHeavyCase(
		func() (bool, error) {
			return true, nil
		},
		func() error {
			return nil
		},
	)
	assert.IsType(t, switching.HeavyCase{}, c)
}

func TestMultipleHeavyCase_ErrorHandling(t *testing.T) {
	expectedError := errors.New("some error")
	t.Run("Panic", func(t *testing.T) {
		called := make([]bool, 3)
		err := switching.MultipleHeavyCase(
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[0] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					panic(expectedError)
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[2] = true
					return nil
				},
			),
		)
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, []bool{true, false, false}, called)
	})
	t.Run("ReturnErrorOnEvaluator", func(t *testing.T) {
		called := make([]bool, 3)
		err := switching.MultipleHeavyCase(
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[0] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, expectedError
				},
				func() error {
					called[1] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[2] = true
					return nil
				},
			),
		)
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, []bool{true, false, false}, called)
	})
	t.Run("ReturnErrorOnCommand", func(t *testing.T) {
		called := make([]bool, 3)
		err := switching.MultipleHeavyCase(
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[0] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					return expectedError
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[2] = true
					return nil
				},
			),
		)
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, []bool{true, false, false}, called)
	})
}

func TestMultipleHeavyCase_OneCaseMatched(t *testing.T) {
	called := make([]bool, 2)
	var err error
	assert.NotPanics(t, func() {
		err = switching.MultipleHeavyCase(
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[0] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return false, nil
				},
				func() error {
					called[1] = true
					return nil
				},
			),
		)
	})
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false}, called)
}

func TestMultipleHeavyCase_AllCaseMatched(t *testing.T) {
	called := make([]bool, 2)
	var err error
	assert.NotPanics(t, func() {
		err = switching.MultipleHeavyCase(
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[0] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return true, nil
				},
				func() error {
					called[1] = true
					return nil
				},
			),
		)
	})
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, true}, called)
}

func TestMultipleHeavyCase_NoCaseMatched(t *testing.T) {
	called := make([]bool, 2)
	var err error
	assert.NotPanics(t, func() {
		err = switching.MultipleHeavyCase(
			switching.NewHeavyCase(
				func() (bool, error) {
					return false, nil
				},
				func() error {
					called[0] = true
					return nil
				},
			),
			switching.NewHeavyCase(
				func() (bool, error) {
					return false, nil
				},
				func() error {
					called[1] = true
					return nil
				},
			),
		)

	})
	assert.NoError(t, err)
	assert.Equal(t, []bool{false, false}, called)
}

func TestMultipleHeavyCase_NoCaseGiven(t *testing.T) {
	var err error
	assert.NotPanics(t, func() {
		err = switching.MultipleHeavyCase()
	})
	assert.NoError(t, err)
}
