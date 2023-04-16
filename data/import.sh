#!/bin/bash
#mongoimport --uri "mongodb://congestion:congestion@localhost:27017/volvo" --collection cities --file "$1"
mongoimport --host mongodb --username congestion --password congestion --db volvo --collection cities --jsonArray --file /mongo_seed/import.json
