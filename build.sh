#!/bin/sh
docker-compose -f docker-compose.json build
if [ $? -ne 0 ]; then
    exit -1
fi