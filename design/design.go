package design // The convention consists of naming the design
// package "design"
import (
	"github.com/goadesign/goa/design" // Use . imports to enable the DSL
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("cellar", func() { // API defines the microservice endpoint and
	apidsl.Title("The virtual wine cellar")    // other global properties. There should be one
	apidsl.Description("A simple goa service") // and exactly one API definition appearing in
	apidsl.Scheme("http")                      // the design.
	apidsl.Host("localhost:8080")
})

var _ = apidsl.Resource("bottle", func() { // Resources group related API endpoints
	apidsl.BasePath("/bottles")      // together. They map to REST resources for REST
	apidsl.DefaultMedia(BottleMedia) // services.

	apidsl.Action("show", func() { // Actions define a single API endpoint together
		apidsl.Description("Get bottle by id")   // with its path, parameters (both path
		apidsl.Routing(apidsl.GET("/:bottleID")) // parameters and querystring values) and payload
		apidsl.Params(func() {                   // (shape of the request body).
			apidsl.Param("bottleID", design.Integer, "Bottle ID")
		})
		apidsl.Response(design.OK)       // Responses define the shape and status code
		apidsl.Response(design.NotFound) // of HTTP responses.
	})
})

var _ = apidsl.Resource("comment", func() {
	apidsl.BasePath("/comments")

	apidsl.Action("show", func() {
		apidsl.Description("Get comment by id")
		apidsl.Routing(apidsl.GET("/comments/:comment_id"))
		apidsl.Params(func() {
			apidsl.Param("comment_id", design.Integer, "Comment ID")
		})
		apidsl.Response(design.OK, CommentMedia)
		apidsl.Response(design.NotFound)

	})

	apidsl.Action("create", func() {
		apidsl.Description("Entry comment")
		apidsl.Routing(apidsl.POST("/comments"))
		apidsl.Response(design.Created)

	})

	apidsl.Action("update", func() {
		apidsl.Description("Update comment")
		apidsl.Routing(apidsl.PUT("/comments/:comment_id"))
		apidsl.Response(design.NoContent)

	})
	apidsl.Action("delete", func() {
		apidsl.Description("Delete comment")
		apidsl.Routing(apidsl.DELETE("/comments/:comment_id"))
		apidsl.Response(design.NoContent)

	})

})
