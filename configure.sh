#!/bin/bash
pooltoycli tx pooltoy create-user $(pooltoycli keys show alice -a) true billy U01246QFA2U --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
pooltoycli tx pooltoy create-user $(pooltoycli keys show bob -a) true sam U011HTNSBSN --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
