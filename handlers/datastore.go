package handlers

import (
	"net/http"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/api/iterator"
)

type DataObject struct {
	client *datastore.Client
	ctx context.Context
}

func NewDataObject(projectId string, r *http.Request) (*DataObject, error) {
	ctx := appengine.NewContext(r)
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	return &DataObject{
		client,
		ctx,
	}, nil
}

type goStore struct {
	Key *datastore.Key `datastore:"__key__"`
	Url string `datastore:"url"`
}

func (d *DataObject) GetURL(name string) (string, error) {
	key := datastore.NameKey("golink", name, nil)
	var golink goStore
	if err := d.client.Get(d.ctx, key, &golink); err != nil {
		return "", err
	}
	return golink.Url, nil
}

type Golink struct {
	Name string
	Url string
}

func (d *DataObject) GetListOfLinks() ([]*Golink, error) {
	query := datastore.NewQuery("golink")
	it := d.client.Run(d.ctx, query)
	res := []*Golink{}
	for {
		var g goStore
		_, err := it.Next(&g)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		res = append(res, &Golink{
			Name: g.Key.Name,
			Url: g.Url,
		})
	}
	return res, nil
}

