# Golang CRUD

First contact with Golang, a simple to-do CRUD using [Gin](https://gin-gonic.com/en/) and [GORM](https://gorm.io)

**MADE FOR STUDY PURPOSES ONLY**

## Table of Contents
  - [Disclaimer](#disclaimer)
  - [Project features](#project-features)
  - [Requirements](#requirements)
  - [How to run](#how-to-run)
  - [Commands](#commands)
    - [Install dependencies](#install-dependencies)
    - [Run database](#run-database)
    - [Run API](#run-api)

## Disclaimer
This is not intended to be professional, or follow the bestest guidelines and resources of Golang. This is a `first time contact` with the language

## Project features
- Basic CRUD operations (GET, POST, PUT, PATCH, DELETE - No pagination or filters)
- DDD-ish approach
- Running with postgres with docker container setup
- Testing (soon)

## Requirements
- Go 1.26.4
- Docker (and docker compose)

## How to run
- [Install dependencies](#install-dependencies)
- Create `.env` file
  - copy `.env.example` > `.env` and it should be all good to go
- [Run the database](#run-database)
- [Run the API](#run-api)

## Commands

### Install dependencies
```bash
go mod tidy
```

### Run database
```bash
docker-compose up -d
```

### Run API
```bash
go run .
```
