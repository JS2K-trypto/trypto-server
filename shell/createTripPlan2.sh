curl --location --request POST 'http://152.69.231.140:1323/v01/trip/myplan' \
--header 'Content-Type: application/json' \
--data-raw '{
    "walletAccount": "0x531946cC77Ae69696456a18CA746B265eB347a15",
    "tripTitle": "Discovering the Energy of New York City",
    "tripCountry": "United States",
    "tripDeparture": "2023-05-29",
    "tripArrival": "2023-05-31",
    "dayItems": [
                {
                    "startDate": "2023-05-29",
                    "endDate": "2023-05-29",
                    "title": "Iconic Landmarks and City Vibes" ,
                    "note": "Start your day with a visit to Times Square, where the dazzling lights and bustling atmosphere will leave you in awe. Explore the renowned attractions such as Central Park, Empire State Building, and Statue of Liberty. Don'\''t forget to try a classic New York-style pizza for lunch!"
                },
                 {
                    "startDate": "2023-05-30",
                    "endDate": "2023-05-30",
                    "title": "Cultural Immersion in the Big Apple",
                    "note": "Dive into the cultural scene of New York City by visiting world-class museums like the Metropolitan Museum of Art or the Museum of Modern Art. Take a stroll through the diverse neighborhoods like SoHo and Greenwich Village, and savor delicious international cuisine along the way."
                },
                 {
                    "startDate": "2023-05-31",
                    "endDate": "2023-05-31",
                    "title": "Hidden Gems and Local Experiences",
                    "note": "On your last day, venture off the beaten path and explore the lesser-known neighborhoods like Brooklyn or Williamsburg. Discover local markets, trendy boutiques, and charming cafes. Wrap up your trip with a scenic walk across the Brooklyn Bridge and soak in the breathtaking skyline views."
                }
    ]
}'