package client

import (
	"net/http"

	"gopkg.in/h2non/gentleman.v2"
)

func New(version, baseURL, token string) *Client {
	if token == "" {
		token = "--"
	}
	return &Client{
		http: gentleman.New().
			BaseURL(baseURL).
			SetHeader("User-Agent", "tog-cli/"+version).
			SetHeader("Authorization", "Bearer "+token),
	}
}

type Client struct {
	http *gentleman.Client
}

func (c *Client) ListFlags(ns string) ([]*Flag, error) {
	res, err := c.http.Get().AddPath("/flags/" + ns).Do()
	if err != nil {
		return nil, err
	}

	if err := c.validateStatus(res); err != nil {
		return nil, err
	}

	var flags []*Flag
	if err := res.JSON(&flags); err != nil {
		return nil, err
	}

	return flags, nil
}

func (c *Client) GetFlag(ns, name string) (*Flag, error) {
	res, err := c.http.Get().AddPath("/flags/" + ns + "/" + name).Do()
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if err := c.validateStatus(res); err != nil {
		return nil, err
	}

	var flag Flag
	if err := res.JSON(&flag); err != nil {
		return nil, err
	}

	return &flag, nil
}

func (c *Client) SaveFlag(flag *Flag) (*Flag, error) {
	clone := *flag
	clone.Namespace = ""
	clone.Name = ""

	res, err := c.http.Put().
		AddPath("/flags/" + flag.Namespace + "/" + flag.Name).
		JSON(clone).
		Do()
	if err != nil {
		return nil, err
	}

	if err := c.validateStatus(res); err != nil {
		return nil, err
	}

	var rFlag *Flag
	if err := res.JSON(&rFlag); err != nil {
		return nil, err
	}

	return rFlag, nil
}

func (c *Client) DeleteFlag(ns, name string) error {
	res, err := c.http.Delete().
		AddPath("/flags/" + ns + "/" + name).
		Do()
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if err := c.validateStatus(res); err != nil {
		return err
	}

	return nil
}

func (c *Client) validateStatus(res *gentleman.Response) error {
	if res.Ok {
		return nil
	}

	if res.StatusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	return &ServerError{res.StatusCode, res.String()}
}
