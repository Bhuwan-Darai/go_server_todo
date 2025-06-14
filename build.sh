#!/bin/bash
set -e

echo "📦 Installing dependencies..."
go mod download

echo "🔧 Installing Prisma client..."
go install github.com/steebchen/prisma-client-go@latest

echo "🚀 Generating Prisma client..."
go run github.com/steebchen/prisma-client-go generate --binary-targets=native,linux-musl

echo "🔨 Building Go binary..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

echo "✅ Build completed successfully!" 