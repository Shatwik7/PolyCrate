#!/bin/bash

go test ./services/user_service/... -coverprofile coverage.out
go tool cover -html coverage.out -o coverage.html
