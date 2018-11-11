#!/usr/bin/env bash

docker image build -t yaser/phonebook-go .
docker image push yaser/phonebook-go