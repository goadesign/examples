package controller

import (
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"

	"github.com/goadesign/goa"

	"github.com/goadesign/examples/appengine/app/test"
)

func TestCreateUser(t *testing.T) {
	// See https://cloud.google.com/appengine/docs/standard/go/tools/localunittesting/#Go_Introducing_the_aetest_package
	opt := &aetest.Options{StronglyConsistentDatastore: true}
	inst, err := aetest.NewInstance(opt)
	if err != nil {
		t.Errorf("NewInstance Error: %v\n", err)
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("NewRequest Error: %v\n", err)
	}
	ctx := appengine.NewContext(req)
	ctx = test.WithNewRequestFunc(ctx, test.NewRequestFunc(inst.NewRequest))

	service := goa.New("appengine")
	ctrl := NewHelloController(service)

	_, example := test.ShowHelloOK(t, ctx, service, ctrl)
	if example.Message == nil {
		t.Errorf("Nil message\n")
	} else if *example.Message != "Hello World" {
		t.Errorf("Invalid message: %v\n", example.Message)
	}
}
