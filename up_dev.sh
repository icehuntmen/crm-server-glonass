#!/bin/bash

swag init -g api/api.go

go build -o ./tmp/crm-server ./cmd/main.go