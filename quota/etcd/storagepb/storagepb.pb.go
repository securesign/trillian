// Copyright 2017 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.20.1
// source: storagepb.proto

// Package storagepb contains definitions for quota storage protos, which are
// recorded in etcd.

package storagepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible states of a quota configuration.
type Config_State int32

const (
	// Unknown quota state. Invalid.
	Config_UNKNOWN_CONFIG_STATE Config_State = 0
	// Quota is enabled.
	Config_ENABLED Config_State = 1
	// Quota is disabled (considered infinite).
	Config_DISABLED Config_State = 2
)

// Enum value maps for Config_State.
var (
	Config_State_name = map[int32]string{
		0: "UNKNOWN_CONFIG_STATE",
		1: "ENABLED",
		2: "DISABLED",
	}
	Config_State_value = map[string]int32{
		"UNKNOWN_CONFIG_STATE": 0,
		"ENABLED":              1,
		"DISABLED":             2,
	}
)

func (x Config_State) Enum() *Config_State {
	p := new(Config_State)
	*p = x
	return p
}

func (x Config_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Config_State) Descriptor() protoreflect.EnumDescriptor {
	return file_storagepb_proto_enumTypes[0].Descriptor()
}

func (Config_State) Type() protoreflect.EnumType {
	return &file_storagepb_proto_enumTypes[0]
}

func (x Config_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Config_State.Descriptor instead.
func (Config_State) EnumDescriptor() ([]byte, []int) {
	return file_storagepb_proto_rawDescGZIP(), []int{2, 0}
}

// Data contained in a quota bucket.
// Stored at each each quota's zero bucket. For example,
// quotas/global/read/0 or quotas/trees/$id/read/0.
type Bucket struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Number of tokens left in the bucket.
	Tokens int64 `protobuf:"varint,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
	// Timestamp of the last time the bucket got replenished.
	LastReplenishMillisSinceEpoch int64 `protobuf:"varint,2,opt,name=last_replenish_millis_since_epoch,json=lastReplenishMillisSinceEpoch,proto3" json:"last_replenish_millis_since_epoch,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	mi := &file_storagepb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_storagepb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_storagepb_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetTokens() int64 {
	if x != nil {
		return x.Tokens
	}
	return 0
}

func (x *Bucket) GetLastReplenishMillisSinceEpoch() int64 {
	if x != nil {
		return x.LastReplenishMillisSinceEpoch
	}
	return 0
}

// Configuration for all quotas.
// Stored at quotas/configs.
type Configs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Known quota configurations.
	Configs       []*Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Configs) Reset() {
	*x = Configs{}
	mi := &file_storagepb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configs) ProtoMessage() {}

func (x *Configs) ProtoReflect() protoreflect.Message {
	mi := &file_storagepb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configs.ProtoReflect.Descriptor instead.
func (*Configs) Descriptor() ([]byte, []int) {
	return file_storagepb_proto_rawDescGZIP(), []int{1}
}

func (x *Configs) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

// Configuration of a quota.
type Config struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the config, eg, “quotas/trees/1234/read/config”.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// State of the config.
	State Config_State `protobuf:"varint,2,opt,name=state,proto3,enum=storagepb.Config_State" json:"state,omitempty"`
	// Max number of tokens available for the config.
	MaxTokens int64 `protobuf:"varint,3,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// Replenishment strategy used by the config.
	//
	// Types that are valid to be assigned to ReplenishmentStrategy:
	//
	//	*Config_SequencingBased
	//	*Config_TimeBased
	ReplenishmentStrategy isConfig_ReplenishmentStrategy `protobuf_oneof:"replenishment_strategy"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_storagepb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_storagepb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_storagepb_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetState() Config_State {
	if x != nil {
		return x.State
	}
	return Config_UNKNOWN_CONFIG_STATE
}

func (x *Config) GetMaxTokens() int64 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *Config) GetReplenishmentStrategy() isConfig_ReplenishmentStrategy {
	if x != nil {
		return x.ReplenishmentStrategy
	}
	return nil
}

func (x *Config) GetSequencingBased() *SequencingBasedStrategy {
	if x != nil {
		if x, ok := x.ReplenishmentStrategy.(*Config_SequencingBased); ok {
			return x.SequencingBased
		}
	}
	return nil
}

func (x *Config) GetTimeBased() *TimeBasedStrategy {
	if x != nil {
		if x, ok := x.ReplenishmentStrategy.(*Config_TimeBased); ok {
			return x.TimeBased
		}
	}
	return nil
}

type isConfig_ReplenishmentStrategy interface {
	isConfig_ReplenishmentStrategy()
}

type Config_SequencingBased struct {
	// Sequencing-based replenishment settings.
	SequencingBased *SequencingBasedStrategy `protobuf:"bytes,4,opt,name=sequencing_based,json=sequencingBased,proto3,oneof"`
}

type Config_TimeBased struct {
	// Time-based replenishment settings.
	TimeBased *TimeBasedStrategy `protobuf:"bytes,5,opt,name=time_based,json=timeBased,proto3,oneof"`
}

