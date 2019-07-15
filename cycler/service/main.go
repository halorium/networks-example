package main

import (
	"strconv"
	"time"

	"github.com/halorium/networks-example/cycler/cycler"
	"github.com/halorium/networks-example/entities"
	"github.com/halorium/networks-example/env-vars"
	"github.com/halorium/networks-example/flaw"
	"github.com/halorium/networks-example/logger"
	"github.com/halorium/networks-example/networks/network"
	"github.com/halorium/networks-example/networks/network-account"
)

func main() {
	URICollection := map[string]bool{}

	// get networks every 5 minutes
	for {
		networkEntity := network.NewEntity()
		networkCollection, flawError := entities.GetEntitiesFromPostgres(networkEntity)

		if flawError != nil {
			logger.Panic("cycler-get-networks", flawError)
		}

		for _, networkEntity := range networkCollection {
			networkEntity, ok := networkEntity.(*network.Entity)

			if !ok {
				logger.Panic("cycler-get-networks", flaw.New("invalid type conversion"))
			}

			networkAccountEntity := networkaccount.NewEntity()
			networkAccountEntity.Head.NetworkID = networkEntity.Head.NetworkID

			networkAccountCollection, flawError := entities.GetEntitiesFromPostgres(networkAccountEntity)

			if flawError != nil {
				logger.Panic("cycler-get-network-accounts", flawError)
			}

			daysToLoad, err := strconv.Atoi(networkEntity.Body.Variations.DaysToLoad)

			if err != nil {
				logger.Critical("cycler-main", flaw.From(err).Wrap("cannot cycle"))
				return
			}

			if daysToLoad <= 0 {
				continue
			}

			daysPerRequest, err := strconv.Atoi(networkEntity.Body.Variations.DaysPerRequest)

			if err != nil {
				logger.Critical("cycler-main", flaw.From(err).Wrap("cannot cycle"))
				return
			}

			sleepDuration, flawError := cycler.CalculateSleep(
				networkEntity.Body.Variations.RateQuantity,
				networkEntity.Body.Variations.RateInterval,
			)

			if flawError != nil {
				logger.Panic("cycler-calculate-sleep", flawError)
			}

			for _, networkAccountEntity := range networkAccountCollection {
				networkAccountEntity, ok := networkAccountEntity.(*networkaccount.Entity)

				if !ok {
					logger.Panic("cycler-get-network-accounts", flaw.New("invalid type conversion"))
				}

				// Do not load Commission Junction SE
				if networkAccountEntity.Head.NetworkID == "9" && networkAccountEntity.Head.NetworkAccountID == "128" {
					continue
				}

				cycle := &cycler.Cycle{
					DaysToLoad:     daysToLoad,
					DaysPerRequest: daysPerRequest,
					Period:         "8h",
					Action:         cycler.Iterate,
					Tag:            "cycler-post-network-account-purchases-load-" + networkEntity.Body.Variant,
					URI:            envvars.NetworksURI + networkEntity.Head.NetworkID + "/accounts/" + networkAccountEntity.Head.NetworkAccountID + "/purchases/load",
					Sleep:          sleepDuration,
					RateInterval:   networkEntity.Body.Variations.RateInterval,
					RateQuantity:   networkEntity.Body.Variations.RateQuantity,
				}

				// if the URI is in the list then it's already running
				if URICollection[cycle.URI] {
					continue
				}

				// if the URI is not in the list we go run it and add it to the list
				go cycle.Run()

				URICollection[cycle.URI] = true
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
