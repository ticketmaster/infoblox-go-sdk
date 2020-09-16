package test

import (
	"log"
	"testing"

	"github.com/tickemaster/infoblox-go-sdk/client"
)

func TestCreateSession(t *testing.T) {
	s, err := client.NewClient(InfobloxServer, Username, Password)
	defer s.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%+v", s)
}
