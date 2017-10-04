package main

import (
	"GoaSample/app"
	"github.com/goadesign/goa"
)

// CommentController implements the comment resource.
type CommentController struct {
	*goa.Controller
}

// NewCommentController creates a comment controller.
func NewCommentController(service *goa.Service) *CommentController {
	return &CommentController{Controller: service.NewController("CommentController")}
}

// Create runs the create action.
func (c *CommentController) Create(ctx *app.CreateCommentContext) error {
	// CommentController_Create: start_implement

	// Put your logic here

	// CommentController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *CommentController) Delete(ctx *app.DeleteCommentContext) error {
	// CommentController_Delete: start_implement

	// Put your logic here

	// CommentController_Delete: end_implement
	return nil
}

// Show runs the show action.
func (c *CommentController) Show(ctx *app.ShowCommentContext) error {
	// CommentController_Show: start_implement

	// Put your logic here

	// CommentController_Show: end_implement
	res := &app.GoaComment{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *CommentController) Update(ctx *app.UpdateCommentContext) error {
	// CommentController_Update: start_implement

	// Put your logic here

	// CommentController_Update: end_implement
	return nil
}
