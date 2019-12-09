#!/bin/bash


if [[ $1 = "run" ]]
then
	echo "running..."

	./dist/gopy

	echo "--done--"
else
	echo "building..."
	
	rm -r ./dist
	cd src
	go build -o ../dist/gopy

	echo "--done--"
fi
