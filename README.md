# Easy CA CLI

[![GitHub](https://img.shields.io/github/license/c3b2a7/easy-ca-cli)](https://github.com/c3b2a7/easy-ca-cli/blob/master/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/c3b2a7/easy-ca-cli/build.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/c3b2a7/easy-ca-cli)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/c3b2a7/easy-ca-cli)

- [What is easy-ca-cli?](#what-is-easy-ca-cli)
- [Features](#features)
  * [Supported Key Algorithms](#supported-key-algorithms)
  * [Certificate Types and Generation](#certificate-types-and-generation)
  * [Customizable Certificate Information](#customizable-certificate-information)
- [Installation](#installation)
  * [Pre-built Binaries](#pre-built-binaries)
  * [Building from Source](#building-from-source)
  * [Verify installation](#verify-installation)
- [Usage](#usage)
  * [Creating a Certificate Authority (CA)](#creating-a-certificate-authority-ca)
  * [Generating an Intermediate CA](#generating-an-intermediate-ca)
  * [Generating a TLS Certificate](#generating-a-tls-certificate)
  * [Others](#others)
- [Troubleshooting](#troubleshooting)
  * [Common Issues](#common-issues)
  * [Getting Help](#getting-help)
- [Dependencies](#dependencies)
- [LICENSE](#license)

## What is easy-ca-cli?

The easy-ca-cli is a fast, simple certificate generation utility built in Go.

It serves as a command-line interface to the core [easy-ca](https://github.com/c3b2a7/easy-ca) library,
providing an accessible way to generate various certificate types with customizable parameters.

## Features

The CLI offers a comprehensive set of features for certificate generation:

### Supported Key Algorithms

| Algorithm | Flag        | Options                                         |
|-----------|-------------|-------------------------------------------------|
| RSA       | `--rsa`     | Key size: `--rsa-keysize` (1024, 2048)          |
| ECDSA     | `--ecdsa`   | Curve: `--ecdsa-curve` (P224, P256, P384, P521) |
| ED25519   | `--ed25519` | No additional options                           |

### Certificate Types and Generation

- Certificate Authority (Root CA)
- Intermediate Certificate Authority (Intermediate CA)
- TLS/Server Certificates
- Support for certificate chains
- PKCS8 format for private keys
- Customizable output paths

### Customizable Certificate Information

- Subject fields: C, O, OU, CN, Serial Number, L, ST, Postal Code
- Subject Alternative Name (SAN)
    - DNS names
    - IP addresses
- Validity period configuration

## Installation

### Pre-built Binaries

You can also download and extract the latest release
from [GitHub Releases Page](https://github.com/c3b2a7/easy-ca-cli/releases)

### Building From Source

```bash
go install github.com/c3b2a7/easy-ca-cli@latest
```

### Verify Installation

```bash
easy-ca-cli --version
```

## Usage

The CLI provider `gen ca` and `gen tls` commands for certificate generation,
here are some basic examples of how to use easy-ca-cli:

### Creating a Certificate Authority (CA)

The following command creates a self-signed root CA certificate using ECDSA with the P384 curve:

```bash
easy-ca-cli gen ca --ecdsa --ecdsa-curve P384 \
  --subject "/C=CN/O=Easy CA/OU=IT Dept./CN=Easy CA Root" \
  --start-date "2023-01-01 12:00:00" --days 3650 \
  --out-key ca_key.pem --out-cert ca_cert.pem
```

This command:

- Uses the ECDSA algorithm with P384 curve
- Sets the certificate subject information
- Makes the certificate valid from January 1, 2023 for 10 years (3650 days)
- Outputs the private key to ca_key.pem and the certificate to ca_cert.pem

### Generating an Intermediate CA

To create an intermediate CA signed by your root CA:

```bash
easy-ca-cli gen ca --ecdsa --ecdsa-curve P384 \
  --subject "/C=CN/O=Easy CA/OU=IT Dept./CN=Easy CA Intermediate" \
  --start-date "2023-01-01 13:00:00" --days 1825 \
  --issuer-key ca_key.pem --issuer-cert ca_cert.pem \
  --out-key intermediate_key.pem --out-cert intermediate_cert.chain.pem
```

This command:

- Creates an intermediate CA using the same algorithm
- References the root CA's private key and certificate for signing
- Outputs a certificate chain in `intermediate_cert.chain.pem` that includes both the intermediate CA and root CA
  certificates

### Generating a TLS Certificate

To create a TLS certificate signed by your intermediate CA:

```bash
easy-ca-cli gen tls --rsa --rsa-keysize 2048 \
  --subject "/C=CN/O=Example Inc/OU=Web Services/CN=example.com" \
  --host "example.com,www.example.com,192.168.1.100" \
  --start-date "2023-01-01 14:00:00" --days 365 \
  --issuer-key intermediate_key.pem --issuer-cert intermediate_cert.chain.pem \
  --out-key server_key.pem --out-cert server_cert.chain.pem
```

This command:

- Uses the RSA algorithm with 2048-bit key size
- Sets the certificate subject and Subject Alternative Names (SANs) for multiple domains and an IP address
- Makes the certificate valid for 1 year
- References the intermediate CA for signing
- Outputs a certificate chain that includes the server certificate, intermediate CA, and root CA

### Others

run `help` for more information about a command and
its flags.

```bash
easy-ca-cli help
```

## Troubleshooting

### Common Issues

| Issue                     | Solution                                                                     |
|---------------------------|------------------------------------------------------------------------------|
| "Command not found" error | Ensure the binary is in a directory listed in your PATH environment variable |
| Permission denied         | Make sure the binary is executable (`chmod +x easy-ca-cli`)                  |
| Incompatible architecture | Download the binary that matches your system architecture (amd64/arm64)      |

### Getting Help

If you encounter issues not covered in this guide, you can:

- Check the GitHub repository for open issues
- Open a new issue if your problem hasn't been reported

## Dependencies

- [easy-ca](https://github.com/c3b2a7/easy-ca): The core library that provides the certificate generation functionality
- [cobar](https://github.com/spf13/cobra): A powerful framework for creating modern CLI applications with nested
  commands

## LICENSE

Easy CA CLI is released under the MIT license. See [LICENSE](https://github.com/c3b2a7/easy-ca-cli/blob/main/LICENSE)
