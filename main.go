package main

import (
	"github.com/MinterTeam/minter-explorer-api/api"
	"github.com/MinterTeam/minter-explorer-api/core"
	"github.com/MinterTeam/minter-explorer-api/database"
	"github.com/MinterTeam/minter-explorer-api/tools/metrics"
)

func main() {
	// init environment
	env := core.NewEnvironment()

	// connect to database
	db := database.Connect(env)
	defer database.Close(db)

	// create explorer
	explorer := core.NewExplorer(db, env)

	// run market price update
	go explorer.MarketService.Run()

	// create ws extender
	extender := core.NewExtenderWsClient(explorer)
	defer extender.Close()

	// subscribe to channel and add cache handler
	sub := extender.CreateSubscription(explorer.Environment.WsBlocksChannel)
	sub.OnPublish(explorer.Cache)
	sub.OnPublish(metrics.NewLastBlockMetric())
	extender.Subscribe(sub)

	// run api
	api.Run(db, explorer)
}
