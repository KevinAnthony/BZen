package model

import "time"

type BoardgameType string

const (
	BGTBase BoardgameType = "base"
	BGTExpansion BoardgameType = "expansion"
	BGTAddon BoardgameType = "addon"
)

type Boardgame struct {
	// nolint: structcheck, unused
	tableName struct{} `pg:"bzen.boardgame,alias:boardgame"`

	ID string
	BGGID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name string
	Type BoardgameType
}
