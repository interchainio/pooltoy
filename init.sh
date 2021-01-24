#!/bin/bash
rm -r ~/.pooltoycli
rm -r ~/.pooltoyd

pooltoyd init mynode --chain-id pooltoy-4

pooltoycli config keyring-backend test

pooltoycli keys add alice
pooltoycli keys add bob
pooltoycli keys add carol
pooltoycli keys add doug

pooltoyd add-genesis-account $(pooltoycli keys show alice -a) 1000token,100000000stake
pooltoyd add-genesis-account $(pooltoycli keys show bob -a) 1token

pooltoycli config chain-id pooltoy-4
pooltoycli config output json
pooltoycli config indent true
pooltoycli config trust-node true

pooltoyd gentx --name alice --keyring-backend test
pooltoyd collect-gentxs