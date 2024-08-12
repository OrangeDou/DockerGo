#!/bin/bash

rm gocker

git pull

go build .

mount -t proc proc /proc