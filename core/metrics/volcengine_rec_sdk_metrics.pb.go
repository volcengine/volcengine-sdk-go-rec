// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: volcengine_rec_sdk_metrics.proto

package metrics

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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric    string            `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	Timestamp uint64            `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64           `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Tags      map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_rec_sdk_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_rec_sdk_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_volcengine_rec_sdk_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *Metric) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Metric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Metric) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type MetricMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricMessage) Reset() {
	*x = MetricMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_rec_sdk_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMessage) ProtoMessage() {}

func (x *MetricMessage) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_rec_sdk_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMessage.ProtoReflect.Descriptor instead.
func (*MetricMessage) Descriptor() ([]byte, []int) {
	return file_volcengine_rec_sdk_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *MetricMessage) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_volcengine_rec_sdk_metrics_proto protoreflect.FileDescriptor

var file_volcengine_rec_sdk_metrics_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x63,
	0x5f, 0x73, 0x64, 0x6b, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x22, 0xd1, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f,
	0x2d, 0x72, 0x65, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volcengine_rec_sdk_metrics_proto_rawDescOnce sync.Once
	file_volcengine_rec_sdk_metrics_proto_rawDescData = file_volcengine_rec_sdk_metrics_proto_rawDesc
)

func file_volcengine_rec_sdk_metrics_proto_rawDescGZIP() []byte {
	file_volcengine_rec_sdk_metrics_proto_rawDescOnce.Do(func() {
		file_volcengine_rec_sdk_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_volcengine_rec_sdk_metrics_proto_rawDescData)
	})
	return file_volcengine_rec_sdk_metrics_proto_rawDescData
}

var file_volcengine_rec_sdk_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_volcengine_rec_sdk_metrics_proto_goTypes = []interface{}{
	(*Metric)(nil),        // 0: volcengine.recommend.metrics.Metric
	(*MetricMessage)(nil), // 1: volcengine.recommend.metrics.MetricMessage
	nil,                   // 2: volcengine.recommend.metrics.Metric.TagsEntry
}
var file_volcengine_rec_sdk_metrics_proto_depIdxs = []int32{
	2, // 0: volcengine.recommend.metrics.Metric.tags:type_name -> volcengine.recommend.metrics.Metric.TagsEntry
	0, // 1: volcengine.recommend.metrics.MetricMessage.metrics:type_name -> volcengine.recommend.metrics.Metric
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_volcengine_rec_sdk_metrics_proto_init() }
func file_volcengine_rec_sdk_metrics_proto_init() {
	if File_volcengine_rec_sdk_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_volcengine_rec_sdk_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_volcengine_rec_sdk_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricMessage); i {
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
			RawDescriptor: file_volcengine_rec_sdk_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_volcengine_rec_sdk_metrics_proto_goTypes,
		DependencyIndexes: file_volcengine_rec_sdk_metrics_proto_depIdxs,
		MessageInfos:      file_volcengine_rec_sdk_metrics_proto_msgTypes,
	}.Build()
	File_volcengine_rec_sdk_metrics_proto = out.File
	file_volcengine_rec_sdk_metrics_proto_rawDesc = nil
	file_volcengine_rec_sdk_metrics_proto_goTypes = nil
	file_volcengine_rec_sdk_metrics_proto_depIdxs = nil
}
