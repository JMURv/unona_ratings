package grpc

import (
	"context"
	pb "github.com/JMURv/unona/ratings/api/pb"
	controller "github.com/JMURv/unona/ratings/internal/controller/rating"
	metrics "github.com/JMURv/unona/ratings/internal/metrics/prometheus"
	"github.com/JMURv/unona/ratings/pkg/model"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Handler struct {
	pb.RatingServiceServer
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetUserRating(ctx context.Context, req *pb.GetUserRatingRequest) (*pb.Rating, error) {
	statusCode := codes.OK
	start := time.Now()

	span := opentracing.GlobalTracer().StartSpan("ratings.GetUserRating.handler")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer func() {
		span.Finish()
		metrics.ObserveRequest(time.Since(start), int(statusCode), "GetUserRating")
	}()

	if req == nil || req.UserUuid == "" {
		statusCode = codes.InvalidArgument
		return nil, status.Errorf(statusCode, "nil req")
	}

	if r, err := h.ctrl.GetUserRating(ctx, uuid.MustParse(req.UserUuid)); err != nil {
		statusCode = codes.Internal
		span.SetTag("error", true)
		return nil, status.Errorf(statusCode, err.Error())
	} else {
		return model.RatingToProto(r), nil
	}
}

func (h *Handler) CreateReport(ctx context.Context, req *pb.CreateReportRequest) (*pb.Rating, error) {
	statusCode := codes.OK
	start := time.Now()

	span := opentracing.GlobalTracer().StartSpan("ratings.CreateReport.handler")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer func() {
		span.Finish()
		metrics.ObserveRequest(time.Since(start), int(statusCode), "CreateReport")
	}()

	if req == nil {
		statusCode = codes.InvalidArgument
		return nil, status.Errorf(statusCode, "nil req")
	}

	r, err := h.ctrl.CreateReport(ctx, model.RatingFromProto(&pb.Rating{
		UserUuid: req.UserUuid,
		Rating:   req.Rating,
		Text:     req.Text,
	}))
	if err != nil {
		span.SetTag("error", true)
		statusCode = codes.Internal
		return nil, status.Errorf(statusCode, err.Error())
	}

	return model.RatingToProto(r), nil
}

func (h *Handler) UpdateReport(ctx context.Context, req *pb.UpdateReportRequest) (*pb.Rating, error) {
	statusCode := codes.OK
	start := time.Now()

	span := opentracing.GlobalTracer().StartSpan("ratings.UpdateReport.handler")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer func() {
		span.Finish()
		metrics.ObserveRequest(time.Since(start), int(statusCode), "UpdateReport")
	}()

	if req == nil || req.RatingId == 0 {
		statusCode = codes.InvalidArgument
		return nil, status.Errorf(statusCode, "nil req or empty id")
	}

	r, err := h.ctrl.UpdateReport(ctx, uint(req.RatingId), model.RatingFromProto(&pb.Rating{
		Rating: req.Rating,
		Text:   req.Text,
	}))
	if err != nil {
		span.SetTag("error", true)
		statusCode = codes.Internal
		return nil, status.Errorf(statusCode, err.Error())
	}

	return model.RatingToProto(r), nil
}

func (h *Handler) DeleteReport(ctx context.Context, req *pb.DeleteReportRequest) (*pb.Empty, error) {
	statusCode := codes.OK
	start := time.Now()

	span := opentracing.GlobalTracer().StartSpan("ratings.DeleteReport.handler")
	ctx = opentracing.ContextWithSpan(ctx, span)
	defer func() {
		span.Finish()
		metrics.ObserveRequest(time.Since(start), int(statusCode), "DeleteReport")
	}()

	if req == nil || req.RatingId == 0 {
		statusCode = codes.InvalidArgument
		return nil, status.Errorf(statusCode, "nil req or empty id")
	}

	if err := h.ctrl.DeleteReport(ctx, uint(req.RatingId)); err != nil {
		span.SetTag("error", true)
		statusCode = codes.Internal
		return nil, status.Errorf(statusCode, err.Error())
	}

	return &pb.Empty{}, nil
}
