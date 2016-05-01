# Types Example

This example shows how to use the DSL to define types. A factor important to remember is that types
defined in the design are not meant to solely represent Go types. The design defines the API data
structures independently of any language or representation. The types must map to Go, JavaScript,
Swagger, JSON, etc. Generators must be able to understand them to generate the artifacts.

The example uses multiple design files, one per type being showcased:

* [primitives.go](design/primitives.go) uses all the primitive type defined by the goa design language.
* [array.go](design/array.go) shows how to define arrays (used for types) and collections (used for media types).
* [hash.go](design/hash.go) shows how to define hashes (a.k.a. maps in Go, objects in JSON).
* [recursive.go](design/recursive.go) shows how to combine arrays and hashes in recursive data structures.

The file [generate.go](generate.go) contains `go generate` directives that uses `goagen` to generate
all the built-in outputs. Run it with:
```go
go generate
```
Inspect the generated files to see how each type maps to each generation output.

Additional information on designing types can be found on
[goa.design](https://goa.design/design/types/).
