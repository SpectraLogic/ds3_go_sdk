#!/bin/bash

for goFile in $(ls *.go) ; do
    mv "$goFile" "$goFile".orig
done

mv consolidated.txt  consolidated.go
