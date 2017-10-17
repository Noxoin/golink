package handlers

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

type DataObject struct {
	Client *datastore.Client
	Ctx context.Context
}

func NewDataObject(projectId string) (*DataObject, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	return &DataObject{
		client,
		ctx,
	}, nil
}

type Golink struct {
	Url string `datastore:"url"`
}

func (d *DataObject) GetURL(name string) (string, error) {
	key := datastore.NameKey("golink", name, nil)
	var golink Golink
	if err := d.Client.Get(d.Ctx, key, &golink); err != nil {
		return "", err
	}
	return golink.Url, nil
}
