curl --location --request POST 'http://152.69.231.140:1323/v01/trip/myplan' \
--header 'Content-Type: application/json' \
--data-raw '{
    "walletAccount": "0x531946cC77Ae69696456a18CA746B265eB347a15",
    "tripTitle": "Spontaneous trip to Gyeongbokgung Palace, South Korea",
    "tripCountry": "Korea",
    "tripDeparture": "2023-05-29",
    "tripArrival": "2023-05-31",
    "dayItems": [
                {
                    "startDate": "2023-05-29",
                    "endDate": "2023-05-29",
                    "title": "Geunjeongjeon Hot Place" ,
                    "note": "If you get a badge here, can you win an event? "
                },
                 {
                    "startDate": "2023-05-30",
                    "endDate": "2023-05-30",
                    "title": "Traveling to the Gyeonghoeru",
                    "note": "Life shot in hanbok at Gyeonghoeru?"
                },
                 {
                    "startDate": "2023-05-31",
                    "endDate": "2023-05-31",
                    "title": "National Palace Museum",
                    "note": "Let'\''s experience Korean history and culture~!"
                }
    ]
}'