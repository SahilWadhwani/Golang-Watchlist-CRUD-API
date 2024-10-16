package business

import (
	"errors"
	"main/src/api/model"
)

func UpdateWatchListService(oldId string, watchlist *model.UpdateWatchList) (*model.UpdateWatchList, error) {

	if _, exists := watchlists[oldId]; !exists {

		return nil, errors.New("watchlist not found")
	}
	// Update the watchlist
	updatedWatchlist := model.CreateWatchList{
		Id:            watchlist.Id,
		WatchListName: watchlist.WatchListName,
		Scripts:       watchlist.Scripts,
	}
	watchlists[watchlist.Id] = updatedWatchlist

	if oldId != watchlist.Id {
		delete(watchlists, oldId)
	}

	return watchlist, nil
}
