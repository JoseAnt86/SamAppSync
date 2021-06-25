package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	resolvers "github.com/sbstjn/appsync-resolvers"
)

// var (
// 	// DefaultHTTPGetAddress Default Address
// 	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

// 	// ErrNoIP No IP found in response
// 	ErrNoIP = errors.New("No IP in HTTP response")

// 	// ErrNon200Response non 200 status code in response
// 	ErrNon200Response = errors.New("Non 200 Response found")
// )

type personEvent struct {
	ID string `json:"id"`
}

type personEvent2 struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type person struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}
type people []*person

func handlePeoples() (people, error) {
	return dataPeople, nil
}

func handlePerson(pe personEvent2) (*person, error) {
	return dataPeople.byId(pe.ID)
}

func handlePersonCreate(pe person) (people, error) {
	return dataPeople.dataPersonCreate(pe)
}

func handlePersonDelete(pe personEvent) (people, error) {
	return dataPeople.dataPersonDelete(pe)
}

func handlePersonUpdate(pe person) (people, error) {
	return dataPeople.dataPersonUpdate(pe)
}

var (
	r = resolvers.New()
)

func init() {

	r.Add("query.people", handlePeoples)
	r.Add("query.getPerson", handlePerson)

	r.Add("mutation.personCreate", handlePersonCreate)
	r.Add("mutation.personDelete", handlePersonDelete)
	r.Add("mutation.personUpdate", handlePersonUpdate)

}

func main() {
	log.Printf("request0")
	log.Printf("request4: %v", r)
	handler, found := r["query.people"]
	log.Printf("request5: %v", handler)
	log.Printf("request6: %v", found)

	lambda.Start(r.Handle)
}
