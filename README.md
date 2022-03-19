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

### Generate private key and certificate

```sh
# Generate private key and certificate
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout server.key -out server.crt -subj /CN=localhost

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

### Fix `Your connection is not private` in Google Chrome

Google Chrome validate the self signed certificate error: `Your connection is not private` ... `NET::ERR_CERT_INVALID` or `NET::ERR_CERT_AUTHORITY_INVALID`.
Browse [chrome://flags/#allow-insecure-localhost](chrome://flags/#allow-insecure-localhost), and Enable the option `Allow invalid certificates for resources loaded from localhost`.
