// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cellar": Application Controllers
//
// Command:
// $ goagen
// --design=GoaSample/design
// --out=$(GOPATH)/src/GoaSample
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// BottleController is the controller interface for the Bottle actions.
type BottleController interface {
	goa.Muxer
	Show(*ShowBottleContext) error
}

// MountBottleController "mounts" a Bottle resource controller on the given service.
func MountBottleController(service *goa.Service, ctrl BottleController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowBottleContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/bottles/:bottleID", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Show", "route", "GET /bottles/:bottleID")
}

// CommentController is the controller interface for the Comment actions.
type CommentController interface {
	goa.Muxer
	Create(*CreateCommentContext) error
	Delete(*DeleteCommentContext) error
	Show(*ShowCommentContext) error
	Update(*UpdateCommentContext) error
}

// MountCommentController "mounts" a Comment resource controller on the given service.
func MountCommentController(service *goa.Service, ctrl CommentController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCommentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/comments/comments", ctrl.MuxHandler("create", h, nil))
	service.LogInfo("mount", "ctrl", "Comment", "action", "Create", "route", "POST /comments/comments")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteCommentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/comments/comments/:comment_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Comment", "action", "Delete", "route", "DELETE /comments/comments/:comment_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowCommentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/comments/comments/:comment_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Comment", "action", "Show", "route", "GET /comments/comments/:comment_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCommentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/comments/comments/:comment_id", ctrl.MuxHandler("update", h, nil))
	service.LogInfo("mount", "ctrl", "Comment", "action", "Update", "route", "PUT /comments/comments/:comment_id")
}
