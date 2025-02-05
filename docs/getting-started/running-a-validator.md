# Running a Validator

This document provides instructions for running a validator node for a [live network](./live-networks.md). With Regen Mainnet, Redwood Testnet, and Hambach Testnet already launched and running, this document will focus on how to become a validator for a network post-genesis.

## Prerequisites

In order to install the `cosmovisor` and `regen` binaries, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.17`

For more information (including hardware recommendations), see [Prerequisites](./prerequisites.md). 

## Quickstart

If you would like to manually set up a validator node, skip to the [next section](#install-regen). Alternatively, you can run the following quickstart script:

```bash
bash <(curl -s https://raw.githubusercontent.com/regen-network/mainnet/blob/main/scripts/mainnet-val-setup.sh)
```

## Install Regen

Clone the `regen-ledger` repository:

```bash
git clone https://github.com/regen-network/regen-ledger
```

Change to the `regen-ledger` directory:

```bash
cd regen-ledger
```

Check out the version that the network launched with.

*For Regen Mainnet:*

```bash
git checkout v1.0.0
```

*For Redwood Testnet:*

```bash
git checkout v1.0.0
```

*For Hambach Testnet:*

```bash
git checkout v2.0.0-beta1
```

Install the `regen` binary (the `EXPERIMENTAL` option enables experimental features).

*For Regen Mainnet:*

```bash
make install
```

*For Redwood Testnet:*

```bash
make install
```

*For Hambach Testnet:*

```bash
EXPERIMENTAL=true make install
```

Check to ensure the install was successful:

```bash
regen version
```

## Initialize Node

Create the configuration files and data directory by initializing the node. In the following command, replace `[moniker]` with a name of your choice. 

*For Regen Mainnet:*

```bash
regen init [moniker] --chain-id regen-1
```

*For Redwood Testnet:*

```bash
regen init [moniker] --chain-id regen-redwood-1
```

*For Hambach Testnet:*

```bash
regen init [moniker] --chain-id regen-hambach-1
```

## Update Genesis

Update the genesis file using a node endpoint.

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*

```bash
curl http://104.131.169.70:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Redwood Testnet:*

```bash
curl http://redwood.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

*For Hambach Testnet:*

```bash
curl http://hambach.regen.network:26657/genesis | jq .result.genesis > ~/.regen/config/genesis.json
```

## Update Peers

Add a seed node for initial peer discovery.

<!-- TODO: update to use dedicated full node operated by RND -->

*For Regen Mainnet:*

```bash
PERSISTENT_PEERS="69975e7afdf731a165e40449fcffc75167a084fc@104.131.169.70:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Redwood Testnet:*

```bash
PERSISTENT_PEERS="a5528d8f5fabd3d50e91e8d6a97e355403c5b842@redwood.regen.network:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

*For Hambach Testnet:*

```bash
PERSISTENT_PEERS="4f5c0be7705bf4acb5b99dcaf93190059ac283a1@hambach.regen.network:26656"
sed -i '/persistent_peers =/c\persistent_peers = "'"$PERSISTENT_PEERS"'"' ~/.regen/config/config.toml
```

## Install Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

To install `cosmovisor`, run the following command:

```bash
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0
```

Check to ensure the install was successful:

```bash
cosmovisor version
```

## Set Genesis Binary

Create the folder for the genesis binary and copy the `regen` binary:

```bash
mkdir -p $HOME/.regen/cosmovisor/genesis/bin
cp $GOBIN/regen $HOME/.regen/cosmovisor/genesis/bin
```

## Cosmovisor Service

The next step will be to configure `cosmovisor` as a `systemd` service. For more information about the environment variables used to configure `cosmovisor`, see [Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor).

Create the `cosmovisor.service` file:

```bash
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.regen"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
User=${USER}
ExecStart=${GOBIN}/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

Move the file to the systemd directory:

```bash
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```

Reload systemctl and start `cosmovisor`:

```bash
sudo systemctl daemon-reload
sudo systemctl start cosmovisor
```

Check the status of the `cosmovisor` service:

```bash
sudo systemctl status cosmovisor
```

## Add Validator Key

As a validator who signs blocks, your node must have a public/private key pair. Regen Ledger keys can be managed with the `regen keys` subcommand. A new key pair can be generated using:

```bash
regen keys add [name]
```

::: warning
If you create a new key, make sure you store the mnemonic phrase in a safe place. You will not be able to recover your new key without the mnemonic phrase.
:::

If you'd like to use an existing key or a custom keyring backend, you can find more information about adding keys and keyring backends in the Cosmos SDK [Keyring](https://docs.cosmos.network/master/run-node/keyring.html) documentation.

## Create Validator

The next step will be to create a validator. You will need to have enough REGEN tokens to stake and to submit the transaction. For more information about the REGEN token, see the [token page](https://www.regen.network/token/). 

::: warning
You'll want to carefully consider the options you set when creating a validator.
:::

Submit a transaction to create a validator:

```bash
regen tx staking create-validator \
  --amount=<stake_amount> \
  --pubkey=$(regen tendermint show-validator) \
  --moniker="<node_moniker>" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --from=<key_name>
```

## Prepare Upgrade

The next step will be to prepare your node for the upgrade process. See [Upgrade Guide v2.0](../migrations/v2.0-upgrade.md) for more information.