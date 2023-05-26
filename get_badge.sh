#!/bin/bash

curl --location --request GET 'https://localhost:1323/v01/badge/user' \
-key 127.0.0.1-key.pem \
-cert 127.0.0.1.pem \
--header 'Content-Type: application/json' \
--data-raw '{
      "walletAccount" :"0x61Facc99384d27fc51C27Bf7C08f4EDf0E53eF48"
}'
