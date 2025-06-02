package design

import (
	. "goa.design/goa/v3/dsl"
)

// API definition
var _ = API("concerts", func() {
	Title("Concert Management API")
	Description("A simple API for managing music concert information")
	Version("1.0")

	Server("concerts", func() {
		Description("Concert management server")
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

// Service definition
var _ = Service("concerts", func() {
	Description("The concerts service manages music concert data. It provides CRUD operations for concert information including artist details, venues, dates, and pricing.")

	Method("list", func() {
		Description("List concerts with optional pagination. Returns an array of concerts sorted by date.")
		Meta("openapi:summary", "List all concerts")

		Payload(func() {
			Attribute("page", Int, "Page number for pagination", func() {
				Minimum(1)
				Default(1)
				Example(1)
				Description("Must be 1 or greater")
			})
			Attribute("limit", Int, "Number of items per page", func() {
				Minimum(1)
				Maximum(100)
				Default(10)
				Example(10)
				Description("Must be between 1 and 100")
			})
		})

		Result(ArrayOf(Concert), func() {
			Description("Array of concerts")
		})

		HTTP(func() {
			GET("/concerts")

			Param("page", Int, "Page number", func() {
				Minimum(1)
				Example(1)
			})
			Param("limit", Int, "Number of items per page", func() {
				Minimum(1)
				Maximum(100)
				Example(10)
			})

			Response(StatusOK)
		})
	})

	Method("create", func() {
		Description("Create a new concert entry. All fields are required to ensure complete concert information.")
		Meta("openapi:summary", "Create a new concert")

		Payload(ConcertPayload, "Concert information to create")

		Result(Concert, "The newly created concert")

		Error("bad_request", ErrorResult, "Invalid input data provided")

		HTTP(func() {
			POST("/concerts")

			Response(StatusCreated)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("show", func() {
		Description("Get a single concert by its unique ID.")
		Meta("openapi:summary", "Get concert by ID")

		Payload(func() {
			Attribute("concertID", String, "Unique concert identifier", func() {
				Format(FormatUUID)
				Example("550e8400-e29b-41d4-a716-446655440000")
			})
			Required("concertID")
		})

		Result(Concert, "The requested concert")

		Error("not_found", ErrorResult, "Concert with the specified ID was not found")

		HTTP(func() {
			GET("/concerts/{concertID}")

			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("update", func() {
		Description("Update an existing concert by ID. Only provided fields will be updated.")
		Meta("openapi:summary", "Update concert")

		Payload(func() {
			Extend(ConcertData)
			Attribute("concertID", String, "ID of the concert to update", func() {
				Format(FormatUUID)
				Example("550e8400-e29b-41d4-a716-446655440000")
			})
			Required("concertID")
		})

		Result(Concert, "The updated concert with all current information")

		Error("not_found", ErrorResult, "Concert with the specified ID was not found")
		Error("bad_request", ErrorResult, "Invalid update data provided")

		HTTP(func() {
			PUT("/concerts/{concertID}")

			Response(StatusOK)
			Response("not_found", StatusNotFound)
			Response("bad_request", StatusBadRequest)
		})
	})

	Method("delete", func() {
		Description("Remove a concert from the system by ID. This operation cannot be undone.")
		Meta("openapi:summary", "Delete concert")

		Payload(func() {
			Attribute("concertID", String, "ID of the concert to remove", func() {
				Format(FormatUUID)
				Example("550e8400-e29b-41d4-a716-446655440000")
			})
			Required("concertID")
		})

		Error("not_found", ErrorResult, "Concert with the specified ID was not found")

		HTTP(func() {
			DELETE("/concerts/{concertID}")

			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})
})

// Data Types
var ConcertData = Type("ConcertData", func() {
	Description("Concert information fields.")

	Attribute("artist", String, "Name of the performing artist or band", func() {
		MinLength(1)
		MaxLength(200)
		Example("The White Stripes")
		Description("The main performer for this concert")
	})

	Attribute("date", String, "Concert date in ISO 8601 format (YYYY-MM-DD)", func() {
		Format(FormatDate)
		Example("2024-12-25")
		Description("The date when the concert will take place")
	})

	Attribute("venue", String, "Name and location of the concert venue", func() {
		MinLength(1)
		MaxLength(300)
		Example("Madison Square Garden, New York, NY")
		Description("The venue where the concert will be held")
	})

	Attribute("price", Int, "Ticket price in US dollars (cents)", func() {
		Minimum(0)
		Maximum(100000) // $1000 max
		Example(7500)   // $75.00
		Description("Base ticket price in cents (e.g., 7500 = $75.00)")
	})
})

var ConcertPayload = Type("ConcertPayload", func() {
	Description("Concert information required for creating a new concert entry.")

	Extend(ConcertData)

	// All fields required for creation
	Required("artist", "date", "venue", "price")
})

var Concert = Type("Concert", func() {
	Description("Complete concert information including system-generated ID.")

	Attribute("id", String, "Unique concert identifier", func() {
		Format(FormatUUID)
		Example("550e8400-e29b-41d4-a716-446655440000")
		Description("System-generated unique identifier")
	})

	Extend(ConcertData)

	Required("id", "artist", "date", "venue", "price")
})
