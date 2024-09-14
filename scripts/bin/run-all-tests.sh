#!/bin/sh
# Run all unit tests, integration tests plus race conditions checking
# Should be marked by // +build integration

echo ""
if command -v grc &> /dev/null
then
    grc go test -tags integration -race -v ./...
else
    go test -tags integration -race -v ./...
fi


