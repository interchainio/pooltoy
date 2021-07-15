#!/bin/bash
#[ -z "$CHAIN_DIR" ] && CHAIN_DIR="~/.pooltoy"
#[ -z "$CHAIN_ID" ] && CHAIN_ID="pooltoy-5"
#[ -z "$KEYRING" ] && CHAIN_ID="test"
#
#rm -rf $CHAIN_DIR
#pooltoy init mynode  --home $CHAIN_DIR --chain-id $CHAIN_ID
#
#jq -c '.balances[]' accounts.json | while read i; do
##<<<<<<< Updated upstream
##
##    echo "y" | pooltoy keys add $(echo "$i" | jq -r ".name") --keyring-backend $KEYRING --home $CHAIN_DIR
##    pooltoy add-genesis-account $(pooltoy keys show $(echo "$i" | jq -r ".name") --address --keyring-backend $KEYRING --home $CHAIN_DIR) $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend $KEYRING --home $CHAIN_DIR
##done
##
##pooltoy gentx alice 100000000stake --home $CHAIN_DIR --chain-id $CHAIN_ID --keyring-backend $KEYRING
##pooltoy collect-gentxs --home $CHAIN_DIR
##pooltoy validate-genesis --home $CHAIN_DIR
##=======
#    echo "y" | pooltoy keys add $(echo "$i" | jq -r ".name") --keyring-backend $KEYRING --home $CHAIN_DIR
#    pooltoy add-genesis-account $(pooltoy keys show $(echo "$i" | jq -r ".name") --address) $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend $KEYRING --home $CHAIN_DIR
#done
#
#pooltoy gentx alice 100000000stake --home $CHAIN_DIR --chain-id $CHAIN_ID --keyring-backend $KEYRING
#pooltoy collect-gentxs --home $CHAIN_DIR
#pooltoy validate-genesis --home $CHAIN_DIR

#pooltoy gentx alice 100000000stake --chain-id pooltoy-5
#pooltoy collect-gentxs
#pooltoy validate-genesis # test to make sure the genesis file is correctly formatted
#>>>>>>> Stashed changes

rm -rf ~/.pooltoy

pooltoy config keyring-backend test
pooltoy config chain-id pooltoy-5

pooltoy config

pooltoy init mynode --chain-id pooltoy-5

jq -c '.balances[]' accounts.json | while read i; do
    echo "y" | pooltoy keys add $(echo "$i" | jq -r ".name") --keyring-backend test
    pooltoy add-genesis-account $(pooltoy keys show $(echo "$i" | jq -r ".name") --address) $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")') --keyring-backend test
done

pooltoy gentx alice 100000000stake --keyring-backend test --chain-id pooltoy-5
pooltoy collect-gentxs
pooltoy validate-genesis # test to make sure the genesis file is correctly formatted
