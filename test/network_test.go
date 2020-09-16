package test

import (
	"testing"

	"github.com/tickemaster/infoblox-go-sdk/client"
)

func TestFetchByNetwork(t *testing.T) {
	var err error
	o := new(client.Network)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	r, err := o.FetchByNetwork(TArgs.CIDR)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", r)
}

func TestFetchNextAvailableIP(t *testing.T) {
	var err error
	o := new(client.Network)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ref, err := o.FetchByNetwork(TArgs.CIDR)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	r, err := o.FetchNextAvailableIP(ref.Result[0].Ref, 3)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", r)
}

func TestFetchTestNextAvailableIP(t *testing.T) {
	var err error
	o := new(client.Network)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ref, err := o.FetchByNetwork(TArgs.CIDR)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	r, err := o.FetchNextAvailableIP(ref.Result[0].Ref, 3)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", r)
}
