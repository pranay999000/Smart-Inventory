#!/bin/bash

MASTER_STATUS=`docker exec mysql_master bash -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "SHOW MASTER STATUS;"'`

MASTER_FILE=`echo $MASTER_STATUS | awk '{print $6}'`
MASTER_POSITION=`echo $MASTER_STATUS | awk '{print $7}'`

echo $MASTER_FILE
echo $MASTER_POSITION

echo "stopping replica 1"
docker exec mysql_replica_1 sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "STOP REPLICA IO_THREAD;"'

echo "stopping replica 2"
docker exec mysql_replica_2 sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "STOP REPLICA IO_THREAD;"'

REPLICA_COMMAND="CHANGE MASTER TO MASTER_HOST='mysql_master', MASTER_USER='root', MASTER_PASSWORD='masterpassword', MASTER_LOG_FILE='$MASTER_FILE', MASTER_LOG_POS=$MASTER_POSITION;"

REPLICA_INITIALIZE_COMMAND='export MYSQL_PWD=masterpassword; mysql -u root -e "'
REPLICA_INITIALIZE_COMMAND+="$REPLICA_COMMAND"
REPLICA_INITIALIZE_COMMAND+='"'

echo "initializing replica 1"
docker exec mysql_replica_1 sh -c "$REPLICA_INITIALIZE_COMMAND"

echo "initializing replica 2"
docker exec mysql_replica_2 sh -c "$REPLICA_INITIALIZE_COMMAND"

echo "stating replica 1"
docker exec mysql_replica_1 sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "START SLAVE;"'

echo "starting replica 2"
docker exec mysql_replica_2 sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "START SLAVE;"'

echo "replica 1 status"
docker exec mysql_replica_1 sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "SHOW SLAVE STATUS\G;"'

echo "replica 2 status"
docker exec mysql_replica_2 sh -c 'export MYSQL_PWD=masterpassword; mysql -u root -e "SHOW SLAVE STATUS\G;"'
