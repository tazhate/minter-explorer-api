package core

import (
	"github.com/MinterTeam/minter-explorer-api/address"
	"github.com/MinterTeam/minter-explorer-api/balance"
	"github.com/MinterTeam/minter-explorer-api/blocks"
	"github.com/MinterTeam/minter-explorer-api/coingecko"
	"github.com/MinterTeam/minter-explorer-api/coins"
	"github.com/MinterTeam/minter-explorer-api/invalid_transaction"
	"github.com/MinterTeam/minter-explorer-api/reward"
	"github.com/MinterTeam/minter-explorer-api/slash"
	"github.com/MinterTeam/minter-explorer-api/stake"
	"github.com/MinterTeam/minter-explorer-api/tools/cache"
	"github.com/MinterTeam/minter-explorer-api/tools/market"
	"github.com/MinterTeam/minter-explorer-api/transaction"
	"github.com/MinterTeam/minter-explorer-api/validator"
	"github.com/go-pg/pg"
)

type Explorer struct {
	CoinRepository               coins.Repository
	BlockRepository              blocks.Repository
	AddressRepository            address.Repository
	TransactionRepository        transaction.Repository
	InvalidTransactionRepository invalid_transaction.Repository
	RewardRepository             reward.Repository
	SlashRepository              slash.Repository
	ValidatorRepository          validator.Repository
	StakeRepository              stake.Repository
	Environment                  Environment
	Cache                        *cache.ExplorerCache
	MarketService                *market.Service
	BalanceService               *balance.Service
}

func NewExplorer(db *pg.DB, env *Environment) *Explorer {
	marketService := market.NewService(coingecko.NewService(env.MarketHost), env.BaseCoin)
	blockRepository := *blocks.NewRepository(db)

	return &Explorer{
		BlockRepository:              blockRepository,
		CoinRepository:               *coins.NewRepository(db, env.BaseCoin),
		AddressRepository:            *address.NewRepository(db),
		TransactionRepository:        *transaction.NewRepository(db),
		InvalidTransactionRepository: *invalid_transaction.NewRepository(db),
		RewardRepository:             *reward.NewRepository(db),
		SlashRepository:              *slash.NewRepository(db),
		ValidatorRepository:          *validator.NewRepository(db),
		StakeRepository:              *stake.NewRepository(db),
		Environment:                  *env,
		Cache:                        cache.NewCache(blockRepository.GetLastBlock()),
		MarketService:                marketService,
		BalanceService:               balance.NewService(env.BaseCoin, marketService),
	}
}
