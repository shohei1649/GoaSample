// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cellar": comment Resource Client
//
// Command:
// $ goagen
// --design=GoaSample/design
// --out=$(GOPATH)/src/GoaSample
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateCommentPath computes a request path to the create action of comment.
func CreateCommentPath() string {

	return fmt.Sprintf("/comments/comments")
}

// Entry comment
func (c *Client) CreateComment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCreateCommentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateCommentRequest create the request corresponding to the create action endpoint of the comment resource.
func (c *Client) NewCreateCommentRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteCommentPath computes a request path to the delete action of comment.
func DeleteCommentPath(commentID string) string {
	param0 := commentID

	return fmt.Sprintf("/comments/comments/%s", param0)
}

// Delete comment
func (c *Client) DeleteComment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCommentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCommentRequest create the request corresponding to the delete action endpoint of the comment resource.
func (c *Client) NewDeleteCommentRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowCommentPath computes a request path to the show action of comment.
func ShowCommentPath(commentID int) string {
	param0 := strconv.Itoa(commentID)

	return fmt.Sprintf("/comments/comments/%s", param0)
}

// Get comment by id
func (c *Client) ShowComment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowCommentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCommentRequest create the request corresponding to the show action endpoint of the comment resource.
func (c *Client) NewShowCommentRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateCommentPath computes a request path to the update action of comment.
func UpdateCommentPath(commentID string) string {
	param0 := commentID

	return fmt.Sprintf("/comments/comments/%s", param0)
}

// Update comment
func (c *Client) UpdateComment(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUpdateCommentRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCommentRequest create the request corresponding to the update action endpoint of the comment resource.
func (c *Client) NewUpdateCommentRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
