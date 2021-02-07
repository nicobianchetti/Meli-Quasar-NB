package model

//SateliteResult .
type SateliteResult struct {
	position Position
	message  string
}

//DTOResult .
type DTOResult struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

//Position .
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

//NewPosition .
func NewPosition(x, y float64) *Position {
	return &Position{X: x, Y: y}
}

//NewDTOResult .
func NewDTOResult(p *Position, message string) *DTOResult {
	return &DTOResult{Position: *p, Message: message}
}
