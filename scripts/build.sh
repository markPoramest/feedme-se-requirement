#!/bin/bash

# Build Script
# Builds the McDonald's Order Management System CLI application

echo "Building McDonald's Order Management System..."
echo ""

# Build the application
cd ..
go build -o bin/feedme main.go

if [ $? -eq 0 ]; then
  echo "Build successful! Binary created at bin/feedme"
  echo "To run the application, execute: ./bin/feedme"
else
  echo "Build failed!"
  exit 1
fi
