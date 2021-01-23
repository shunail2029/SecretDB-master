#!/bin/bash

secretdbd unsafe-reset-all

rm $HOME/.secretdbd/config/genesis.json
rm $HOME/.secretdbd/config/gentx/*
rm $HOME/.secretdbcli/config/config.toml

secretdbd init mynode --chain-id test1

secretdbcli config keyring-backend test
secretdbcli config chain-id test1
secretdbcli config output json
secretdbcli config indent true
secretdbcli config trust-node true

secretdbcli keys add user1
secretdbcli keys add user2

secretdbd add-genesis-account $(secretdbcli keys show user1 -a) 1000token,100000000stake
secretdbd add-genesis-account $(secretdbcli keys show user2 -a) 1000token

secretdbd gentx --name user1 --keyring-backend test
secretdbd collect-gentxs
