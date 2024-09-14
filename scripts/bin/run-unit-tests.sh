#!/bin/sh

echo ""
if command -v grc &> /dev/null
then
    grc go test -v ./...
else
    go test -v ./...
fi

