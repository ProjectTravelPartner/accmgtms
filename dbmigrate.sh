#!/bin/bash

# Uses: https://github.com/golang-migrate/migrate/tree/master/cli

action=$1

if [ "action" == "" ];then
    action="up"
fi

migrate -source file://dbmigrations -database 'mysql://root:root@tcp(localhost:3306)/accmgtms' $action
