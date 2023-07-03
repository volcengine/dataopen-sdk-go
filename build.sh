#!/bin/sh

rm -rf release/dataopen-sdk-go*
mkdir release/dataopen-sdk-go

cp -rf LICENSE release/dataopen-sdk-go/
cp -rf README.md release/dataopen-sdk-go/
cp -rf client.go release/dataopen-sdk-go/
cp -rf go.mod release/dataopen-sdk-go/

cd release
zip -r dataopen-sdk-go.zip dataopen-sdk-go/*

rm -rf dataopen-sdk-go

cd ../