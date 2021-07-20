package database

import (
	"cavall.in/syn/syn"
	"cloud.google.com/go/firestore"
	"context"
)

// Client is a wrapper around Firestore to more easily store data
type Client struct {
	ctx        context.Context
	Connection *firestore.Client
}

func NewClient(ctx context.Context, projectId string) (*Client, error) {
	connection, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &Client{ctx, connection}, nil
}

func (c *Client) AddEvent(event *syn.Event) (string, error) {
	doc, _, err := c.Connection.Collection("Events").Add(c.ctx, event)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
