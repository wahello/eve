// Copyright(c) 2017-2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.6.1
// source: config/appconfig.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type InstanceOpsCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter uint32 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	OpsTime string `protobuf:"bytes,4,opt,name=opsTime,proto3" json:"opsTime,omitempty"` // Not currently used
}

func (x *InstanceOpsCmd) Reset() {
	*x = InstanceOpsCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_appconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceOpsCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceOpsCmd) ProtoMessage() {}

func (x *InstanceOpsCmd) ProtoReflect() protoreflect.Message {
	mi := &file_config_appconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceOpsCmd.ProtoReflect.Descriptor instead.
func (*InstanceOpsCmd) Descriptor() ([]byte, []int) {
	return file_config_appconfig_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceOpsCmd) GetCounter() uint32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *InstanceOpsCmd) GetOpsTime() string {
	if x != nil {
		return x.OpsTime
	}
	return ""
}

// The complete configuration for an Application Instance
// When changing key fields such as the drives/volumeRefs or the number
// of interfaces, the controller is required to issue a purge command i.e.,
// increase the purge counter. Otherwise there wil be an error (The controller
// can also issue a purge command to re-construct the content of the first
// drive/volumeRef without any changes.)
// Some changes such as ACL changes in the interfaces do not require a restart,
// but all other changes (such as fixedresources and adapters) require a
// restart command i.e., an increase to the restart counter. The restart counter
// can also be increased to cause an application instance restart without
// any other change to the application instance.
type AppInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	Displayname    string          `protobuf:"bytes,2,opt,name=displayname,proto3" json:"displayname,omitempty"` // User-friendly name
	Fixedresources *VmConfig       `protobuf:"bytes,3,opt,name=fixedresources,proto3" json:"fixedresources,omitempty"`
	// VolumeRefs, if supported by EVE, will supercede drives. Drives still
	// exist for backward compatibility.
	// Drives will be deprecated in the future.
	// The order here is critical because they are presented to the VM or
	// container in the order they are listed, e.g., the first VM image
	// will be the root disk.
	Drives []*Drive `protobuf:"bytes,4,rep,name=drives,proto3" json:"drives,omitempty"`
	// Set activate to start the application instance; clear it to stop it.
	Activate bool `protobuf:"varint,5,opt,name=activate,proto3" json:"activate,omitempty"`
	// NetworkAdapter are virtual adapters assigned to the application
	// The order here is critical because they are presented to the VM or
	// container in the order they are listed, e.g., the first NetworkAdapter
	// will appear in a Linux VM as eth0. Also, the MAC address is determined
	// based on the order in the list.
	Interfaces []*NetworkAdapter `protobuf:"bytes,6,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	// Physical adapters such as eth1 or USB controllers and GPUs assigned
	// to the application instance.
	// The Name in Adapter should be set to PhysicalIO.assigngrp
	Adapters []*Adapter `protobuf:"bytes,7,rep,name=adapters,proto3" json:"adapters,omitempty"`
	// The device behavior for a restart command (if counter increased)
	// is to restart the application instance
	// Increasing this multiple times does not imply the application instance
	// will restart more than once.
	// EVE can assume that the adapters did not change.
	Restart *InstanceOpsCmd `protobuf:"bytes,9,opt,name=restart,proto3" json:"restart,omitempty"`
	// The EVE behavior for a purge command is to restart the application instance
	// with the first drive/volumeRef recreated from its origin.
	Purge *InstanceOpsCmd `protobuf:"bytes,10,opt,name=purge,proto3" json:"purge,omitempty"`
	// App Instance initialization configuration data provided by user
	// This will be used as "user-data" in cloud-init
	// Empty string will indicate that cloud-init is not required
	// It is also used to carry environment variables for containers.
	// XXX will be deprecated and replaced by the cipherData below.
	UserData string `protobuf:"bytes,11,opt,name=userData,proto3" json:"userData,omitempty"`
	// Config flag if the app-instance should be made accessible
	// through a remote console session established by the device.
	RemoteConsole bool `protobuf:"varint,12,opt,name=remoteConsole,proto3" json:"remoteConsole,omitempty"`
	// contains the encrypted userdata
	CipherData *CipherBlock `protobuf:"bytes,13,opt,name=cipherData,proto3" json:"cipherData,omitempty"`
	// The static IP address assigned on the NetworkAdapter which App Container
	// stats collection uses. If the 'collectStatsIPAddr' is not empty and valid,
	// it enables the container stats collection for this App.
	// During App instance creation, after user enables the collection of stats
	// from App, cloud needs to make sure at least one 'Local' type of Network-Instance
	// is assigned to the App interface, and based on the subnet of the NI, statically
	// assign an IP address on the same subnet, e.g. 10.1.0.100
	CollectStatsIPAddr string `protobuf:"bytes,15,opt,name=collectStatsIPAddr,proto3" json:"collectStatsIPAddr,omitempty"`
	// The volumes to be attached to the app-instance.
	// The order here is critical because they are presented to the VM or
	// container in the order they are listed, e.g., the first VM image
	// will be the root disk.
	// Note that since the name volumeRef was used before and deprecated
	// python protobuf seems to require that we use a different name.
	VolumeRefList []*VolumeRef `protobuf:"bytes,16,rep,name=volumeRefList,proto3" json:"volumeRefList,omitempty"`
}

func (x *AppInstanceConfig) Reset() {
	*x = AppInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_appconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInstanceConfig) ProtoMessage() {}

func (x *AppInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_appconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInstanceConfig.ProtoReflect.Descriptor instead.
func (*AppInstanceConfig) Descriptor() ([]byte, []int) {
	return file_config_appconfig_proto_rawDescGZIP(), []int{1}
}

func (x *AppInstanceConfig) GetUuidandversion() *UUIDandVersion {
	if x != nil {
		return x.Uuidandversion
	}
	return nil
}

func (x *AppInstanceConfig) GetDisplayname() string {
	if x != nil {
		return x.Displayname
	}
	return ""
}

func (x *AppInstanceConfig) GetFixedresources() *VmConfig {
	if x != nil {
		return x.Fixedresources
	}
	return nil
}

func (x *AppInstanceConfig) GetDrives() []*Drive {
	if x != nil {
		return x.Drives
	}
	return nil
}

func (x *AppInstanceConfig) GetActivate() bool {
	if x != nil {
		return x.Activate
	}
	return false
}

func (x *AppInstanceConfig) GetInterfaces() []*NetworkAdapter {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

func (x *AppInstanceConfig) GetAdapters() []*Adapter {
	if x != nil {
		return x.Adapters
	}
	return nil
}

func (x *AppInstanceConfig) GetRestart() *InstanceOpsCmd {
	if x != nil {
		return x.Restart
	}
	return nil
}

func (x *AppInstanceConfig) GetPurge() *InstanceOpsCmd {
	if x != nil {
		return x.Purge
	}
	return nil
}

func (x *AppInstanceConfig) GetUserData() string {
	if x != nil {
		return x.UserData
	}
	return ""
}

func (x *AppInstanceConfig) GetRemoteConsole() bool {
	if x != nil {
		return x.RemoteConsole
	}
	return false
}

func (x *AppInstanceConfig) GetCipherData() *CipherBlock {
	if x != nil {
		return x.CipherData
	}
	return nil
}

func (x *AppInstanceConfig) GetCollectStatsIPAddr() string {
	if x != nil {
		return x.CollectStatsIPAddr
	}
	return ""
}

func (x *AppInstanceConfig) GetVolumeRefList() []*VolumeRef {
	if x != nil {
		return x.VolumeRefList
	}
	return nil
}

// Reference to a Volume specified separately in the API
// If a volume is purged (re-created from scratch) it will either have a new
// UUID or a new generationCount
type VolumeRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid            string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // Volume UUID
	GenerationCount int64  `protobuf:"varint,2,opt,name=generationCount,proto3" json:"generationCount,omitempty"`
}

func (x *VolumeRef) Reset() {
	*x = VolumeRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_appconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeRef) ProtoMessage() {}

func (x *VolumeRef) ProtoReflect() protoreflect.Message {
	mi := &file_config_appconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeRef.ProtoReflect.Descriptor instead.
func (*VolumeRef) Descriptor() ([]byte, []int) {
	return file_config_appconfig_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeRef) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *VolumeRef) GetGenerationCount() int64 {
	if x != nil {
		return x.GenerationCount
	}
	return 0
}

var File_config_appconfig_proto protoreflect.FileDescriptor

var file_config_appconfig_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x61, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x70, 0x73, 0x43, 0x6d, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xd8, 0x04, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x0e, 0x75, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x64,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x55, 0x55, 0x49, 0x44, 0x61, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x75, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x64, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x31, 0x0a, 0x0e, 0x66, 0x69, 0x78, 0x65, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x56, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0e, 0x66, 0x69, 0x78, 0x65, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12,
	0x2f, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x24, 0x0a, 0x08, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x08, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4f, 0x70, 0x73, 0x43, 0x6d, 0x64, 0x52, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x70, 0x73, 0x43, 0x6d,
	0x64, 0x52, 0x05, 0x70, 0x75, 0x72, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x0a, 0x63, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x12, 0x30, 0x0a, 0x0d, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x66, 0x52, 0x0d, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x09, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64,
	0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_appconfig_proto_rawDescOnce sync.Once
	file_config_appconfig_proto_rawDescData = file_config_appconfig_proto_rawDesc
)

func file_config_appconfig_proto_rawDescGZIP() []byte {
	file_config_appconfig_proto_rawDescOnce.Do(func() {
		file_config_appconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_appconfig_proto_rawDescData)
	})
	return file_config_appconfig_proto_rawDescData
}

var file_config_appconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_appconfig_proto_goTypes = []interface{}{
	(*InstanceOpsCmd)(nil),    // 0: InstanceOpsCmd
	(*AppInstanceConfig)(nil), // 1: AppInstanceConfig
	(*VolumeRef)(nil),         // 2: VolumeRef
	(*UUIDandVersion)(nil),    // 3: UUIDandVersion
	(*VmConfig)(nil),          // 4: VmConfig
	(*Drive)(nil),             // 5: Drive
	(*NetworkAdapter)(nil),    // 6: NetworkAdapter
	(*Adapter)(nil),           // 7: Adapter
	(*CipherBlock)(nil),       // 8: CipherBlock
}
var file_config_appconfig_proto_depIdxs = []int32{
	3, // 0: AppInstanceConfig.uuidandversion:type_name -> UUIDandVersion
	4, // 1: AppInstanceConfig.fixedresources:type_name -> VmConfig
	5, // 2: AppInstanceConfig.drives:type_name -> Drive
	6, // 3: AppInstanceConfig.interfaces:type_name -> NetworkAdapter
	7, // 4: AppInstanceConfig.adapters:type_name -> Adapter
	0, // 5: AppInstanceConfig.restart:type_name -> InstanceOpsCmd
	0, // 6: AppInstanceConfig.purge:type_name -> InstanceOpsCmd
	8, // 7: AppInstanceConfig.cipherData:type_name -> CipherBlock
	2, // 8: AppInstanceConfig.volumeRefList:type_name -> VolumeRef
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_config_appconfig_proto_init() }
func file_config_appconfig_proto_init() {
	if File_config_appconfig_proto != nil {
		return
	}
	file_config_acipherinfo_proto_init()
	file_config_devcommon_proto_init()
	file_config_storage_proto_init()
	file_config_vm_proto_init()
	file_config_netconfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_appconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceOpsCmd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_appconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInstanceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_appconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_appconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_appconfig_proto_goTypes,
		DependencyIndexes: file_config_appconfig_proto_depIdxs,
		MessageInfos:      file_config_appconfig_proto_msgTypes,
	}.Build()
	File_config_appconfig_proto = out.File
	file_config_appconfig_proto_rawDesc = nil
	file_config_appconfig_proto_goTypes = nil
	file_config_appconfig_proto_depIdxs = nil
}
