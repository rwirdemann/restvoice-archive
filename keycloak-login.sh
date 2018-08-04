#!/bin/bash

RESULT=`curl --data "grant_type=password&client_id=restvoice&username=user&password=start1234" \
http://localhost:8180/auth/realms/master/protocol/openid-connect/token`

TOKEN=`echo $RESULT | sed 's/.*access_token":"//g' | sed 's/".*//g'`
echo ${TOKEN}