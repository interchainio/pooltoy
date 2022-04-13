#!/bin/bash
# pooltoy tx faucet mintfor $(pooltoy keys show bob -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
# pooltoy tx send bob $(pooltoy keys show alice -a) 1🚀 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx
# pooltoy tx send you $(pooltoy keys show me -a) 1token --from you -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx
# pooltoy tx send me $(pooltoy keys show who -a) 1token --from me | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx

echo "make alice admin by alice"
pooltoy tx pooltoy create-user $(pooltoy keys show alice -a) true billy billy@interchain.berlin --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

echo "make bob admin by alice"
pooltoy tx pooltoy create-user $(pooltoy keys show bob -a) true sam sam@interchain.berlin --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
# echo "make carol non-admin by bob"
# pooltoy tx pooltoy create-user $(pooltoy keys show carol -a) false marko marko@interchain.berlin --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

# # this one should fail
# echo "should fail"
# pooltoy tx pooltoy create-user $(pooltoy keys show doug -a) false dieter dieter@interchain.berlin --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

echo "how long til alice can mint?"
pooltoy q faucet when-brrr $(pooltoy keys show alice -a)

echo "alice mints 🚀 for bob"
pooltoy tx faucet mintfor $(pooltoy keys show bob -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

echo "how long til alice can mint?"
pooltoy q faucet when-brrr $(pooltoy keys show alice -a)

echo "can bob send to alice?"
pooltoy tx send bob $(pooltoy keys show alice -a) 1🚀 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

# echo "should fail"
# pooltoy tx faucet mintfor $(pooltoy keys show carol -a) 🚀 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
# echo "bob mints 🌝 for carol"
# pooltoy tx faucet mintfor $(pooltoy keys show carol -a) 🌝 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
# echo "carol mints 💸 for alice"
# pooltoy tx faucet mintfor $(pooltoy keys show alice -a) 💸 --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

# # this one should fail
# echo "should fail"
# pooltoy tx faucet mintfor $(pooltoy keys show alice -a) 💸 --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

# pooltoy q account $(pooltoy keys show alice -a)
# pooltoy q account $(pooltoy keys show bob -a)
# pooltoy q account $(pooltoy keys show carol -a)

# pooltoy tx send alice $(pooltoy keys show bob -a) 1💸 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
# pooltoy tx send bob $(pooltoy keys show carol -a) 1🚀 --from bob -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"
# pooltoy tx send carol $(pooltoy keys show alice -a) 1🌝 --from carol -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

# pooltoy q account $(pooltoy keys show alice -a)
# pooltoy q account $(pooltoy keys show bob -a)
# pooltoy q account $(pooltoy keys show carol -a)

# echo "should fail"
# pooltoy tx send alice $(pooltoy keys show doug -a) 1🌝 --from alice -y | jq ".txhash" |  xargs $(sleep 6) pooltoy q tx | jq ".raw_log"

