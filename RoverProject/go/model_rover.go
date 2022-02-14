package mars

type Rover struct {

	XCoordinate int `json:"x_coordinate,omitempty"`

	YCoordinate int `json:"y_coordinate,omitempty"`

	Direction string `json:"direction,omitempty"`

	Moves string `json:"moves,omitempty"`

}