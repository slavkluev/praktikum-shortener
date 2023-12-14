package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/slavkluev/praktikum-shortener/internal/app/domain"
	pb "github.com/slavkluev/praktikum-shortener/internal/app/record/delivery/grpc/proto"
)

type RecordsServer struct {
	pb.UnimplementedRecordsServer
	recordUsecase domain.RecordUsecase
}

func NewRecordsServer(recordUsecase domain.RecordUsecase) *RecordsServer {
	return &RecordsServer{
		recordUsecase: recordUsecase,
	}
}

func (s *RecordsServer) GetAllUrls(ctx context.Context, in *pb.GetAllUrlsRequest) (*pb.GetAllUrlsResponse, error) {
	records, err := s.recordUsecase.GetByUserID(ctx, in.User)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(records) == 0 {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	var urls []*pb.ShortenURL
	for _, record := range records {
		urls = append(urls, &pb.ShortenURL{
			OriginalUrl: record.URL,
			UniqueId:    record.ID,
		})
	}

	return &pb.GetAllUrlsResponse{Urls: urls}, nil
}

func (s *RecordsServer) GetOriginalURL(ctx context.Context, in *pb.GetOriginalURLRequest) (*pb.GetOriginalURLResponse, error) {
	record, err := s.recordUsecase.GetByID(ctx, in.UniqueId)
	if err != nil && record.Deleted {
		return nil, status.Error(codes.NotFound, "Not found")
	}

	return &pb.GetOriginalURLResponse{OriginalUrl: record.URL}, nil
}

func (s *RecordsServer) ShortenURL(ctx context.Context, in *pb.ShortenURLRequest) (*pb.ShortenURLResponse, error) {
	record := &domain.Record{
		User: in.User,
		URL:  in.OriginalUrl,
	}

	err := s.recordUsecase.Store(ctx, record)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrConflict):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.ShortenURLResponse{UniqueId: record.ID}, nil
}
