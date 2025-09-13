// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dynamicpb_test

import (
	"fmt"
	"testing"

	"github.com/vedadiyan/protobuf/proto"
	"github.com/vedadiyan/protobuf/reflect/protoreflect"
	"github.com/vedadiyan/protobuf/reflect/protoregistry"
	"github.com/vedadiyan/protobuf/testing/prototest"
	"github.com/vedadiyan/protobuf/types/dynamicpb"

	testpb "github.com/vedadiyan/protobuf/internal/testprotos/test"
	test3pb "github.com/vedadiyan/protobuf/internal/testprotos/test3"
)

func TestConformance(t *testing.T) {
	for _, message := range []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testpb.TestAllExtensions)(nil),
	} {
		t.Run(fmt.Sprintf("%T", message), func(t *testing.T) {
			mt := dynamicpb.NewMessageType(message.ProtoReflect().Descriptor())
			prototest.Message{}.Test(t, mt)
		})
	}
}

func TestDynamicExtensions(t *testing.T) {
	for _, message := range []proto.Message{
		(*testpb.TestAllExtensions)(nil),
	} {
		t.Run(fmt.Sprintf("%T", message), func(t *testing.T) {
			mt := dynamicpb.NewMessageType(message.ProtoReflect().Descriptor())
			prototest.Message{
				Resolver: extResolver{},
			}.Test(t, mt)
		})
	}
}

func TestDynamicEnums(t *testing.T) {
	for _, enum := range []protoreflect.Enum{
		testpb.TestAllTypes_FOO,
		test3pb.TestAllTypes_FOO,
	} {
		t.Run(fmt.Sprintf("%v", enum), func(t *testing.T) {
			et := dynamicpb.NewEnumType(enum.Descriptor())
			prototest.Enum{}.Test(t, et)
		})
	}
}

type extResolver struct{}

func (extResolver) FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error) {
	xt, err := protoregistry.GlobalTypes.FindExtensionByName(field)
	if err != nil {
		return nil, err
	}
	return dynamicpb.NewExtensionType(xt.TypeDescriptor().Descriptor()), nil
}

func (extResolver) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
	xt, err := protoregistry.GlobalTypes.FindExtensionByNumber(message, field)
	if err != nil {
		return nil, err
	}
	return dynamicpb.NewExtensionType(xt.TypeDescriptor().Descriptor()), nil
}

func (extResolver) RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool) {
	protoregistry.GlobalTypes.RangeExtensionsByMessage(message, func(xt protoreflect.ExtensionType) bool {
		return f(dynamicpb.NewExtensionType(xt.TypeDescriptor().Descriptor()))
	})
}
