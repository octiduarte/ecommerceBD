package entities

type SocialMedia struct {
	SocialMediaID int64  `json:"social_media_id,omitempty"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	StoreID       int64  `json:"store_id,omitempty"`
}
