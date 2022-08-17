package services

import (
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/rating"
	"github.com/nikolapetrovic1/ratemymusic/rating/pkg/models"
)

func mapListRatingRatingData(ratingsSong []models.SongRating, ratingsAlbum []models.AlbumRating) []*pb.RatingData {
	var ratingDatas []*pb.RatingData
	for _, rating := range ratingsSong {
		ratingDatas = append(ratingDatas, mapSongRatingRatingData(&rating))
	}
	for _, rating := range ratingsAlbum {
		ratingDatas = append(ratingDatas, mapAlbumRatingRatingData(&rating))
	}
	return ratingDatas
}
func mapListSongRatingRatingData(ratingsSong []models.SongRating) []*pb.RatingData {
	var ratingDatas []*pb.RatingData

	for _, rating := range ratingsSong {
		ratingDatas = append(ratingDatas, mapSongRatingRatingData(&rating))
	}
	return ratingDatas
}

func mapListAlbumRatingRatingData(ratingsAlbum []models.AlbumRating) []*pb.RatingData {
	var ratingDatas []*pb.RatingData

	for _, rating := range ratingsAlbum {
		ratingDatas = append(ratingDatas, mapAlbumRatingRatingData(&rating))
	}
	return ratingDatas
}
func mapSongRatingRatingData(rating *models.SongRating) *pb.RatingData {
	return &pb.RatingData{
		RatingValue: float32(rating.RatingValue),
		UserId:      rating.UserID,
		RatingType: &pb.RatingData_SongId{
			SongId: rating.SongID,
		},
	}
}
func mapAlbumRatingRatingData(rating *models.AlbumRating) *pb.RatingData {
	return &pb.RatingData{
		RatingValue: float32(rating.RatingValue),
		UserId:      rating.UserID,
		RatingType: &pb.RatingData_AlbumId{
			AlbumId: rating.AlbumId,
		},
	}
}

func mapRatingDataRatingAlbum(ratingData *pb.RateRequest) *models.AlbumRating {
	return &models.AlbumRating{
		ID:          ratingData.Id,
		UserID:      ratingData.UserId,
		RatingValue: float64(ratingData.RatingValue),
		AlbumId:     ratingData.GetAlbumId(),
	}
}
func mapRatingDataRatingSong(ratingData *pb.RateRequest) *models.SongRating {
	return &models.SongRating{
		ID:          ratingData.Id,
		UserID:      ratingData.UserId,
		RatingValue: float64(ratingData.RatingValue),
		SongID:      ratingData.GetSongId(),
	}
}
