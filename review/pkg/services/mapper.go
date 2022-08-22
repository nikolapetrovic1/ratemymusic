package services

import (
	pb "github.com/nikolapetrovic1/ratemymusic/common/pkg/review"
	"github.com/nikolapetrovic1/ratemymusic/review/pkg/models"
)

func mapListReviewToReviewData(reviews []models.Review) []*pb.ReviewData {
	var reviewData []*pb.ReviewData
	for _, review := range reviews {
		reviewData = append(reviewData, mapReviewToReviewData(review))
	}
	return reviewData
}
func mapReviewToReviewData(review models.Review) *pb.ReviewData {

	return &pb.ReviewData{
		Id:     review.ID,
		UserId: review.UserId,
		TypeId: review.TypeId,
		Text:   review.Text,
		Type:   review.Type,
	}
}

func mapReviewDataToReview(reviewData *pb.ReviewData) *models.Review {
	return &models.Review{
		ID:     reviewData.Id,
		UserId: reviewData.UserId,
		Text:   reviewData.Text,
		Type:   reviewData.Type,
		TypeId: reviewData.TypeId,
	}
}
