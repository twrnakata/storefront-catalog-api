package apidocs

import "embed"

//go:embed index.html openapi.en.yaml openapi.th.yaml
var Files embed.FS
