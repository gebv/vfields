package vfields

import (
	"encoding/json"
	"math"
)

type NullPoint struct {
	X     float64 `json:"lng"`
	Y     float64 `json:"lat"`
	Valid bool    `json:"valid"` // Valid is true if not NULL
}

func (p NullPoint) Value() (interface{}, error) {
	return p, nil
}

func (p NullPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			X float64 `json:"lng"`
			Y float64 `json:"lat"`
		}{
			p.X,
			p.Y,
		},
	)
}

func (p *NullPoint) UnmarshalJSON(_b []byte) error {
	dto := &struct {
		X float64 `json:"lng"`
		Y float64 `json:"lat"`
	}{}
	if err := json.Unmarshal(_b, dto); err != nil {
		return err
	}
	p.X = dto.X
	p.Y = dto.Y
	if p.X == 0 && p.Y == 0 {
		return nil
	}

	p.Valid = true
	return nil
}

const earthRadius = 6371e3

// Distance return distance in m
// calculate distane in meters between two geo points http://www.movable-type.co.uk/scripts/latlong.html
func (p *NullPoint) Distance(_p NullPoint) float64 {
	if !p.Valid {
		return 0
	}

	phiOne := degreeToRadians(p.Y)
	phiTwo := degreeToRadians(_p.Y)

	deltaPhi := degreeToRadians(_p.Y - p.Y)
	deltaLambda := degreeToRadians(_p.X - p.X)

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phiOne)*math.Cos(phiTwo)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// degreeToRadians convert degree to radians
func degreeToRadians(degree float64) float64 {
	return degree * (math.Pi / 180)
}
