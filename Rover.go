package codecraft

type Rover struct {
	X, Y      int
	Direction string
}

const (
	NORTH = "N"
	SOUTH = "S"
	EAST = "E"
	WEST = "W"

	)

func (r *Rover) Rotate(rotation string) {
	if r.Direction == NORTH && rotation == "L" {
		r.Direction = WEST
	}

	if r.Direction == NORTH && rotation == "R" {
		r.Direction = EAST
	}
}

func (r *Rover) Move(FoB string) {

	if FoB == "F" {
		if r.Direction == NORTH{
			r.Y++
		}else if r.Direction == SOUTH{
			r.Y--
		}

	} else {
		r.Y--
	}
}
