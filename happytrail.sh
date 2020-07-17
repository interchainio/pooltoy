#!/bin/bash
# pooltoycli tx faucet mintfor $(pooltoycli keys show bob -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
# pooltoycli tx send bob $(pooltoycli keys show alice -a) 1🚀 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx
# pooltoycli tx send you $(pooltoycli keys show me -a) 1token --from you -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx
# pooltoycli tx send me $(pooltoycli keys show who -a) 1token --from me | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx

echo "make alice admin by alice"
pooltoycli tx pooltoy create-user $(pooltoycli keys show alice -a) true billy billy@interchain.berlin --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
echo "make bob admin by alice"
pooltoycli tx pooltoy create-user $(pooltoycli keys show bob -a) true sam sam@interchain.berlin --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
echo "make carol non-admin by bob"
pooltoycli tx pooltoy create-user $(pooltoycli keys show carol -a) false marko marko@interchain.berlin --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"

# this one should fail
echo "should fail"
pooltoycli tx pooltoy create-user $(pooltoycli keys show doug -a) false dieter dieter@interchain.berlin --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"

echo "alice mints 🚀 for bob"
pooltoycli tx faucet mintfor $(pooltoycli keys show bob -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
echo "should fail"
pooltoycli tx faucet mintfor $(pooltoycli keys show carol -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
echo "bob mints 🌝 for carol"
pooltoycli tx faucet mintfor $(pooltoycli keys show carol -a) 🌝 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
echo "carol mints 💸 for alice"
pooltoycli tx faucet mintfor $(pooltoycli keys show alice -a) 💸 --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"

# this one should fail
echo "should fail"
pooltoycli tx faucet mintfor $(pooltoycli keys show alice -a) 💸 --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"

pooltoycli q account $(pooltoycli keys show alice -a)
pooltoycli q account $(pooltoycli keys show bob -a)
pooltoycli q account $(pooltoycli keys show carol -a)

pooltoycli tx send alice $(pooltoycli keys show bob -a) 1💸 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
pooltoycli tx send bob $(pooltoycli keys show carol -a) 1🚀 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"
pooltoycli tx send carol $(pooltoycli keys show alice -a) 1🌝 --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"

pooltoycli q account $(pooltoycli keys show alice -a)
pooltoycli q account $(pooltoycli keys show bob -a)
pooltoycli q account $(pooltoycli keys show carol -a)

echo "should fail"
pooltoycli tx send alice $(pooltoycli keys show doug -a) 1🌝 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoycli q tx | jq ".raw_log"

