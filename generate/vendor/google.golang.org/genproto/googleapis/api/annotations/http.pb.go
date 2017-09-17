// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/http.proto

package annotations

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the HTTP configuration for a service. It contains a list of
// [HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method
// to one or more HTTP REST API methods.
type Http struct {
	// A list of HTTP configuration rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*HttpRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *Http) Reset()                    { *m = Http{} }
func (m *Http) String() string            { return proto.CompactTextString(m) }
func (*Http) ProtoMessage()               {}
func (*Http) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Http) GetRules() []*HttpRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// `HttpRule` defines the mapping of an RPC method to one or more HTTP
// REST APIs.  The mapping determines what portions of the request
// message are populated from the path, query parameters, or body of
// the HTTP request.  The mapping is typically specified as an
// `google.api.http` annotation, see "google/api/annotations.proto"
// for details.
//
// The mapping consists of a field specifying the path template and
// method kind.  The path template can refer to fields in the request
// message, as in the example below which describes a REST GET
// operation on a resource collection of messages:
//
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http).get = "/v1/messages/{message_id}/{sub.subfield}";
//       }
//     }
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // mapped to the URL
//       SubMessage sub = 2;    // `sub.subfield` is url-mapped
//     }
//     message Message {
//       string text = 1; // content of the resource
//     }
//
// The same http annotation can alternatively be expressed inside the
// `GRPC API Configuration` YAML file.
//
//     http:
//       rules:
//         - selector: <proto_package_name>.Messaging.GetMessage
//           get: /v1/messages/{message_id}/{sub.subfield}
//
// This definition enables an automatic, bidrectional mapping of HTTP
// JSON to RPC. Example:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456/foo`  | `GetMessage(message_id: "123456" sub: SubMessage(subfield: "foo"))`
//
// In general, not only fields but also field paths can be referenced
// from a path pattern. Fields mapped to the path pattern cannot be
// repeated and must have a primitive (non-message) type.
//
// Any fields in the request message which are not bound by the path
// pattern automatically become (optional) HTTP query
// parameters. Assume the following definition of the request message:
//
//
//     message GetMessageRequest {
//       message SubMessage {
//         string subfield = 1;
//       }
//       string message_id = 1; // mapped to the URL
//       int64 revision = 2;    // becomes a parameter
//       SubMessage sub = 3;    // `sub.subfield` becomes a parameter
//     }
//
//
// This enables a HTTP JSON to RPC mapping as below:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456?revision=2&sub.subfield=foo` | `GetMessage(message_id: "123456" revision: 2 sub: SubMessage(subfield: "foo"))`
//
// Note that fields which are mapped to HTTP parameters must have a
// primitive type or a repeated primitive type. Message types are not
// allowed. In the case of a repeated type, the parameter can be
// repeated in the URL, as in `...?param=A&param=B`.
//
// For HTTP method kinds which allow a request body, the `body` field
// specifies the mapping. Consider a REST update method on the
// message resource collection:
//
//
//     service Messaging {
//       rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           put: "/v1/messages/{message_id}"
//           body: "message"
//         };
//       }
//     }
//     message UpdateMessageRequest {
//       string message_id = 1; // mapped to the URL
//       Message message = 2;   // mapped to the body
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled, where the
// representation of the JSON in the request body is determined by
// protos JSON encoding:
//
// HTTP | RPC
// -----|-----
// `PUT /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id: "123456" message { text: "Hi!" })`
//
// The special name `*` can be used in the body mapping to define that
// every field not bound by the path template should be mapped to the
// request body.  This enables the following alternative definition of
// the update method:
//
//     service Messaging {
//       rpc UpdateMessage(Message) returns (Message) {
//         option (google.api.http) = {
//           put: "/v1/messages/{message_id}"
//           body: "*"
//         };
//       }
//     }
//     message Message {
//       string message_id = 1;
//       string text = 2;
//     }
//
//
// The following HTTP JSON to RPC mapping is enabled:
//
// HTTP | RPC
// -----|-----
// `PUT /v1/messages/123456 { "text": "Hi!" }` | `UpdateMessage(message_id: "123456" text: "Hi!")`
//
// Note that when using `*` in the body mapping, it is not possible to
// have HTTP parameters, as all fields not bound by the path end in
// the body. This makes this option more rarely used in practice of
// defining REST APIs. The common usage of `*` is in custom methods
// which don't use the URL at all for transferring data.
//
// It is possible to define multiple HTTP methods for one RPC by using
// the `additional_bindings` option. Example:
//
//     service Messaging {
//       rpc GetMessage(GetMessageRequest) returns (Message) {
//         option (google.api.http) = {
//           get: "/v1/messages/{message_id}"
//           additional_bindings {
//             get: "/v1/users/{user_id}/messages/{message_id}"
//           }
//         };
//       }
//     }
//     message GetMessageRequest {
//       string message_id = 1;
//       string user_id = 2;
//     }
//
//
// This enables the following two alternative HTTP JSON to RPC
// mappings:
//
// HTTP | RPC
// -----|-----
// `GET /v1/messages/123456` | `GetMessage(message_id: "123456")`
// `GET /v1/users/me/messages/123456` | `GetMessage(user_id: "me" message_id: "123456")`
//
// # Rules for HTTP mapping
//
// The rules for mapping HTTP path, query parameters, and body fields
// to the request message are as follows:
//
// 1. The `body` field specifies either `*` or a field path, or is
//    omitted. If omitted, it assumes there is no HTTP body.
// 2. Leaf fields (recursive expansion of nested messages in the
//    request) can be classified into three types:
//     (a) Matched in the URL template.
//     (b) Covered by body (if body is `*`, everything except (a) fields;
//         else everything under the body field)
//     (c) All other fields.
// 3. URL query parameters found in the HTTP request are mapped to (c) fields.
// 4. Any body sent with an HTTP request can contain only (b) fields.
//
// The syntax of the path template is as follows:
//
//     Template = "/" Segments [ Verb ] ;
//     Segments = Segment { "/" Segment } ;
//     Segment  = "*" | "**" | LITERAL | Variable ;
//     Variable = "{" FieldPath [ "=" Segments ] "}" ;
//     FieldPath = IDENT { "." IDENT } ;
//     Verb     = ":" LITERAL ;
//
// The syntax `*` matches a single path segment. It follows the semantics of
// [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
// Expansion.
//
// The syntax `**` matches zero or more path segments. It follows the semantics
// of [RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.3 Reserved
// Expansion. NOTE: it must be the last segment in the path except the Verb.
//
// The syntax `LITERAL` matches literal text in the URL path.
//
// The syntax `Variable` matches the entire path as specified by its template;
// this nested template must not contain further variables. If a variable
// matches a single path segment, its template may be omitted, e.g. `{var}`
// is equivalent to `{var=*}`.
//
// NOTE: the field paths in variables and in the `body` must not refer to
// repeated fields or map fields.
//
// Use CustomHttpPattern to specify any HTTP method that is not included in the
// `pattern` field, such as HEAD, or "*" to leave the HTTP method unspecified for
// a given URL path rule. The wild-card rule is useful for services that provide
// content to Web (HTML) clients.
type HttpRule struct {
	// Selects methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// Determines the URL pattern is matched by this rules. This pattern can be
	// used with any of the {get|put|post|delete|patch} methods. A custom method
	// can be defined using the 'custom' field.
	//
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern isHttpRule_Pattern `protobuf_oneof:"pattern"`
	// The name of the request field whose value is mapped to the HTTP body, or
	// `*` for mapping all fields not captured by the path pattern to the HTTP
	// body. NOTE: the referred field must not be a repeated field and must be
	// present at the top-level of request message type.
	Body string `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	// Additional HTTP bindings for the selector. Nested bindings must
	// not contain an `additional_bindings` field themselves (that is,
	// the nesting may only be one level deep).
	AdditionalBindings []*HttpRule `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings" json:"additional_bindings,omitempty"`
}

