#!/bin/bash

rm -rf ~/.pooltoy
pooltoy init mynode  --home ~/.pooltoy --chain-id pooltoy-5

jq -c '.balances[]' accounts.json | while read i; do

    echo "y" | pooltoy keys add $(echo "$i" | jq -r ".name") --keyring-backend test --home ~/.pooltoy 
    pooltoy add-genesis-account $(pooltoy keys show $(echo "$i" | jq -r ".name") --address --keyring-backend test --home ~/.pooltoy) $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend test --home ~/.pooltoy 
done

pooltoy gentx alice 100000000stake --home ~/.pooltoy --chain-id pooltoy-5 --keyring-backend test 
pooltoy collect-gentxs --home ~/.pooltoy
pooltoy validate-genesis --home ~/.pooltoy
