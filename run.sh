#!/bin/zsh

set -xe

rm -rf ./app

go build -o app ./

./app