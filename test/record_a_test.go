package test

import (
	"log"
	"testing"

	"github.com/tickemaster/infoblox-go-sdk/client"
	"github.com/tickemaster/infoblox-go-sdk/model"
)

func TestCreateRecordA(t *testing.T) {
	var err error
	o := new(client.RecordA)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	refO, err := o.FetchByName(TArgs.Host)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	if len(refO.Result) > 0 && refO.Result[0].Ref != "" {
		_, err := o.Delete(refO.Result[0].Ref)
		if err != nil {
			log.Print(err)
			t.Fail()
		}
	}
	req := model.RecordACreateRequest{}
	req.Name = TArgs.Host
	req.Ipv4Addr = TArgs.IP
	r, err := o.Create(req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestGetAllRecordA(t *testing.T) {
	var err error
	o := new(client.RecordA)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	o.Filter = TArgs.FilterAndRegex
	r, err := o.Fetch()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Print(len(r.Result))
}
func TestGetRecordAByName(t *testing.T) {
	var err error
	o := new(client.RecordA)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.FetchByName(TArgs.Host)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestGetRecordAByRef(t *testing.T) {
	var err error
	o := new(client.RecordA)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.FetchByRef(TArgs.ARef)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestUpdateRecordA(t *testing.T) {
	var err error
	o := new(client.RecordA)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	req := model.RecordAUpdateRequest{}
	req.Name = TArgs.Host
	req.Ipv4Addr = TArgs.IP
	r, err := o.Modify(TArgs.ARef, req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestDeleteRecordA(t *testing.T) {
	var err error
	o := new(client.RecordA)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.Delete(TArgs.ARef)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