func (m *HttpRule) Reset()                    { *m = HttpRule{} }
func (m *HttpRule) String() string            { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()               {}
func (*HttpRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,oneof"`
}
type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,oneof"`
}
type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,oneof"`
}
type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,oneof"`
}
type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,oneof"`
}
type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern()    {}
func (*HttpRule_Put) isHttpRule_Pattern()    {}
func (*HttpRule_Post) isHttpRule_Pattern()   {}
func (*HttpRule_Delete) isHttpRule_Pattern() {}
func (*HttpRule_Patch) isHttpRule_Pattern()  {}
func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpRule_OneofMarshaler, _HttpRule_OneofUnmarshaler, _HttpRule_OneofSizer, []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

func _HttpRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Get)
	case *HttpRule_Put:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Put)
	case *HttpRule_Post:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Post)
	case *HttpRule_Delete:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Delete)
	case *HttpRule_Patch:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Patch)
	case *HttpRule_Custom:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HttpRule.Pattern has unexpected type %T", x)
	}
	return nil
}

func _HttpRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpRule)
	switch tag {
	case 2: // pattern.get
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Get{x}
		return true, err
	case 3: // pattern.put
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Put{x}
		return true, err
	case 4: // pattern.post
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Post{x}
		return true, err
	case 5: // pattern.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Delete{x}
		return true, err
	case 6: // pattern.patch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Patch{x}
		return true, err
	case 8: // pattern.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CustomHttpPattern)
		err := b.DecodeMessage(msg)
		m.Pattern = &HttpRule_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HttpRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Get)))
		n += len(x.Get)
	case *HttpRule_Put:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Put)))
		n += len(x.Put)
	case *HttpRule_Post:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Post)))
		n += len(x.Post)
	case *HttpRule_Delete:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Delete)))
		n += len(x.Delete)
	case *HttpRule_Patch:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Patch)))
		n += len(x.Patch)
	case *HttpRule_Custom:
		s := proto.Size(x.Custom)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A custom pattern is used for defining custom HTTP verb.
