#!/bin/bash

set -e

# Hermes config
CONFIG_FILE=./scripts/hermes/config.toml

# Create IBC client identifiers on each of the chains
echo "Configuring clients..."
hermes -c $CONFIG_FILE tx raw create-client test-0 test-1 
hermes -c $CONFIG_FILE tx raw create-client test-1 test-0 

### Connection Handshake
echo "Initiating connection handshake..."
# conn-init
hermes -c $CONFIG_FILE tx raw conn-init test-0 test-1 07-tendermint-0 07-tendermint-0
# conn-try
echo "trying connection handshake..."
hermes -c $CONFIG_FILE tx raw conn-try test-1 test-0 07-tendermint-0 07-tendermint-0 -s connection-0
# conn-ack
echo "Acknowledging connection handshake..."
hermes -c $CONFIG_FILE tx raw conn-ack test-0 test-1 07-tendermint-0 07-tendermint-0 -d connection-0 -s connection-0
# conn-confirm
echo "Confirming connection handshake..."
hermes -c $CONFIG_FILE tx raw conn-confirm test-1 test-0 07-tendermint-0 07-tendermint-0 -d connection-0 -s connection-0

echo "Establishing transfer channel..."
# now that the handshake was successful, we can establish a transfer channel
hermes -c $CONFIG_FILE create channel test-0 test-1 --port-a transfer --port-b transfer

echo "Listening to transfer channel..."
hermes -c $CONFIG_FILE start test-0 test-1 -p transfer -c channel-0