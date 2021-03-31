package repos

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/lndaquino/scratchpay-challenge/pkg/entity"
	"github.com/lndaquino/scratchpay-challenge/utils"
)

// Seed populate the empty database
func (repo *PracticeInMemory) Seed() error {
	if len(repo.m) == 0 {
		log.Println("==> Seeding database...")

		seedStates()

		if err := seedDentalPractice(repo); err != nil {
			return err
		}

		if err := seedVetPractice(repo); err != nil {
			return err
		}

		for i, v := range repo.m {
			log.Printf("%s ==> %v", i, v)
		}
	} else {
		log.Println("==> Database already populated, continuing...")
	}

	log.Println("==> Seeding completed!")
	return nil
}

func seedDentalPractice(repo *PracticeInMemory) error {
	dentalURL := os.Getenv("DENTAL_PRACTICE_URL")
	resp, err := http.Get(dentalURL)
	if err != nil {
		return errors.New("Error retrieving information from " + dentalURL)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	dentalPractices := []dentalPractice{}
	err = json.Unmarshal(body, &dentalPractices)
	if err != nil {
		return errors.New("Cannot convert to json: " + err.Error())
	}

	for _, v := range dentalPractices {
		id := uuid.New().String()

		var stateCode string
		for key, state := range states {
			if state == v.StateName {
				stateCode = key
			}
		}

		openingHour, openingMinute, err := utils.GetHourAndMinute(v.Availability.From)
		if err != nil {
			return err
		}

		closureHour, closureMinute, err := utils.GetHourAndMinute(v.Availability.To)
		if err != nil {
			return err
		}

		practice := &entity.Practice{
			ID:            id,
			Name:          v.Name,
			Category:      "dental",
			StateName:     v.StateName,
			StateCode:     stateCode,
			OpeningHour:   *openingHour,
			OpeningMinute: *openingMinute,
			ClosureHour:   *closureHour,
			ClosureMinute: *closureMinute,
		}

		repo.m[id] = practice
	}

	return nil
}

func seedVetPractice(repo *PracticeInMemory) error {
	vetURL := os.Getenv("VET_PRACTICE_URL")
	resp, err := http.Get(vetURL)
	if err != nil {
		return errors.New("Error retrieving information from " + vetURL)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	vetPractices := []vetPractice{}
	err = json.Unmarshal(body, &vetPractices)
	if err != nil {
		return errors.New("Cannot convert to json: " + err.Error())
	}

	for _, v := range vetPractices {
		id := uuid.New().String()

		openingHour, openingMinute, err := utils.GetHourAndMinute(v.Opening.From)
		if err != nil {
			return err
		}

		closureHour, closureMinute, err := utils.GetHourAndMinute(v.Opening.To)
		if err != nil {
			return err
		}

		practice := &entity.Practice{
			ID:            id,
			Name:          v.ClinicName,
			Category:      "vet",
			StateName:     states[v.StateCode],
			StateCode:     v.StateCode,
			OpeningHour:   *openingHour,
			OpeningMinute: *openingMinute,
			ClosureHour:   *closureHour,
			ClosureMinute: *closureMinute,
		}
		repo.m[id] = practice
	}

	return nil
}

type opening struct {
	From string
	To   string
}

type vetPractice struct {
	ClinicName string `json:"clinicName"`
	StateCode  string `json:"stateCode"`
	Opening    opening
}

type dentalPractice struct {
	Name         string `json:"name"`
	StateName    string `json:"stateName"`
	Availability opening
}

func seedStates() {
	states = map[string]string{
		"AK": "Alaska",
		"AL": "Alabama",
		"AR": "Arkansas",
		"AZ": "Arizona",
		"CA": "California",
		"CO": "Colorado",
		"CT": "Connecticut",
		"DC": "	Washington DC",
		"DE": "Delaware",
		"FL": "Florida",
		"GA": "Georgia",
		"GU": "Guam",
		"HI": "Hawaii",
		"IA": "Iowa",
		"ID": "Idaho",
		"IL": "Illinois",
		"IN": "Indiana",
		"KS": "Kansas",
		"KY": "Kentucky",
		"LA": "Louisiana",
		"MA": "Massachusetts",
		"MD": "Maryland",
		"ME": "Maine",
		"MI": "Michigan",
		"MN": "Minnesota",
		"MO": "Missouri",
		"MS": "Mississippi",
		"MT": "Montana",
		"NC": "North Carolina",
		"ND": "North Dakota",
		"NE": "Nebraska",
		"NH": "New Hampshire",
		"NJ": "New Jersey",
		"NM": "New Mexico",
		"NV": "Nevada",
		"NY": "New York",
		"OH": "Ohio",
		"OK": "Oklahoma",
		"OR": "Oregon",
		"PA": "Pennsylvania",
		"PR": "Puerto Rico",
		"RI": "Rhode Island",
		"SC": "South Carolina",
		"SD": "South Dakota",
		"TN": "Tennessee",
		"TX": "Texas",
		"UT": "Utah",
		"VA": "Virginia",
		"VI": "Virgin Islands",
		"VT": "Vermont",
		"WA": "Washington",
		"WI": "Wisconsin",
		"WV": "West Virginia",
		"WY": "Wyoming",
	}
}
