package vfields

import (
	"log"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	MoscowCenterPoint = NullPoint{
		X:     37.6155600,
		Y:     55.7522200,
		Valid: true,
	}
	MoscowRadius = 100000 // 100km
)

func ValidationDistanceByPoint(center NullPoint, district int) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		log.Printf("ValidationDistanceByPoint: %T", fl.Field())
		return false
		// log.Println("v++")
		// point, ok := fl.Field().Interface().(NullPoint)
		// if !ok {
		// 	return false
		// }
		// // TODO: fake check
		// return int(point.Distance(center)) <= district
	}
}

func ValidationInMoscow(fl validator.FieldLevel) bool {
	// TODO: implement
	log.Printf("DEBUG: check call validator, got data type %T\n", fl.Field().Interface())
	return false
}
