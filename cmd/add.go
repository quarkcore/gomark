package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"quarkcore/gomark/types"
	"quarkcore/gomark/utils"
	"regexp"
	"strings"
	"unicode"

	"golang.design/x/clipboard"
)

func Add() {
    commandArgs := os.Args[2:]

    target := ""
    clipTarget := ""
    group := ""
    shortcut := ""

    for i, arg := range commandArgs {
        switch arg {
            case "-t":
                target = utils.ValidateArgOpt(commandArgs, i)
                fmt.Println("target", target)
            case "-c":
                err := clipboard.Init()
                if err != nil {
                    panic(err)
                }

                cBytes := clipboard.Read(clipboard.FmtText)
                clipTarget = string(cBytes[:])
                fmt.Println("clipboard")
            case "-g":
                group = utils.ValidateArgOpt(commandArgs, i)
                fmt.Println("group", group)
            case "-s":
                shortcut = utils.ValidateArgOpt(commandArgs, i)
                fmt.Println("shortcut", shortcut)
        }
    }


    if target == "" && clipTarget == "" {
        fmt.Println("No target specified")
        return
    }

    if target == "" {
        target = clipTarget
    }

    fmt.Println("Adding bookmark...", target)

    groupPath := utils.RootPath() + "/" + group

    err := os.MkdirAll(groupPath, os.ModePerm)
    if err != nil {
        fmt.Println("Failed to create group directory")
        return
    }

    fileName := sanitizeForFilename(target)
    targetPath := groupPath + "/" + fileName + ".url"

    fmt.Println("Target path", targetPath)

    bookmark := types.Bookmark {
        Name: target,
        Url: target,
    }

    file, err := os.Create(targetPath)
    if err != nil {
        fmt.Println("Failed to create bookmark file")
        return
    }

    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.Encode(bookmark)
}

func sanitizeForFilename(input string) string {
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	safe := re.ReplaceAllString(input, "_")

	safe = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) && !unicode.IsControl(r) {
			return r
		}
		return -1
	}, safe)

	safe = strings.ReplaceAll(safe, " ", "_")
	safe = regexp.MustCompile(`_+`).ReplaceAllString(safe, "_")

	return strings.Trim(safe, "_")
}
