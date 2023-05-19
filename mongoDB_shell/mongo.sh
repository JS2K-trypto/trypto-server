#!/bin/bash

#/home/robertseo/mongo/data/db


docker run --name mongodb -v /home/robertseo/mongo/data/db:/data/db -d -p 27017:27017 mongo
