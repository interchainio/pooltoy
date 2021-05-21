#!/bin/bash
rm -rf ~/.pooltoy

pooltoy init mynode --chain-id pooltoy-5

# pooltoycli config keyring-backend test

# pooltoy keys add alice --keyring-backend test
# pooltoy keys add bob --keyring-backend test

# pooltoy add-genesis-account $(pooltoy keys show alice -a --keyring-backend test) 1000token,100000000stake --keyring-backend test
# pooltoy add-genesis-account $(pooltoy keys show bob -a --keyring-backend test) 1token --keyring-backend test

jq -c '.balances[]' accounts.json | while read i; do

    echo "y" | pooltoy keys add $(echo "$i" | jq -r ".name") --keyring-backend test
    pooltoy add-genesis-account $(pooltoy keys show $(echo "$i" | jq -r ".name") --address --keyring-backend test) $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend test
    #echo $(echo "$i" | jq -r '.address') $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")')
done

pooltoy gentx alice 100000000stake --chain-id pooltoy-5 --keyring-backend test
pooltoy collect-gentxs
pooltoy validate-genesis
