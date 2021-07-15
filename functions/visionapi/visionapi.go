package visionapi

import (
	"cavall.in/syn/syn"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"github.com/thoas/go-funk"
	visionTypes "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"io"
)

type Client struct {
	ctx        context.Context
	Connection *vision.ImageAnnotatorClient
}

func NewClient(ctx context.Context) (*Client, error) {
	connection, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{ctx, connection}, nil
}

func (c *Client) DetectImageLabels(rc io.Reader) ([]syn.Label, error) {
	image, err := vision.NewImageFromReader(rc)
	if err != nil {
		return nil, err
	}

	res, err := c.Connection.DetectLabels(c.ctx, image, nil, 5)
	if err != nil {
		return nil, err
	}

	labels := funk.Map(res, func(l *visionTypes.EntityAnnotation) syn.Label {
		return syn.NewLabel(l.Description, l.Score)
	}).([]syn.Label)

	return labels, nil
}
