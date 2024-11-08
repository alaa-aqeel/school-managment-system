#!/bin/bash

case "$1" in
    "up")
        sudo docker compose -f docker-compose.dev.yml up --build
        ;;
    "test")
        # Run migrate down to undo the last migration
        sudo docker compose -f docker-compose.test.yml up --build
        ;;
    *)
        echo "Usage: $0 {start|test}"
        exit 1
        ;;
esac