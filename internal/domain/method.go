package domain

import (
	"math/rand"
	"sort"
	"time"

	"github.com/robstave/meowmorize/internal/domain/types"
)

// selectRandomCards selects random cards from the deck
func selectRandomCards(cards []types.Card, count int) []types.Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards[:count]
}

// selectFailsCards selects top N cards based on fail rate (percentage)
func selectFailsCards(cards []types.Card, count int) []types.Card {
	// Calculate fail rates
	sort.Slice(cards, func(i, j int) bool {
		return calculateFailRate(cards[i]) > calculateFailRate(cards[j])
	})
	if count > len(cards) {
		count = len(cards)
	}
	return cards[:count]
}

// selectSkipsCards selects top N cards based on skip rate (percentage)
func selectSkipsCards(cards []types.Card, count int) []types.Card {

	// Calculate skip rates
	sort.Slice(cards, func(i, j int) bool {
		return calculateSkipRate(cards[i]) > calculateSkipRate(cards[j])
	})
	if count > len(cards) {
		count = len(cards)
	}
	return cards[:count]
}

// selectWorstCards selects top N cards based on combined fail and skip rates
func selectWorstCards(cards []types.Card, count int) []types.Card {

	// Calculate combined fail and skip rates
	sort.Slice(cards, func(i, j int) bool {
		return calculateCombinedRate(cards[i]) > calculateCombinedRate(cards[j])
	})
	if count > len(cards) {
		count = len(cards)
	}
	return cards[:count]
}

// calculateFailRate computes the fail rate percentage for a card
func calculateFailRate(card types.Card) float64 {
	if card.PassCount == 0 {
		return 100.0 // If no successes, highest priority
	}
	return (float64(card.FailCount) / float64(card.PassCount)) * 100.0
}

// calculateSkipRate computes the skip rate percentage for a card
func calculateSkipRate(card types.Card) float64 {
	if card.PassCount == 0 {
		return 100.0 // If no successes, highest priority
	}
	return (float64(card.SkipCount) / float64(card.PassCount)) * 100.0
}

// calculateCombinedRate computes the combined fail and skip rate percentage for a card
func calculateCombinedRate(card types.Card) float64 {
	if card.PassCount == 0 {
		return 100.0 // If no successes, highest priority
	}
	return ((float64(card.FailCount) + float64(card.SkipCount)) / float64(card.PassCount)) * 100.0
}

// selectStarsCards selects top N cards based on star rating with some randomization
func selectStarsCards(cards []types.Card, count int) []types.Card {
	// Sort cards by StarRating descending
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].StarRating > cards[j].StarRating
	})

	// Group cards by StarRating
	starGroups := make(map[int][]types.Card)
	for _, card := range cards {
		starGroups[card.StarRating] = append(starGroups[card.StarRating], card)
	}

	// Get unique star ratings in descending order
	var starRatings []int
	for star := range starGroups {
		starRatings = append(starRatings, star)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(starRatings)))

	selectedCards := []types.Card{}
	remaining := count

	for _, star := range starRatings {
		group := starGroups[star]
		if remaining <= 0 {
			break
		}

		// Determine how many cards to pick from this group
		pick := remaining
		if len(group) < remaining {
			pick = len(group)
		}

		// Shuffle the group to add randomness
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(group), func(i, j int) { group[i], group[j] = group[j], group[i] })

		selectedCards = append(selectedCards, group[:pick]...)
		remaining -= pick
	}

	return selectedCards
}

// selectAdjustedRandomCards selects cards using the adjusted random strategy.
// For each card, it calculates a score as follows:
//   - If star rating is 0, treat it as 2.5; otherwise use the actual star rating
//   - Multiply the star rating by 2 (baseline score between 0 and 10)
//   - Add a random number between -2.5 and 2.5
//   - If FailCount > PassCount, add a random number between 0 and 2
//
// Then it sorts the cards ascending by the score and returns the top 'count' cards.
func selectAdjustedRandomCards(cards []types.Card, count int) []types.Card {
	type scoredCard struct {
		card  types.Card
		score float64
	}

	scoredCards := make([]scoredCard, len(cards))
	for i, card := range cards {
		var baseline float64
		if card.StarRating == 0 {
			baseline = 2.5 * 2 // equals 5
		} else {
			baseline = float64(card.StarRating) * 2
		}
		randomAdj := (rand.Float64() * 5) - 2.5 // random from -2.5 to 2.5
		extra := 0.0
		if card.FailCount > card.PassCount {
			extra = rand.Float64() * 2 // random from 0 to 2
		}
		score := baseline + randomAdj + extra
		scoredCards[i] = scoredCard{
			card:  card,
			score: score,
		}
	}

	// Sort by ascending score (lower score means higher priority)
	sort.Slice(scoredCards, func(i, j int) bool {
		return scoredCards[i].score > scoredCards[j].score
	})

	n := count
	if n > len(scoredCards) {
		n = len(scoredCards)
	}
	result := make([]types.Card, n)
	for i := 0; i < n; i++ {
		result[i] = scoredCards[i].card
	}
	return result
}

// selectUnratedCards selects top N unrated cards first, then randomize
func selectUnratedCards(cards []types.Card, count int) []types.Card {
	unratedCards := []types.Card{}
	ratedCards := []types.Card{}

	for _, card := range cards {
		if card.StarRating == 0 {
			unratedCards = append(unratedCards, card)
		} else {
			ratedCards = append(ratedCards, card)
		}
	}

	selectedCards := []types.Card{}

	// Add unrated cards first
	if len(unratedCards) >= count {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(unratedCards), func(i, j int) { unratedCards[i], unratedCards[j] = unratedCards[j], unratedCards[i] })
		selectedCards = append(selectedCards, unratedCards[:count]...)
		return selectedCards
	}

	selectedCards = append(selectedCards, unratedCards...)
	remaining := count - len(unratedCards)

	// Shuffle rated cards and add the remaining
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ratedCards), func(i, j int) { ratedCards[i], ratedCards[j] = ratedCards[j], ratedCards[i] })

	if remaining > len(ratedCards) {
		remaining = len(ratedCards)
	}

	selectedCards = append(selectedCards, ratedCards[:remaining]...)
	return selectedCards
}
