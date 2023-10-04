#!/bin/bash

host="$1"
user="$2"
database="$3"
port="$4"
password="$5"

PGPASSWORD="$password" createdb -h "$host" -U "$user" -p "$port" "$database"

