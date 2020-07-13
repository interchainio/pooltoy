#!/bin/bash
rm -r ~/.pooltoycli
rm -r ~/.pooltoyd

pooltoyd init mynode --chain-id pooltoy

pooltoycli config keyring-backend test

pooltoycli keys add me
pooltoycli keys add you

pooltoyd add-genesis-account $(pooltoycli keys show me -a) 1000token,100000000stake
pooltoyd add-genesis-account $(pooltoycli keys show you -a) 1token

pooltoycli config chain-id pooltoy
pooltoycli config output json
pooltoycli config indent true
pooltoycli config trust-node true

pooltoyd gentx --name me --keyring-backend test
pooltoyd collect-gentxs