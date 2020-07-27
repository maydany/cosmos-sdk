// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/tx/sig_desc.proto

package tx

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SignatureDescriptors wraps multiple SignatureDescriptor's
type SignatureDescriptors struct {
	Signatures []*SignatureDescriptor `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *SignatureDescriptors) Reset()         { *m = SignatureDescriptors{} }
func (m *SignatureDescriptors) String() string { return proto.CompactTextString(m) }
func (*SignatureDescriptors) ProtoMessage()    {}
func (*SignatureDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_d560eb6bd3d62ef7, []int{0}
}
func (m *SignatureDescriptors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureDescriptors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureDescriptors.Merge(m, src)
}
func (m *SignatureDescriptors) XXX_Size() int {
	return m.Size()
}
func (m *SignatureDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureDescriptors proto.InternalMessageInfo

func (m *SignatureDescriptors) GetSignatures() []*SignatureDescriptor {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// SignatureDescriptor is a convenience type which represents the full data for a
// signature including the public key of the signer, signing modes and the signature itself
type SignatureDescriptor struct {
	// public_key is the public key of the signer
	PublicKey *types.PublicKey `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// data represents the signature data
	Data *SignatureDescriptor_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SignatureDescriptor) Reset()         { *m = SignatureDescriptor{} }
func (m *SignatureDescriptor) String() string { return proto.CompactTextString(m) }
func (*SignatureDescriptor) ProtoMessage()    {}
func (*SignatureDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_d560eb6bd3d62ef7, []int{1}
}
func (m *SignatureDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureDescriptor.Merge(m, src)
}
func (m *SignatureDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *SignatureDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureDescriptor proto.InternalMessageInfo

func (m *SignatureDescriptor) GetPublicKey() *types.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignatureDescriptor) GetData() *SignatureDescriptor_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type SignatureDescriptor_Data struct {
	// sum is the oneof that specifies whether this represents single or multi-signature data
	//
	// Types that are valid to be assigned to Sum:
	//	*SignatureDescriptor_Data_Single_
	//	*SignatureDescriptor_Data_Multi_
	Sum isSignatureDescriptor_Data_Sum `protobuf_oneof:"sum"`
}

func (m *SignatureDescriptor_Data) Reset()         { *m = SignatureDescriptor_Data{} }
func (m *SignatureDescriptor_Data) String() string { return proto.CompactTextString(m) }
func (*SignatureDescriptor_Data) ProtoMessage()    {}
func (*SignatureDescriptor_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_d560eb6bd3d62ef7, []int{1, 0}
}
func (m *SignatureDescriptor_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureDescriptor_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureDescriptor_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureDescriptor_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureDescriptor_Data.Merge(m, src)
}
func (m *SignatureDescriptor_Data) XXX_Size() int {
	return m.Size()
}
func (m *SignatureDescriptor_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureDescriptor_Data.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureDescriptor_Data proto.InternalMessageInfo

type isSignatureDescriptor_Data_Sum interface {
	isSignatureDescriptor_Data_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SignatureDescriptor_Data_Single_ struct {
	Single *SignatureDescriptor_Data_Single `protobuf:"bytes,1,opt,name=single,proto3,oneof" json:"single,omitempty"`
}
type SignatureDescriptor_Data_Multi_ struct {
	Multi *SignatureDescriptor_Data_Multi `protobuf:"bytes,2,opt,name=multi,proto3,oneof" json:"multi,omitempty"`
}

func (*SignatureDescriptor_Data_Single_) isSignatureDescriptor_Data_Sum() {}
func (*SignatureDescriptor_Data_Multi_) isSignatureDescriptor_Data_Sum()  {}

func (m *SignatureDescriptor_Data) GetSum() isSignatureDescriptor_Data_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *SignatureDescriptor_Data) GetSingle() *SignatureDescriptor_Data_Single {
	if x, ok := m.GetSum().(*SignatureDescriptor_Data_Single_); ok {
		return x.Single
	}
	return nil
}

func (m *SignatureDescriptor_Data) GetMulti() *SignatureDescriptor_Data_Multi {
	if x, ok := m.GetSum().(*SignatureDescriptor_Data_Multi_); ok {
		return x.Multi
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignatureDescriptor_Data) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignatureDescriptor_Data_Single_)(nil),
		(*SignatureDescriptor_Data_Multi_)(nil),
	}
}

