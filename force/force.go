//go:build !prod
// +build !prod

package force

import "github.com/lp117937/go-rebuild"

func init() {
	rebuild.Run(true)
}
