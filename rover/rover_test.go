package rover

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_CoordinatesRover(t *testing.T) {
	r := Rover{
		X:         110,
		Y:         220,
		Direction: "",
	}

	assert.Equal(t, r.X, 100, "X coordinate should be 100")
	assert.Equal(t, r.Y, 200, "Y coordinate should be 200")
}

func Test_Rotation(t *testing.T) {
	r := Rover{}

	t.Run("Rotate to left", func(t *testing.T) {
		r.Direction = NORTH
		r.Rotate(LEFT)
		assert.Equal(t, r.Direction, WEST)
	})

	t.Run("Rotate to right", func(t *testing.T) {
		r.Direction = NORTH
		r.Rotate(RIGHT)
		assert.Equal(t, r.Direction, EAST)
	})
}

func Test_Movement(t *testing.T) {

	t.Run("forward", func(t *testing.T) {

		t.Run("facing north", func(t *testing.T) {
			evaluateMovement(t, NORTH, FORWARD, 0, 1)
		})

		t.Run("south side", func(t *testing.T) {
			evaluateMovement(t, SOUTH, FORWARD, 0, -1)
		})

		t.Run("east side", func(t *testing.T) {
			evaluateMovement(t, EAST, FORWARD, 1, 0)
		})

		t.Run("west side", func(t *testing.T) {
			evaluateMovement(t, WEST, FORWARD, -1, 0)
		})
	})

	t.Run("backward", func(t *testing.T) {

		t.Run("facing north", func(t *testing.T) {
			evaluateMovement(t, NORTH, BACKWARD, 0, -1)
		})

		t.Run("south side", func(t *testing.T) {
			evaluateMovement(t, SOUTH, BACKWARD, 0, 1)
		})

		t.Run("east side", func(t *testing.T) {
			evaluateMovement(t, EAST, BACKWARD, -1, 0)
		})

		t.Run("west side", func(t *testing.T) {
			evaluateMovement(t, WEST, BACKWARD, 1, 0)
		})
	})
}

func evaluateMovement(t *testing.T, direction string, movement string, x int, y int) {
	r := Rover{}
	r.Direction = direction
	r.Move(movement)
	assert.Equal(t, r.X, x)
	assert.Equal(t, r.Y, y)
}
