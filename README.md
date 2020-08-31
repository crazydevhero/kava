<p align="center">
  <img src="./kava-logo.svg" width="300">
</p>
<h3 align="center">DeFi for Crypto.</h3>

<div align="center">

[![version](https://img.shields.io/github/tag/kava-labs/kava.svg)](https://github.com/kava-labs/kava/releases/latest)
[![CircleCI](https://circleci.com/gh/Kava-Labs/kava/tree/master.svg?style=shield)](https://circleci.com/gh/Kava-Labs/kava/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kava-labs/kava)](https://goreportcard.com/report/github.com/kava-labs/kava)
[![API Reference](https://godoc.org/github.com/Kava-Labs/kava?status.svg)](https://godoc.org/github.com/Kava-Labs/kava)
[![GitHub](https://img.shields.io/github/license/kava-labs/kava.svg)](https://github.com/Kava-Labs/kava/blob/master/LICENSE.md)
[![Twitter Follow](https://img.shields.io/twitter/follow/kava_labs.svg?label=Follow&style=social)](https://twitter.com/kava_labs)
[![Discord Chat](https://img.shields.io/discord/704389840614981673.svg)](https://discord.com/invite/kQzh3Uv)

</div>

<div align="center">

### [Telegram](https://t.me/kavalabs) | [Medium](https://medium.com/kava-labs) | [Discord](https://discord.com/invite/kQzh3Uv)

</div>

Reference implementation of Kava, a blockchain for cross-chain DeFi. Built using the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk).

## Mainnet

Note, the current recommended version of the software for mainnet is v0.10.0. The master branch of this repository often contains considerable development work since the last mainnet release and is __not__ runnable on mainnet.

### Installation

```bash
git checkout v0.10.0
make install
```

### Upgrade

The scheduled mainnet upgrade to `kava-3` took place on June 10th, 2020 at 14:00 UTC. The current version of Kava for `kava-3` is [__v0.10.0__](https://github.com/Kava-Labs/kava/releases/tag/v0.10.0).

Migration instructions can be found [here](https://github.com/Kava-Labs/kava/blob/master/contrib/kava-3/migration.md).

The canonical genesis file can be found [here](https://github.com/Kava-Labs/launch/tree/master/kava-3)

The canonical genesis file hash is

```
jq -S -c -M '' genesis.json | shasum -a 256
# f73628abfab82601c9af97a023d357a95507b9c630c5331564f48c4acab97b85
```

## Testnet

The recommended version of the software for kava-testnet-8000 is 0.9.0-rc1. For further information on joining the testnet, head over to the [testnet repo](https://github.com/Kava-Labs/kava-testnets).

## Docs

Kava protocol and client documentation can be found in the [Kava docs](https://docs.kava.io).

If you have technical questions or concerns, ask a developer or community member in the [Kava discord](https://discord.com/invite/kQzh3Uv).

## License

Copyright © Kava Labs, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
