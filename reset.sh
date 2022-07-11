#!/bin/bash

rm -rf ~/.pooltoy
pooltoy init mynode --chain-id pooltoy-7
pooltoy config chain-id pooltoy-7
pooltoy config keyring-backend test
cp ~/.pooltoy-6/config/priv_validator_key.json ~/.pooltoy/config/priv_validator_key.json
cp ~/.pooltoy-6/data/priv_validator_state.json ~/.pooltoy/data/priv_validator_state.json
cp -r ~/.pooltoy-6/keyring-test ~/.pooltoy
cp real-fresh.json ~/.pooltoy/config/genesis.json
pooltoy validate-genesis