package usecase

import (
	"context"
	"time"

	"github.com/petruspierre/cartola-consolidation/internal/domain/entity"
	"github.com/petruspierre/cartola-consolidation/internal/domain/repository"
	"github.com/petruspierre/cartola-consolidation/pkg/uow"
)

type AddMatchInput struct {
	ID      string
	Date    time.Time
	TeamAID string
	TeamBID string
}

type AddMatchUseCase struct {
	Uow uow.UowInterface
}

func (a *AddMatchUseCase) Execute(ctx context.Context, input AddMatchInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		matchRepository := a.getMatchRepository(ctx)
		teamRepository := a.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}

		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)
		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddMatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")

	if err != nil {
		panic(err)
	}

	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *AddMatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := a.Uow.GetRepository(ctx, "TeamRepository")

	if err != nil {
		panic(err)
	}

	return teamRepository.(repository.TeamRepositoryInterface)
}
