package model

import "time"

type (
	BoardgameType string
	ContainerType string
)

const (
	BGTBase      BoardgameType = "base"
	BGTExpansion BoardgameType = "expansion"
	BGTAddon     BoardgameType = "addon"
)

const (
	CTShelf      ContainerType = "shelf"
	CTSnapTopBox ContainerType = "snap_top_box"
	CTBag        ContainerType = "bag"
)

type Boardgame struct {
	// nolint: structcheck, unused
	tableName struct{} `pg:"bzen.boardgame,alias:boardgame"`

	ID        string        `pg:"id,pk,type:uuid"`
	CreatedAt time.Time     `pg:"created_at"`
	UpdatedAt time.Time     `pg:"updated_at"`
	DeletedAt *time.Time    `pg:"deleted_at,soft_delete"`
	Name      string        `pg:"name"`
	Type      BoardgameType `pg:"type,type:bzen.boardgame_type_t"`
	UPC       *string       `pg:"upc"`
	BggID     string        `pg:"bgg_id"`
}

type Container struct {
	// nolint: structcheck, unused
	tableName struct{} `pg:"bzen.container,alias:container"`

	ID        string        `pg:"id,pk,type:uuid"`
	CreatedAt time.Time     `pg:"created_at"`
	UpdatedAt time.Time     `pg:"updated_at"`
	DeletedAt *time.Time    `pg:"deleted_at,soft_delete"`
	Name      string        `pg:"name"`
	Type      ContainerType `pg:"type,type:bzen.container_type_t"`
	Location  Area          `pg:"rel:has-one,join_fk:area_id"`
}

type Area struct {
	// nolint: structcheck, unused
	tableName struct{} `pg:"bzen.area,alias:area,"`

	ID        string     `pg:"id,pk,type:uuid"`
	CreatedAt time.Time  `pg:"created_at"`
	UpdatedAt time.Time  `pg:"updated_at"`
	DeletedAt *time.Time `pg:"deleted_at,soft_delete"`
	Name      string     `pg:"name"`
}
