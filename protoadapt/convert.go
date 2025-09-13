// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protoadapt bridges the original and new proto APIs.
package protoadapt

import (
	"github.com/vedadiyan/protobuf/proto"
	"github.com/vedadiyan/protobuf/runtime/protoiface"
	"github.com/vedadiyan/protobuf/runtime/protoimpl"
)

// MessageV1 is the original [github.com/golang/protobuf/proto.Message] type.
type MessageV1 = protoiface.MessageV1

// MessageV2 is the [github.com/vedadiyan/protobuf/proto.Message] type used by the
// current [github.com/vedadiyan/protobuf] module, adding support for reflection.
type MessageV2 = proto.Message

// MessageV1Of converts a v2 message to a v1 message.
// It returns nil if m is nil.
func MessageV1Of(m MessageV2) MessageV1 {
	return protoimpl.X.ProtoMessageV1Of(m)
}

// MessageV2Of converts a v1 message to a v2 message.
// It returns nil if m is nil.
func MessageV2Of(m MessageV1) MessageV2 {
	return protoimpl.X.ProtoMessageV2Of(m)
}
