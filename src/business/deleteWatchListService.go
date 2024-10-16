package business

import (
	"errors"
)

func DeleteWatchListService(id string) error {
	if _, exists := watchlists[id]; !exists {
		return errors.New("watchlist not found")
	}
	delete(watchlists, id)
	return nil
}
