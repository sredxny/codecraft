package codecraft

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// current coordenadas
// rotacion
// movimientos = front / back

func Test_CoordinatesRover(t *testing.T) {
	r := Rover{}
	r.X = 100
	r.Y = 200

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
			r := Rover{}
			r.Direction = NORTH
			r.Move(FORWARD)
			assert.Equal(t, r.X, 0)
			assert.Equal(t, r.Y, 1)
		})

		t.Run("south side", func(t *testing.T) {
			r := Rover{}
			r.Direction = SOUTH

			r.Move(FORWARD)
			assert.Equal(t, r.X, 0)
			assert.Equal(t, r.Y, -1)
		})

	})

	t.Run("backward", func(t *testing.T) {
		r := Rover{}
		r.Direction = NORTH
		r.Move(BACKWARD)
		assert.Equal(t, r.X, 0)
		assert.Equal(t, r.Y, -1)
	})
}
