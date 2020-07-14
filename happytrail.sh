#!/bin/bash
# pooltoycli tx send me $(pooltoycli keys show you -a) 1token --from me -y -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx

# pooltoycli tx send you $(pooltoycli keys show me -a) 1token --from you -y -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx

# pooltoycli tx send me $(pooltoycli keys show who -a) 1token --from me -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx

pooltoycli tx pooltoy create-user $(pooltoycli keys show alice -a) true billy billy@interchain.berlin --from alice -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx
pooltoycli tx pooltoy create-user $(pooltoycli keys show bob -a) true sam sam@interchain.berlin --from alice -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx
pooltoycli tx pooltoy create-user $(pooltoycli keys show carol -a) false marko marko@interchain.berlin --from bob -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx

pooltoycli tx faucet mintfor $(pooltoycli keys show bob -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx
pooltoycli tx faucet mintfor $(pooltoycli keys show carol -a) 🌝 --from bob -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx
pooltoycli tx faucet mintfor $(pooltoycli keys show alice -a) 💸 --from carol -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx
pooltoycli tx faucet mintfor $(pooltoycli keys show alice -a) 💸 --from carol -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx