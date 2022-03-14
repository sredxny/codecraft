package codecraft

type Rover struct {
	X, Y      int
	Direction string
}

func (r *Rover) SetX(i int) {
	r.X = i
}

func (r *Rover) SetY(i int) {
	r.Y = i
}

func (r *Rover) Rotate(rotation string) {
	if r.Direction == "N" && rotation == "L" {
		r.Direction = "W"
	}

	if r.Direction == "N" && rotation == "R" {
		r.Direction = "E"
	}
}

func (r *Rover) Move(FoB string) {

	if FoB == "F" {
		if r.Direction == "N"{
			r.Y++
		}else if r.Direction == "S"{
			r.Y--
		}

	} else {
		r.Y--
	}
}
