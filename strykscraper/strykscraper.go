package strykscraper

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

func FetchData() ([]Bet, []Odds, error) {
	url := "https://spela.svenskaspel.se/stryktipset"
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var jsObject string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptText := s.Text()
		re := regexp.MustCompile(`_svs\.tipsen\.data\.preloadedState\s*=\s*({.*?});`)
		matches := re.FindStringSubmatch(scriptText)
		if len(matches) > 1 {
			jsObject = matches[1]
		}
	})

	if jsObject == "" {
		return nil, nil, fmt.Errorf("JavaScript object not found")
	}

	jsObject = strings.ReplaceAll(jsObject, `\,`, `,`)
	betEvents := gjson.Get(jsObject, "BetEvents")
	odds := gjson.Get(jsObject, "EventTypeStatistic")

	var Matches []string
	betEvents.ForEach(func(key, value gjson.Result) bool {
		Matches = append(Matches, key.String())
		return true
	})

	var MatchData []Bet
	for _, match := range Matches {
		for _, bet := range betEvents.Get(match).Array() {
			eventDesc := bet.Get("eventDescription").String()
			eventTypeStatId := bet.Get("eventTypeStatisticId").String()
			matchid := bet.Get("matchId").String()

			MatchData = append(MatchData, Bet{
				MatchId:              matchid,
				EventDescription:     eventDesc,
				EventTypeStatisticId: eventTypeStatId,
			})
		}
	}

	var oddsData []Odds
	for _, bet := range MatchData {
		oddsObj := odds.Get(bet.EventTypeStatisticId)

		var currentOdds []float64
		oddsObj.Get("odds.current.value").ForEach(func(key, value gjson.Result) bool {
			oddsValue := value.Float()
			currentOdds = append(currentOdds, oddsValue)
			return true
		})

		var distributions []int
		oddsObj.Get("distributions").ForEach(func(key, value gjson.Result) bool {
			value.ForEach(func(_, v gjson.Result) bool {
				v.Get("current.value").ForEach(func(_, vv gjson.Result) bool {
					distributionValue := int(vv.Int())
					distributions = append(distributions, distributionValue)
					return true
				})
				return false
			})
			return false
		})

		var favourites []float64
		oddsObj.Get("favourites.current.value").ForEach(func(key, value gjson.Result) bool {
			favouriteValue := value.Float()
			favourites = append(favourites, favouriteValue)
			return true
		})

		oddsData = append(oddsData, Odds{
			MatchId:      bet.MatchId,
			Odds:         currentOdds,
			Distribution: distributions,
			Favourites:   favourites,
		})
	}

	return MatchData, oddsData, nil
}
