package practice

import (
	"github.com/lndaquino/scratchpay-challenge/pkg/entity"
	er "github.com/lndaquino/scratchpay-challenge/pkg/errors"
	"github.com/lndaquino/scratchpay-challenge/utils"
)

// PracticeUseCase struct models a practice usecase
type PracticeUseCase struct {
	repo PracticeRepo
}

// PracticeRepo interface defines the methods a practice repository must have
type PracticeRepo interface {
	Seed() error
	SearchByName(string) ([]*entity.Practice, error)
	SearchByState(string) ([]*entity.Practice, error)
	SearchByNameAndState(string, string) ([]*entity.Practice, error)
	SearchByAvailability(*int, *int, *int, *int) ([]*entity.Practice, error)
	SearchByNameAndAvailability(string, *int, *int, *int, *int) ([]*entity.Practice, error)
	SearchByStateAndAvailability(string, *int, *int, *int, *int) ([]*entity.Practice, error)
	SearchByNameAndStateAndAvailability(string, string, *int, *int, *int, *int) ([]*entity.Practice, error)
}

// NewPracticeUseCase returns a PracticeUseCase instance
func NewPracticeUseCase(repo PracticeRepo) *PracticeUseCase {
	return &PracticeUseCase{repo: repo}
}

// Search practice usecase method handles business rules when retrieving practices
func (usecase *PracticeUseCase) Search(name, state, from, to string) (list []*entity.Practice, err error) {
	var fromHour, fromMinute, toHour, toMinute *int
	hasFrom := false
	hasTo := false
	if from != "" {
		hasFrom = true
		fromHour, fromMinute, err = utils.GetHourAndMinute(from)
		if err != nil {
			return nil, er.ErrInvalidFrom
		}
	}

	if to != "" {
		hasTo = true
		toHour, toMinute, err = utils.GetHourAndMinute(to)
		if err != nil {
			return nil, er.ErrInvalidTo
		}
	}

	if hasFrom && hasTo {
		if *fromHour > *toHour || (*fromHour == *toHour && *fromMinute >= *toMinute) {
			return nil, er.ErrInvalidParameters
		}
	}

	if name != "" && state == "" && !hasFrom && !hasTo {
		list, err = usecase.repo.SearchByName(name)
	}

	if state != "" && name == "" && !hasFrom && !hasTo {
		list, err = usecase.repo.SearchByState(state)
	}

	if name != "" && state != "" && !hasFrom && !hasTo {
		list, err = usecase.repo.SearchByNameAndState(name, state)
	}

	if (hasFrom || hasTo) && name == "" && state == "" {
		list, err = usecase.repo.SearchByAvailability(fromHour, fromMinute, toHour, toMinute)
	}

	if name != "" && (hasFrom || hasTo) && state == "" {
		list, err = usecase.repo.SearchByNameAndAvailability(name, fromHour, fromMinute, toHour, toMinute)
	}

	if state != "" && (hasFrom || hasTo) && name == "" {
		list, err = usecase.repo.SearchByStateAndAvailability(state, fromHour, fromMinute, toHour, toMinute)
	}

	if name != "" && state != "" && (hasFrom || hasTo) {
		list, err = usecase.repo.SearchByNameAndStateAndAvailability(name, state, fromHour, fromMinute, toHour, toMinute)
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}
