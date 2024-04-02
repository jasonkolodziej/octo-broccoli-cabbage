package redact

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsGo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG_PGR")).
		RegisterModule(redactor()).
		RegisterPostProcessor(pgsGo.GoFmt()).
		Render()
}
