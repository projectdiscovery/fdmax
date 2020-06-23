package autofdmax

import (
	"github.com/projectdiscovery/fdmax"
)

func init() {
	fdmax.Set(fdmax.Max)
}
