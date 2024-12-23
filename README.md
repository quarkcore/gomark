![gomark-icon](https://github.com/user-attachments/assets/b68c671c-944a-4a61-9f7e-ebc5bbd73bdf)

# GOMARK

With Gomark, you can effortlessly add, search, and open your bookmarks directly from the terminal without ever touching a browser's cluttered bookmark manager. Perfect for power users and developers who value speed and simplicity.

## PREREQUISITES

Fzf: https://github.com/junegunn/fzf

## USAGE
### Add a bookmark

From clipboard using `-c` option
```
gomark add -c
```

Directly passing a url
```
gomark add -t https://something.com/bookmarked/path
```

Assigning a group (aka subdir(s) of `~/.gomark`)
```
gomark add -c -g some/sub/group
```

Note: If setting `-c` and `-t` togehter latter just wins.

### Open
```
gomark open <some-pattern>
```

This internally launches fzf with the specified pattern. Fzf is set to use the option `--select-1` so if there is only one result found it will be opened directly.

If no pattern is passed all bookmarks would be found (Preferd way to have an alternative to a `gomark list` command which is currently missing)

## TROUBLEHSOOT

On Macos: either sell your mac and get a linux box or resign the binary after build:
```
codesign -f -s - /Users/<username>/go/bin/gomark
```
