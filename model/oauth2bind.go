package model

type Oauth2Bind struct {
	Common

	//UserID   uint64 `gorm:"uniqueIndex:u_p_o" json:"user_id,omitempty"` //diy
	UserID   uint64   `gorm:"uniqueIndex:u_p_o"`
	//Provider string `gorm:"uniqueIndex:u_p_o" json:"provider,omitempty"` //diy
	Provider string `gorm:"type:varchar(64);not null;uniqueIndex:u_p_o"`
	//OpenID   string `gorm:"uniqueIndex:u_p_o" json:"open_id,omitempty"` //diy
	OpenID   string `gorm:"type:varchar(191);not null;uniqueIndex:u_p_o"`
}

type Oauth2LoginType uint8

const (
	_ Oauth2LoginType = iota
	RTypeLogin
	RTypeBind
)

type Oauth2State struct {
	Action      Oauth2LoginType
	Provider    string
	State       string
	RedirectURL string
}
