package test

import (
	"log"
	"testing"

	"github.com/ticketmaster/infoblox-go-sdk/client"
	"github.com/ticketmaster/infoblox-go-sdk/model"
)

func TestCreateRecordHost(t *testing.T) {
	var err error
	o := new(client.RecordHost)
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
	req := model.RecordHostCreateRequest{}
	req.Name = TArgs.Host
	req.Ipv4Addrs = append(req.Ipv4Addrs, model.Ipv4Addr{Ipv4Addr: TArgs.IP})
	r, err := o.Create(req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestGetAllRecordHost(t *testing.T) {
	var err error
	o := new(client.RecordHost)
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
func TestGetRecordHostByName(t *testing.T) {
	var err error
	o := new(client.RecordHost)
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
func TestGetRecordHostByRef(t *testing.T) {
	var err error
	o := new(client.RecordHost)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.FetchByRef(TArgs.HostRef)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestUpdateRecordHost(t *testing.T) {
	var err error
	o := new(client.RecordHost)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	req := model.RecordHostUpdateRequest{}
	req.Name = TArgs.HostRef
	req.Ipv4Addrs = []model.Ipv4Addr{model.Ipv4Addr{Ipv4Addr: TArgs.IP}}
	req.Aliases = []string{TArgs.Cname}
	r, err := o.Modify(TArgs.HostRef, req)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
func TestDeleteRecordHost(t *testing.T) {
	var err error
	o := new(client.RecordHost)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	r, err := o.Delete(TArgs.HostRef)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	log.Printf("%s", client.ToPrettyJSON(r))
}
