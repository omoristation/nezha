package model

import "time"

type StreamServer struct {
	ID           uint64 `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Note         string `json:"note,omitempty"`          // 管理员可见备注 //diy 发送到前端的端口号
	PublicNote   string `json:"public_note,omitempty"`   // 公开备注，只第一个数据包有值
	DisplayIndex int    `json:"display_index,omitempty"` // 展示排序，越大越靠前

	Host        *Host      `json:"host,omitempty"`
	State       *HostState `json:"state,omitempty"`
	CountryCode string     `json:"country_code,omitempty"`
	Ipv4        string     `json:"ipv4,omitempty"` //diy
	Ipv4Geo     string     `json:"ipv4_geo,omitempty"` //diy
	Ipv6        string     `json:"ipv6,omitempty"` //diy
	Ipv6Geo     string     `json:"ipv6_geo,omitempty"` //diy
	LastActive  time.Time  `json:"last_active,omitempty"`
}

type StreamServerData struct {
	Now     int64          `json:"now,omitempty"`
	Online  int            `json:"online,omitempty"`
	Servers []StreamServer `json:"servers,omitempty"`
}

type ServerForm struct {
	Name         string   `json:"name,omitempty"`
	Note         string   `json:"note,omitempty" validate:"optional"`                   // 管理员可见备注
	PublicNote   string   `json:"public_note,omitempty" validate:"optional"`            // 公开备注
	DisplayIndex int      `json:"display_index,omitempty" default:"0"`                  // 展示排序，越大越靠前
	HideForGuest bool     `json:"hide_for_guest,omitempty" validate:"optional"`         // 对游客隐藏
	EnableDDNS   bool     `json:"enable_ddns,omitempty" validate:"optional"`            // 启用DDNS
	DDNSProfiles []uint64 `gorm:"-" json:"ddns_profiles,omitempty" validate:"optional"` // DDNS配置
}

type ForceUpdateResponse struct {
	Success []uint64 `json:"success,omitempty" validate:"optional"`
	Failure []uint64 `json:"failure,omitempty" validate:"optional"`
	Offline []uint64 `json:"offline,omitempty" validate:"optional"`
}
