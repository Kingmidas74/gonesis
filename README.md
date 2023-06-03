![Github CI/CD](https://img.shields.io/github/workflow/status/kingmidas74/gonesis_engine/Publish%20Web)
![GitHub last commit](https://img.shields.io/github/last-commit/kingmidas74/gonesis_engine)
![Go Report](https://goreportcard.com/badge/github.com/kingmidas74/gonesis_engine)
![Repository Top Language](https://img.shields.io/github/languages/top/kingmidas74/gonesis_engine)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kingmidas74/gonesis_engine)
![Github Repository Size](https://img.shields.io/github/repo-size/kingmidas74/gonesis_engine)
![Github Open Issues](https://img.shields.io/github/issues/kingmidas74/gonesis_engine)
![Lines of code](https://img.shields.io/tokei/lines/github/kingmidas74/gonesis_engine)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub contributors](https://img.shields.io/github/contributors/kingmidas74/gonesis_engine)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

# Gonesis

## Description

This is zero-player game. 

### Core

- Several types of cells
- Several types of terrains

### Optional

- Ability to parametrize many settings

## Solution notes

- :book: standard Go project layout (or not :neutral_face:)
- :cd: github CI/CD + Makefile included
- :card_file_box: WebAssembly support

## HOWTO

- run with `make` (rebuild binaries) and go to [localhost:9091](http://localhost:9091)
- rebuild WebAssembly with `make wasm`
- start with `make server` (in docker mode) and go to [localhost:8989](http://localhost:8989)
- test with `make test`

## A picture is worth a thousand words...

## ... but the live demo is even better!

[Play](https://kingmidas74.github.io/gonesis_engine/)