// Single is the signature data for a single signer
type SignatureDescriptor_Data_Single struct {
	// mode is the signing mode of the single signer
	Mode signing.SignMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cosmos.tx.signing.SignMode" json:"mode,omitempty"`
	// signature is the raw signature bytes
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureDescriptor_Data_Single) Reset()         { *m = SignatureDescriptor_Data_Single{} }
func (m *SignatureDescriptor_Data_Single) String() string { return proto.CompactTextString(m) }
func (*SignatureDescriptor_Data_Single) ProtoMessage()    {}
func (*SignatureDescriptor_Data_Single) Descriptor() ([]byte, []int) {
	return fileDescriptor_d560eb6bd3d62ef7, []int{1, 0, 0}
}
func (m *SignatureDescriptor_Data_Single) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureDescriptor_Data_Single) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureDescriptor_Data_Single.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureDescriptor_Data_Single) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureDescriptor_Data_Single.Merge(m, src)
}
func (m *SignatureDescriptor_Data_Single) XXX_Size() int {
	return m.Size()
}
func (m *SignatureDescriptor_Data_Single) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureDescriptor_Data_Single.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureDescriptor_Data_Single proto.InternalMessageInfo

func (m *SignatureDescriptor_Data_Single) GetMode() signing.SignMode {
	if m != nil {
		return m.Mode
	}
	return signing.SignMode_SIGN_MODE_UNSPECIFIED
}

