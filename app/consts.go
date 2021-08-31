package app

type ThingType string

const (
	ThingBoardGame          ThingType = "boardgame"
	ThingBoardGameExpansion ThingType = "boardgameexpansion"
	ThingBoardGameAccessory ThingType = "boardgameaccessory"
	ThingVideoGame          ThingType = "videogame"
	ThingRPGItem            ThingType = "rpgitem"
	ThingRPGIssue           ThingType = "rpgissue"
)

const (
	BGGBase = "https://www.boardgamegeek.com/xmlapi2/"
)
