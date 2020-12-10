package types

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {

	testUsers := []User{
		User{},

		User{
			Creator:     bytes.Repeat([]byte{1}, sdk.AddrLen),
			UserAccount: bytes.Repeat([]byte{2}, sdk.AddrLen),
			IsAdmin:     false,
		},

		User{
			Creator:     bytes.Repeat([]byte{3}, sdk.AddrLen),
			UserAccount: bytes.Repeat([]byte{4}, sdk.AddrLen),
			IsAdmin:     true,
		},

		User{
			Creator: bytes.Repeat([]byte{3}, sdk.AddrLen),
			IsAdmin: true,
		},

		User{
			UserAccount: bytes.Repeat([]byte{4}, sdk.AddrLen),
			IsAdmin:     true,
		},
	}

	tests := []struct {
		desc        string
		genesis     GenesisState
		shouldError bool
	}{
		{desc: "User is Alice",
			genesis:     DefaultGenesisState(),
			shouldError: false,
		},
		{
			desc:        "Empty user and creator",
			genesis:     NewGenesisState(testUsers[0]),
			shouldError: true,
		},
		{
			desc:        "Creator is not an admin",
			genesis:     NewGenesisState(testUsers[1]),
			shouldError: true,
		},
		{
			desc:        "Legit user",
			genesis:     NewGenesisState(testUsers[2]),
			shouldError: false,
		},
		{
			desc:        "Empty user",
			genesis:     NewGenesisState(testUsers[3]),
			shouldError: true,
		},
		{
			desc:        "Empty creator",
			genesis:     NewGenesisState(testUsers[4]),
			shouldError: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			if test.shouldError {
				require.Error(t, ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, ValidateGenesis(test.genesis))
			}
		})
	}
}
