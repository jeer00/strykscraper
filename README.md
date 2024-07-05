# Strykscraper

Strykscraper is a Go package designed to fetch and parse data from Stryktipset. It extracts betting odds (from "oddset"), percentage distrubtions from the players, aswell as the percentage equvalent of the odds.

Keep in mind this is just scraping data from the website, and could break at any time.

## Installation

To install the package, run:

```bash
go get "github.com/jeer00/strykscraper/strykscraper"
```

## Usage

To use the package, import it in your Go code:

```go
import "github.com/jeer00/strykscraper/strykscraper"
```

Then, use the `FetchData` function to fetch the data:

```go
odds, bets, err := strykscraper.FetchData()
```

The function returns three values:

- `odds`: A slice of `Odds` structs, containing the betting odds and related information.
- `bets`: A slice of `Bet` structs, containing information about the betting events.
- `err`: An error object, if any error occurred during the fetching process.

## Functions 

```go
func FetchData() ([]Bet, []Odds, error)
```

## Types
```go
type Bet struct {
    MatchId              string
    EventDescription     string
    EventTypeStatisticId string
}

type Odds struct {
    MatchId      string
    Odds         []float64
    Distribution []int
    Favourites   []float64
}
```


