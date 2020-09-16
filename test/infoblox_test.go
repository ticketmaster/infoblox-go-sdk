package test

import (
	"log"
	"testing"

	"github.com/ticketmaster/infoblox-go-sdk"
	"github.com/ticketmaster/infoblox-go-sdk/client"
)

func TestInfobloxFetch(t *testing.T) {
	var err error
	filter := TArgs.HostRegex
	host := new(client.Host)
	host.Name = InfobloxServer
	host.UserName = Username
	host.Password = Password
	o := infoblox.Set(host)
	defer o.Unset()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.Fetch(filter)
	if err != nil {
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}

func TestInfobloxFetchByIP(t *testing.T) {
	var err error
	host := new(client.Host)
	host.Name = InfobloxServer
	host.UserName = Username
	host.Password = Password
	o := infoblox.Set(host)
	defer o.Unset()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	filter := TArgs.IP
	r, err := o.FetchByIP(filter)
	if err != nil {
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}

func TestCreateA(t *testing.T) {
	var err error
	host := new(client.Host)
	host.Name = InfobloxServer
	host.UserName = Username
	host.Password = Password
	o := infoblox.Set(host)
	defer o.Unset()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	req := infoblox.Record{}
	req.Name = TArgs.Host
	req.Alias = append(req.Alias, infoblox.Alias{Name: TArgs.Cname})
	req.Ipv4Addr = TArgs.IP
	err = o.DeleteA(req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.CreateA(req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}

func TestModify(t *testing.T) {
	///////////////////////////////////////
	TestCreateA(t)
	///////////////////////////////////////
	var err error
	host := new(client.Host)
	host.Name = InfobloxServer
	host.UserName = Username
	host.Password = Password
	o := infoblox.Set(host)
	defer o.Unset()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	///////////////////////////////////////
	req := infoblox.Record{}
	req.Name = TArgs.Host
	req.Ipv4Addr = TArgs.IP

	r, err := o.ModifyA(req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
