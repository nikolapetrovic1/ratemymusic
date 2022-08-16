package util

type SongRatingBody struct {
	UserID      string  `json:"user_id"`
	RatingValue float32 `json:"rating_value"`
	SongId      int64   `json:"song_id"`
}

type AlbumRatingBody struct {
}
