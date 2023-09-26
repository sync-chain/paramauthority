package tests

import (
	"context"
	"encoding/json"
	"testing"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/icza/dyno"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	// Authority
	authorityTypes "github.com/noble-assets/paramauthority/x/authority/types"
	// IBC Transfer
	ibcTransferTypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	// Params
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

func TestLegacyParams(t *testing.T) {
	ctx := context.Background()
	reporter := testreporter.NewNopReporter()
	client, network := interchaintest.DockerSetup(t)

	var authority ibc.Wallet

	factory := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:    "authority-simd",
			Version: "local",
			ChainConfig: ibc.ChainConfig{
				UsingNewGenesisCommand: true,
				PreGenesis: func(cfg ibc.ChainConfig) error {
					authority = RandomAccount(cfg)
					return nil
				},
				ModifyGenesis: func(cfg ibc.ChainConfig, bz []byte) ([]byte, error) {
					genesis := make(map[string]interface{})
					_ = json.Unmarshal(bz, &genesis)

					_ = dyno.Set(genesis, authority.FormattedAddress(), "app_state", "authority", "authority")

					newGenesis, _ := json.Marshal(genesis)
					return newGenesis, nil
				},
			},
		},
	})

	chains, _ := factory.Chains(t.Name())
	simapp := chains[0].(*cosmos.CosmosChain)

	interchain := interchaintest.NewInterchain().AddChain(simapp)
	t.Cleanup(func() {
		_ = interchain.Close()
	})

	require.NoError(t, interchain.Build(ctx, reporter.RelayerExecReporter(t), interchaintest.InterchainBuildOptions{
		TestName:  t.Name(),
		Client:    client,
		NetworkID: network,
	}))

	authority, err := interchaintest.GetAndFundTestUserWithMnemonic(ctx, t.Name(), authority.Mnemonic(), 1000000, simapp)
	require.NoError(t, err)

	//
	t.Run("Update IBC Transfer Legacy Params", func(t *testing.T) {
		content := params.NewParameterChangeProposal("", "", []params.ParamChange{
			{
				Subspace: ibcTransferTypes.ModuleName,
				Key:      string(ibcTransferTypes.KeySendEnabled),
				Value:    "false",
			},
			{
				Subspace: ibcTransferTypes.ModuleName,
				Key:      string(ibcTransferTypes.KeyReceiveEnabled),
				Value:    "false",
			},
		})
		rawContent, _ := codecTypes.NewAnyWithValue(content)

		legacyMsg := &authorityTypes.MsgExecuteLegacyContent{
			Authority: authorityTypes.ModuleAddress.String(),
			Content:   rawContent,
		}
		rawLegacyMsg, _ := codecTypes.NewAnyWithValue(legacyMsg)

		msg := &authorityTypes.MsgExecute{
			Authority: authority.FormattedAddress(),
			Messages:  []*codecTypes.Any{rawLegacyMsg},
		}

		_, err := cosmos.BroadcastTx(ctx, cosmos.NewBroadcaster(t, simapp), authority, msg)
		require.NoError(t, err)

		//
		sendEnabled, err := simapp.QueryParam(ctx, ibcTransferTypes.ModuleName, string(ibcTransferTypes.KeySendEnabled))
		require.NoError(t, err)
		require.Equal(t, "false", sendEnabled.Value)

		receiveEnabled, err := simapp.QueryParam(ctx, ibcTransferTypes.ModuleName, string(ibcTransferTypes.KeyReceiveEnabled))
		require.NoError(t, err)
		require.Equal(t, "false", receiveEnabled.Value)
	})
}
