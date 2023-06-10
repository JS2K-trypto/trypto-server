#!/bin/bash

curl --location --request POST 'http://152.69.231.140:1323/v01/acc/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "walletAccount":"0x531946cC77Ae69696456a18CA746B265eB347a15",
    "nickName":"Noname"
}'