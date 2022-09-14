package ante_test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/tx"
// 	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
// 	"github.com/cosmos/cosmos-sdk/testutil/testdata"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/tx/signing"
// 	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
// 	"github.com/cosmos/cosmos-sdk/x/params/types"
// 	"github.com/stretchr/testify/suite"
// 	tmrand "github.com/tendermint/tendermint/libs/rand"
// 	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

// 	app "github.com/eve-network/eve/app"
// 	gaiahelpers "github.com/eve-network/eve/app/helpers"
// 	"github.com/eve-network/eve/x/globalfee"
// 	globfeetypes "github.com/eve-network/eve/x/globalfee/types"
// )

// type IntegrationTestSuite struct {
// 	suite.Suite

// 	app         *app.EveApp
// 	anteHandler sdk.AnteHandler
// 	ctx         sdk.Context
// 	clientCtx   client.Context
// 	txBuilder   client.TxBuilder
// }

// func TestIntegrationTestSuite(t *testing.T) {
// 	suite.Run(t, new(IntegrationTestSuite))
// }

// func (s *IntegrationTestSuite) SetupTest() {
// 	app := gaiahelpers.Setup(s.T(), false, 1)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{
// 		ChainID: fmt.Sprintf("test-chain-%s", tmrand.Str(4)),
// 		Height:  1,
// 	})

// 	encodingConfig := app.MakeTestEncodingConfig()
// 	encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
// 	testdata.RegisterInterfaces(encodingConfig.InterfaceRegistry)

// 	s.app = app
// 	s.ctx = ctx
// 	s.clientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
// }

// func (s *IntegrationTestSuite) setupTestGlobalFeeStoreAndMinGasPrice(minGasPrice []sdk.DecCoin, globalFeeParams *globfeetypes.Params) types.Subspace {
// 	subspace := s.app.GetSubspace(globalfee.ModuleName)
// 	subspace.SetParamSet(s.ctx, globalFeeParams)
// 	s.ctx = s.ctx.WithMinGasPrices(minGasPrice).WithIsCheckTx(true)

// 	return subspace
// }

// func (s *IntegrationTestSuite) CreateTestTx(privs []cryptotypes.PrivKey, accNums []uint64, accSeqs []uint64, chainID string) (xauthsigning.Tx, error) {
// 	var sigsV2 []signing.SignatureV2
// 	for i, priv := range privs {
// 		sigV2 := signing.SignatureV2{
// 			PubKey: priv.PubKey(),
// 			Data: &signing.SingleSignatureData{
// 				SignMode:  s.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
// 				Signature: nil,
// 			},
// 			Sequence: accSeqs[i],
// 		}

// 		sigsV2 = append(sigsV2, sigV2)
// 	}

// 	if err := s.txBuilder.SetSignatures(sigsV2...); err != nil {
// 		return nil, err
// 	}

// 	sigsV2 = []signing.SignatureV2{}
// 	for i, priv := range privs {
// 		signerData := xauthsigning.SignerData{
// 			ChainID:       chainID,
// 			AccountNumber: accNums[i],
// 			Sequence:      accSeqs[i],
// 		}
// 		sigV2, err := tx.SignWithPrivKey(
// 			s.clientCtx.TxConfig.SignModeHandler().DefaultMode(),
// 			signerData,
// 			s.txBuilder,
// 			priv,
// 			s.clientCtx.TxConfig,
// 			accSeqs[i],
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		sigsV2 = append(sigsV2, sigV2)
// 	}

// 	if err := s.txBuilder.SetSignatures(sigsV2...); err != nil {
// 		return nil, err
// 	}

// 	return s.txBuilder.GetTx(), nil
// }
