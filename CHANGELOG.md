# Changelog

All notable changes to this project will be documented in this file.

## [1.5.0](https://github.com/c3b2a7/easy-ca-cli/compare/v1.4.0..v1.5.0) - 2025-08-11

### ✨ Features

- Add installation script for easier setup and verification ([#30](https://github.com/c3b2a7/easy-ca-cli/issues/30)) - ([82d985d](https://github.com/c3b2a7/easy-ca-cli/commit/82d985d03ad981ea3c6e5b4ef48879fc85a80a96))

### 📚 Documentation

- *(readme)* Add examples for using docker images ([#39](https://github.com/c3b2a7/easy-ca-cli/issues/39)) - ([ac11c56](https://github.com/c3b2a7/easy-ca-cli/commit/ac11c56c2023ebcf50323f305dfbfbdbe09cf6d6))
- *(readme)* Update installation section with manual steps ([#31](https://github.com/c3b2a7/easy-ca-cli/issues/31)) - ([4251298](https://github.com/c3b2a7/easy-ca-cli/commit/4251298eba6021fa04900da3b5bc98c0d1e1e44b))
- *(readme)* Add a badge linking to DeepWiki ([#27](https://github.com/c3b2a7/easy-ca-cli/issues/27)) - ([dbaa223](https://github.com/c3b2a7/easy-ca-cli/commit/dbaa223b92f6115a0917df36eb09b80d6c135d75))

### ⚙️ Miscellaneous Tasks

- Build and release docker images ([#38](https://github.com/c3b2a7/easy-ca-cli/issues/38)) - ([f69080e](https://github.com/c3b2a7/easy-ca-cli/commit/f69080e93abb6dbd9cd1344996d95e6ebaa8edc9))
- Rename dependentbot.yml to dependabot.yml ([#28](https://github.com/c3b2a7/easy-ca-cli/issues/28)) - ([51ab130](https://github.com/c3b2a7/easy-ca-cli/commit/51ab130229d1b15402ab6b42d627bd8a26c338ef))
- Add dependabot.yml and fix typos in documentation ([#25](https://github.com/c3b2a7/easy-ca-cli/issues/25)) - ([2dd7177](https://github.com/c3b2a7/easy-ca-cli/commit/2dd7177df9baf85e08e5f74716e4b8602e38c143))
- Add community guidelines and contribution templates ([#24](https://github.com/c3b2a7/easy-ca-cli/issues/24)) - ([37e1d97](https://github.com/c3b2a7/easy-ca-cli/commit/37e1d9772c112798b308df17c1ca9c8033c9133a))


## [1.4.0](https://github.com/c3b2a7/easy-ca-cli/compare/v1.3.0..v1.4.0) - 2025-05-07

### ✨ Features

- Add support for generating PKCS#12 files ([#13](https://github.com/c3b2a7/easy-ca-cli/issues/13)) - ([b02bce1](https://github.com/c3b2a7/easy-ca-cli/commit/b02bce1ca81a6bff582e4b5ed08f69abadbe83f6))

### 🐛 Bug Fixes

- Avoid infinite loop when decoding invalid or malformed PEM files ([#14](https://github.com/c3b2a7/easy-ca-cli/issues/14)) - ([a31db88](https://github.com/c3b2a7/easy-ca-cli/commit/a31db885326c560a29a6bf43a10aa6b87578a0f7))
- Resolve `days` flag conflict caused by shared `cli.CertConfig` ([#12](https://github.com/c3b2a7/easy-ca-cli/issues/12)) - ([1b66bbb](https://github.com/c3b2a7/easy-ca-cli/commit/1b66bbbca06066720a65521ba988e7514dd51a3b))

### 🚜 Refactor

- Use generic type parameters for better flexibility ([#16](https://github.com/c3b2a7/easy-ca-cli/issues/16)) - ([c1d74f7](https://github.com/c3b2a7/easy-ca-cli/commit/c1d74f7e73507ea6a769e08980764398b2d2f5e5))

### 📚 Documentation

- *(changelog)* Configure `git-cliff` for changelog generation ([#19](https://github.com/c3b2a7/easy-ca-cli/issues/19)) - ([d335b59](https://github.com/c3b2a7/easy-ca-cli/commit/d335b59c298038c60311aa6cbc43f12948047994))
- Add security policy ([#20](https://github.com/c3b2a7/easy-ca-cli/issues/20)) - ([5d89614](https://github.com/c3b2a7/easy-ca-cli/commit/5d89614677eb14a8e2cd3cf08c199f5f6d203c1c))

### 🎨 Styling

- Add .editorconfig and reformat code for consistency ([#23](https://github.com/c3b2a7/easy-ca-cli/issues/23)) - ([678cc5f](https://github.com/c3b2a7/easy-ca-cli/commit/678cc5f9cd3ee82596e92fc9a3d5d750929318fe))

### ⚙️ Miscellaneous Tasks

- *(cli)* Update version output template for improved readability ([#17](https://github.com/c3b2a7/easy-ca-cli/issues/17)) - ([a84a025](https://github.com/c3b2a7/easy-ca-cli/commit/a84a025fbd5aae5d4ae437f7c7abf1479e8de6ab))
- Update the workflow to include the correct files - ([59445f9](https://github.com/c3b2a7/easy-ca-cli/commit/59445f94309cee33c3ac361775e1ce03869ca0db))
- Add actions for checks and resolve lint issues ([#21](https://github.com/c3b2a7/easy-ca-cli/issues/21)) - ([179346c](https://github.com/c3b2a7/easy-ca-cli/commit/179346c41a9a25769e247e7914d12ac6032cdddd))
- Integrate GoReleaser with changelog and release automation ([#18](https://github.com/c3b2a7/easy-ca-cli/issues/18)) - ([ab8322b](https://github.com/c3b2a7/easy-ca-cli/commit/ab8322b1f6e620aeb683dd828a83c5fe9de62cf6))
- Update README.md - ([c61db40](https://github.com/c3b2a7/easy-ca-cli/commit/c61db40af29dd9744e365f3374ee21bec9379973))

## 👏 New Contributors

* @dependabot[bot] made their first contribution in [#15](https://github.com/c3b2a7/easy-ca-cli/pull/15)

## [1.3.0](https://github.com/c3b2a7/easy-ca-cli/compare/v1.1.0..v1.3.0) - 2025-05-03

### ✨ Features

- Generate default output path for the certificate and private key ([#10](https://github.com/c3b2a7/easy-ca-cli/issues/10)) - ([f465b1a](https://github.com/c3b2a7/easy-ca-cli/commit/f465b1a81ad2761cbb6f659e81b0a393defbfd8b))

### 🐛 Bug Fixes

- Key algorithm flags do not work ([#8](https://github.com/c3b2a7/easy-ca-cli/issues/8)) - ([3d4a713](https://github.com/c3b2a7/easy-ca-cli/commit/3d4a7133100098408bb4f6171437b2f889e25b28))
- Typo in ed25519 algorithm ([#3](https://github.com/c3b2a7/easy-ca-cli/issues/3)) - ([064a229](https://github.com/c3b2a7/easy-ca-cli/commit/064a229750f96c49ae63cf37788e00c09c053e2d))


## [1.0.0] - 2023-03-19

## 👏 New Contributors

* @c3b2a7 made their first contribution

<!-- generated by git-cliff -->
