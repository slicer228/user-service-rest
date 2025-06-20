package suppose_age

import (
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"user-service/internal/lib/api/client"
)

func RequestPredictedAge(log *slog.Logger, name string) int {
	s := client.NewHttpClientSession()

	reqUrl := strings.Builder{}
	reqUrl.WriteString("https://api.agify.io/?name=")
	reqUrl.WriteString(name)

	resp, err := s.Get(reqUrl.String())

	if err != nil {
		log.Error("RequestPredictedAge Get Error", "name", name, "err", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("RequestPredictedAge Read Error", "name", name, "err", err)
		return 0
	}

	var encodedBody struct {
		Age int `json:"age"`
	}
	json.Unmarshal(body, &encodedBody)

	return encodedBody.Age
}
