// File: internal/embed/embed.go
package embed

import _ "embed"

//go:embed ../templates/server.tmpl
var ServerTmpl []byte
