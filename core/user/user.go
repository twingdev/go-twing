package user

import (
	"github.com/twingdev/go-twing/core/common"
	"github.com/twingdev/go-twing/core/medias"
)

type User struct {
	ID           string `json:"id"`
	privateData  *PersonalData
	PublicHandle string
	Avatar       *medias.Media
	HeaderBg     *medias.Media
	Timestamp    int64
	LastSeen     int64
	Flags        map[string]bool
}

type PersonalData struct {
	Name         string `json:"full_name"`
	First        string `json:"first_name"`
	Last         string `json:"last_name"`
	Email        string `json:"email_address"`
	BackupEmail  string `json:"email_backup"`
	MobileNumber *PhoneData
	DOB          *common.Date
}

type PhoneData struct {
	CountryCode int  `json:"country_code"`
	AreaCode    int  `json:"area_code"`
	Number      int  `json:"phone_number"`
	Verified    bool `json:"phone_verified"`
}
