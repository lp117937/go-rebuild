//go:build !prod
// +build !prod

package warning

import "github.com/lp117937/go-rebuild"

func init() {
	rebuild.Run(false)
}
