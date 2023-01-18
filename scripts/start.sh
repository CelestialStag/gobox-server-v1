#!/bin/bash

$(cd src && go build -o ./dist/gopy) && echo 'fin: build'
./dist/gopy