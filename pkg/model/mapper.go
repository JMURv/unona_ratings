package model

import (
	pb "github.com/JMURv/unona/ratings/api/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func RatingFromProto(r *pb.Rating) *Rating {
	userUUID, _ := uuid.Parse(r.UserUuid)
	return &Rating{
		UserUUID:  userUUID,
		Rating:    uint8(r.Rating),
		Text:      r.Text,
		CreatedAt: r.CreatedAt.AsTime(),
	}
}

func RatingToProto(r *Rating) *pb.Rating {
	return &pb.Rating{
		UserUuid: r.UserUUID.String(),
		Rating:   uint32(r.Rating),
		Text:     r.Text,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: r.CreatedAt.Unix(),
			Nanos:   int32(r.CreatedAt.Nanosecond()),
		},
	}
}
