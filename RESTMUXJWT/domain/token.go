package domain

type AccessDetails struct {
	AccessUuid string `json:"access_uuid" bson:"access_uuid,omitempty"`
	UserId     uint64 `json:"user_id" bson:"user_id,omitempty"`
}

type TokenDetails struct {
	AccessToken  string `json:"access_token" bson:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token,omitempty"`
	AccessUuid   string `json:"access_uuid" bson:"access_uuid,omitempty"`
	RefreshUuid  string `json:"refresh_uuid" bson:"refresh_uuid,omitempty"`
	AtExpires    int64  `json:"at_expires" bson:"at_expires,omitempty"`
	RtExpires    int64  `json:"rt_expires" bson:"rt_expires,omitempty"`
}
