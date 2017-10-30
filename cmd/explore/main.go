package main

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/carb/certainty/config"
	"github.com/carb/certainty/controllers"
	"github.com/carb/certainty/entities"
)

func main() {
	fx.New(
		config.Module,
		entities.Module,
		controllers.Module,
		fx.Invoke(ExploreBoard),
	).Run()
}

func ExploreBoard(lc fx.Lifecycle, board controllers.BoardController, logger *zap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			var err error

			// Fake Red turn
			err = board.PlaceHQ(entities.RedFaction, entities.Alaska)
			if err != nil {
				return err
			}

			// Fake Blue turn that throws an erro
			err = board.PlaceHQ(entities.BlueFaction, entities.Alberta)
			if err != nil {
				return err
			}

			// Fake Blue turn that is never executed
			err = board.PlaceHQ(entities.BlueFaction, entities.CentralAmerica)
			if err != nil {
				return err
			}

			return nil
		},
		OnStop: func(context.Context) error {
			logger.Sync()
			return nil
		},
	})
}
