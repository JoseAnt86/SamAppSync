package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	resolvers "github.com/sbstjn/appsync-resolvers"
)

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

	lambda.Start(r.Handle)
}
