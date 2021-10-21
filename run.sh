#!/bin/bash

go build -o udemy-api cmd/web/*.go
./udemy-api -dbuser=root -dbpass=root -dbport=3306 -production=false -dbname=todoapp