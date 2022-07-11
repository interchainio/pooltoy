#!/bin/bash



rm -rf ~/.pooltoy
pooltoy init mynode --chain-id pooltoy-7
pooltoy config chain-id pooltoy-7
pooltoy config keyring-backend test
cp alice-priv_validator_key.json ~/.pooltoy/config/priv_validator_key.json
unzip keys.zip -d ~/.pooltoy/keyring-test
# will yell about empty accounts but that's ok
jq -c '.app_state.bank.balances[] | select((.coins | length) > 2)' real-fresh.json | while read i; do
  # ignore alice's account and add manually at the end
  if [ $(echo "$i" | jq -r ".address") != "cosmos1hf8x9kgx5djqz7c9e23v6lp8up97h5kcgp0cl3" ]; then
    echo "$i" | jq -r ".address"
    pooltoy add-genesis-account $(echo "$i" | jq -r ".address") $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")')
  fi
done
pooltoy add-genesis-account alice 10000000000000000stake --keyring-backend test
pooltoy gentx alice  1000000000000stake  --chain-id pooltoy-7
pooltoy collect-gentxs
