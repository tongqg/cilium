// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v3/outlier_detection.proto

package envoy_config_cluster_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
// [#next-free-field: 21]
type OutlierDetection struct {
	// The number of consecutive 5xx responses or local origin errors that are mapped
	// to 5xx error codes before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *duration.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status codes)
	// before a consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	// Determines whether to distinguish local origin failures from external errors. If set to true
	// the following configuration parameters are taken into account:
	// :ref:`consecutive_local_origin_failure<envoy_api_field_config.cluster.v3.OutlierDetection.consecutive_local_origin_failure>`,
	// :ref:`enforcing_consecutive_local_origin_failure<envoy_api_field_config.cluster.v3.OutlierDetection.enforcing_consecutive_local_origin_failure>`
	// and
	// :ref:`enforcing_local_origin_success_rate<envoy_api_field_config.cluster.v3.OutlierDetection.enforcing_local_origin_success_rate>`.
	// Defaults to false.
	SplitExternalLocalOriginErrors bool `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	ConsecutiveLocalOriginFailure *wrappers.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive locally originated failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingConsecutiveLocalOriginFailure *wrappers.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics for locally originated errors.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingLocalOriginSuccessRate *wrappers.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	// The failure percentage to use when determining failure percentage-based outlier detection. If
	// the failure percentage of a given host is greater than or equal to this value, it will be
	// ejected. Defaults to 85.
	FailurePercentageThreshold *wrappers.UInt32Value `protobuf:"bytes,16,opt,name=failure_percentage_threshold,json=failurePercentageThreshold,proto3" json:"failure_percentage_threshold,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status is detected through
	// failure percentage statistics. This setting can be used to disable ejection or to ramp it up
	// slowly. Defaults to 0.
	//
	// [#next-major-version: setting this without setting failure_percentage_threshold should be
	// invalid in v4.]
	EnforcingFailurePercentage *wrappers.UInt32Value `protobuf:"bytes,17,opt,name=enforcing_failure_percentage,json=enforcingFailurePercentage,proto3" json:"enforcing_failure_percentage,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status is detected through
	// local-origin failure percentage statistics. This setting can be used to disable ejection or to
	// ramp it up slowly. Defaults to 0.
	EnforcingFailurePercentageLocalOrigin *wrappers.UInt32Value `protobuf:"bytes,18,opt,name=enforcing_failure_percentage_local_origin,json=enforcingFailurePercentageLocalOrigin,proto3" json:"enforcing_failure_percentage_local_origin,omitempty"`
	// The minimum number of hosts in a cluster in order to perform failure percentage-based ejection.
	// If the total number of hosts in the cluster is less than this value, failure percentage-based
	// ejection will not be performed. Defaults to 5.
	FailurePercentageMinimumHosts *wrappers.UInt32Value `protobuf:"bytes,19,opt,name=failure_percentage_minimum_hosts,json=failurePercentageMinimumHosts,proto3" json:"failure_percentage_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one interval (as defined by the
	// interval duration above) to perform failure percentage-based ejection for this host. If the
	// volume is lower than this setting, failure percentage-based ejection will not be performed for
	// this host. Defaults to 50.
	FailurePercentageRequestVolume *wrappers.UInt32Value `protobuf:"bytes,20,opt,name=failure_percentage_request_volume,json=failurePercentageRequestVolume,proto3" json:"failure_percentage_request_volume,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}              `json:"-"`
	XXX_unrecognized               []byte                `json:"-"`
	XXX_sizecache                  int32                 `json:"-"`
}

func (m *OutlierDetection) Reset()         { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()    {}
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_31a4392bf965e39a, []int{0}
}

func (m *OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetection.Unmarshal(m, b)
}
func (m *OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetection.Marshal(b, m, deterministic)
}
func (m *OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetection.Merge(m, src)
}
func (m *OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_OutlierDetection.Size(m)
}
func (m *OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetection proto.InternalMessageInfo

func (m *OutlierDetection) GetConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *duration.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if m != nil {
		return m.SplitExternalLocalOriginErrors
	}
	return false
}

func (m *OutlierDetection) GetConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetFailurePercentageThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.FailurePercentageThreshold
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingFailurePercentage() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingFailurePercentage
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingFailurePercentageLocalOrigin() *wrappers.UInt32Value {
	if m != nil {
		return m.EnforcingFailurePercentageLocalOrigin
	}
	return nil
}

func (m *OutlierDetection) GetFailurePercentageMinimumHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.FailurePercentageMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetFailurePercentageRequestVolume() *wrappers.UInt32Value {
	if m != nil {
		return m.FailurePercentageRequestVolume
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.config.cluster.v3.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v3/outlier_detection.proto", fileDescriptor_31a4392bf965e39a)
}

var fileDescriptor_31a4392bf965e39a = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x49, 0xd8, 0xed, 0x86, 0x59, 0xd8, 0x86, 0xa1, 0x6c, 0x9c, 0xee, 0x6e, 0xe8, 0x16,
	0x15, 0x96, 0x0a, 0xd9, 0x52, 0xa3, 0xbd, 0xe1, 0x06, 0x29, 0x6c, 0x4a, 0x41, 0xa0, 0x86, 0xb4,
	0xb4, 0x42, 0x45, 0x1a, 0x4d, 0x9d, 0x13, 0x77, 0x90, 0x3d, 0x63, 0x66, 0xc6, 0xae, 0x7b, 0x05,
	0xdc, 0xf1, 0x0c, 0x3c, 0x02, 0x8f, 0xc0, 0x3d, 0x12, 0xb7, 0xbc, 0x07, 0x4f, 0xd0, 0x2b, 0xe4,
	0x7f, 0x89, 0x9d, 0xb8, 0x60, 0xdf, 0x45, 0x3a, 0xe7, 0xfb, 0x7d, 0xe7, 0x9c, 0x19, 0x9f, 0x0c,
	0xb2, 0x80, 0x87, 0xe2, 0xc6, 0xb2, 0x05, 0x9f, 0x33, 0xc7, 0xb2, 0xdd, 0x40, 0x69, 0x90, 0x56,
	0x38, 0xb4, 0x44, 0xa0, 0x5d, 0x06, 0x92, 0xcc, 0x40, 0x83, 0xad, 0x99, 0xe0, 0xa6, 0x2f, 0x85,
	0x16, 0xb8, 0x97, 0x08, 0xcc, 0x54, 0x60, 0x66, 0x02, 0x33, 0x1c, 0x6e, 0x0f, 0x1c, 0x21, 0x1c,
	0x17, 0xac, 0x24, 0xed, 0x32, 0x98, 0x5b, 0xb3, 0x40, 0xd2, 0xa5, 0x70, 0x3d, 0x7e, 0x2d, 0xa9,
	0xef, 0x83, 0x54, 0x59, 0xfc, 0x59, 0x30, 0xf3, 0xa9, 0x45, 0x39, 0x17, 0x3a, 0x91, 0x29, 0x4b,
	0x69, 0xaa, 0x83, 0x3c, 0xfc, 0x7c, 0x2d, 0x1c, 0x82, 0x54, 0x4c, 0x70, 0xc6, 0x9d, 0x2c, 0xa5,
	0x17, 0x52, 0x97, 0xcd, 0xa8, 0x06, 0x2b, 0xff, 0x91, 0x06, 0x76, 0xff, 0xd9, 0x44, 0xdd, 0xe3,
	0xb4, 0x9f, 0x57, 0x79, 0x3b, 0x78, 0x8c, 0x36, 0x6d, 0xc1, 0x15, 0xd8, 0x81, 0x66, 0x21, 0x90,
	0x97, 0x51, 0x64, 0xb4, 0x76, 0x5a, 0x2f, 0x1e, 0x1e, 0x3c, 0x35, 0xd3, 0x4a, 0xcd, 0xbc, 0x52,
	0xf3, 0xdb, 0x2f, 0xb8, 0x1e, 0x1e, 0x9c, 0x51, 0x37, 0x80, 0xe9, 0xa3, 0x82, 0xe8, 0x65, 0x14,
	0xe1, 0x4f, 0x51, 0x87, 0x71, 0x0d, 0x32, 0xa4, 0xae, 0xd1, 0x4e, 0xf4, 0xfd, 0x35, 0xfd, 0xab,
	0x6c, 0x12, 0xa3, 0xce, 0xed, 0xe8, 0xfe, 0xef, 0xad, 0xf6, 0xfe, 0x6b, 0xd3, 0x85, 0x08, 0x7f,
	0x83, 0xf0, 0x25, 0x55, 0x40, 0xe0, 0x87, 0xb4, 0x30, 0xa2, 0x99, 0x07, 0xc6, 0xeb, 0xf5, 0x51,
	0xdd, 0x58, 0x3e, 0xce, 0xd4, 0xa7, 0xcc, 0x03, 0x7c, 0x8e, 0xb6, 0x3c, 0x1a, 0x2d, 0x89, 0x3e,
	0x48, 0x1b, 0xb8, 0x36, 0xee, 0xfd, 0x7f, 0x7f, 0xa3, 0x07, 0xb7, 0xa3, 0x7b, 0xfb, 0x6d, 0x63,
	0x36, 0xc5, 0x1e, 0x8d, 0x72, 0xea, 0x24, 0x05, 0x60, 0x8a, 0xfa, 0xc0, 0xe7, 0x42, 0xda, 0x8c,
	0x3b, 0x64, 0x75, 0x7a, 0xf7, 0x9b, 0xd0, 0x7b, 0x0b, 0xce, 0x67, 0xe5, 0x79, 0x5e, 0xa0, 0xc7,
	0x4b, 0x0b, 0x15, 0xd8, 0x36, 0x28, 0x45, 0x24, 0xd5, 0x60, 0x6c, 0x34, 0xe1, 0x6f, 0x2d, 0x20,
	0x27, 0x29, 0x63, 0x4a, 0x35, 0xe0, 0xef, 0xd0, 0x76, 0x11, 0x49, 0x3c, 0xc6, 0x99, 0x17, 0x78,
	0xe4, 0x4a, 0x28, 0xad, 0x8c, 0x07, 0x35, 0x8e, 0xbf, 0xa7, 0x96, 0xb8, 0xaf, 0x53, 0xf5, 0x51,
	0x2c, 0xc6, 0x17, 0xe8, 0x49, 0x09, 0x2d, 0xe1, 0xc7, 0x00, 0x94, 0x26, 0xa1, 0x70, 0x03, 0x0f,
	0x8c, 0x4e, 0x0d, 0xb6, 0x51, 0x60, 0x4f, 0x53, 0xf9, 0x59, 0xa2, 0xc6, 0xe7, 0xa8, 0x5f, 0x82,
	0x2b, 0x3d, 0x83, 0x90, 0xcc, 0xa9, 0xad, 0x85, 0x34, 0xde, 0xa8, 0x81, 0x7e, 0x5c, 0x40, 0x9f,
	0xc4, 0xe2, 0xc3, 0x44, 0x8b, 0xbf, 0x47, 0x4f, 0x8a, 0xc7, 0xe8, 0x50, 0x0d, 0xd7, 0xf4, 0x86,
	0xcc, 0x29, 0x73, 0x03, 0x09, 0x06, 0xaa, 0x81, 0xee, 0x17, 0x00, 0x9f, 0xa7, 0xfa, 0xc3, 0x54,
	0x8e, 0x23, 0xb4, 0x57, 0x7d, 0x5d, 0x56, 0x7d, 0x1e, 0x36, 0x39, 0xda, 0xdd, 0xaa, 0xab, 0xb3,
	0xe2, 0xfc, 0x25, 0xda, 0x55, 0xbe, 0xcb, 0x34, 0x81, 0x48, 0x83, 0xe4, 0xd4, 0x25, 0xae, 0xb0,
	0xa9, 0x4b, 0x84, 0x64, 0x0e, 0xe3, 0x04, 0xa4, 0x14, 0x52, 0x19, 0x6f, 0xee, 0xb4, 0x5e, 0x74,
	0xa6, 0x83, 0x24, 0x73, 0x9c, 0x25, 0x7e, 0x15, 0xe7, 0x1d, 0x27, 0x69, 0xe3, 0x24, 0x0b, 0x03,
	0xda, 0x29, 0xd6, 0x5e, 0x02, 0xe5, 0x0d, 0xbc, 0x55, 0x63, 0x50, 0xcf, 0x0a, 0x94, 0x82, 0x4b,
	0x5e, 0xf2, 0x2f, 0x2d, 0xb4, 0x5f, 0x3d, 0xad, 0x4a, 0xc7, 0x47, 0x4d, 0x46, 0xf6, 0x41, 0xd5,
	0xc8, 0x2a, 0x6a, 0x50, 0xe8, 0xfd, 0x65, 0x09, 0x25, 0xdb, 0xd2, 0x97, 0xb8, 0xd9, 0xc4, 0xfb,
	0xbd, 0x05, 0xb1, 0x60, 0x58, 0xfc, 0x28, 0x1d, 0xf4, 0x34, 0x6b, 0x2a, 0x5f, 0x54, 0xd4, 0x01,
	0xa2, 0xaf, 0x24, 0xa8, 0x2b, 0xe1, 0xce, 0x8c, 0x6e, 0x13, 0xb7, 0xed, 0x0c, 0x35, 0x59, 0x90,
	0x4e, 0x73, 0x50, 0x6c, 0xb4, 0xec, 0x6e, 0xdd, 0xd2, 0x78, 0xbb, 0x91, 0xd1, 0x02, 0x75, 0xb8,
	0xea, 0x88, 0x7f, 0x42, 0x1f, 0xfd, 0x97, 0x51, 0x69, 0xb2, 0x06, 0x6e, 0xe2, 0xba, 0x77, 0xb7,
	0x6b, 0x61, 0xba, 0xf1, 0x95, 0xad, 0xb0, 0x2d, 0x6f, 0xbb, 0x77, 0xea, 0x5c, 0xd9, 0xb5, 0x69,
	0x96, 0x76, 0x9e, 0x83, 0x9e, 0x57, 0xd8, 0xac, 0x6c, 0xbe, 0xad, 0x1a, 0x3e, 0x83, 0x35, 0x9f,
	0xd2, 0xfe, 0xfb, 0xe4, 0xe3, 0xdf, 0xfe, 0xfc, 0x75, 0xf0, 0x61, 0xbc, 0x4d, 0xe2, 0xb7, 0x07,
	0xf5, 0x99, 0x19, 0x1e, 0x2c, 0xde, 0x1e, 0xab, 0xff, 0xec, 0xa3, 0xa3, 0x3f, 0x7e, 0xfe, 0xeb,
	0xef, 0x8d, 0x76, 0xb7, 0x8d, 0xf6, 0x98, 0x30, 0x13, 0x8d, 0x2f, 0x45, 0x74, 0x63, 0xde, 0xf1,
	0x74, 0x19, 0xbd, 0xbb, 0x8a, 0x98, 0xc4, 0x45, 0x4e, 0x5a, 0x97, 0x1b, 0x49, 0xb5, 0xc3, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xa1, 0xe3, 0x91, 0x26, 0x09, 0x00, 0x00,
}
