#!/bin/bash
# rm -rf ~/.pooltoycli
rm -rf ~/.pooltoyd

pooltoy init mynode --chain-id pooltoy-0

# pooltoycli config keyring-backend test

pooltoy keys add alice
pooltoy keys add bob

pooltoy add-genesis-account $(pooltoy keys show alice -a) 1000token,100000000stake
pooltoy add-genesis-account $(pooltoy keys show bob -a) 1token

jq -c '.accounts[]' accounts.json | while read i; do
    pooltoy add-genesis-account $(echo "$i" | jq -r '.value.address') $(echo $i | jq -r '.value.coins | map(.amount) | join(",")')
done

pooltoy collect-gentxs