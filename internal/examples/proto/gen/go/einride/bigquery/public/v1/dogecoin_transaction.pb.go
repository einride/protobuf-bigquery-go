// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: einride/bigquery/public/v1/dogecoin_transaction.proto

package publicv1

import (
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Protobuf schema for the BigQuery public table:
//
//	bigquery-public-data.crypto_dogecoin.transactions
type DogecoinTransaction struct {
	state               protoimpl.MessageState        `protogen:"open.v1"`
	Hash                string                        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                                                            // STRING REQUIRED
	Size                int64                         `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`                                                           // INTEGER NULLABLE
	VirtualSize         int64                         `protobuf:"varint,3,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`                          // INTEGER NULLABLE
	Version             int64                         `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`                                                     // INTEGER NULLABLE
	LockTime            int64                         `protobuf:"varint,5,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`                                   // INTEGER NULLABLE
	BlockHash           string                        `protobuf:"bytes,6,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`                                 // STRING REQUIRED
	BlockNumber         int64                         `protobuf:"varint,7,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`                          // INTEGER REQUIRED
	BlockTimestamp      *timestamppb.Timestamp        `protobuf:"bytes,8,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`                  // TIMESTAMP REQUIRED
	BlockTimestampMonth *date.Date                    `protobuf:"bytes,9,opt,name=block_timestamp_month,json=blockTimestampMonth,proto3" json:"block_timestamp_month,omitempty"` // DATE REQUIRED
	InputCount          int64                         `protobuf:"varint,10,opt,name=input_count,json=inputCount,proto3" json:"input_count,omitempty"`                            // INTEGER NULLABLE
	OutputCount         int64                         `protobuf:"varint,11,opt,name=output_count,json=outputCount,proto3" json:"output_count,omitempty"`                         // INTEGER NULLABLE
	IsCoinbase          bool                          `protobuf:"varint,14,opt,name=is_coinbase,json=isCoinbase,proto3" json:"is_coinbase,omitempty"`                            // BOOLEAN NULLABLE
	Inputs              []*DogecoinTransaction_Input  `protobuf:"bytes,16,rep,name=inputs,proto3" json:"inputs,omitempty"`                                                       // RECORD REPEATED
	Outputs             []*DogecoinTransaction_Output `protobuf:"bytes,17,rep,name=outputs,proto3" json:"outputs,omitempty"`                                                     // RECORD REPEATED
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *DogecoinTransaction) Reset() {
	*x = DogecoinTransaction{}
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DogecoinTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogecoinTransaction) ProtoMessage() {}

func (x *DogecoinTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogecoinTransaction.ProtoReflect.Descriptor instead.
func (*DogecoinTransaction) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *DogecoinTransaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *DogecoinTransaction) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DogecoinTransaction) GetVirtualSize() int64 {
	if x != nil {
		return x.VirtualSize
	}
	return 0
}

func (x *DogecoinTransaction) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DogecoinTransaction) GetLockTime() int64 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *DogecoinTransaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *DogecoinTransaction) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *DogecoinTransaction) GetBlockTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.BlockTimestamp
	}
	return nil
}

func (x *DogecoinTransaction) GetBlockTimestampMonth() *date.Date {
	if x != nil {
		return x.BlockTimestampMonth
	}
	return nil
}

func (x *DogecoinTransaction) GetInputCount() int64 {
	if x != nil {
		return x.InputCount
	}
	return 0
}

func (x *DogecoinTransaction) GetOutputCount() int64 {
	if x != nil {
		return x.OutputCount
	}
	return 0
}

func (x *DogecoinTransaction) GetIsCoinbase() bool {
	if x != nil {
		return x.IsCoinbase
	}
	return false
}

func (x *DogecoinTransaction) GetInputs() []*DogecoinTransaction_Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *DogecoinTransaction) GetOutputs() []*DogecoinTransaction_Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

