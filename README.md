# Easy CA CLI

[![GitHub](https://img.shields.io/github/license/c3b2a7/easy-ca-cli)](https://github.com/c3b2a7/easy-ca-cli/blob/master/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/c3b2a7/easy-ca-cli/build.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/c3b2a7/easy-ca-cli)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/c3b2a7/easy-ca-cli)

> A fast, simple and easy-to-use certificate generator.

## Overview

key algorithm support

- RSA
    - with key size: `1024`、`2048`
- ECDSA
    - ecdsa curve: `P224`、`P256`、`P384`、`P521`
- ED25591

generate certificate support

- support certificate Authority、middle certificate authority、general TLS certificate
- support certificate chain and private key in pkcs8 format
- support specified output path

certificate info custom support

- Subject
    - C、O、OU、CN、SERIALNUMBER、L、ST、POSTALCODE
- Subject Alternative Name
    - dns names
    - ip addresses
- valid from and valid for

## Usage

Provider `gen ca` and `gen tls` commands for certificate generation, run `help` for more information about a command and
its flags.

```bash
easy-ca-cli help
```

## Examples

1、Generate a certificate authority with specified subject info、valid time and output path with ecdsa algorithm.

```bash
easy-ca-cli gen ca --ecdsa --ecdsa-curve P512 \
 --subject "C=CN,O=Easy CA,OU=IT Dept.,CN=Easy CA Root" \
 --start-date "Jan 1 15:00:00 2022" --days 3650 \
 --out-key ca_key.pem --out-cert ca_cert.pem
```

2、Generate a middle certificate authority using the certificate authority generated above

```bash
easy-ca-cli gen ca --ecdsa --ecdsa-curve P384 \
 --subject "C=CN,O=Easy CA,OU=IT Dept.,CN=Easy CA Authority R1" \
 --start-date "Jan 1 15:05:00 2022" --days 1800 \
 --issuer-key ca_key.pem --issuer-cert ca_cert.pem \
 --out-key mca_key.pem --out-cert mca_cert.chain.pem
```

3、Generate a TLS certificate using the certificate authority generated above

```bash
easy-ca-cli gen tls --rsa --rsa-keysize 2048 \
  --subject "C=CN,O=Easy CA,OU=IT Dept.,CN=easy-ca.com" \
  --host "easy-ca.com,www.easy-ca.com,cli.easy-ca.com" \
  --start-date "Jan 1 15:10:00 2022" --days 365 \
  --issuer-key mca_key.pem --issuer-cert mca_cert.chain.pem \
  --out-key easyca_key.pem --out-cert easyca_cert.chain.pem
```

## Dependencies

- [easy-ca](https://github.com/c3b2a7/easy-ca)
- [cobar](https://github.com/spf13/cobra)

## LICENSE

Easy CA CLI is released under the MIT license. See [LICENSE](https://github.com/c3b2a7/easy-ca-cli/blob/main/LICENSE)