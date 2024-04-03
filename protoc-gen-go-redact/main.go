package main

import (
	"github.com/jasonkolodziej/protoc-gen-go-redact/protoc-gen-go-redact/internal"
	pgs "github.com/lyft/protoc-gen-star"
	pgsGo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG_PGR")).
		RegisterModule(internal.Redactor()).
		RegisterPostProcessor(pgsGo.GoFmt()).
		Render()
}
