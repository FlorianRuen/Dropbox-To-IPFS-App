#!/bin/sh

cd ../backend/database/liquibase
bash liquibase --changeLogFile=master.xml --log-level=DEBUG update
