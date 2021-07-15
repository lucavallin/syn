package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"log"
	"time"
)

// Client is a wrapper around Cloud Storage to more easily read/write objects
type Client struct {
	ctx        context.Context
	Connection *storage.Client
}

type Object struct {
	Bucket  string
	Name    string
	Created time.Time
	URI     string
}

func NewClient(ctx context.Context) (*Client, error) {
	connection, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{ctx, connection}, nil
}

func (c *Client) GetObject(bucket string, name string) (*Object, error) {
	object := c.Connection.Bucket(bucket).Object(name)
	attrs, err := object.Attrs(c.ctx)
	if err != nil {
		return nil, err
	}

	return &Object{
		Bucket:  object.BucketName(),
		Name:    object.ObjectName(),
		Created: attrs.Created,
		URI:     fmt.Sprintf("gs://%s/%s", object.BucketName(), object.ObjectName()),
	}, nil
}

func (c *Client) Delete(object *Object) error {
	if err := c.Connection.Bucket(object.Bucket).Object(object.Name).Delete(c.ctx); err != nil {
		log.Printf("Failed to delete upload: %s", object.Name)
		return err
	}

	return nil
}
