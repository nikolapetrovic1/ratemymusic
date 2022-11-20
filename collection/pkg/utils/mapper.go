package utils

import (
	"github.com/nikolapetrovic1/ratemymusic/collection/pkg/models"
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/collection"
)

func MapFavoriteToFavoriteData(favorite models.Favorite) *pb.FavoriteData {
	return &pb.FavoriteData{
		Id:       favorite.ID,
		UserId:   favorite.UserId,
		Kind:     favorite.Kind,
		KindId:   favorite.KindId,
		Name:     favorite.Name,
		Favorite: favorite.Favorite,
	}
}

func MapListCollectionToCollectionData(favorites []models.Favorite) []*pb.FavoriteData {
	var favoritesData []*pb.FavoriteData

	for _, favorite := range favorites {
		favoritesData = append(favoritesData, MapFavoriteToFavoriteData(favorite))
	}
	return favoritesData
}

func MapFavoriteDataToFavorite(favoriteData *pb.FavoriteData) models.Favorite {
	return models.Favorite{
		ID:       favoriteData.Id,
		UserId:   favoriteData.UserId,
		Kind:     favoriteData.Kind,
		KindId:   favoriteData.KindId,
		Name:     favoriteData.Name,
		Favorite: favoriteData.Favorite,
	}
}
