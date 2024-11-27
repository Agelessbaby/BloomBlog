#!/bin/bash
# Please run this script under project root path
export Root_Path=$(pwd)

# 验证环境变量
echo $Root_Path

cd $Root_Path/cmd/user

kitex -module github.com/Agelessbaby/BloomBlog -type protobuf -I ../../idl -service user  ../../idl/user.proto

cd $Root_Path/cmd/relation

kitex -module github.com/Agelessbaby/BloomBlog -type protobuf -I ../../idl -service relation -protobuf Muser.proto=github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user ../../idl/relation.proto

cd $Root_Path/cmd/feed

kitex -module github.com/Agelessbaby/BloomBlog -type protobuf -I ../../idl -service feed -protobuf Muser.proto=github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user  ../../idl/feed.proto

cd $Root_Path/cmd/publish

kitex -module github.com/Agelessbaby/BloomBlog -type protobuf -I ../../idl -service publish -protobuf Mfeed.proto=github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed -protobuf Muser.proto=github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user ../../idl/publish.proto

cd $Root_Path/cmd/favorite

kitex -module github.com/Agelessbaby/BloomBlog -type protobuf -I ../../idl -service favorite -protobuf Mfeed.proto=github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed -protobuf Muser.proto=github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user ../../idl/favorite.proto

cd $Root_Path/cmd/comment

kitex -module github.com/Agelessbaby/BloomBlog -type protobuf -I ../../idl -service comment -protobuf Muser.proto=github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user ../../idl/comment.proto