#!/bin/bash
# pooltoycli tx send me $(pooltoycli keys show you -a) 1token --from me -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx

# pooltoycli tx send you $(pooltoycli keys show me -a) 1token --from you -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx

pooltoycli tx send me $(pooltoycli keys show who -a) 1token --from me -y | jq ".txhash" |  xargs $(sleep 5) pooltoycli q tx