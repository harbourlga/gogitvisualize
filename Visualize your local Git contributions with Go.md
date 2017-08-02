Run `go install` and

- `gogitlocalstats -add /path/to/folder` will scan that folder and its subdirectories for repositories to scan
- `gogitlocalstats -email your@email.com` will generate a CLI stats graph representing the last 6 months of activity for the passed email. You can configure the default in `main.go`, so you can run `gogitlocalstats` without parameters.

Being able to pass an email as param makes it possible to scan repos for collaborators activity as well.

![](https://flaviocopes.com/img/go-git-contributions/output.png)

[![License: CC BY-SA 4.0](https://img.shields.io/badge/License-CC%20BY--SA%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by-sa/4.0/)
