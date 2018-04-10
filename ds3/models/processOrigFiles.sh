#!/bin/bash

mv consolidated.go consolidated.txt

for origFile in $(ls *.orig) ; do
    origFileBase=$(echo "$origFile" | cut -d "." -f 1)
    mv "$origFile" "$origFileBase.go"
done
