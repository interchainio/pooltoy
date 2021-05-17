#!/bin/bash
rm -rf ~/.pooltoy

pooltoy init mynode --chain-id pooltoy-4

# pooltoycli config keyring-backend test

pooltoy keys add alice --keyring-backend test
pooltoy keys add bob --keyring-backend test

pooltoy add-genesis-account $(pooltoy keys show alice -a --keyring-backend test) 1000token,100000000stake --keyring-backend test
pooltoy add-genesis-account $(pooltoy keys show bob -a --keyring-backend test) 1token --keyring-backend test

jq -c '.accounts[]' accounts.json | while read i; do
    pooltoy add-genesis-account $(echo "$i" | jq -r '.address') $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend test
    #echo $(echo "$i" | jq -r '.address') $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")')
done

pooltoy collect-gentxs
