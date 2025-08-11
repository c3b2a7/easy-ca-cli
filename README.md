# Easy CA CLI

[![GitHub](https://img.shields.io/github/license/c3b2a7/easy-ca-cli)](https://github.com/c3b2a7/easy-ca-cli/blob/master/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/c3b2a7/easy-ca-cli/ci.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/c3b2a7/easy-ca-cli)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/c3b2a7/easy-ca-cli)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/c3b2a7/easy-ca-cli)

## Table Of Contents

- [What is easy-ca-cli?](#what-is-easy-ca-cli)
- [Features](#features)
  - [Supported Key Algorithms](#supported-key-algorithms)
  - [Certificate Types and Generation](#certificate-types-and-generation)
  - [Customizable Certificate Information](#customizable-certificate-information)
- [Installation](#installation)
  - [Recommended](#recommended)
  - [Manual](#manual)
  - [Building From Source](#building-from-source)
- [Docker image](#docker-image)
- [Usage](#usage)
  - [Creating a Certificate Authority (CA)](#creating-a-certificate-authority-ca)
  - [Generating an Intermediate CA](#generating-an-intermediate-ca)
  - [Generating a TLS Certificate](#generating-a-tls-certificate)
- [Troubleshooting](#troubleshooting)
  - [Common Issues](#common-issues)
  - [Getting Help](#getting-help)
- [Related Projects](#related-projects)
- [CHANGELOG](#changelog)
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
| RSA       | `--rsa`     | Key size: `--rsa-keysize` (2048, 3072, 4096)    |
| ECDSA     | `--ecdsa`   | Curve: `--ecdsa-curve` (P224, P256, P384, P521) |
| ED25519   | `--ed25519` | No additional options                           |

### Certificate Types and Generation

- Certificate Authority (Root CA)
- Intermediate Certificate Authority (Intermediate CA)
- TLS/Server Certificates
- Support for certificate chains
- PKCS8 format for private keys
- PKCS12 format for the key pair
- Customizable output paths

### Customizable Certificate Information

- Subject fields: C, O, OU, CN, Serial Number, L, ST, Postal Code
- Subject Alternative Name (SAN)
  - DNS names
  - IP addresses
- Validity period configuration

## Installation

easy-ca-cli binaries are provided for Linux, macOS and Windows.

### Recommended

```sh
curl -sSfL https://get.lolico.me/easy-ca-cli | sh -s -- -b /usr/local/bin
```

Install script options:

- `-b`: Specify a custom installation directory (defaults to `./bin`)
- `-d`: More verbose logging levels (`-d` for debug, `-dd` for trace)
- `-v`: Verify the signature of the downloaded artifact before installation (requires [cosign](https://github.com/sigstore/cosign) to be installed)

### Manual

You can download and extract the latest release
from [GitHub Releases Page](https://github.com/c3b2a7/easy-ca-cli/releases)

<details>
<summary><b>Guided installation steps</b></summary>

Here we use macOS (Darwin) on the arm64 architecture as an example for installation.

1. Download these files for your version from the [GitHub Releases Page](https://github.com/c3b2a7/easy-ca-cli/releases)

    - tarball: `easy-ca-cli_$version_darwin_arm64.tar.gz`
    - checksums (optional): `easy-ca-cli_$version_sha256_checksums.txt`
    - signature (optional): `easy-ca-cli_$version_sha256_checksums.txt.sig`
    - certificate (optional): `easy-ca-cli_$version_sha256_checksums.txt.pem`

2. **(Optional) Verifying the signature of your downloaded checksums file**

    If you chose to manually download the tarball file in the above steps, you can use the following steps to verify the signatures and checksums by using the [cosign](https://github.com/sigstore/cosign) and sha256sum tool.

    ```bash
    # Your downloaded version (without prefix `v`)
    # Example:
    version=1.4.0

    # 1. Verify the signature of checksum file
    cosign verify-blob easy-ca-cli_${version}_sha256_checksums.txt \
      --certificate easy-ca-cli_${version}_sha256_checksums.txt.pem \
      --signature easy-ca-cli_${version}_sha256_checksums.txt.sig \
      --certificate-identity "https://github.com/c3b2a7/easy-ca-cli/.github/workflows/cd.yml@refs/tags/v${version}" \
      --certificate-oidc-issuer "https://token.actions.githubusercontent.com"

    # 2. Verify the tarball’s checksums
    sha256sum --ignore-missing --check easy-ca-cli_${version}_sha256_checksums.txt
    ```

3. Unpack tarball
4. Confirm the version with the following command

    ```bash
    easy-ca-cli --version
    ```

</details>

### Building From Source

```bash
go install github.com/c3b2a7/easy-ca-cli@latest
```

## Docker image

easy-ca-cli are also available as a Docker image. For example:

```bash
docker run --rm -i ghcr.io/c3b2a7/easy-ca-cli:latest
```

> [!TIP]
> You can mount volumes using the `-v` option, e.g. `-v <host_path>:<container_path>`, so it's easier to access the generated certificates and keys.

```bash
docker run -v ./data:/data --rm -i ghcr.io/c3b2a7/easy-ca-cli:v1.4.0-amd64 gen ca --rsa \
  --subject "/C=CN/O=Easy CA/OU=IT Dept./CN=Easy CA Root" \
  --out-key /data/ca_key.pem --out-cert /data/ca_cert.pem
```

Multiple versions are available on [ghcr.io](https://github.com/c3b2a7/easy-ca-cli/pkgs/container/easy-ca-cli).

## Usage

The CLI provider `gen ca` and `gen tls` commands for certificate generation, you can run `easy-ca-cli help [command]` for more information about a command and its flags.

Here are some basic examples of how to use easy-ca-cli.

### Creating a Certificate Authority (CA)

The following command creates a self-signed root CA certificate using ECDSA with the P384 curve:

```bash
easy-ca-cli gen ca --ecdsa --ecdsa-curve P384 \
  --subject "/C=CN/O=Easy CA/OU=IT Dept./CN=Easy CA Root" \
  --start-date "2025-01-01 12:00:00" --days 3650 \
  --out-key ca_key.pem --out-cert ca_cert.pem
```

This command:

- Uses the `ECDSA` algorithm with `P384` curve
- Sets the certificate subject information
- Makes the certificate valid from January 1, 2025 for 10 years (3650 days)
- Outputs the private key to `ca_key.pem` and the certificate to `ca_cert.pem`

### Generating an Intermediate CA

To create an intermediate CA signed by your root CA:

```bash
easy-ca-cli gen ca --ed25519 \
  --subject "/C=CN/O=Easy CA/OU=IT Dept./CN=Easy CA Intermediate" \
  --start-date "2025-01-01 13:00:00" --days 1825 \
  --issuer-key ca_key.pem --issuer-cert ca_cert.pem \
  --out-key intermediate_key.pem --out-cert intermediate_cert.chain.pem
```

This command:

- Creates an intermediate CA using the `ed25519` algorithm
- References the root CA's private key and certificate for signing
- Outputs a certificate chain in `intermediate_cert.chain.pem` that includes both the intermediate CA and root CA
  certificates

### Generating a TLS Certificate

To create a TLS certificate signed by your intermediate CA:

```bash
easy-ca-cli gen tls --rsa --rsa-keysize 4096 \
  --subject "/C=CN/O=Example Inc/OU=Web Services/CN=example.com" \
  --host "example.com,www.example.com,192.168.1.100" \
  --start-date "2025-01-01 14:00:00" --days 365 \
  --issuer-key intermediate_key.pem --issuer-cert intermediate_cert.chain.pem \
```

This command:

- Uses the `RSA` algorithm with 4096-bit key size
- Sets the certificate subject and Subject Alternative Names (SANs) for multiple domains and an IP address
- Makes the certificate valid for 1 year
- References the intermediate CA for signing
- Outputs generated certificate and private key by default file name, if flags are not specified.

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

## Related Projects

- [easy-ca](https://github.com/c3b2a7/easy-ca): The core library that provides the certificate generation functionality
- [cobar](https://github.com/spf13/cobra): A powerful framework for creating modern CLI applications with nested
  commands

## CHANGELOG

See [CHANGELOG.md](./CHANGELOG.md)

## LICENSE

Easy CA CLI is released under the MIT license. See [LICENSE](./LICENSE)