type CustomHttpPattern struct {
	// The name of this custom HTTP verb.
	Kind string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	// The path matched by this custom verb.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *CustomHttpPattern) Reset()                    { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string            { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()               {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Http)(nil), "google.api.Http")
	proto.RegisterType((*HttpRule)(nil), "google.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "google.api.CustomHttpPattern")
}

func init() { proto.RegisterFile("google/api/http.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xd7, 0x89, 0xe3, 0x24, 0x13, 0x58, 0x58, 0x6d, 0x76, 0x11, 0x85, 0x42, 0xc8, 0x29,
	0xf4, 0x60, 0x43, 0x7a, 0xe8, 0x21, 0xa7, 0xb8, 0x94, 0xa6, 0xb7, 0xe0, 0x63, 0x2f, 0x45, 0xb1,
	0x85, 0xa2, 0xd6, 0x91, 0x84, 0x3d, 0x3e, 0xf4, 0x75, 0xfa, 0x0e, 0x7d, 0xb7, 0x1e, 0x8b, 0xfe,
	0xa4, 0x09, 0x14, 0x7a, 0x9b, 0xef, 0x37, 0x9f, 0x34, 0xa3, 0x19, 0xc1, 0x3f, 0xa1, 0xb5, 0xa8,
	0x79, 0xc6, 0x8c, 0xcc, 0xf6, 0x88, 0x26, 0x35, 0x8d, 0x46, 0x4d, 0xc0, 0xe3, 0x94, 0x19, 0x39,
	0x5f, 0x42, 0xbc, 0x41, 0x34, 0xe4, 0x0a, 0x06, 0x4d, 0x57, 0xf3, 0x96, 0x46, 0xb3, 0xfe, 0x62,
	0xb2, 0x9c, 0xa6, 0x27, 0x4f, 0x6a, 0x0d, 0x45, 0x57, 0xf3, 0xc2, 0x5b, 0xe6, 0xef, 0x3d, 0x18,
	0x1d, 0x19, 0xb9, 0x80, 0x51, 0xcb, 0x6b, 0x5e, 0xa2, 0x6e, 0x68, 0x34, 0x8b, 0x16, 0xe3, 0xe2,
	0x4b, 0x13, 0x02, 0x7d, 0xc1, 0x91, 0xf6, 0x2c, 0xde, 0xfc, 0x2a, 0xac, 0xb0, 0xcc, 0x74, 0x48,
	0xfb, 0x47, 0x66, 0x3a, 0x24, 0x53, 0x88, 0x8d, 0x6e, 0x91, 0xc6, 0x01, 0x3a, 0x45, 0x28, 0x24,
	0x15, 0xaf, 0x39, 0x72, 0x3a, 0x08, 0x3c, 0x68, 0xf2, 0x1f, 0x06, 0x86, 0x61, 0xb9, 0xa7, 0x49,
	0x48, 0x78, 0x49, 0x6e, 0x20, 0x29, 0xbb, 0x16, 0xf5, 0x81, 0x8e, 0x66, 0xd1, 0x62, 0xb2, 0xbc,
	0x3c, 0x7f, 0xc5, 0xad, 0xcb, 0xd8, 0xbe, 0xb7, 0x0c, 0x91, 0x37, 0xca, 0x5e, 0xe8, 0xed, 0x84,
	0x40, 0xbc, 0xd3, 0xd5, 0x2b, 0x1d, 0xba, 0x07, 0xb8, 0x98, 0xdc, 0xc1, 0x5f, 0x56, 0x55, 0x12,
	0xa5, 0x56, 0xac, 0x7e, 0xda, 0x49, 0x55, 0x49, 0x25, 0x5a, 0x3a, 0xf9, 0x61, 0x3e, 0xe4, 0x74,
	0x20, 0x0f, 0xfe, 0x7c, 0x0c, 0x43, 0xe3, 0xeb, 0xcd, 0x57, 0xf0, 0xe7, 0x5b, 0x13, 0xb6, 0xf4,
	0x8b, 0x54, 0x55, 0x98, 0x9d, 0x8b, 0x2d, 0x33, 0x0c, 0xf7, 0x7e, 0x70, 0x85, 0x8b, 0xf3, 0x67,
	0xf8, 0x5d, 0xea, 0xc3, 0x59, 0xd9, 0x7c, 0xec, 0xae, 0xb1, 0x1b, 0xdd, 0x46, 0x8f, 0xeb, 0x90,
	0x10, 0xba, 0x66, 0x4a, 0xa4, 0xba, 0x11, 0x99, 0xe0, 0xca, 0xed, 0x3b, 0xf3, 0x29, 0x66, 0x64,
	0xeb, 0x7e, 0x02, 0x53, 0x4a, 0x23, 0xb3, 0x6d, 0xb6, 0xab, 0xb3, 0xf8, 0x23, 0x8a, 0xde, 0x7a,
	0xf1, 0xfd, 0x7a, 0xfb, 0xb0, 0x4b, 0xdc, 0xb9, 0xeb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68,
	0x15, 0x60, 0x5b, 0x40, 0x02, 0x00, 0x00,
}
