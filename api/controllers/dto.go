package controllers

import (
	"fmt"

	"github.com/lndaquino/scratchpay-challenge/pkg/entity"
)

type opening struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type searchResponse struct {
	ID        string  `json:"id"`
	Category  string  `json:"category"`
	Name      string  `json:"name"`
	StateName string  `json:"stateName"`
	StateCode string  `json:"stateCode"`
	Opening   opening `json:"opening"`
}

func parseSearchResponse(practiceList []*entity.Practice) (resp []*searchResponse) {
	for _, practice := range practiceList {
		r := new(searchResponse)
		r.ID = practice.ID
		r.Category = practice.Category
		r.Name = practice.Name
		r.StateName = practice.StateName
		r.StateCode = practice.StateCode
		r.Opening.From = fmt.Sprintf("%02d:%02d", practice.OpeningHour, practice.OpeningMinute)
		r.Opening.To = fmt.Sprintf("%02d:%02d", practice.ClosureHour, practice.ClosureMinute)
		resp = append(resp, r)
	}

	return
}
