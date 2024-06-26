// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: chromecast/device/v1/device.proto

package devicev1

import (
	context "context"
	redact "github.com/jasonkolodziej/protoc-gen-go-redact/gen/redact"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
	_ emptypb.Empty
)

// RegisterRedactedDeviceControllerServer wraps the DeviceControllerServer with the redacted server and registers the service in GRPC
func RegisterRedactedDeviceControllerServer(s grpc.ServiceRegistrar, srv DeviceControllerServer, bypass redact.Bypass) {
	RegisterDeviceControllerServer(s, RedactedDeviceControllerServer(srv, bypass))
}

func RedactedDeviceControllerServer(srv DeviceControllerServer, bypass redact.Bypass) DeviceControllerServer {
	if bypass == nil {
		bypass = redact.Falsy
	}
	return &redactedDeviceControllerServer{srv: srv, bypass: bypass}
}

type redactedDeviceControllerServer struct {
	UnsafeDeviceControllerServer
	srv    DeviceControllerServer
	bypass redact.Bypass
}

// Play is the redacted wrapper for the actual DeviceControllerServer.Play method
func (s *redactedDeviceControllerServer) Play(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	res, err := s.srv.Play(ctx, in)
	if !s.bypass.CheckInternal(ctx) {
		// Apply redaction to the response
		redact.Apply(res)
	}
	return res, err
}

// Redact method implementation for Device
func (x *Device) Redact() {
	if x == nil {
		return
	}
}
