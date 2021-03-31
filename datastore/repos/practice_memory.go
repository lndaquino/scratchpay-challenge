package repos

import (
	"strings"

	"github.com/lndaquino/scratchpay-challenge/pkg/entity"
	er "github.com/lndaquino/scratchpay-challenge/pkg/errors"
)

var states map[string]string

// PracticeInMemory struct models a practice repository
type PracticeInMemory struct {
	m map[string]*entity.Practice
}

// NewPracticeInMemoryRepo returns a PracticeInMemory instance
func NewPracticeInMemoryRepo() *PracticeInMemory {
	return &PracticeInMemory{map[string]*entity.Practice{}}
}

// SearchByName searchs practices by its name
func (repo *PracticeInMemory) SearchByName(name string) (practiceList []*entity.Practice, err error) {
	for _, practice := range repo.m {
		if strings.Contains(strings.ToLower(practice.Name), name) {
			practiceList = append(practiceList, practice)
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}

// SearchByState searchs practices by its state name or code
func (repo *PracticeInMemory) SearchByState(state string) (practiceList []*entity.Practice, err error) {
	stateCode := strings.ToUpper(state)
	stateName := strings.ToLower(state)

	for _, practice := range repo.m {
		if strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode) {
			practiceList = append(practiceList, practice)
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}

// SearchByNameAndState searchs practices by its name, state name and state code
func (repo *PracticeInMemory) SearchByNameAndState(name, state string) (practiceList []*entity.Practice, err error) {
	stateCode := strings.ToUpper(state)
	stateName := strings.ToLower(state)

	for _, practice := range repo.m {
		if (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) && strings.Contains(strings.ToLower(practice.Name), name) {
			practiceList = append(practiceList, practice)
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}

// SearchByAvailability searchs practices by availability
func (repo *PracticeInMemory) SearchByAvailability(fromHour, fromMinute, toHour, toMinute *int) (practiceList []*entity.Practice, err error) {
	if fromHour != nil && toHour == nil {
		for _, practice := range repo.m {
			if (*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*fromHour < practice.ClosureHour || (*fromHour == practice.ClosureHour && *fromMinute < practice.ClosureMinute)) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if toHour != nil && fromHour == nil {
		for _, practice := range repo.m {
			if (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute)) && (*toHour > practice.OpeningHour || (*toHour == practice.OpeningHour && *toMinute > practice.OpeningMinute)) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if fromHour != nil && toHour != nil {
		for _, practice := range repo.m {
			if (*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute)) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}

// SearchByNameAndAvailability searchs practices by name and availability
func (repo *PracticeInMemory) SearchByNameAndAvailability(name string, fromHour, fromMinute, toHour, toMinute *int) (practiceList []*entity.Practice, err error) {

	if fromHour != nil && toHour == nil {
		for _, practice := range repo.m {
			if (*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*fromHour < practice.ClosureHour || (*fromHour == practice.ClosureHour && *fromMinute < practice.ClosureMinute)) && strings.Contains(strings.ToLower(practice.Name), name) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if toHour != nil && fromHour == nil {
		for _, practice := range repo.m {
			if (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute)) && (*toHour > practice.OpeningHour || (*toHour == practice.OpeningHour && *toMinute > practice.OpeningMinute)) && strings.Contains(strings.ToLower(practice.Name), name) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if fromHour != nil && toHour != nil {
		for _, practice := range repo.m {
			if ((*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute))) && strings.Contains(strings.ToLower(practice.Name), name) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}

// SearchByStateAndAvailability searchs practices by state name, state code and availability
func (repo *PracticeInMemory) SearchByStateAndAvailability(state string, fromHour, fromMinute, toHour, toMinute *int) (practiceList []*entity.Practice, err error) {
	stateCode := strings.ToUpper(state)
	stateName := strings.ToLower(state)

	if fromHour != nil && toHour == nil {
		for _, practice := range repo.m {
			if (*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*fromHour < practice.ClosureHour || (*fromHour == practice.ClosureHour && *fromMinute < practice.ClosureMinute)) && (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if toHour != nil && fromHour == nil {
		for _, practice := range repo.m {
			if (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute)) && (*toHour > practice.OpeningHour || (*toHour == practice.OpeningHour && *toMinute > practice.OpeningMinute)) && (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if fromHour != nil && toHour != nil {
		for _, practice := range repo.m {
			if ((*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute))) && (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}

// SearchByNameAndStateAndAvailability searchs practices by name, state name, state code and availability
func (repo *PracticeInMemory) SearchByNameAndStateAndAvailability(name, state string, fromHour, fromMinute, toHour, toMinute *int) (practiceList []*entity.Practice, err error) {
	stateCode := strings.ToUpper(state)
	stateName := strings.ToLower(state)

	if fromHour != nil && toHour == nil {
		for _, practice := range repo.m {
			if (*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*fromHour < practice.ClosureHour || (*fromHour == practice.ClosureHour && *fromMinute < practice.ClosureMinute)) && (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) && strings.Contains(strings.ToLower(practice.Name), name) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if toHour != nil && fromHour == nil {
		for _, practice := range repo.m {
			if (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute)) && (*toHour > practice.OpeningHour || (*toHour == practice.OpeningHour && *toMinute > practice.OpeningMinute)) && (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) && strings.Contains(strings.ToLower(practice.Name), name) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if fromHour != nil && toHour != nil {
		for _, practice := range repo.m {
			if ((*fromHour > practice.OpeningHour || (*fromHour == practice.OpeningHour && *fromMinute >= practice.OpeningMinute)) && (*toHour < practice.ClosureHour || (*toHour == practice.ClosureHour && *toMinute <= practice.ClosureMinute))) && (strings.Contains(strings.ToLower(practice.StateName), stateName) || strings.Contains(practice.StateCode, stateCode)) && strings.Contains(strings.ToLower(practice.Name), name) {
				practiceList = append(practiceList, practice)
			}
		}
	}

	if len(practiceList) == 0 {
		return nil, er.ErrNotFound
	}

	return practiceList, nil
}
