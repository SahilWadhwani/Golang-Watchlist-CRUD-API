package business

import (
	"fmt"
	"main/src/api/model"
)

var watchlists = make(map[string]model.CreateWatchList)

func CreateWatchListService(watchlist *model.CreateWatchList) {
	watchlists[watchlist.Id] = *watchlist
	fmt.Println("watchlist created", watchlist)
}
