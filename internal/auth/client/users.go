// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": users Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-auth/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-auth
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateUsersPayload is the users create action payload.
type CreateUsersPayload struct {
	Data *CreateUserData `form:"data" json:"data" xml:"data"`
}

// CreateUsersPath computes a request path to the create action of users.
func CreateUsersPath() string {

	return fmt.Sprintf("/api/users")
}

// create a user using a service account
func (c *Client) CreateUsers(ctx context.Context, path string, payload *CreateUsersPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUsersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUsersRequest create the request corresponding to the create action endpoint of the users resource.
func (c *Client) NewCreateUsersRequest(ctx context.Context, path string, payload *CreateUsersPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListUsersPath computes a request path to the list action of users.
func ListUsersPath() string {

	return fmt.Sprintf("/api/users")
}

// List all users.
func (c *Client) ListUsers(ctx context.Context, path string, filterEmail *string, filterUsername *string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewListUsersRequest(ctx, path, filterEmail, filterUsername, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUsersRequest create the request corresponding to the list action endpoint of the users resource.
func (c *Client) NewListUsersRequest(ctx context.Context, path string, filterEmail *string, filterUsername *string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if filterEmail != nil {
		values.Set("filter[email]", *filterEmail)
	}
	if filterUsername != nil {
		values.Set("filter[username]", *filterUsername)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if ifModifiedSince != nil {

		header.Set("If-Modified-Since", *ifModifiedSince)
	}
	if ifNoneMatch != nil {

		header.Set("If-None-Match", *ifNoneMatch)
	}
	return req, nil
}

// SendEmailVerificationCodeUsersPath computes a request path to the sendEmailVerificationCode action of users.
func SendEmailVerificationCodeUsersPath() string {

	return fmt.Sprintf("/api/users/verificationcode")
}

// Send a verification code to the user's email address
func (c *Client) SendEmailVerificationCodeUsers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSendEmailVerificationCodeUsersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSendEmailVerificationCodeUsersRequest create the request corresponding to the sendEmailVerificationCode action endpoint of the users resource.
func (c *Client) NewSendEmailVerificationCodeUsersRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowUsersPath computes a request path to the show action of users.
func ShowUsersPath(id string) string {
	param0 := id

	return fmt.Sprintf("/api/users/%s", param0)
}

// Retrieve user for the given ID.
func (c *Client) ShowUsers(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewShowUsersRequest(ctx, path, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUsersRequest create the request corresponding to the show action endpoint of the users resource.
func (c *Client) NewShowUsersRequest(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if ifModifiedSince != nil {

		header.Set("If-Modified-Since", *ifModifiedSince)
	}
	if ifNoneMatch != nil {

		header.Set("If-None-Match", *ifNoneMatch)
	}
	return req, nil
}

// UpdateUsersPayload is the users update action payload.
type UpdateUsersPayload struct {
	Data *UpdateUserData `form:"data" json:"data" xml:"data"`
}

// UpdateUsersPath computes a request path to the update action of users.
func UpdateUsersPath() string {

	return fmt.Sprintf("/api/users")
}

// update the authenticated user
func (c *Client) UpdateUsers(ctx context.Context, path string, payload *UpdateUsersPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUsersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUsersRequest create the request corresponding to the update action endpoint of the users resource.
func (c *Client) NewUpdateUsersRequest(ctx context.Context, path string, payload *UpdateUsersPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// VerifyEmailUsersPath computes a request path to the verifyEmail action of users.
func VerifyEmailUsersPath() string {

	return fmt.Sprintf("/api/users/verifyemail")
}

// Verify if the new email updated by the user is a valid email
func (c *Client) VerifyEmailUsers(ctx context.Context, path string, code string) (*http.Response, error) {
	req, err := c.NewVerifyEmailUsersRequest(ctx, path, code)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewVerifyEmailUsersRequest create the request corresponding to the verifyEmail action endpoint of the users resource.
func (c *Client) NewVerifyEmailUsersRequest(ctx context.Context, path string, code string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("code", code)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
