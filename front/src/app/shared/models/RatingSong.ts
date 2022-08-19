export interface RateSongBody {
    id: number;
    user_id: number;
    rating_value: number;
    song_id: number;
  }
  

// type RateSongBody struct {
// 	Id          int64   `json:"id"`
// 	UserId      int64   `json:"user_id"`
// 	RatingValue float32 `json:"rating_value"`
// 	SongId      int64   `json:"song_id"`
// }
