package capi

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func loadEnvFile()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(0)
	}
}

var cma = NewTestCMAClient()

func NewTestCMAClient() cmaClient {
	loadEnvFile()
	var conf = CMAConfig{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		SpaceID: os.Getenv("SPACE_ID"),
		Environment: os.Getenv("ENVIRONMENT"),
		OrganisationID: os.Getenv("ORGANISATION_ID"),
	}

	return NewCMAClient(conf)
}

func TestSpace_List(t *testing.T) {
	s := NewSpace(cma)
	result := s.List()

	fmt.Printf("%#v", result.Total)
}

func TestSpace_Create(t *testing.T) {
	s := NewSpace(cma)
	result := s.Create("Another fine blog", "en")

	fmt.Println(result)
}

func TestSpace_Delete(t *testing.T) {
	s := NewSpace(cma)
	result := s.Delete("rdadmup1fbvp")

	fmt.Println(result)
}
