#!/bin/bash

# Script para rebuildar e reiniciar com docker-compose

echo "🔨 Rebuilding with docker-compose..."
docker-compose up -d --build

if [ $? -ne 0 ]; then
    echo "❌ Build with docker-compose failed!"
    exit 1
fi

echo ""
echo "✅ All services started!"
echo ""
echo "📍 API Instances:"
echo "  - api-1: http://localhost:8080"
echo "  - api-2: http://localhost:8081"
echo "  - api-3: http://localhost:8082"
echo ""
docker ps --filter "name=alp-api"
