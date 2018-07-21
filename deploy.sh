#!/bin/bash
git pull origin master

cp bin/deploy bin/deploy_`date +%y%m%d%H%M`
go build -o bin/deploy
