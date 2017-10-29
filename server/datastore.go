package server

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type goStore struct {
	Key *datastore.Key `datastore:"__key__"`
	Url string `datastore:"url"`
}

func getURL(ctx context.Context, name string) (string, error) {
	client, err := datastore.NewClient(ctx, ctx.Value("projectId").(string))
	if err != nil {
		return "", err
	}
	key := datastore.NameKey("golink", name, nil)
	var val goStore
	if err := client.Get(ctx, key, &val); err != nil {
		return "", err
	}
	return val.Url, nil
}

type Golink struct {
	Name string
	Url string
}

func getListOfLinks(ctx context.Context) ([]*Golink, error) {
	client, err := datastore.NewClient(ctx, ctx.Value("projectId").(string))
	if err != nil {
		return nil, err
	}
	query := datastore.NewQuery("golink")
	it := client.Run(ctx, query)
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

func updateLink(ctx context.Context, golink Golink) (error) {
	client, err := datastore.NewClient(ctx, ctx.Value("projectId").(string))
	if err != nil {
		return err
	}
	key := datastore.NameKey("golink", golink.Name, nil)
	val := goStore{
		Key: key,
		Url: golink.Url,
	}
	if _, err := client.Put(ctx, key, &val); err != nil {
		return err
	}
	return nil
}

