#!/bin/sh
echo "Starting.."
port_number=$port_number
shared_folder=$shared_folder
exec app -p $port_number -f $shared_folder