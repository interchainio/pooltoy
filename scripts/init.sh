#!/bin/bash
rm -rf ~/.pooltoy

pooltoy config keyring-backend test
pooltoy config chain-id pooltoy-6

pooltoy config

pooltoy init mynode --chain-id pooltoy-6

jq -c '.balances[]' accounts.json | while read i; do
    echo "y" | pooltoy keys add $(echo "$i" | jq -r ".name") --keyring-backend test
    pooltoy add-genesis-account $(pooltoy keys show $(echo "$i" | jq -r ".name") --address) $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend test
done

pooltoy gentx alice 100000000stake --keyring-backend test --chain-id pooltoy-6
pooltoy collect-gentxs
pooltoy validate-genesis # test to make sure the genesis file is correctly formatted
