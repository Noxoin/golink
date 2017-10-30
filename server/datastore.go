package server

import (
	"time"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type goStore struct {
	Key *datastore.Key `datastore:"__key__"`
	Url string `datastore:"url"`
	Timestamp int64 `datastore:"timestamp"`
}

type DataStore struct {
	client *datastore.Client
}

type Golink struct {
	Name string
	Url string
}

func NewDataStore(ctx context.Context, projectId string) (*DataStore, error) {
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	return &DataStore{
		client: client,
	}, nil
}

func (d *DataStore) GetURL(ctx context.Context, name string) (string, error) {
	key := datastore.NameKey("golink", name, nil)
	var val goStore
	if err := d.client.Get(ctx, key, &val); err != nil {
		return "", err
	}
	return val.Url, nil
}

func (d *DataStore) GetListOfLinks(ctx context.Context) ([]*Golink, error) {
	query := datastore.NewQuery("golink").Order("-timestamp")
	it := d.client.Run(ctx, query)
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

func (d *DataStore) UpdateLink(ctx context.Context, golink Golink) (error) {
	key := datastore.NameKey("golink", golink.Name, nil)
	val := goStore{
		Key: key,
		Url: golink.Url,
		Timestamp: time.Now().UnixNano(),
	}
	if _, err := d.client.Put(ctx, key, &val); err != nil {
		return err
	}
	return nil
}

