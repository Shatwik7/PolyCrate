#!/bin/bash

# TO START DEVELOPMENT ENV USING DOCKER-COMPOSE


# Path to your docker-compose file
DOCKER_COMPOSE_FILE="./infra/dev/docker-compose.yml"

# Check if docker-compose file exists
if [ ! -f "$DOCKER_COMPOSE_FILE" ]; then
  echo "❌ Docker Compose file not found at: $DOCKER_COMPOSE_FILE"
  exit 1
fi

echo "🚀 Starting development environment using Docker Compose..."

# Run Docker Compose in detached mode
docker-compose -f "$DOCKER_COMPOSE_FILE" up -d

if [ $? -eq 0 ]; then
  echo "✅ Development environment started successfully."
else
  echo "❌ Failed to start development environment."
  exit 1
fi