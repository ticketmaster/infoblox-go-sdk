package test

import (
	"log"
	"testing"

	"github.com/tickemaster/infoblox-go-sdk/client"
	"github.com/tickemaster/infoblox-go-sdk/model"
)

func TestCreateRecordCname(t *testing.T) {
	var err error
	o := new(client.RecordCname)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	//////////////////////////////////////////////////////////
	TestCreateRecordA(t)
	//////////////////////////////////////////////////////////
	req := model.RecordCnameCreateRequest{}
	req.Name = TArgs.Host
	req.Canonical = TArgs.Cname
	r, err := o.Create(req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}

func TestFetchCname(t *testing.T) {
	var err error
	o := new(client.RecordCname)
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
	log.Printf("%s", client.ToPrettyJSON(r))
}

func TestGetRecordCnameByName(t *testing.T) {
	var err error
	o := new(client.RecordCname)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.FetchByName(TArgs.Cname)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestGetRecordCnameByRef(t *testing.T) {
	var err error
	o := new(client.RecordCname)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.FetchByRef(TArgs.Cname)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestUpdateRecordCname(t *testing.T) {
	var err error
	o := new(client.RecordCname)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	req := model.RecordCnameUpdateRequest{}
	req.Name = TArgs.Cname
	r, err := o.Modify(TArgs.CnameRef, req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestDeleteRecordCname(t *testing.T) {
	var err error
	o := new(client.RecordCname)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.Delete(TArgs.CnameRef)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	TestDeleteRecordA(t)
	log.Printf("%s", client.ToPrettyJSON(r))
}
