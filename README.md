[![Release Version](https://img.shields.io/github/v/release/ankitcharolia/gitlab-artifacts-cleaner?label=gitlab-artifacts-cleaner)](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases/latest)
![Build CI](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/actions/workflows/build-publish.yaml/badge.svg)
![CodeQL](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/actions/workflows/codeql-analysis.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ankitcharolia/gitlab-artifacts-cleaner)](https://goreportcard.com/report/github.com/ankitcharolia/gitlab-artifacts-cleaner)
[![License](https://img.shields.io/badge/License-MIT%20-blue.svg)](https://github.com/ankitcharolia/gitlab-artifacts-cleaner/blob/master/LICENSE)
[![Releases](https://img.shields.io/github/downloads/ankitcharolia/gitlab-artifacts-cleaner/total.svg)]()

# Gitlab Artifacts Cleaner
Gitlab Artifacts Cleaner is a command-line tool to cleanup Gitlab stored Artifacts

**NOTE:** If you need to wipe out a Gitlab project's artifacts, this is the tool for you. **It will delete all the artifacts of the project, but It will NOT delete the project.**

## Installation
**Releases**: https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases

**LINUX**
```shell
sudo wget -O - https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases/latest/download/gitlab-artifacts-cleaner-linux-amd64.tar.gz | sudo tar -xz -C /usr/local/bin
```

**MAC**
```shell
wget -O - https://github.com/ankitcharolia/gitlab-artifacts-cleaner/releases/latest/download/gitlab-artifacts-cleaner-darwin-amd64.tar.gz | sudo tar -xz -C /usr/local/bin
```

## Usage
```shell
$ gitlab-artifacts-cleaner --help
Usage: gitlab-artifacts-cleaner [flags] [<options>]
Flags:
  --help          gitlab-artifacts-cleaner help command
  --server        GitLab Server URL. Default: https://gitlab.com
  --token         Your Personal Access Token.
  --project-id    Project ID or Project Path.
```

```bash
gitlab-artifacts-cleaner --server https://gitlab.com --token <token> --project_id <project_id>
```
**NOTE:** You must have at least the Maintainer role for the project.

## Support
Feel free to create an Issue/Pull Request if you find any bug with `gitlab-artifacts-cleaner`