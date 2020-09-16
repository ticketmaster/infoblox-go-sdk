package test

import (
	"testing"

	"github.com/tickemaster/infoblox-go-sdk/client"
)

func TestFetchByRange(t *testing.T) {
	var err error
	o := new(client.Range)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	r, err := o.FetchByRange(TArgs.StartIP)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", r)
}
func TestFetchByComment(t *testing.T) {
	var err error
	o := new(client.Range)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	o.Filter = TArgs.Comment
	r, err := o.Fetch()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", r)
}
func TestFetchNextAvailableIPRange(t *testing.T) {
	var err error
	o := new(client.Range)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ref, err := o.FetchByRange(TArgs.StartIP)
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

func TestFetchTestNextAvailableIPRange(t *testing.T) {
	var err error
	o := new(client.Range)
	o.Client, err = client.NewClient(InfobloxServer, Username, Password)
	defer o.Client.LogOut()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	ref, err := o.FetchByRange(TArgs.StartIP)
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
