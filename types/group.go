package types

import (
	"os"
	"quarkcore/gomark/utils"
)

type Group struct {
    path string
}

func (g *Group) groupExists() bool {
    _, err := os.Open(utils.RootPath() + g.path)
    if err != nil {
        return false
    }

    return true
}
