package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/h2non/filetype"
	"io"
	"io/ioutil"
)

type Client struct {
	ctx        context.Context
	Connection *storage.Client
}

func NewClient(ctx context.Context) (*Client, error) {
	connection, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{ctx, connection}, nil
}

func (c *Client) GetObject(bucket string, name string) (*storage.ObjectHandle, *storage.ObjectAttrs, *storage.Reader, error) {
	object := c.Connection.Bucket(bucket).Object(name)
	attrs, err := object.Attrs(c.ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	rc, err := object.NewReader(c.ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	return object, attrs, rc, nil
}

func (c *Client) IsImage(reader io.Reader) (bool, error) {
	image, err := ioutil.ReadAll(reader)
	if err != nil {
		return false, err
	}

	if !filetype.IsImage(image) {
		return false, nil
	}

	return true, nil
}
