#!/bin/bash

shouldEcho=false

while IFS='' read -r line || [[ -n "$line" ]]; do
    firstField=$(echo "$line" | cut -d " " -f 1)

    if [ "$firstField" == "type" ] || [ "$firstField" == "func" ] || [ "$firstField" == "const" ] ; then
	shouldEcho=true
    fi

    if $shouldEcho ; then
	echo "$line"
    fi
done < "$1"
