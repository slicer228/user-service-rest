package suppose_gender

import (
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"user-service/internal/lib/api/client"
)

func RequestPredictedGender(log *slog.Logger, name string) string {
	s := client.NewHttpClientSession(log)

	reqUrl := strings.Builder{}
	reqUrl.WriteString("https://api.genderize.io/?name=")
	reqUrl.WriteString(name)

	log.Debug("Requesting predicted gender", "url", reqUrl.String())
	resp, err := s.Get(reqUrl.String())

	if err != nil {
		log.Error("RequestPredictedGender Get Error", "name", name, "err", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("RequestPredictedGender Read Error", "name", name, "err", err)
		return ""
	}

	var encodedBody struct {
		Gender string `json:"gender"`
	}
	json.Unmarshal(body, &encodedBody)

	return encodedBody.Gender
}
