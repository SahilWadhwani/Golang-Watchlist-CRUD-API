package business

import (
	"errors"
	"main/src/api/model"
)

func GetWatchListService(id string) (*model.CreateWatchList, error) {
	watchlist, exists := watchlists[id]
	if !exists {
		return nil, errors.New("watchlist not found")
	}
	return &watchlist, nil
}
