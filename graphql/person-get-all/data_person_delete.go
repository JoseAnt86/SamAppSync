package main

import (
	"log"
)

func (p people) dataPersonDelete(pe personEvent) (people, error) {
	var newPeople []*person
	log.Printf("request30: %v", pe)
	log.Printf("request31: %v", pe.ID)

	for _, p := range dataPeople {
		if p.Id != pe.ID {
			log.Printf("request31: %v", p.Id)

			newPeople = append(newPeople, p)
		}
	}
	dataPeople = newPeople
	return dataPeople, nil
}
