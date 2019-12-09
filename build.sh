#!/bin/bash


if [[ $1 = "run" ]]
then
	echo "running..."

	./dist/gopy

	echo "done..."
elif [[ $1 = "build" ]]
then
	echo "building..."
	
	rm -r ./dist
	cd src
	go build -o ../dist/gopy

	echo "done..."
else
	./build.sh build
	./build.sh run
fi