func (*Config_SequencingBased) isConfig_ReplenishmentStrategy() {}

func (*Config_TimeBased) isConfig_ReplenishmentStrategy() {}

// Sequencing-based replenishment strategy settings.
type SequencingBasedStrategy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SequencingBasedStrategy) Reset() {
	*x = SequencingBasedStrategy{}
	mi := &file_storagepb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SequencingBasedStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequencingBasedStrategy) ProtoMessage() {}

func (x *SequencingBasedStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_storagepb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequencingBasedStrategy.ProtoReflect.Descriptor instead.
func (*SequencingBasedStrategy) Descriptor() ([]byte, []int) {
	return file_storagepb_proto_rawDescGZIP(), []int{3}
}

// Time-based replenishment strategy settings.
type TimeBasedStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Number of tokens to replenish at every replenish_interval_seconds.
	TokensToReplenish int64 `protobuf:"varint,1,opt,name=tokens_to_replenish,json=tokensToReplenish,proto3" json:"tokens_to_replenish,omitempty"`
	// Interval at which tokens_to_replenish get replenished.
	ReplenishIntervalSeconds int64 `protobuf:"varint,2,opt,name=replenish_interval_seconds,json=replenishIntervalSeconds,proto3" json:"replenish_interval_seconds,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *TimeBasedStrategy) Reset() {
	*x = TimeBasedStrategy{}
	mi := &file_storagepb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeBasedStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeBasedStrategy) ProtoMessage() {}

func (x *TimeBasedStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_storagepb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeBasedStrategy.ProtoReflect.Descriptor instead.
func (*TimeBasedStrategy) Descriptor() ([]byte, []int) {
	return file_storagepb_proto_rawDescGZIP(), []int{4}
}

func (x *TimeBasedStrategy) GetTokensToReplenish() int64 {
	if x != nil {
		return x.TokensToReplenish
	}
	return 0
}

func (x *TimeBasedStrategy) GetReplenishIntervalSeconds() int64 {
	if x != nil {
		return x.ReplenishIntervalSeconds
	}
	return 0
}

var File_storagepb_proto protoreflect.FileDescriptor

var file_storagepb_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x22, 0x6a, 0x0a, 0x06,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x48,
	0x0a, 0x21, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68,
	0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x6c, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x53, 0x69,
	0x6e, 0x63, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x36, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x22, 0xd2, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x4f, 0x0a,
	0x10, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x61,
	0x73, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x73, 0x65, 0x64, 0x12, 0x3d,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x22, 0x3c, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x18, 0x0a, 0x16, 0x72,
	0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x22, 0x81, 0x01, 0x0a, 0x11, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x65, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x3c, 0x0a, 0x1a, 0x72, 0x65, 0x70, 0x6c, 0x65, 0x6e,
	0x69, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x72, 0x65, 0x70, 0x6c,
	0x65, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69,
	0x61, 0x6e, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storagepb_proto_rawDescOnce sync.Once
	file_storagepb_proto_rawDescData = file_storagepb_proto_rawDesc
)

func file_storagepb_proto_rawDescGZIP() []byte {
	file_storagepb_proto_rawDescOnce.Do(func() {
		file_storagepb_proto_rawDescData = protoimpl.X.CompressGZIP(file_storagepb_proto_rawDescData)
	})
	return file_storagepb_proto_rawDescData
}

var file_storagepb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storagepb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storagepb_proto_goTypes = []any{
	(Config_State)(0),               // 0: storagepb.Config.State
	(*Bucket)(nil),                  // 1: storagepb.Bucket
	(*Configs)(nil),                 // 2: storagepb.Configs
	(*Config)(nil),                  // 3: storagepb.Config
	(*SequencingBasedStrategy)(nil), // 4: storagepb.SequencingBasedStrategy
	(*TimeBasedStrategy)(nil),       // 5: storagepb.TimeBasedStrategy
}
var file_storagepb_proto_depIdxs = []int32{
	3, // 0: storagepb.Configs.configs:type_name -> storagepb.Config
	0, // 1: storagepb.Config.state:type_name -> storagepb.Config.State
	4, // 2: storagepb.Config.sequencing_based:type_name -> storagepb.SequencingBasedStrategy
	5, // 3: storagepb.Config.time_based:type_name -> storagepb.TimeBasedStrategy
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storagepb_proto_init() }
func file_storagepb_proto_init() {
	if File_storagepb_proto != nil {
		return
	}
	file_storagepb_proto_msgTypes[2].OneofWrappers = []any{
		(*Config_SequencingBased)(nil),
		(*Config_TimeBased)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storagepb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storagepb_proto_goTypes,
		DependencyIndexes: file_storagepb_proto_depIdxs,
		EnumInfos:         file_storagepb_proto_enumTypes,
		MessageInfos:      file_storagepb_proto_msgTypes,
	}.Build()
	File_storagepb_proto = out.File
	file_storagepb_proto_rawDesc = nil
	file_storagepb_proto_goTypes = nil
	file_storagepb_proto_depIdxs = nil
}
