package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

var CommentMedia = apidsl.MediaType("application/vnd.goa.comment+json", func() {
	apidsl.Description("Comment")
	apidsl.Attributes(func() {
		apidsl.Attribute("id", design.Integer, "comment ID")
		apidsl.Attribute("content", design.String, "comment content")
		apidsl.Attribute("created_dt", design.DateTime, "comment created datetime")
		apidsl.Attribute("updated_dt", design.DateTime, "comment updated datetime")
		apidsl.Required("id", "content")

	})
	apidsl.View("default", func() {
		apidsl.Attribute("id")
		apidsl.Attribute("content")
		apidsl.Attribute("created_dt")
		apidsl.Attribute("updated_dt")
	})

})

// BottleMedia defines the media type used to render bottles.
var BottleMedia = apidsl.MediaType("application/vnd.goa.example.bottle+json", func() {
	apidsl.Description("A bottle of wine")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("id", design.Integer, "Unique bottle ID")
		apidsl.Attribute("href", design.String, "API href for making requests on the bottle")
		apidsl.Attribute("name", design.String, "Name of wine")
		apidsl.Required("id", "href", "name")
	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("id")   // Media types may have multiple views and must
		apidsl.Attribute("href") // have a "default" view.
		apidsl.Attribute("name")
	})
})