type DogecoinTransaction_Input struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Index                int64                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                                                            // INTEGER REQUIRED
	SpentTransactionHash string                 `protobuf:"bytes,2,opt,name=spent_transaction_hash,json=spentTransactionHash,proto3" json:"spent_transaction_hash,omitempty"` // STRING NULLABLE
	SpentOutputIndex     int64                  `protobuf:"varint,3,opt,name=spent_output_index,json=spentOutputIndex,proto3" json:"spent_output_index,omitempty"`            // INTEGER NULLABLE
	ScriptAsm            string                 `protobuf:"bytes,4,opt,name=script_asm,json=scriptAsm,proto3" json:"script_asm,omitempty"`                                    // STRING NULLABLE
	ScriptHex            string                 `protobuf:"bytes,5,opt,name=script_hex,json=scriptHex,proto3" json:"script_hex,omitempty"`                                    // STRING NULLABLE
	Sequence             int64                  `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence,omitempty"`                                                      // INTEGER NULLABLE
	RequiredSignatures   int64                  `protobuf:"varint,7,opt,name=required_signatures,json=requiredSignatures,proto3" json:"required_signatures,omitempty"`        // INTEGER NULLABLE
	Type                 string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`                                                               // STRING NULLABLE
	Addresses            []string               `protobuf:"bytes,9,rep,name=addresses,proto3" json:"addresses,omitempty"`                                                     // STRING REPEATED
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *DogecoinTransaction_Input) Reset() {
	*x = DogecoinTransaction_Input{}
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DogecoinTransaction_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogecoinTransaction_Input) ProtoMessage() {}

func (x *DogecoinTransaction_Input) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogecoinTransaction_Input.ProtoReflect.Descriptor instead.
func (*DogecoinTransaction_Input) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DogecoinTransaction_Input) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetSpentTransactionHash() string {
	if x != nil {
		return x.SpentTransactionHash
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetSpentOutputIndex() int64 {
	if x != nil {
		return x.SpentOutputIndex
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetScriptAsm() string {
	if x != nil {
		return x.ScriptAsm
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetScriptHex() string {
	if x != nil {
		return x.ScriptHex
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetRequiredSignatures() int64 {
	if x != nil {
		return x.RequiredSignatures
	}
	return 0
}

func (x *DogecoinTransaction_Input) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DogecoinTransaction_Input) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type DogecoinTransaction_Output struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Index              int64                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`                                                     // INTEGER REQUIRED
	ScriptAsm          string                 `protobuf:"bytes,2,opt,name=script_asm,json=scriptAsm,proto3" json:"script_asm,omitempty"`                             // STRING NULLABLE
	ScriptHex          string                 `protobuf:"bytes,3,opt,name=script_hex,json=scriptHex,proto3" json:"script_hex,omitempty"`                             // STRING NULLABLE
	RequiredSignatures int64                  `protobuf:"varint,4,opt,name=required_signatures,json=requiredSignatures,proto3" json:"required_signatures,omitempty"` // INTEGER NULLABLE
	Type               string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`                                                        // STRING NULLABLE
	Addresses          []string               `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`                                              // STRING REPEATED
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DogecoinTransaction_Output) Reset() {
	*x = DogecoinTransaction_Output{}
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DogecoinTransaction_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogecoinTransaction_Output) ProtoMessage() {}

func (x *DogecoinTransaction_Output) ProtoReflect() protoreflect.Message {
	mi := &file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogecoinTransaction_Output.ProtoReflect.Descriptor instead.
func (*DogecoinTransaction_Output) Descriptor() ([]byte, []int) {
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DogecoinTransaction_Output) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DogecoinTransaction_Output) GetScriptAsm() string {
	if x != nil {
		return x.ScriptAsm
	}
	return ""
}

func (x *DogecoinTransaction_Output) GetScriptHex() string {
	if x != nil {
		return x.ScriptHex
	}
	return ""
}

func (x *DogecoinTransaction_Output) GetRequiredSignatures() int64 {
	if x != nil {
		return x.RequiredSignatures
	}
	return 0
}

func (x *DogecoinTransaction_Output) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DogecoinTransaction_Output) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_einride_bigquery_public_v1_dogecoin_transaction_proto protoreflect.FileDescriptor

const file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc = "" +
	"\n" +
	"5einride/bigquery/public/v1/dogecoin_transaction.proto\x12\x1aeinride.bigquery.public.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x16google/type/date.proto\"\xee\b\n" +
	"\x13DogecoinTransaction\x12\x12\n" +
	"\x04hash\x18\x01 \x01(\tR\x04hash\x12\x12\n" +
	"\x04size\x18\x02 \x01(\x03R\x04size\x12!\n" +
	"\fvirtual_size\x18\x03 \x01(\x03R\vvirtualSize\x12\x18\n" +
	"\aversion\x18\x04 \x01(\x03R\aversion\x12\x1b\n" +
	"\tlock_time\x18\x05 \x01(\x03R\blockTime\x12\x1d\n" +
	"\n" +
	"block_hash\x18\x06 \x01(\tR\tblockHash\x12!\n" +
	"\fblock_number\x18\a \x01(\x03R\vblockNumber\x12C\n" +
	"\x0fblock_timestamp\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\x0eblockTimestamp\x12E\n" +
	"\x15block_timestamp_month\x18\t \x01(\v2\x11.google.type.DateR\x13blockTimestampMonth\x12\x1f\n" +
	"\vinput_count\x18\n" +
	" \x01(\x03R\n" +
	"inputCount\x12!\n" +
	"\foutput_count\x18\v \x01(\x03R\voutputCount\x12\x1f\n" +
	"\vis_coinbase\x18\x0e \x01(\bR\n" +
	"isCoinbase\x12M\n" +
	"\x06inputs\x18\x10 \x03(\v25.einride.bigquery.public.v1.DogecoinTransaction.InputR\x06inputs\x12P\n" +
	"\aoutputs\x18\x11 \x03(\v26.einride.bigquery.public.v1.DogecoinTransaction.OutputR\aoutputs\x1a\xbe\x02\n" +
	"\x05Input\x12\x14\n" +
	"\x05index\x18\x01 \x01(\x03R\x05index\x124\n" +
	"\x16spent_transaction_hash\x18\x02 \x01(\tR\x14spentTransactionHash\x12,\n" +
	"\x12spent_output_index\x18\x03 \x01(\x03R\x10spentOutputIndex\x12\x1d\n" +
	"\n" +
	"script_asm\x18\x04 \x01(\tR\tscriptAsm\x12\x1d\n" +
	"\n" +
	"script_hex\x18\x05 \x01(\tR\tscriptHex\x12\x1a\n" +
	"\bsequence\x18\x06 \x01(\x03R\bsequence\x12/\n" +
	"\x13required_signatures\x18\a \x01(\x03R\x12requiredSignatures\x12\x12\n" +
	"\x04type\x18\b \x01(\tR\x04type\x12\x1c\n" +
	"\taddresses\x18\t \x03(\tR\taddresses\x1a\xbf\x01\n" +
	"\x06Output\x12\x14\n" +
	"\x05index\x18\x01 \x01(\x03R\x05index\x12\x1d\n" +
	"\n" +
	"script_asm\x18\x02 \x01(\tR\tscriptAsm\x12\x1d\n" +
	"\n" +
	"script_hex\x18\x03 \x01(\tR\tscriptHex\x12/\n" +
	"\x13required_signatures\x18\x04 \x01(\x03R\x12requiredSignatures\x12\x12\n" +
	"\x04type\x18\x05 \x01(\tR\x04type\x12\x1c\n" +
	"\taddresses\x18\x06 \x03(\tR\taddressesB\xaa\x02\n" +
	"\x1ecom.einride.bigquery.public.v1B\x18DogecoinTransactionProtoP\x01Zago.einride.tech/protobuf-bigquery/internal/examples/proto/gen/einride/bigquery/public/v1;publicv1\xa2\x02\x03EBP\xaa\x02\x1aEinride.Bigquery.Public.V1\xca\x02\x1bEinride\\Bigquery\\Public_\\V1\xe2\x02'Einride\\Bigquery\\Public_\\V1\\GPBMetadata\xea\x02\x1dEinride::Bigquery::Public::V1b\x06proto3"

var (
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescOnce sync.Once
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescData []byte
)

func file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescGZIP() []byte {
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescOnce.Do(func() {
		file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc), len(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc)))
	})
	return file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDescData
}

var file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_bigquery_public_v1_dogecoin_transaction_proto_goTypes = []any{
	(*DogecoinTransaction)(nil),        // 0: einride.bigquery.public.v1.DogecoinTransaction
	(*DogecoinTransaction_Input)(nil),  // 1: einride.bigquery.public.v1.DogecoinTransaction.Input
	(*DogecoinTransaction_Output)(nil), // 2: einride.bigquery.public.v1.DogecoinTransaction.Output
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(*date.Date)(nil),                  // 4: google.type.Date
}
var file_einride_bigquery_public_v1_dogecoin_transaction_proto_depIdxs = []int32{
	3, // 0: einride.bigquery.public.v1.DogecoinTransaction.block_timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: einride.bigquery.public.v1.DogecoinTransaction.block_timestamp_month:type_name -> google.type.Date
	1, // 2: einride.bigquery.public.v1.DogecoinTransaction.inputs:type_name -> einride.bigquery.public.v1.DogecoinTransaction.Input
	2, // 3: einride.bigquery.public.v1.DogecoinTransaction.outputs:type_name -> einride.bigquery.public.v1.DogecoinTransaction.Output
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_einride_bigquery_public_v1_dogecoin_transaction_proto_init() }
func file_einride_bigquery_public_v1_dogecoin_transaction_proto_init() {
	if File_einride_bigquery_public_v1_dogecoin_transaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc), len(file_einride_bigquery_public_v1_dogecoin_transaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_bigquery_public_v1_dogecoin_transaction_proto_goTypes,
		DependencyIndexes: file_einride_bigquery_public_v1_dogecoin_transaction_proto_depIdxs,
		MessageInfos:      file_einride_bigquery_public_v1_dogecoin_transaction_proto_msgTypes,
	}.Build()
	File_einride_bigquery_public_v1_dogecoin_transaction_proto = out.File
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_goTypes = nil
	file_einride_bigquery_public_v1_dogecoin_transaction_proto_depIdxs = nil
}
