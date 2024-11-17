#!/bin/bash
# Please run this script under project root path
export Root_Path=$(pwd)


cd $Root_Path/cmd/user

./output/bin/user&

cd $Root_Path/cmd/relation

./output/bin/relation&

cd $Root_Path/cmd/publish

./output/bin/publish&