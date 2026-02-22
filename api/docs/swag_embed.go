// Package docs 嵌入式Swagger文档
package docs

import _ "embed"

//go:embed swagger.yaml
var SwaggerYAML string

//go:embed swagger.json
var SwaggerJSON string
