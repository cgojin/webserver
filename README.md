# webserver

Simple static file web server with go.

## Building

```sh
git clone https://github.com/cgojin/webserver
cd webserver
go build
```

## Run webserver with ``http``

```sh
# Simple running
./webserver

# Specify port and directory
./webserver -addr 8080 -dir ./public
```

## Run webserver with ``https``

### Generate certificate

```sh
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key

# Generation of self-signed(x509) public key (PEM-encodings `.pem`|`.crt`) based on the private (`.key`)
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

### Run webserver with certificate

```sh
./webserver -cert server.crt -key server.key
```

### Test webserver

```sh
curl https://localhost:8080
    curl: (60) SSL certificate problem: self signed certificate

# need use the -k (or –insecure) option
curl -k https://localhost:8080
    ok
```

Note: In Google Chrome if there is an issue validating the certificate the error will show as ```“your connection is not private”```
You can browse [chrome://flags/#allow-insecure-localhost](chrome://flags/#allow-insecure-localhost), and Enable the option "Allow invalid certificates for resources loaded from localhost."
