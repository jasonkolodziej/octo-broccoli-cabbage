// Code generated by protoc-gen-redact. DO NOT EDIT.
// source: redact/v1/redact.proto

package v1

import (
	context "context"
	redact "github.com/jasonkolodziej/protoc-gen-go-redact/gen/redact"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ grpc.Server
	_ context.Context
	_ redact.Redactor
	_ codes.Code
	_ status.Status
	_ descriptorpb.FileDescriptorSet
)

// Redact method implementation for FieldRules
func (x *FieldRules) Redact() {
	if x == nil {
		return
	}

	// Safe field: Float

	// Safe field: Double

	// Safe field: Int32

	// Safe field: Int64

	// Safe field: Uint32

	// Safe field: Uint64

	// Safe field: Sint32

	// Safe field: Sint64

	// Safe field: Fixed32

	// Safe field: Fixed64

	// Safe field: Sfixed32

	// Safe field: Sfixed64

	// Safe field: Bool

	// Safe field: String_

	// Safe field: Bytes

	// Safe field: Enum

	// Safe field: Message

	// Safe field: Element
}

// Redact method implementation for MessageRules
func (x *MessageRules) Redact() {
	if x == nil {
		return
	}

	// Safe field: Skip

	// Safe field: Empty

	// Safe field: Nil

	// Safe field: Apply
}

// Redact method implementation for ElementRules
func (x *ElementRules) Redact() {
	if x == nil {
		return
	}

	// Safe field: Empty

	// Safe field: Nested

	// Safe field: Item
}
