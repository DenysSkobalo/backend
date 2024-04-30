#!/bin/bash

cd ../cmd/authservice

echo "Building the auth service..."
go build -o ../../bin/authservice

cd ../..

echo "Starting the auth service ..."
./bin/authservice