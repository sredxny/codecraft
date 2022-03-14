package codecraft

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// current coordenadas
// rotacion
// movimientos = front / back

func Test_CoordinatesRover(t *testing.T){
	r := Rover{}
	r.SetX(100)
	r.SetY(200)

	assert.Equal(t,r.X,100,"X coordinate should be 100")
	assert.Equal(t, r.Y, 200, "Y coordinate should be 200")
}

func Test_Rotation(t *testing.T){
	r := Rover{}

	t.Run("Rotate to left", func(t *testing.T) {
		r.Direction = "N"
		r.Rotate("L")
		assert.Equal(t,r.Direction,"W")
	})

	t.Run("Rotate to right", func(t *testing.T) {
		r.Direction = "N"
		r.Rotate("R")
		assert.Equal(t,r.Direction,"E")
	})
}

func Test_Movement(t *testing.T){

	t.Run("forward", func(t *testing.T) {

		t.Run("facing north", func(t *testing.T) {
			r := Rover{}
			r.Direction = "N"
			r.Move("F")
			assert.Equal(t,r.X,0)
			assert.Equal(t,r.Y,1)
		})

		t.Run("south side", func(t *testing.T) {
			r := Rover{}
			r.Direction = "S"

			r.Move("F")
			assert.Equal(t,r.X,0)
			assert.Equal(t,r.Y,-1)
		})

	})

	t.Run("backward", func(t *testing.T) {
		r := Rover{}
		r.Direction = "N"
		r.Move("B")
		assert.Equal(t,r.X,0)
		assert.Equal(t,r.Y,-1)
	})
}





