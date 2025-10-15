package models

// Account : Account Model
type Account struct {
	ID      int64  `bun:",pk,autoincrement"`
	UserID  int64  `bun:",notnull"`
	User    *User  `bun:"rel:belongs-to,join:user_id=id"`
	AssetID string `bun:",notnull"`
	Asset   *Asset `bun:"rel:has-one,join:asset_id=asset_id"`
	Type    string `bun:",notnull"`
}
