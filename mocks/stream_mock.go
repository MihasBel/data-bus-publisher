package mocks

import (
	"context"
	"github.com/MihasBel/data-bus-publisher/delivery/grpc/gen/v1/publisher"
	"google.golang.org/grpc/metadata"
)

type StreamMock struct {
	SendFunc func(*publisher.Message) error
}

func (s *StreamMock) SetHeader(md metadata.MD) error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamMock) SendHeader(md metadata.MD) error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamMock) SetTrailer(md metadata.MD) {
	//TODO implement me
	panic("implement me")
}

func (s *StreamMock) Context() context.Context {
	//TODO implement me
	panic("implement me")
}

func (s *StreamMock) SendMsg(m interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamMock) RecvMsg(m interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *StreamMock) Send(msg *publisher.Message) error {
	return s.SendFunc(msg)
}
