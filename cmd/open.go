package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"quarkcore/gomark/types"
	"quarkcore/gomark/utils"
	"strings"

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

    outPath := rootPath + "/" + outStr
    outPath = strings.TrimSpace(outPath)
    fmt.Println("Selected", outPath)

    file, err := os.Open(outPath)
    if err != nil {
        fmt.Println("FileError", err)
        return
    }

    defer file.Close()

    bookmark := types.Bookmark{}

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&bookmark)
    if err != nil {
        fmt.Println("Decoding err:", err)
        return
    }

    fmt.Println("Opening", bookmark.Url)
    browser.OpenURL(bookmark.Url)
}

func Open() {
    args := os.Args[2:]

    if len(args) != 1 {
        fmt.Println("Invalid number of arguments gomark open <bookmark-pattern>")
        return
    }

    findBookmark(args[0])
}

