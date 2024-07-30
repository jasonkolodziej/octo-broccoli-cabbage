// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: chromecast/core_api/v1/cast_channel.proto

package core_apiv1

import (
	context "context"
	redact "github.com/jasonkolodziej/protoc-gen-go-redact/gen/redact"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
)

// Redact method implementation for CastMessage
func (x *CastMessage) Redact() {
	if x == nil {
		return
	}

	// Safe field: ProtocolVersion

	// Safe field: SourceId

	// Safe field: DestinationId

	// Safe field: Namespace

	// Safe field: PayloadType

	// Safe field: PayloadUtf8

	// Safe field: PayloadBinary
}

// Redact method implementation for AuthChallenge
func (x *AuthChallenge) Redact() {
	if x == nil {
		return
	}
}

// Redact method implementation for AuthResponse
func (x *AuthResponse) Redact() {
	if x == nil {
		return
	}

	// Safe field: Signature

	// Safe field: ClientAuthCertificate
}

// Redact method implementation for AuthError
func (x *AuthError) Redact() {
	if x == nil {
		return
	}

	// Safe field: ErrorType
}

// Redact method implementation for DeviceAuthMessage
func (x *DeviceAuthMessage) Redact() {
	if x == nil {
		return
	}

	// Safe field: Challenge

	// Safe field: Response

	// Safe field: Error
}