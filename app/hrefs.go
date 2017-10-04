// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cellar": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=GoaSample/design
// --out=$(GOPATH)/src/GoaSample
// --version=v1.3.0

package app

import (
	"fmt"
	"strings"
)

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/bottles/%v", parambottleID)
}

// CommentHref returns the resource href.
func CommentHref(commentID interface{}) string {
	paramcommentID := strings.TrimLeftFunc(fmt.Sprintf("%v", commentID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/comments/comments/%v", paramcommentID)
}