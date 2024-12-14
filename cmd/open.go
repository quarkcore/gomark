package cmd

import (
	"fmt"
	"os"
	"quarkcore/gomark/utils"

	"github.com/pkg/browser"
)

func findBookmark(pattern string) {
    fmt.Println("Finding bookmark", pattern)

    rootPath := utils.RootPath()

    err, outStr := utils.TriggerFzf(rootPath, pattern)

    if err != nil {
        fmt.Println("Error", err)
        return
    }

    fmt.Println("Selected", outStr)

    browser.OpenURL(outStr)
}
func Open() {
    args := os.Args[2:]

    if len(args) != 1 {
        fmt.Println("Invalid number of arguments gomark open <bookmark-pattern>")
        return
    }

    findBookmark(args[0])
}

