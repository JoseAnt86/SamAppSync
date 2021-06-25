package main

import (
	"log"
)

func (p people) dataPersonCreate(pe person) (people, error) {
	log.Printf("request20: %v", dataPeople)

	dataPeople = append(dataPeople, &pe)
	log.Printf("request21: %v", dataPeople)

	return dataPeople, nil
}
