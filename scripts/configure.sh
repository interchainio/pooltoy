#!/bin/bash
pooltoy tx pooltoy create-user $(pooltoy keys show alice -a --keyring-backend test) true alice alice --from alice -y --keyring-backend test --chain-id pooltoy-5 
# | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
