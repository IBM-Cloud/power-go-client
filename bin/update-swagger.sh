#!/bin/bash

# get swagger directory
printf "Enter swagger file directory, i.e., ~/User/Projects/service-broker/swagger\n   "
read swaggerdir
if [[ "$swaggerdir" != *"swagger.yaml" ]]; then
    if [[ "$swaggerdir" != *"/" ]]; then
        swaggerdir+="/"
    fi
    swaggerdir+="swagger.yaml"
fi

# remove old power directory
if [ -d "power" ]; then
    printf "Removing old Power-Go-Client swagger files\n"
    rm -rf power
fi
mkdir power

# generate swagger files
printf "Generating new Power-Go-Client swagger file\n"
generatecmd="swagger generate client -f $swaggerdir -t power"
eval $generatecmd

# remove unused files
printf "\nRemoving unused Power-Go-Client swagger files\n"
rm power/models/user_access.go
rm power/models/principal.go