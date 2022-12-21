# Comms


## Needs redis


```console
spin build
spin up --follow-all
```

```
curl --location --request POST 'localhost:3000/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from" : "user1",
    "to" : "user2",
    "message" : "Hi there!" 
}'
```

```
curl --location --request POST 'localhost:3000/encrypt' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message" : "Hi there!" 
}'
```

```
curl --location --request GET 'localhost:3000/translate/text to traslate'
```
