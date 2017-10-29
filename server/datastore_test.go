package server

import (
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
)

func TestGetLink(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	ctx = context.WithValue(ctx, "projectId", "api-project-377888563324")
	key := datastore.NewKey(ctx, "golink", "foo", 0, nil)
	entity := &struct{
		Url string `datastore:"url"`
	}{
		Url: "https://www.google.com/",
	}
	if _, err := datastore.Put(ctx, key, entity); err != nil {
		t.Fatalf("Failed Setup: %v", err)
	}
	url, err := getURL(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if url != "https://www.google.com" {
		t.Errorf("TestGetLink failed: got: %v, want: https://www.google.com/", url)
	}
}

/*
func TestUpdateLink(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	ctx = context.WithValue(ctx, "projectId", "api-project-377888563324")
	golink := Golink {
		Name: "testing",
		Url: "https://www.noxoin.com/",
	}
	if err := updateLink(ctx, golink); err != nil {
		t.Fatal(err)
	}
}
*/

