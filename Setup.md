
Install spin

curl -fsSL https://developer.fermyon.com/downloads/install.sh | bash
sudo mv ./spin /usr/local/bin/spin

Install js and typescript 

spin plugin install js2wasm
spin templates install --git https://github.com/fermyon/spin-js-sdk

Instal tiny go (latest version doesnt work 0.26.0)
    curl -fsSL https://github.com/tinygo-org/tinygo/releases/download/v0.25.0/tinygo0.25.0.darwin-amd64.tar.gz -o tinygo.tar.gz 
    tar xvzf tinygo.tar.gz 
    export PATH=<extract location>/tinygo/bin:$PATH

redis
    redis-server

send component:
    spin add
    http-ts
    /send/.. path

    NOTES: spin build fails first time but works if run again!

curl --location --request POST 'localhost:3000/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from" : "user1",
    "to" : "user2",
    "message" : "Hi there!" 
}'


encrypt compenent:

    spin add 
    http-rust
    /encrypt/...

    cargo add serde_json    
    cargo add serde --features derive
    cargo install encrypt

curl --location --request POST 'localhost:3000/encrypt' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message" : "Hi there!" 
}'


Translate component:

    spin add
    http-go
    /translate/...
    add to spin.toml    allowed_http_hosts = ["https://api.funtranslations.com"]

curl --location --request GET 'localhost:3000/translate'