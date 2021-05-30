# apigolinkaja
apigolinkaja

GO & MySQL

Untuk menjalankannya
cd src/apilinkaja
linux / macos
1. go build main.go
2. ./main

Windows
1. go build main.go
2. main.exe

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


