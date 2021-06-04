#!/bin/bash

set -e

# Hermes config
CONFIG_FILE=./scripts/hermes/config.toml

# Send a clover from test-0 to test-1
hermes -c $CONFIG_FILE  tx raw ft-transfer test-1 test-0 transfer channel-0 1 -o 1 -d  🍀

# Send a sushi back from test-1 to test-0
hermes -c $CONFIG_FILE  tx raw ft-transfer test-1 test-0 transfer channel-0 1 -o 1 -d  🍣