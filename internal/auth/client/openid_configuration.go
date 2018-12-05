// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": openid_configuration Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-auth/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-auth
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowOpenidConfigurationPath computes a request path to the show action of openid_configuration.
func ShowOpenidConfigurationPath() string {

	return fmt.Sprintf("/api/.well-known/openid-configuration")
}

// Show Indentity Provider Configuration. It lists all endpoints supported by Auth Service
func (c *Client) ShowOpenidConfiguration(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowOpenidConfigurationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowOpenidConfigurationRequest create the request corresponding to the show action endpoint of the openid_configuration resource.
func (c *Client) NewShowOpenidConfigurationRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
