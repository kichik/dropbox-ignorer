# dropbox-ignorer

Quickly set the [Dropbox ignore flag](https://help.dropbox.com/files-folders/restore-delete/ignored-files) on all
folders matching a given pattern under `$HOME/Dropbox`. Can be used to easily remove all `node_modules` folders from
your Dropbox account.

* Works on Windows, Linux and Mac OS X
* CLI only
* Very fast
* 0% fluff

```
C:\> dropbox-ignorer node_modules
2022/02/21 16:51:24 Scanning C:\Users\kichik\Dropbox for node_modules...
2022/02/21 16:51:24 Excluded C:\Users\kichik\Dropbox\some-node-project\node_modules
2022/02/21 16:51:24 Excluded C:\Users\kichik\Dropbox\another-node-project\node_modules
2022/02/21 16:51:24 Excluded C:\Users\kichik\Dropbox\sub\node\node_modules
```

[![CI](https://github.com/kichik/dropbox-ignorer/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/kichik/dropbox-ignorer/actions/workflows/goreleaser.yml) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/kichik/dropbox-ignorer.svg)](https://github.com/kichik/dropbox-ignorer)
[![GoReportCard](https://goreportcard.com/badge/github.com/kichik/dropbox-ignorer)](https://goreportcard.com/report/github.com/kichik/dropbox-ignorer) [![GitHub license](https://img.shields.io/github/license/kichik/dropbox-ignorer.svg)](https://github.com/kichik/dropbox-ignorer/blob/main/LICENSE) [![GitHub release](https://img.shields.io/github/release/kichik/dropbox-ignorer.svg)](https://GitHub.com/kichik/dropbox-ignorer/releases/)