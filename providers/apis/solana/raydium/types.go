package raydium

import (
	"encoding/json"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/skip-mev/slinky/providers/apis/solana/types/raydium"
)

var (
	raydiumDexProgramAccount, _ = solana.PublicKeyFromBase58("675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8")
)

const (
	Name = "raydium_api"
	coinVaultSeed = "coin_vault_associated_seed"
	pcVaultSeed = "pc_vault_associated_seed"
	ammInfoAssociatedSeed = "amm_associated_seed"
)

// PoolAccountDataPerTicker represents the account-data necessary per account, as they pertain
// to the derivation of a price for a given market via raydium.
type PoolAccountDataPerTicker struct {
	// AmmInfo is the account data for the raydium AmmInfo account
	AmmInfo *raydium.AmmInfo

	// QuoteTokenAccount is the token account for the quote token
	QuoteTokenAccount *token.Account

	// BaseTokenAccount is the token account for the base token
	BaseTokenAccount *token.Account
}

// PoolAssociatedAccounts represents the set of solana addresses that should be queried
// when determining the price of a given BASE/QUOTE pair, according to that market's raydium
// pool state
type PoolAssociatedAccounts struct {
	// PoolAccountInfo is the address of the account storing the raydium AmmInfo for the pool
	PoolAccountInfo solana.PublicKey

	// QuoteTokenPoolVaultAccount is the address of the account storing the quote token pool vault
	QuoteTokenPoolVaultAccount solana.PublicKey

	// BaseTokenPoolVaultAccount is the address of the account storing the base token pool vault
	BaseTokenPoolVaultAccount solana.PublicKey
}

// RaydiumTickerMetadata represents the metadata associated with a ticker's corresponding
// raydium pool. Specifically, we require the open-book dex's market-account address for
// the market, so that we can derive program derived addresses for the pool's associated
// accounts.
type RaydiumTickerMetadata struct {
	// MarketAccount is the base58 encoded address of the open-book dex's market account
	MarketAccount string `json:"market_account"`

	// QuoteTokenAddress is the base58 encoded address of the serum token corresponding
	// to this market's quote address
	QuoteTokenAddress string `json:"quote_token_address"`

	// BaseTokenAddress is the base58 encoded address of the serum token corresponding
	// to this market's base address
	BaseTokenAddress string `json:"base_token_address"`
}

// unmarshalMetadataJSON unmarshals the given metadata string into a RaydiumTickerMetadata,
// this method assumes that the metadata string is valid json, otherwise an error is returned.
func unmarshalMetadataJSON(metadata string) (RaydiumTickerMetadata, error) {
	// unmarshal the metadata string into a RaydiumTickerMetadata
	var tickerMetadata RaydiumTickerMetadata
	if err := json.Unmarshal([]byte(metadata), &tickerMetadata); err != nil {
		return RaydiumTickerMetadata{}, err
	}

	return tickerMetadata, nil
}

// derivePoolAssociatedAccounts derives the set of solana addresses that should be queried for a given pool for a ticker. The program 
// account derivations are performed in accordance with the raydium [spec](https://github.com/raydium-io/raydium-docs/blob/master/dev-resources/raydium-hybrid-amm-dev-doc.pdf)
func derivePoolAssociatedAccounts(raydiumDEXProgramAccount solana.PublicKey, tickerMetadata RaydiumTickerMetadata) (PoolAssociatedAccounts, error) {
	// derive the marketAccount
	marketAccount, err := solana.PublicKeyFromBase58(tickerMetadata.MarketAccount)
	if err != nil {
		return PoolAssociatedAccounts{}, err
	}

	// derive the poolAccountInfo (ignore the bump seed)
	poolAccountInfo, _, err := solana.FindProgramAddress(
		[][]byte{
			raydiumDEXProgramAccount[:],
			marketAccount[:],
			[]byte(ammInfoAssociatedSeed),
		},
		raydiumDEXProgramAccount,
	)
	if err != nil {
		return PoolAssociatedAccounts{}, err
	}

	// derive the quoteTokenPoolVaultAccount (ignore the bump seed)
	quoteTokenPoolVaultAccount, _, err := solana.FindProgramAddress(
		[][]byte{
			raydiumDEXProgramAccount[:],
			marketAccount[:],
			[]byte(pcVaultSeed),
		},
		raydiumDEXProgramAccount,
	)
	if err != nil {
		return PoolAssociatedAccounts{}, err
	}
	
	// derive the baseTokenPoolVaultAccount (ignore the bump seed)
	baseTokenPoolVaultAccount, _, err := solana.FindProgramAddress(
		[][]byte{
			raydiumDEXProgramAccount[:],
			marketAccount[:],
			[]byte(coinVaultSeed),
		},
		raydiumDEXProgramAccount,
	)
	if err != nil {
		return PoolAssociatedAccounts{}, err
	}

	return PoolAssociatedAccounts{
		PoolAccountInfo: poolAccountInfo,
		QuoteTokenPoolVaultAccount: quoteTokenPoolVaultAccount,
		BaseTokenPoolVaultAccount: baseTokenPoolVaultAccount,
	}, nil
}
