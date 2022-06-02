package v1

import (
	"github.com/pinax-network/dtypes/proto/v1/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// Event represents a metering event that needs to be recorded
type Event struct {
	Source         string     `json:"source"`
	Kind           string     `json:"kind"`
	Method         string     `json:"method"`
	Network        string     `json:"network"`
	UserId         string     `json:"user_id"`
	ApiKeyId       string     `json:"api_key_id"`
	IpAddress      string     `json:"ip_address"`
	RequestsCount  int64      `json:"requests_count,omitempty"`
	ResponsesCount int64      `json:"responses_count,omitempty"`
	IngressBytes   int64      `json:"ingress_bytes,omitempty"`
	EgressBytes    int64      `json:"egress_bytes,omitempty"`
	Time           *time.Time `json:"timestamp,omitempty"`
}

func (e *Event) ToProtobuf() *pb.Event {

	res := &pb.Event{
		Source:         e.Source,
		Kind:           e.Kind,
		Method:         e.Method,
		Network:        e.Network,
		UserId:         e.UserId,
		ApiKeyId:       e.ApiKeyId,
		IpAddress:      e.IpAddress,
		RequestsCount:  e.RequestsCount,
		ResponsesCount: e.ResponsesCount,
		IngressBytes:   e.IngressBytes,
		EgressBytes:    e.EgressBytes,
	}

	if e.Time != nil {
		res.Time = timestamppb.New(*e.Time)
	}

	return res
}

func (e *Event) FromProtobuf(pb *pb.Event) {

	e.Source = pb.Source
	e.Kind = pb.Kind
	e.Method = pb.Method
	e.Network = pb.Network
	e.UserId = pb.UserId
	e.ApiKeyId = pb.ApiKeyId
	e.IpAddress = pb.IpAddress
	e.RequestsCount = pb.RequestsCount
	e.ResponsesCount = pb.ResponsesCount
	e.IngressBytes = pb.IngressBytes
	e.EgressBytes = pb.EgressBytes

	if pb.Time.IsValid() {
		pbTime := pb.Time.AsTime()
		e.Time = &pbTime
	}
}
