package strykscraper

// Odds holds the odds and related information for a match.
type Odds struct {
	MatchId      string
	Odds         []float64
	Distribution []int
	Favourites   []float64
}

// Bet holds information about the betting event.
type Bet struct {
	MatchId              string
	EventDescription     string
	EventTypeStatisticId string
}
