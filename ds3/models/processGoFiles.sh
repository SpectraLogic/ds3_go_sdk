#!/bin/bash

outFile="consolidated.txt"
echo "package models" > "$outFile"
echo >> "$outFile"

echo "import (" >> "$outFile"
echo "\"net/http\"" >> "$outFile"
echo "\"io\"" >> "$outFile"
echo "\"encoding/xml\"" >> "$outFile"
echo "\"fmt\"" >> "$outFile"
echo "\"bytes\"" >> "$outFile"
echo "\"errors\"" >> "$outFile"
echo "\"log\"" >> "$outFile"
echo "\"io/ioutil\"" >> "$outFile"
echo "\"strings\"" >> "$outFile"
echo "\"strconv\"" >> "$outFile"
echo ")" >> "$outFile"

for goFile in $(ls *.go) ; do
    ./getGoTypeInfo.sh "$goFile" >> "$outFile"
    mv "$goFile" "$goFile".orig
done

mv "$outFile" consolidated.go
