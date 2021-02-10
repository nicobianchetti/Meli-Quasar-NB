package model

//Satellite .
type Satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

//DTORequestSatellites .
type DTORequestSatellites struct {
	Satellites []Satellite `json:"satellites"`
}

// //DTORequestSplit .
// type DTORequestSplit struct {
// 	distance
// }
