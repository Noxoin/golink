package server

import (
	"testing"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

func TestGetLink(t *testing.T) {
	ctx := context.WithValue(context.Background(), "projectId", "default")
	client, err := datastore.NewClient(ctx, ctx.Value("projectId").(string))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	key := datastore.NameKey("golink", "foo", nil)
	defer client.Delete(ctx, key)
	entity := &struct{
		Url string `datastore:"url"`
	}{
		Url: "https://www.google.com/",
	}
	if _, err := client.Put(ctx, key, entity); err != nil {
		t.Fatalf("Failed Setup: %v", err)
	}
	url, err := getURL(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}
	if url != "https://www.google.com/" {
		t.Errorf("TestGetLink failed: got: %v, want: https://www.google.com/", url)
	}
}

func TestUpdateLink(t *testing.T) {
	ctx := context.WithValue(context.Background(), "projectId", "default")
	client, err := datastore.NewClient(ctx, ctx.Value("projectId").(string))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	key := datastore.NameKey("golink", "testing", nil)
	defer client.Delete(ctx, key)
	golink := Golink {
		Name: "testing",
		Url: "https://www.noxoin.com/",
	}
	if err := updateLink(ctx, golink); err != nil {
		t.Fatal(err)
	}
	url, err := getURL(ctx, "testing")
	if err != nil {
		t.Fatal(err)
	}
	if url != "https://www.noxoin.com/" {
		t.Errorf("TestGetLink failed: got: %v, want: https://www.noxoin.com/", url)
	}
}

func TestGetListOfLinks(t *testing.T) {
	ctx := context.WithValue(context.Background(), "projectId", "default")
	client, err := datastore.NewClient(ctx, ctx.Value("projectId").(string))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	golink := Golink {
		Name: "testing",
		Url: "https://www.noxoin.com/",
	}
	if err := updateLink(ctx, golink); err != nil {
		t.Fatal(err)
	}
	defer client.Delete(ctx, datastore.NameKey("golink", "testing", nil))

	golink = Golink {
		Name: "foo",
		Url: "https://www.google.com/",
	}
	if err := updateLink(ctx, golink); err != nil {
		t.Fatal(err)
	}
	defer client.Delete(ctx, datastore.NameKey("golink", "foo", nil))

	golinks, err := getListOfLinks(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(golinks) != 2 {
		t.Errorf("TestGetListOfLinks failed: got: %v results, wanted: 2", len(golinks))
	}
	if golinks[0].Url != "https://www.google.com/" {
		t.Errorf("TestGetListOfLinks first result failed: got: %v, want: https://www.google.com/", golinks[0].Url)
	}
	if golinks[1].Url != "https://www.noxoin.com/" {
		t.Errorf("TestGetListOfLinks second result failed: got: %v, want: https://www.noxoin.com/", golinks[1].Url)
	}
}

