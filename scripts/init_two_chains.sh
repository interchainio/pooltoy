#!/bin/bash

set -e

# Configure predetermined settings
# val0: cosmos1mjk79fjjgpplak5wq838w0yd982gzkyfrk07am
# val1: cosmos17dtl0mjt3t77kpuhg2edqzjpszulwhgzuj9ljs
BINARY=pooltoy
CHAIN_DIR=./data
CHAINID_0=test-0
CHAINID_1=test-1
MNEMONIC_0="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
MNEMONIC_1="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"
P2PPORT_0=16656
P2PPORT_1=26656
RPCPORT_0=16657
RPCPORT_1=26657
GRPCPORT_0=8090
GRPCPORT_1=9090
RESTPORT_0=1316
RESTPORT_1=1317
# Hermes config
CONFIG_FILE=./scripts/hermes/config.toml

# Stop if it is already running 
if pgrep -x "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall $BINARY
fi

echo "Removing previous data..."
rm -rf $CHAIN_DIR/$CHAINID_0 &> /dev/null
rm -rf $CHAIN_DIR/$CHAINID_1 &> /dev/null

# Add directories for both chains, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAINID_0 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

if ! mkdir -p $CHAIN_DIR/$CHAINID_1 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

echo "Initializing $CHAINID_0..."
echo "Initializing $CHAINID_1..."
$BINARY init test --home $CHAIN_DIR/$CHAINID_0 --chain-id=$CHAINID_0
$BINARY init test --home $CHAIN_DIR/$CHAINID_1 --chain-id=$CHAINID_1

echo "Adding genesis accounts..."
echo $MNEMONIC_0 | $BINARY keys add val0 --home $CHAIN_DIR/$CHAINID_0 --recover --keyring-backend=test 
echo $MNEMONIC_1 | $BINARY keys add val1 --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test 

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_0 keys show val0 --keyring-backend test -a) 100000000000stake,1🍀  --home $CHAIN_DIR/$CHAINID_0
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show val1 --keyring-backend test -a) 100000000000stake,1🍣   --home $CHAIN_DIR/$CHAINID_1

echo "Creating and collecting gentx..."
$BINARY gentx val0 7000000000stake --home $CHAIN_DIR/$CHAINID_0 --chain-id $CHAINID_0 --keyring-backend test
$BINARY gentx val1 7000000000stake --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test

$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID_0
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID_1

echo "Changing defaults and ports in app.toml and config.toml files..."
sed -i -e 's#"tcp://0.0.0.0:26656"#"tcp://0.0.0.0:'"$P2PPORT_0"'"#g' $CHAIN_DIR/$CHAINID_0/config/config.toml
sed -i -e 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:'"$RPCPORT_0"'"#g' $CHAIN_DIR/$CHAINID_0/config/config.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' $CHAIN_DIR/$CHAINID_0/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $CHAIN_DIR/$CHAINID_0/config/config.toml
sed -i -e 's/index_all_keys = false/index_all_keys = true/g' $CHAIN_DIR/$CHAINID_0/config/config.toml
sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAINID_0/config/app.toml
sed -i -e 's/swagger = false/swagger = true/g' $CHAIN_DIR/$CHAINID_0/config/app.toml
sed -i -e 's#"tcp://0.0.0.0:1317"#"tcp://0.0.0.0:'"$RESTPORT_0"'"#g' $CHAIN_DIR/$CHAINID_0/config/app.toml

sed -i -e 's#"tcp://0.0.0.0:26656"#"tcp://0.0.0.0:'"$P2PPORT_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:'"$RPCPORT_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/index_all_keys = false/index_all_keys = true/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's/swagger = false/swagger = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's#"tcp://0.0.0.0:1317"#"tcp://0.0.0.0:'"$RESTPORT_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/app.toml

echo "Starting $CHAINID_0 in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/$CHAINID_0.log"
$BINARY start --home $CHAIN_DIR/$CHAINID_0 --pruning=nothing --grpc.address="0.0.0.0:$GRPCPORT_0" > $CHAIN_DIR/$CHAINID_0.log 2>&1 &

echo "Starting $CHAINID_1 in $CHAIN_DIR..."
echo "Creating log file at $CHAIN_DIR/$CHAINID_1.log"
$BINARY start --home $CHAIN_DIR/$CHAINID_1 --pruning=nothing --grpc.address="0.0.0.0:$GRPCPORT_1" > $CHAIN_DIR/$CHAINID_1.log 2>&1 &

# Add the key seeds to the keyring of each chain
### Sleep is needed otherwise the relayer crashes when trying to init
echo "Importing keys..."
sleep 1s
hermes -c $CONFIG_FILE keys restore $CHAINID_0 -m "$MNEMONIC_0"
sleep 5s
hermes -c $CONFIG_FILE keys restore $CHAINID_1 -m "$MNEMONIC_1"
sleep 5s
echo "Done!"