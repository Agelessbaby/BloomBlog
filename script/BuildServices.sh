#!/bin/bash
# Please run this script under project root path
export Root_Path=$(pwd)

cd $Root_Path/cmd/api

go build

cd $Root_Path/cmd/user

./build.sh

cd $Root_Path/cmd/relation

./build.sh

cd $Root_Path/cmd/publish

./build.sh

cd $Root_Path/cmd/feed

./build.sh

cd $Root_Path/cmd/favorite

./build.sh

cd $Root_Path/cmd/comment

./build.sh