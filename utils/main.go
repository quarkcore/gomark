package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RootPath() string {
    rootPath := os.Getenv("GOMARK_ROOT")

    if rootPath == "" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            panic("Could not get user home directory")
        }

        rootPath = homeDir + "/.gomarks"
    }

    _, err := os.Stat(rootPath)
    if err != nil {
        os.Mkdir(rootPath, 0775)
    }

    return rootPath
}

func TriggerFzf(rootPath string, pattern string) (error, string) {
    cmd := exec.Command("fzf", "--query", pattern)

    cmd.Dir = rootPath

    cmd.Stdin = os.Stdin

    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }

    outStr := out.String()
    fmt.Println(outStr)

    return err, outStr
}

func ValidateArgOpt(args []string, idx int) string {
    optIdx := idx + 1
    if len(args) <= optIdx {
        panic("Missing argument for option " + args[idx])
    }

    return args[optIdx]

}
