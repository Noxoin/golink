package server

import (
	"time"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type goStore struct {
	Key       *datastore.Key `datastore:"__key__"`
	Url       string         `datastore:"url"`
	Count     int64          `datastore:"count"`
	Timestamp int64          `datastore:"timestamp"`
}

type DataStore struct {
	client    *datastore.Client
	projectId string
	kind      string
}

type Golink struct {
	Name  string
	Url   string
	Count int64
}

func NewDataStore(ctx context.Context, projectId, kind string) (*DataStore, error) {
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	return &DataStore{
		client:    client,
		projectId: projectId,
		kind:      kind,
	}, nil
}

func (d *DataStore) GetURL(ctx context.Context, name string) (string, error) {
	key := datastore.NameKey(d.kind, name, nil)
	var val goStore
	if err := d.client.Get(ctx, key, &val); err != nil {
		return "", err
	}
	return val.Url, nil
}

func (d *DataStore) GetListOfLinks(ctx context.Context) ([]*Golink, error) {
	query := datastore.NewQuery(d.kind).Order("-timestamp")
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
			Name:  g.Key.Name,
			Url:   g.Url,
			Count: g.Count,
		})
	}
	return res, nil
}

func (d *DataStore) UpdateLink(ctx context.Context, golink Golink) error {
	key := datastore.NameKey(d.kind, golink.Name, nil)
	val := goStore{
		Key:       key,
		Url:       golink.Url,
		Count:     0,
		Timestamp: time.Now().UnixNano(),
	}
	if _, err := d.client.Put(ctx, key, &val); err != nil {
		return err
	}
	return nil
}

func (d *DataStore) IncrementCount(ctx context.Context, name string) error {
	key := datastore.NameKey(d.kind, name, nil)
	var val goStore
	if err := d.client.Get(ctx, key, &val); err != nil {
		return err
	}
	val.Count++
	if _, err := d.client.Put(ctx, key, &val); err != nil {
		return err
	}
	return nil
}
