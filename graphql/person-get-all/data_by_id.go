package main

import (
	"fmt"
	"log"
)

func (p people) byId(id string) (*person, error) {
	log.Printf("request2: %v", id)

	for _, dp := range dataPeople {
		if dp.Id == id {
			log.Printf("request13: %v", dp)
			log.Printf("request14: %v", dp.Name)

			return dp, nil
		}
	}

	return nil, fmt.Errorf("cannot find person with ID: %s", id)
}
