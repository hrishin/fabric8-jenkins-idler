// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": authorize Resource Client
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

// AuthorizeAuthorizePath computes a request path to the authorize action of authorize.
func AuthorizeAuthorizePath() string {

	return fmt.Sprintf("/api/authorize")
}

// Authorize service client
func (c *Client) AuthorizeAuthorize(ctx context.Context, path string, clientID string, redirectURI string, responseType string, state string, apiClient *string, responseMode *string, scope *string) (*http.Response, error) {
	req, err := c.NewAuthorizeAuthorizeRequest(ctx, path, clientID, redirectURI, responseType, state, apiClient, responseMode, scope)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAuthorizeAuthorizeRequest create the request corresponding to the authorize action endpoint of the authorize resource.
func (c *Client) NewAuthorizeAuthorizeRequest(ctx context.Context, path string, clientID string, redirectURI string, responseType string, state string, apiClient *string, responseMode *string, scope *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("client_id", clientID)
	values.Set("redirect_uri", redirectURI)
	values.Set("response_type", responseType)
	values.Set("state", state)
	if apiClient != nil {
		values.Set("api_client", *apiClient)
	}
	if responseMode != nil {
		values.Set("response_mode", *responseMode)
	}
	if scope != nil {
		values.Set("scope", *scope)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CallbackAuthorizePath computes a request path to the callback action of authorize.
func CallbackAuthorizePath() string {

	return fmt.Sprintf("/api/authorize/callback")
}

// Authorize service client callback
func (c *Client) CallbackAuthorize(ctx context.Context, path string, code string, state string) (*http.Response, error) {
	req, err := c.NewCallbackAuthorizeRequest(ctx, path, code, state)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCallbackAuthorizeRequest create the request corresponding to the callback action endpoint of the authorize resource.
func (c *Client) NewCallbackAuthorizeRequest(ctx context.Context, path string, code string, state string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("code", code)
	values.Set("state", state)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
