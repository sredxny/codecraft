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
	RIGHT = "R"
	LEFT = "L"
	)

func (r *Rover) Rotate(rotation string) {
	if r.Direction == NORTH && rotation == LEFT {
		r.Direction = WEST
	}

	if r.Direction == NORTH && rotation == RIGHT {
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
