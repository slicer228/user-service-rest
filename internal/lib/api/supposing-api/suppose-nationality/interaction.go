package suppose_nationality

import (
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"user-service/internal/lib/api/client"
)

func RequestPredictedNationality(log *slog.Logger, name string) string {
	s := client.NewHttpClientSession()

	reqUrl := strings.Builder{}
	reqUrl.WriteString("https://api.nationalize.io/?name=")
	reqUrl.WriteString(name)

	resp, err := s.Get(reqUrl.String())

	if err != nil {
		log.Error("RequestPredictedNationality Get Error", "name", name, "err", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("RequestPredictedNationality Read Error", "name", name, "err", err)
		return ""
	}

	var encodedBody struct {
		CountryProbs []struct {
			CountryId   string  `json:"country_id"`
			Probability float64 `json:"probability"`
		} `json:"country"`
	}
	json.Unmarshal(body, &encodedBody)

	max := 0
	for i := 1; i < len(encodedBody.CountryProbs); i++ {
		if encodedBody.CountryProbs[max].Probability < encodedBody.CountryProbs[i].Probability {
			max = i
		}
	}

	return encodedBody.CountryProbs[max].CountryId
}