func (m *SignatureDescriptor_Data_Single) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Multi is the signature data for a multisig public key
type SignatureDescriptor_Data_Multi struct {
	// bitarray specifies which keys within the multisig are signing
	Bitarray *types.CompactBitArray `protobuf:"bytes,1,opt,name=bitarray,proto3" json:"bitarray,omitempty"`
	// signatures is the signatures of the multi-signature
	Signatures []*SignatureDescriptor_Data `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *SignatureDescriptor_Data_Multi) Reset()         { *m = SignatureDescriptor_Data_Multi{} }
func (m *SignatureDescriptor_Data_Multi) String() string { return proto.CompactTextString(m) }
func (*SignatureDescriptor_Data_Multi) ProtoMessage()    {}
func (*SignatureDescriptor_Data_Multi) Descriptor() ([]byte, []int) {
	return fileDescriptor_d560eb6bd3d62ef7, []int{1, 0, 1}
}
func (m *SignatureDescriptor_Data_Multi) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureDescriptor_Data_Multi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureDescriptor_Data_Multi.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureDescriptor_Data_Multi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureDescriptor_Data_Multi.Merge(m, src)
}
func (m *SignatureDescriptor_Data_Multi) XXX_Size() int {
	return m.Size()
}
func (m *SignatureDescriptor_Data_Multi) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureDescriptor_Data_Multi.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureDescriptor_Data_Multi proto.InternalMessageInfo

func (m *SignatureDescriptor_Data_Multi) GetBitarray() *types.CompactBitArray {
	if m != nil {
		return m.Bitarray
	}
	return nil
}

func (m *SignatureDescriptor_Data_Multi) GetSignatures() []*SignatureDescriptor_Data {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*SignatureDescriptors)(nil), "cosmos.tx.SignatureDescriptors")
	proto.RegisterType((*SignatureDescriptor)(nil), "cosmos.tx.SignatureDescriptor")
	proto.RegisterType((*SignatureDescriptor_Data)(nil), "cosmos.tx.SignatureDescriptor.Data")
	proto.RegisterType((*SignatureDescriptor_Data_Single)(nil), "cosmos.tx.SignatureDescriptor.Data.Single")
	proto.RegisterType((*SignatureDescriptor_Data_Multi)(nil), "cosmos.tx.SignatureDescriptor.Data.Multi")
}

func init() { proto.RegisterFile("cosmos/tx/sig_desc.proto", fileDescriptor_d560eb6bd3d62ef7) }

var fileDescriptor_d560eb6bd3d62ef7 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xeb, 0xd3, 0x40,
	0x10, 0xc5, 0x93, 0x7e, 0xd3, 0x62, 0xa7, 0xe2, 0x61, 0xf5, 0x10, 0xa2, 0xac, 0xa5, 0x22, 0x54,
	0xc1, 0x04, 0xea, 0xa1, 0xe0, 0x41, 0xe9, 0x8f, 0x43, 0x41, 0x0a, 0xb2, 0x05, 0x05, 0x2f, 0x65,
	0x93, 0x2c, 0x71, 0x69, 0x93, 0x0d, 0xd9, 0x0d, 0x34, 0x77, 0x0f, 0x1e, 0xfd, 0xb3, 0x3c, 0xf6,
	0xe8, 0x51, 0xda, 0xff, 0xc2, 0x93, 0x64, 0x93, 0xd4, 0x56, 0x8a, 0xf4, 0x34, 0xed, 0xe6, 0x7d,
	0xde, 0x3c, 0x66, 0x06, 0xec, 0x40, 0xc8, 0x58, 0x48, 0x4f, 0xed, 0x3c, 0xc9, 0xa3, 0x75, 0xc8,
	0x64, 0xe0, 0xa6, 0x99, 0x50, 0x02, 0x75, 0xab, 0x2f, 0xae, 0xda, 0x39, 0x4e, 0x2d, 0x0a, 0xb2,
	0x22, 0x55, 0xa2, 0x2e, 0x95, 0xcc, 0x79, 0x7a, 0x61, 0x90, 0xf0, 0x24, 0x6a, 0x6a, 0x25, 0x18,
	0x7c, 0x84, 0x47, 0x2b, 0x1e, 0x25, 0x54, 0xe5, 0x19, 0x9b, 0x33, 0x19, 0x64, 0x3c, 0x55, 0x22,
	0x93, 0xe8, 0x2d, 0x80, 0x6c, 0xde, 0xa5, 0x6d, 0xf6, 0xef, 0x86, 0xbd, 0x11, 0x76, 0x4f, 0x4d,
	0xdd, 0x2b, 0x10, 0x39, 0x23, 0x06, 0x5f, 0x2d, 0x78, 0x78, 0x45, 0x83, 0xc6, 0x00, 0x69, 0xee,
	0x6f, 0x79, 0xb0, 0xde, 0xb0, 0xc2, 0x36, 0xfb, 0xe6, 0xb0, 0x37, 0xb2, 0x1b, 0xdf, 0x3a, 0xfa,
	0x07, 0x2d, 0x78, 0xcf, 0x0a, 0xd2, 0x4d, 0x9b, 0x9f, 0x68, 0x0c, 0x56, 0x48, 0x15, 0xb5, 0x5b,
	0x1a, 0x79, 0xf6, 0xff, 0x28, 0xee, 0x9c, 0x2a, 0x4a, 0x34, 0xe0, 0xfc, 0x6e, 0x81, 0x55, 0xfe,
	0x45, 0x73, 0xe8, 0x48, 0x9e, 0x44, 0x5b, 0x56, 0xb7, 0x7d, 0x79, 0x83, 0x87, 0xbb, 0xd2, 0xc4,
	0xc2, 0x20, 0x35, 0x8b, 0x26, 0xd0, 0x8e, 0xf3, 0xad, 0xe2, 0x75, 0x90, 0x17, 0xb7, 0x98, 0x2c,
	0x4b, 0x60, 0x61, 0x90, 0x8a, 0x74, 0x3e, 0x41, 0xa7, 0xb2, 0x45, 0x1e, 0x58, 0xb1, 0x08, 0xab,
	0x40, 0x0f, 0x46, 0x8f, 0xcf, 0xbc, 0x9a, 0x2d, 0x95, 0x9e, 0x4b, 0x11, 0x32, 0xa2, 0x85, 0xe8,
	0x09, 0x74, 0x4f, 0x43, 0xd6, 0x09, 0xee, 0x93, 0xbf, 0x0f, 0xce, 0x37, 0x13, 0xda, 0xba, 0x17,
	0x7a, 0x03, 0xf7, 0x7c, 0xae, 0x68, 0x96, 0xd1, 0x66, 0xc8, 0xf8, 0x9f, 0x21, 0xcf, 0x44, 0x9c,
	0xd2, 0x40, 0x4d, 0xb9, 0x9a, 0x94, 0x2a, 0x72, 0xd2, 0xa3, 0xd9, 0xc5, 0xea, 0x5b, 0x7a, 0xf5,
	0x37, 0xcd, 0xfb, 0x0c, 0x9b, 0xb6, 0xe1, 0x4e, 0xe6, 0xf1, 0xf4, 0xdd, 0x8f, 0x03, 0x36, 0xf7,
	0x07, 0x6c, 0xfe, 0x3a, 0x60, 0xf3, 0xfb, 0x11, 0x1b, 0xfb, 0x23, 0x36, 0x7e, 0x1e, 0xb1, 0xf1,
	0xf9, 0x79, 0xc4, 0xd5, 0x97, 0xdc, 0x77, 0x03, 0x11, 0x7b, 0xcd, 0x01, 0xeb, 0xf2, 0x4a, 0x86,
	0x1b, 0x4f, 0x15, 0x29, 0x2b, 0xaf, 0xd6, 0xef, 0xe8, 0x33, 0x7d, 0xfd, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x10, 0xb1, 0xf1, 0xc8, 0x0a, 0x03, 0x00, 0x00,
}

func (m *SignatureDescriptors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureDescriptors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signatures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSigDesc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SignatureDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSigDesc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.PublicKey != nil {
		{
			size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSigDesc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignatureDescriptor_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureDescriptor_Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptor_Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *SignatureDescriptor_Data_Single_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptor_Data_Single_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Single != nil {
		{
			size, err := m.Single.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSigDesc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SignatureDescriptor_Data_Multi_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptor_Data_Multi_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Multi != nil {
		{
			size, err := m.Multi.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSigDesc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *SignatureDescriptor_Data_Single) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureDescriptor_Data_Single) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptor_Data_Single) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintSigDesc(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.Mode != 0 {
		i = encodeVarintSigDesc(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignatureDescriptor_Data_Multi) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureDescriptor_Data_Multi) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureDescriptor_Data_Multi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signatures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSigDesc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Bitarray != nil {
		{
			size, err := m.Bitarray.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSigDesc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSigDesc(dAtA []byte, offset int, v uint64) int {
	offset -= sovSigDesc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignatureDescriptors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovSigDesc(uint64(l))
		}
	}
	return n
}

func (m *SignatureDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovSigDesc(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovSigDesc(uint64(l))
	}
	return n
}

func (m *SignatureDescriptor_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *SignatureDescriptor_Data_Single_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Single != nil {
		l = m.Single.Size()
		n += 1 + l + sovSigDesc(uint64(l))
	}
	return n
}
func (m *SignatureDescriptor_Data_Multi_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multi != nil {
		l = m.Multi.Size()
		n += 1 + l + sovSigDesc(uint64(l))
	}
	return n
}
func (m *SignatureDescriptor_Data_Single) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mode != 0 {
		n += 1 + sovSigDesc(uint64(m.Mode))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSigDesc(uint64(l))
	}
	return n
}

func (m *SignatureDescriptor_Data_Multi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bitarray != nil {
		l = m.Bitarray.Size()
		n += 1 + l + sovSigDesc(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovSigDesc(uint64(l))
		}
	}
	return n
}

func sovSigDesc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSigDesc(x uint64) (n int) {
	return sovSigDesc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignatureDescriptors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigDesc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignatureDescriptors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureDescriptors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, &SignatureDescriptor{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigDesc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignatureDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigDesc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignatureDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &types.PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &SignatureDescriptor_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigDesc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignatureDescriptor_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigDesc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Single", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignatureDescriptor_Data_Single{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &SignatureDescriptor_Data_Single_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multi", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignatureDescriptor_Data_Multi{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &SignatureDescriptor_Data_Multi_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigDesc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignatureDescriptor_Data_Single) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigDesc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Single: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Single: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= signing.SignMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigDesc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignatureDescriptor_Data_Multi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSigDesc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Multi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Multi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bitarray", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bitarray == nil {
				m.Bitarray = &types.CompactBitArray{}
			}
			if err := m.Bitarray.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSigDesc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSigDesc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, &SignatureDescriptor_Data{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSigDesc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSigDesc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSigDesc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSigDesc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSigDesc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSigDesc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSigDesc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSigDesc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSigDesc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSigDesc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSigDesc = fmt.Errorf("proto: unexpected end of group")
)