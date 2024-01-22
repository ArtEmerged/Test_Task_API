package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ageApi         = "https://api.agify.io/"
	genderApi      = "https://api.genderize.io/"
	nationalizeApi = "https://api.nationalize.io/"
)

type ageEnrichment struct {
	Count uint  `json:"count"`
	Age   uint8 `json:"age"`
}

type genderEnrichment struct {
	Count  uint   `json:"count"`
	Gender string `json:"gender"`
}

type nationalizeEnrichment struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"-"`
	} `json:"country"`
}

func EnrichmentOfDataOnAge(name string) (uint8, error) {
	query := fmt.Sprintf("%s?name=%s", ageApi, name)
	resp, err := http.Get(query)
	if err != nil {
		return 0, err
	}

	var data ageEnrichment

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	if data.Count == 0 {
		return 0, nil
	}
	return data.Age, nil
}

func EnrichingDataOnGender(name string) (string, error) {
	query := fmt.Sprintf("%s?name=%s", genderApi, name)
	resp, err := http.Get(query)
	if err != nil {
		return "", err
	}

	var data genderEnrichment

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	if data.Count == 0 {
		return "", nil
	}

	return data.Gender, nil

}

func EnrichmentOfDataOnNationality(name string) ([]string, error) {
	query := fmt.Sprintf("%s?name=%s", nationalizeApi, name)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var data nationalizeEnrichment

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Count == 0 {
		return nil, nil
	}

	countryID := make([]string, len(data.Country))
	for ind, counrty := range data.Country {
		countryID[ind] = counrty.CountryID
	}

	return countryID, nil

}
