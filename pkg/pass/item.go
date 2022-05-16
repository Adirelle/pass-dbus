package pass

import (
	"time"
)

type (
	ItemPath string

	Item struct {
		Path     ItemPath
		Created  time.Time
		Modified time.Time

		pass *Pass
	}
)
