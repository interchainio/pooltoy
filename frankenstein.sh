#!/bin/bash

# will yell about empty accounts but that's ok
jq -c '.app_state.bank.balances[]' real-fresh.json | while read i; do
pooltoy add-genesis-account $(echo "$i" | jq -r ".address") $(echo $i | jq -r '.coins | map(.amount+.denom) | join(",")')
done