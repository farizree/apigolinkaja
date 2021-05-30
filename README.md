# apigolinkaja
apigolinkaja

GO & MySQL
{
    Untuk menjalankannya di local jangan lupa untuk merubah function LoggingActivity di main.go menjadi
    var filename string = "../apiloglinkaja/loglinkaja/log" + date + ".log" 
}
{
    Untuk menjalankannya di docker jangan lupa untuk merubah function di main.go menjadi
    var filename string = "src/apiloglinkaja/loglinkaja/log" + date + ".log" 
}
Running di local computer
{
    cd src/apilinkaja
    linux / macos
    1. go build main.go
    2. ./main

    Windows
    1. go build main.go
    2. main.exe
}
Running di docker 
{
    docker build . -t apigolinkaja.latest
    docker run --name apigolinkaja -p 2021:2021 -d apigolinkaja:latest
}

export db to mysql phpmyadmin

# API

ENDPOINT

- Check Saldo

1. http://localhost:2021/api/v1/linkaja/checksaldo/account
{
    "account_number":555001
}

- Transfer

2. http://localhost:2021/api/v1/linkaja/transfer
{
    "from_account_number": 555004,
    "to_account_number": 555003,
    "amount": 1000
}


