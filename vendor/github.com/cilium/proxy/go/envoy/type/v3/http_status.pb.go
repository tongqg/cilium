// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/v3/http_status.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// HTTP response codes supported in Envoy.
// For more details: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int32

const (
	// Empty - This code not part of the HTTP status code specification, but it is needed for proto
	// `enum` type.
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}

var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_588aaadec77e6b51, []int{0}
}

// HTTP status.
type HttpStatus struct {
	// Supplies HTTP response code.
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.type.v3.StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpStatus) Reset()         { *m = HttpStatus{} }
func (m *HttpStatus) String() string { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()    {}
func (*HttpStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_588aaadec77e6b51, []int{0}
}

func (m *HttpStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpStatus.Unmarshal(m, b)
}
func (m *HttpStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpStatus.Marshal(b, m, deterministic)
}
func (m *HttpStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStatus.Merge(m, src)
}
func (m *HttpStatus) XXX_Size() int {
	return xxx_messageInfo_HttpStatus.Size(m)
}
func (m *HttpStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStatus proto.InternalMessageInfo

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterEnum("envoy.type.v3.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.v3.HttpStatus")
}

func init() { proto.RegisterFile("envoy/type/v3/http_status.proto", fileDescriptor_588aaadec77e6b51) }

var fileDescriptor_588aaadec77e6b51 = []byte{
	// 975 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x73, 0x1b, 0xc5,
	0x17, 0xcf, 0x68, 0xb2, 0xb9, 0x13, 0xdb, 0x2f, 0x9d, 0xf8, 0xef, 0xc4, 0x49, 0xfe, 0x98, 0x9c,
	0x28, 0x0e, 0x56, 0x15, 0xbe, 0x00, 0x37, 0xdb, 0xb1, 0x63, 0x17, 0x96, 0xa2, 0x92, 0xa5, 0x5c,
	0xa9, 0xf6, 0xf4, 0x93, 0xd4, 0x95, 0x51, 0xbf, 0x49, 0xcf, 0x93, 0xec, 0xe1, 0x44, 0x71, 0xe2,
	0xc8, 0xbe, 0x84, 0xfd, 0xc0, 0x52, 0x54, 0x42, 0xa0, 0x80, 0x0b, 0x77, 0xaa, 0xc2, 0x0e, 0x5f,
	0x81, 0xcf, 0xc0, 0x1a, 0x28, 0xa0, 0xba, 0xb5, 0xd8, 0x29, 0x8a, 0x9b, 0xf4, 0xfa, 0x2d, 0xbf,
	0xe5, 0xcd, 0x13, 0xf7, 0xa0, 0xed, 0x53, 0x51, 0xe6, 0x22, 0xc3, 0x72, 0x7f, 0xb1, 0xdc, 0x61,
	0xce, 0x1e, 0xcd, 0x59, 0x71, 0x2f, 0x5f, 0xc8, 0x1c, 0x31, 0xc9, 0xc9, 0x90, 0xb0, 0xe0, 0x13,
	0x16, 0xfa, 0x8b, 0x73, 0xe7, 0x7b, 0x3a, 0x53, 0x65, 0x65, 0x2d, 0xb1, 0x62, 0x43, 0x36, 0x2f,
	0xef, 0xcf, 0x9e, 0xbb, 0xf7, 0x5f, 0xcf, 0x7d, 0x74, 0xb9, 0x21, 0x6b, 0x6c, 0x7b, 0x98, 0x32,
	0xdb, 0x57, 0xa9, 0xd1, 0x8a, 0xb1, 0x3c, 0xfa, 0x31, 0x78, 0xb8, 0x80, 0x42, 0xac, 0x33, 0x67,
	0x5b, 0xa1, 0x9f, 0x7c, 0x48, 0x1c, 0x4c, 0x48, 0xe3, 0xe9, 0x68, 0x3e, 0xba, 0x6f, 0xea, 0x81,
	0x33, 0x0b, 0x77, 0xc1, 0x58, 0x18, 0x24, 0xad, 0x90, 0xc6, 0x65, 0x71, 0x67, 0xf9, 0xc8, 0x13,
	0xd1, 0x41, 0x88, 0xe6, 0x0f, 0xd4, 0x43, 0xc9, 0xc3, 0xe7, 0xae, 0x7f, 0xfe, 0xe4, 0xff, 0x67,
	0xc5, 0xcc, 0xbe, 0x92, 0xbd, 0xc6, 0xf7, 0x7f, 0x3a, 0x21, 0xc4, 0x5e, 0xb9, 0x9c, 0x10, 0x87,
	0x56, 0xbb, 0x19, 0x17, 0x70, 0x40, 0x1e, 0x17, 0x47, 0x57, 0xc8, 0xb2, 0xb1, 0x3d, 0x04, 0x2d,
	0x8f, 0x88, 0xd2, 0xe5, 0x47, 0xe0, 0x76, 0x24, 0x8f, 0x8b, 0x23, 0x2b, 0x0e, 0x15, 0xa3, 0x86,
	0x2f, 0x22, 0x39, 0x29, 0x8e, 0x2e, 0x25, 0x09, 0x66, 0xfe, 0xef, 0x97, 0x91, 0x9c, 0x17, 0x67,
	0xab, 0x64, 0x97, 0x7a, 0xdc, 0x21, 0x67, 0x3c, 0xe7, 0x3e, 0x6e, 0xd8, 0x16, 0xb9, 0x6e, 0xa0,
	0x0f, 0x5f, 0x45, 0x72, 0x4a, 0x4c, 0x54, 0xc9, 0xf7, 0x45, 0xcb, 0xf0, 0x75, 0x24, 0x4f, 0x88,
	0xe3, 0x75, 0xcc, 0x91, 0x47, 0xa1, 0x6f, 0x22, 0x79, 0x52, 0x4c, 0xd5, 0x94, 0x63, 0xa3, 0xd2,
	0x51, 0xf0, 0xdb, 0x48, 0x82, 0x38, 0x56, 0xe9, 0xa5, 0x6c, 0x06, 0x58, 0xe1, 0xbb, 0x48, 0x9e,
	0x12, 0xd3, 0x4b, 0xa9, 0x43, 0xa5, 0x8b, 0x3a, 0x66, 0xe4, 0x3c, 0x82, 0xef, 0x23, 0x79, 0x4c,
	0x1c, 0xde, 0xa8, 0x34, 0x73, 0xd4, 0xf0, 0x63, 0x48, 0x09, 0x45, 0x59, 0x8a, 0x2b, 0x1d, 0x32,
	0x09, 0xe6, 0x70, 0xa3, 0x24, 0x67, 0x04, 0x54, 0xa8, 0x8f, 0xba, 0x86, 0xae, 0xab, 0x2c, 0x5a,
	0x4e, 0x0b, 0xb8, 0x59, 0x92, 0x42, 0x1c, 0x5a, 0xa3, 0x9e, 0xd5, 0xf0, 0x41, 0xc9, 0xd3, 0xda,
	0x42, 0xbc, 0xcc, 0x1d, 0x74, 0x70, 0xab, 0xe4, 0x87, 0x57, 0x89, 0x2b, 0xa4, 0x4d, 0xcb, 0xa0,
	0x86, 0x0f, 0x43, 0x42, 0x33, 0xc7, 0x9a, 0xa3, 0xdd, 0x02, 0x3e, 0x2a, 0xc9, 0xff, 0x89, 0x13,
	0x0d, 0xec, 0x66, 0xe4, 0x94, 0x2b, 0xea, 0xa8, 0x8d, 0xc3, 0x84, 0xe1, 0xe3, 0x10, 0x1f, 0x4f,
	0x19, 0xc7, 0x3f, 0x29, 0xc9, 0x69, 0x21, 0x96, 0x95, 0xae, 0xe3, 0xb5, 0x1e, 0xe6, 0x0c, 0x4f,
	0xc5, 0x5e, 0x86, 0xa6, 0x55, 0x03, 0xdd, 0x1e, 0x43, 0x0d, 0x4f, 0xc7, 0x1e, 0x7c, 0x4d, 0x15,
	0xdd, 0x50, 0x79, 0xad, 0x67, 0x1c, 0x6a, 0x78, 0x26, 0xf6, 0xfa, 0xad, 0x91, 0xdb, 0x36, 0x5a,
	0xa3, 0x85, 0x67, 0x63, 0x0f, 0xa4, 0x4a, 0x3c, 0x00, 0xfe, 0x5c, 0x1c, 0xb8, 0x21, 0x77, 0x48,
	0x57, 0x89, 0x97, 0xd2, 0x94, 0x76, 0x50, 0xc3, 0xf3, 0xb1, 0x94, 0x62, 0xd2, 0x07, 0x82, 0x53,
	0x6a, 0x3b, 0x45, 0x78, 0x21, 0xf6, 0x5e, 0x05, 0xfc, 0xde, 0x2d, 0xb4, 0x6c, 0x92, 0xe0, 0xd1,
	0x78, 0xd6, 0x8b, 0xb1, 0x37, 0x62, 0x08, 0xb1, 0x61, 0xba, 0x48, 0x3d, 0x86, 0x97, 0xc2, 0xc0,
	0x15, 0xb2, 0xad, 0xd4, 0x24, 0x0c, 0x2f, 0xc7, 0x72, 0x42, 0x1c, 0xbc, 0x44, 0x16, 0xe1, 0x7a,
	0x48, 0xdf, 0x44, 0xdb, 0xe6, 0xce, 0xb8, 0xc7, 0x2b, 0xb1, 0x9c, 0x15, 0xb2, 0xe6, 0x30, 0x21,
	0xab, 0x8d, 0x6f, 0xbf, 0xa6, 0x4c, 0x8a, 0x1a, 0x5e, 0x1d, 0xd1, 0x4b, 0x49, 0xe9, 0x06, 0xd1,
	0xa6, 0x72, 0x6d, 0x84, 0xd7, 0x62, 0x2f, 0x4c, 0xb3, 0xbe, 0xe1, 0x23, 0x64, 0xdb, 0xf0, 0x7a,
	0x2c, 0xcf, 0x88, 0x53, 0x4d, 0x9b, 0xf7, 0xb2, 0x81, 0xc3, 0x15, 0xd4, 0x46, 0x35, 0x8a, 0x0c,
	0xe1, 0x8d, 0x58, 0x9e, 0x16, 0x27, 0xeb, 0xca, 0xb6, 0xb1, 0x4a, 0xbc, 0xa5, 0xd8, 0xe4, 0x2d,
	0x13, 0xa8, 0xbd, 0x19, 0x7b, 0xd9, 0x57, 0x77, 0x33, 0x4c, 0x06, 0x5f, 0xdd, 0x70, 0xe6, 0x5b,
	0x01, 0x4c, 0xc5, 0xe4, 0x03, 0x1b, 0x70, 0x2c, 0xff, 0xdb, 0xa1, 0x55, 0xd3, 0x66, 0x8e, 0x12,
	0xcc, 0x73, 0xdf, 0x64, 0xd5, 0xb2, 0xe1, 0x02, 0xde, 0x89, 0xfd, 0x3e, 0x6d, 0x52, 0x72, 0x15,
	0x35, 0xbc, 0x1b, 0xd4, 0x1d, 0x34, 0xbb, 0x88, 0x19, 0x5a, 0x8d, 0x36, 0x29, 0xe0, 0xbd, 0x40,
	0xa5, 0x99, 0xb5, 0x9d, 0xd2, 0x38, 0x66, 0xfe, 0x7e, 0x40, 0xbe, 0x9f, 0xf9, 0xf8, 0xe9, 0x46,
	0x28, 0x68, 0x10, 0x55, 0x94, 0x2d, 0x86, 0x18, 0x72, 0xb8, 0x19, 0x0c, 0x19, 0xfe, 0x5d, 0x47,
	0xa5, 0xd1, 0xad, 0x19, 0x4c, 0x75, 0x3e, 0x56, 0xe7, 0x56, 0x80, 0xb9, 0x61, 0x19, 0x9d, 0x55,
	0xe9, 0x16, 0xba, 0x3e, 0xba, 0x55, 0xe7, 0xc8, 0xc1, 0x4f, 0x41, 0xfb, 0x2a, 0xf1, 0x46, 0x37,
	0x4b, 0xd1, 0x6f, 0x0c, 0x6a, 0xf8, 0x39, 0x1e, 0x6e, 0xd9, 0x25, 0xc5, 0xb8, 0xa3, 0x0a, 0xf8,
	0x25, 0xf0, 0xf7, 0x75, 0x26, 0xc1, 0xa6, 0x55, 0x7d, 0x65, 0xd2, 0x20, 0xd8, 0xaf, 0xa1, 0x7c,
	0x98, 0x36, 0x72, 0xfa, 0xb7, 0x58, 0x9e, 0x13, 0xb3, 0xeb, 0x8d, 0x46, 0xed, 0xca, 0xe0, 0x64,
	0x79, 0x95, 0x47, 0x36, 0xc0, 0xef, 0xb1, 0x9c, 0x13, 0x33, 0x57, 0x94, 0x33, 0xca, 0xf2, 0x52,
	0x9a, 0x53, 0x15, 0xdb, 0xc4, 0x46, 0x31, 0xe6, 0x70, 0x67, 0x88, 0x33, 0xef, 0xb5, 0x5a, 0x26,
	0x31, 0x68, 0x79, 0x8b, 0xc9, 0xa9, 0x36, 0xc2, 0x1f, 0x61, 0xcf, 0x37, 0x89, 0xb2, 0x8b, 0xc8,
	0xc1, 0x02, 0xf8, 0x33, 0x1e, 0x7e, 0x5c, 0xab, 0xbb, 0xec, 0x15, 0xd5, 0xf0, 0x57, 0x2c, 0x2f,
	0x88, 0xf3, 0x55, 0xe4, 0x1d, 0x72, 0x57, 0xff, 0x63, 0x37, 0xff, 0x8e, 0x97, 0x1f, 0xfc, 0xec,
	0xf1, 0xdb, 0x3f, 0x1c, 0x2e, 0x41, 0x49, 0x9c, 0x35, 0x34, 0x38, 0x87, 0x99, 0xdf, 0xe6, 0xbb,
	0x2f, 0xe3, 0xf2, 0xf4, 0xde, 0xa9, 0xab, 0xf9, 0xb3, 0x5a, 0x8b, 0xb6, 0x0f, 0x87, 0xfb, 0xba,
	0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x01, 0x4d, 0xee, 0xec, 0x05, 0x00, 0x00,
}
