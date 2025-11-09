#!/bin/bash

json_data=$(curl -s https://acad.learn2earn.ng/assets/superhero/all.json)

#To find Id to call

name=$(echo "$json_data" | jq -r '.[] | select(.id == 70) | .name')

output this result
echo "\"$name\"" 

