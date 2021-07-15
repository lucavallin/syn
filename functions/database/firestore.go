package database

import (
	"cavall.in/syn/syn"
	"cloud.google.com/go/firestore"
	"context"
)

type Client struct {
	ctx        context.Context
	connection *firestore.Client
}

func NewClient(ctx context.Context, projectId string) (*Client, error) {
	connection, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &Client{ctx, connection}, nil
}

func (c *Client) AddUpload(upload *syn.Upload) (string, error) {
	doc, _, err := c.connection.Collection("Uploads").Add(c.ctx, upload)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
