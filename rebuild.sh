#!/bin/bash

# Script para rebuildar a imagem Docker e reiniciar os containers

echo "🔨 Rebuilding Docker image..."
docker build -t alp-api:latest .

if [ $? -ne 0 ]; then
    echo "❌ Build failed!"
    exit 1
fi

echo "✅ Build successful!"
echo ""
echo "🛑 Stopping old containers..."
docker stop alp-api-1 alp-api-2 alp-api-3 2>/dev/null || true
docker rm alp-api-1 alp-api-2 alp-api-3 2>/dev/null || true

echo "✅ Old containers removed!"
echo ""
echo "🚀 Starting new containers..."
docker run -d --name alp-api-1 -p 8080:8080 alp-api:latest
docker run -d --name alp-api-2 -p 8081:8080 alp-api:latest
docker run -d --name alp-api-3 -p 8082:8080 alp-api:latest

echo ""
echo "✅ All containers started!"
echo ""
echo "📍 API Instances:"
echo "  - api-1: http://localhost:8080"
echo "  - api-2: http://localhost:8081"
echo "  - api-3: http://localhost:8082"
echo ""
docker ps --filter "name=alp-api"
