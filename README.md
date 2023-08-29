[![Release Version](https://img.shields.io/github/v/release/ankitcharolia/gitlab-artifacts-cleaner?label=gitlab-artifacts-cleaner)](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases/latest)
![Build CI](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/actions/workflows/build-publish.yaml/badge.svg)
![CodeQL](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/actions/workflows/codeql-analysis.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ankitcharolia/gitlab-artifacts-cleaner)](https://goreportcard.com/report/github.com/ankitcharolia/gitlab-artifacts-cleaner)
[![License](https://img.shields.io/badge/License-MIT%20-blue.svg)](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/downloads/ankitcharolia/gitlab-artifacts-cleaner/total.svg)]()

# Gitlab Artifacts Cleaner
Gitlab Artifacts Cleaner is a command-line tool to cleanup Gitlab stored Artifacts

## Installation
**Releases**: https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases

**LINUX**
```shell
wget -O - https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases/latest/download/gitlab-artifacts-cleaner-linux-amd64.tar.gz | tar -xz -C /usr/local/bin
```

**MAC**
```shell
wget -O - https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases/latest/download/gitlab-artifacts-cleaner-darwin-amd64.tar.gz | tar -xz -C /usr/local/bin
```

## Usage
```shell
$ gitlab-artifacts-cleaner --help
Usage: gitlab-artifacts-cleaner [flags] [<options>]
Flags:
  --help          goenv help command
  --install       Install a specific version of GOLANG
  --list          List all installed GOLANG versions
  --list-remote   List all remote versions of GOLANG
  --uninstall     Uninstall a specific version of GOLANG
  --use           Use a specific version of GOLANG
```

## Support
Feel free to create an Issue/Pull Request if you find any bug with `goenv`