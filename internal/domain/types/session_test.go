package types // Adjust the package name as needed

import (
	"sync"
	"testing"
)

// Helper function to create a sample session
func createSampleSession() Session {
	return Session{
		DeckID: "test-deck",
		CardStats: []CardStats{
			{CardID: "card1"},
			{CardID: "card2"},
			{CardID: "card3"},
		},
		Method: "rando",
		Index:  0,
		mu:     sync.Mutex{},
		Stats: SessionStats{
			TotalCards:   3,
			ViewedCount:  0,
			Remaining:    3,
			CurrentIndex: 0,
		},
	}
}

func TestGetNextCard(t *testing.T) {
	t.Run("Normal Sequence", func(t *testing.T) {
		session := createSampleSession()

		expected := []string{"card1", "card2", "card3"}
		for i, exp := range expected {
			got := session.GetNextCard()
			if got != exp {
				t.Errorf("Iteration %d: expected %s, got %s", i, exp, got)
			}
		}
	})

	t.Run("Wrap Around", func(t *testing.T) {
		session := createSampleSession()

		// Exhaust all cards
		for i := 0; i < 3; i++ {
			session.GetNextCard()
		}

		// Should wrap around to the first card
		got := session.GetNextCard()
		if got != "card1" {
			t.Errorf("Expected card1 after wrap around, got %s", got)
		}
	})

	t.Run("Stats Update", func(t *testing.T) {
		session := createSampleSession()

		session.GetNextCard()

		if session.Stats.ViewedCount != 1 {
			t.Errorf("Expected ViewedCount to be 1, got %d", session.Stats.ViewedCount)
		}
		if session.Stats.Remaining != 2 {
			t.Errorf("Expected Remaining to be 2, got %d", session.Stats.Remaining)
		}
		if session.Stats.CurrentIndex != 1 {
			t.Errorf("Expected CurrentIndex to be 1, got %d", session.Stats.CurrentIndex)
		}
	})

	t.Run("Empty CardStats", func(t *testing.T) {
		session := &Session{
			CardStats: []CardStats{},
			mu:        sync.Mutex{},
		}

		got := session.GetNextCard()
		if got != "" {
			t.Errorf("Expected empty string for empty CardStats, got %s", got)
		}
	})

	t.Run("Concurrent Access", func(t *testing.T) {
		session := createSampleSession()

		const numGoroutines = 100
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				session.GetNextCard()
			}()
		}

		wg.Wait()

		if session.Stats.ViewedCount != numGoroutines {
			t.Errorf("Expected ViewedCount to be %d, got %d", numGoroutines, session.Stats.ViewedCount)
		}
	})
}

// Helper function to create a sample session
func createSampleSession2() *Session {
	return &Session{
		DeckID: "test-deck",
		CardStats: []CardStats{
			{CardID: "card1", Viewed: false},
			{CardID: "card2", Viewed: false},
			{CardID: "card3", Viewed: false},
			{CardID: "card4", Viewed: false},
		},
		Method: SessionMethod("some-method"),
		Index:  0,
		mu:     sync.Mutex{},
	}
}

func TestGetSessionStats(t *testing.T) {
	t.Run("Initial State", func(t *testing.T) {
		session := createSampleSession2()
		stats := session.GetSessionStats()

		if stats.TotalCards != 4 {
			t.Errorf("Expected TotalCards to be 4, got %d", stats.TotalCards)
		}
		if stats.ViewedCount != 0 {
			t.Errorf("Expected ViewedCount to be 0, got %d", stats.ViewedCount)
		}
		if stats.Remaining != 4 {
			t.Errorf("Expected Remaining to be 4, got %d", stats.Remaining)
		}
		if stats.CurrentIndex != 0 {
			t.Errorf("Expected CurrentIndex to be 0, got %d", stats.CurrentIndex)
		}
	})

	t.Run("After Viewing Some Cards", func(t *testing.T) {
		session := createSampleSession2()
		session.CardStats[0].Viewed = true
		session.CardStats[2].Viewed = true
		session.Index = 3

		stats := session.GetSessionStats()

		if stats.TotalCards != 4 {
			t.Errorf("Expected TotalCards to be 4, got %d", stats.TotalCards)
		}
		if stats.ViewedCount != 2 {
			t.Errorf("Expected ViewedCount to be 2, got %d", stats.ViewedCount)
		}
		if stats.Remaining != 2 {
			t.Errorf("Expected Remaining to be 2, got %d", stats.Remaining)
		}
		if stats.CurrentIndex != 3 {
			t.Errorf("Expected CurrentIndex to be 3, got %d", stats.CurrentIndex)
		}
	})

	t.Run("All Cards Viewed", func(t *testing.T) {
		session := createSampleSession2()
		for i := range session.CardStats {
			session.CardStats[i].Viewed = true
		}
		session.Index = 4

		stats := session.GetSessionStats()

		if stats.TotalCards != 4 {
			t.Errorf("Expected TotalCards to be 4, got %d", stats.TotalCards)
		}
		if stats.ViewedCount != 4 {
			t.Errorf("Expected ViewedCount to be 4, got %d", stats.ViewedCount)
		}
		if stats.Remaining != 0 {
			t.Errorf("Expected Remaining to be 0, got %d", stats.Remaining)
		}
		if stats.CurrentIndex != 4 {
			t.Errorf("Expected CurrentIndex to be 4, got %d", stats.CurrentIndex)
		}
	})

	t.Run("Empty Session", func(t *testing.T) {
		session := &Session{
			CardStats: []CardStats{},
			mu:        sync.Mutex{},
		}

		stats := session.GetSessionStats()

		if stats.TotalCards != 0 {
			t.Errorf("Expected TotalCards to be 0, got %d", stats.TotalCards)
		}
		if stats.ViewedCount != 0 {
			t.Errorf("Expected ViewedCount to be 0, got %d", stats.ViewedCount)
		}
		if stats.Remaining != 0 {
			t.Errorf("Expected Remaining to be 0, got %d", stats.Remaining)
		}
		if stats.CurrentIndex != 0 {
			t.Errorf("Expected CurrentIndex to be 0, got %d", stats.CurrentIndex)
		}
	})

	t.Run("Concurrent Access", func(t *testing.T) {
		session := createSampleSession()
		const numGoroutines = 100
		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				defer wg.Done()
				session.GetSessionStats()
			}()
		}

		wg.Wait()
		// If this test completes without deadlock or race conditions, it passes
	})
}
