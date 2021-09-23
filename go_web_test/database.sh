#!/usr/bin/env bash

# All confiuration and data of database need to be passed (for tests) to folder databse/database.go into var connection

NAME="postgres_to_go" # The choice of name is free

############################# Docker container of Postgres database instructions #########################################

# Uncomment this line below just if the database is not running as a container service (create database as a container using docker)
# docker run --name $NAME -p 5432:5432 -e POSTGRES_PASSWORD=1234 -d postgres

# After create the database, comment again the line of docker run

# You are connected to database "postgres" as user "postgres" via socket in "/var/run/postgresql" at port "5432"

############################# Instructions to run database after create it ###############################################

# After database has been created, uncomment the line of "docker start"
# Uncomment this line below if database is not running as a container service but already exists 
docker start $NAME

############################# Postgres database instructions #############################################################

# After create and run the database, enter into it and create a table called products and its fields
# id (id is primary key), product_name, product_description, product_price and product_quantity
