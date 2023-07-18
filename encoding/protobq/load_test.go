package protobq

import (
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	examplev1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/example/v1"
	publicv1 "go.einride.tech/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/public/v1"
	expr "google.golang.org/genproto/googleapis/api/expr/v1beta1"
	"google.golang.org/genproto/googleapis/example/library/v1"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gotest.tools/v3/assert"
)

func TestUnmarshalOptions_Load(t *testing.T) {
	for _, tt := range []struct {
		name          string
		row           []bigquery.Value
		schema        bigquery.Schema
		opt           UnmarshalOptions
		expected      proto.Message
		errorContains string
	}{
		{
			name: "library.Book",
			row: []bigquery.Value{
				"name",
				"author",
				"title",
				true,
			},
			schema: bigquery.Schema{
				{Name: "name", Type: bigquery.StringFieldType},
				{Name: "author", Type: bigquery.StringFieldType},
				{Name: "title", Type: bigquery.StringFieldType},
				{Name: "read", Type: bigquery.BooleanFieldType},
			},
			expected: &library.Book{
				Name:   "name",
				Author: "author",
				Title:  "title",
				Read:   true,
			},
		},
		{
			name: "publicv1.WhosOnFirstGeoJson",
			row: []bigquery.Value{
				"geoid",
				int64(1192980459),
				`{"lat": 76.16119999999999, "lon": 14.03075}`,
				"Point",
				"76.1612,14.03075,76.1612,14.03075",
				"geonames",
				int64(1536368481),
				time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
			},
			schema: bigquery.Schema{
				{Name: "geoid", Type: bigquery.StringFieldType},
				{Name: "id", Type: bigquery.IntegerFieldType},
				{Name: "body", Type: bigquery.JSONFieldType},
				{Name: "geometry_type", Type: bigquery.StringFieldType},
				{Name: "bounding_box", Type: bigquery.StringFieldType},
				{Name: "geom", Type: bigquery.StringFieldType},
				{Name: "last_modified", Type: bigquery.IntegerFieldType},
				{Name: "last_modified_timestamp", Type: bigquery.TimestampFieldType},
			},
			expected: &publicv1.WhosOnFirstGeoJson{
				Geoid: "geoid",
				Id:    1192980459,
				Body: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"lat": {
							Kind: &structpb.Value_NumberValue{
								NumberValue: 76.16119999999999,
							},
						},
						"lon": {
							Kind: &structpb.Value_NumberValue{
								NumberValue: 14.03075,
							},
						},
					},
				},
				GeometryType: "Point",
				BoundingBox:  "76.1612,14.03075,76.1612,14.03075",
				Geom:         "geonames",
				LastModified: 1536368481,
				LastModifiedTimestamp: &timestamppb.Timestamp{
					Seconds: 1673308800,
				},
			},
		},
		{
			name: "library.UpdateBookRequest",
			row: []bigquery.Value{
				[]bigquery.Value{
					"name",
					"author",
					"title",
					true,
				},
			},
			schema: bigquery.Schema{
				{
					Name: "book",
					Type: bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "name", Type: bigquery.StringFieldType},
						{Name: "author", Type: bigquery.StringFieldType},
						{Name: "title", Type: bigquery.StringFieldType},
						{Name: "read", Type: bigquery.BooleanFieldType},
					},
				},
			},
			expected: &library.UpdateBookRequest{
				Book: &library.Book{
					Name:   "name",
					Author: "author",
					Title:  "title",
					Read:   true,
				},
			},
		},

		{
			name: "expr.Value (bool)",
			row: []bigquery.Value{
				true,
				nil,
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
			},
			expected: &expr.Value{
				Kind: &expr.Value_BoolValue{
					BoolValue: true,
				},
			},
		},

		{
			name: "expr.Value (double)",
			row: []bigquery.Value{
				nil,
				float64(42),
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
			},
			expected: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
		},

		{
			name: "error on unknown fields",
			row: []bigquery.Value{
				nil,
				float64(42),
				"bar",
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "foo", Type: bigquery.StringFieldType},
			},
			expected:      &expr.Value{},
			errorContains: "unknown field: foo",
		},

		{
			name: "discard unknown fields",
			row: []bigquery.Value{
				nil,
				float64(42),
				"bar",
			},
			schema: bigquery.Schema{
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "foo", Type: bigquery.StringFieldType},
			},
			opt: UnmarshalOptions{DiscardUnknown: true},
			expected: &expr.Value{
				Kind: &expr.Value_DoubleValue{
					DoubleValue: 42,
				},
			},
		},

		{
			name: "enum values",
			row: []bigquery.Value{
				"ENUM_VALUE1",
			},
			schema: bigquery.Schema{
				{Name: "enum_value", Type: bigquery.StringFieldType},
			},
			expected: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
		},

		{
			name: "enum numbers",
			row: []bigquery.Value{
				int64(1),
			},
			schema: bigquery.Schema{
				{Name: "enum_value", Type: bigquery.IntegerFieldType},
			},
			opt: UnmarshalOptions{
				Schema: SchemaOptions{
					UseEnumNumbers: true,
				},
			},
			expected: &examplev1.ExampleEnum{
				EnumValue: examplev1.ExampleEnum_ENUM_VALUE1,
			},
		},

		{
			name: "wrappers",
			row: []bigquery.Value{
				float64(1),
				float64(2),
				"foo",
				[]byte("bar"),
				int64(3),
				int64(4),
				uint64(5),
				uint64(6),
				true,
			},
			schema: bigquery.Schema{
				{Name: "float_value", Type: bigquery.FloatFieldType},
				{Name: "double_value", Type: bigquery.FloatFieldType},
				{Name: "string_value", Type: bigquery.StringFieldType},
				{Name: "bytes_value", Type: bigquery.BytesFieldType},
				{Name: "int32_value", Type: bigquery.IntegerFieldType},
				{Name: "int64_value", Type: bigquery.IntegerFieldType},
				{Name: "uint32_value", Type: bigquery.IntegerFieldType},
				{Name: "uint64_value", Type: bigquery.IntegerFieldType},
				{Name: "bool_value", Type: bigquery.BooleanFieldType},
			},
			expected: &examplev1.ExampleWrappers{
				FloatValue:  wrapperspb.Float(1),
				DoubleValue: wrapperspb.Double(2),
				StringValue: wrapperspb.String("foo"),
				BytesValue:  wrapperspb.Bytes([]byte("bar")),
				Int32Value:  wrapperspb.Int32(3),
				Int64Value:  wrapperspb.Int64(4),
				Uint32Value: wrapperspb.UInt32(5),
				Uint64Value: wrapperspb.UInt64(6),
				BoolValue:   wrapperspb.Bool(true),
			},
		},

		{
			name: "primitive lists",
			row: []bigquery.Value{
				[]bigquery.Value{int64(1), int64(2)},
				[]bigquery.Value{"a", "b"},
				[]bigquery.Value{"ENUM_VALUE1", "ENUM_VALUE2"},
			},
			schema: bigquery.Schema{
				{Name: "int64_list", Type: bigquery.IntegerFieldType, Repeated: true},
				{Name: "string_list", Type: bigquery.StringFieldType, Repeated: true},
				{Name: "enum_list", Type: bigquery.StringFieldType, Repeated: true},
			},
			expected: &examplev1.ExampleList{
				Int64List:  []int64{1, 2},
				StringList: []string{"a", "b"},
				EnumList: []examplev1.ExampleList_Enum{
					examplev1.ExampleList_ENUM_VALUE1,
					examplev1.ExampleList_ENUM_VALUE2,
				},
			},
		},

		{
			name: "well-known-type lists",
			row: []bigquery.Value{
				[]bigquery.Value{float32(1), float32(2)},
			},
			schema: bigquery.Schema{
				{Name: "float_value_list", Type: bigquery.FloatFieldType, Repeated: true},
			},
			expected: &examplev1.ExampleList{
				FloatValueList: []*wrapperspb.FloatValue{
					wrapperspb.Float(1), wrapperspb.Float(2),
				},
			},
		},

		{
			name: "lists",
			row: []bigquery.Value{
				[]bigquery.Value{int64(1), int64(2)},
				[]bigquery.Value{"a", "b"},
				[]bigquery.Value{"ENUM_VALUE1", "ENUM_VALUE2"},
				[]bigquery.Value{
					[]bigquery.Value{
						[]bigquery.Value{"a", "b"},
					},
					[]bigquery.Value{
						[]bigquery.Value{"c", "d"},
					},
				},
			},
			schema: bigquery.Schema{
				{Name: "int64_list", Type: bigquery.IntegerFieldType, Repeated: true},
				{Name: "string_list", Type: bigquery.StringFieldType, Repeated: true},
				{Name: "enum_list", Type: bigquery.StringFieldType, Repeated: true},
				{
					Name:     "nested_list",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "string_list", Type: bigquery.StringFieldType, Repeated: true},
					},
				},
			},
			expected: &examplev1.ExampleList{
				Int64List:  []int64{1, 2},
				StringList: []string{"a", "b"},
				EnumList: []examplev1.ExampleList_Enum{
					examplev1.ExampleList_ENUM_VALUE1,
					examplev1.ExampleList_ENUM_VALUE2,
				},
				NestedList: []*examplev1.ExampleList_Nested{
					{StringList: []string{"a", "b"}},
					{StringList: []string{"c", "d"}},
				},
			},
		},

		{
			name: "primitive maps",
			row: []bigquery.Value{
				[]bigquery.Value{
					map[string]bigquery.Value{"key": "a", "value": "b"},
				},
				[]bigquery.Value{
					map[string]bigquery.Value{"key": "a", "value": "ENUM_VALUE1"},
				},
				[]bigquery.Value{
					map[string]bigquery.Value{"key": int64(1), "value": "a"},
				},
				[]bigquery.Value{
					map[string]bigquery.Value{"key": int64(2), "value": "a"},
				},
				[]bigquery.Value{
					map[string]bigquery.Value{"key": uint64(3), "value": "a"},
				},
				[]bigquery.Value{
					map[string]bigquery.Value{"key": true, "value": "a"},
				},
			},
			schema: bigquery.Schema{
				{
					Name:     "string_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "string_to_enum",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "int32_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.IntegerFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "int64_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.IntegerFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "uint32_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.IntegerFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},

				{
					Name:     "bool_to_string",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.BooleanFieldType},
						{Name: "value", Type: bigquery.StringFieldType},
					},
				},
			},
			expected: &examplev1.ExampleMap{
				StringToString: map[string]string{"a": "b"},
				StringToEnum:   map[string]examplev1.ExampleMap_Enum{"a": examplev1.ExampleMap_ENUM_VALUE1},
				Int32ToString:  map[int32]string{1: "a"},
				Int64ToString:  map[int64]string{2: "a"},
				Uint32ToString: map[uint32]string{3: "a"},
				BoolToString:   map[bool]string{true: "a"},
			},
		},

		{
			name: "well-known-type maps",
			row: []bigquery.Value{
				[]bigquery.Value{
					map[string]bigquery.Value{"key": "a", "value": float64(1)},
				},
			},
			schema: bigquery.Schema{
				{
					Name:     "string_to_float_value",
					Repeated: true,
					Type:     bigquery.RecordFieldType,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{Name: "value", Type: bigquery.FloatFieldType},
					},
				},
			},
			expected: &examplev1.ExampleMap{
				StringToFloatValue: map[string]*wrapperspb.FloatValue{
					"a": wrapperspb.Float(1),
				},
			},
		},

		{
			name: "nested maps",
			row: []bigquery.Value{
				[]bigquery.Value{
					map[string]bigquery.Value{
						"key": "a",
						"value": []bigquery.Value{
							[]bigquery.Value{
								map[string]bigquery.Value{"key": "a", "value": "b"},
							},
						},
					},
				},
			},
			schema: bigquery.Schema{
				{
					Name:     "string_to_nested",
					Type:     bigquery.RecordFieldType,
					Repeated: true,
					Schema: bigquery.Schema{
						{Name: "key", Type: bigquery.StringFieldType},
						{
							Name: "value",
							Type: bigquery.RecordFieldType,
							Schema: bigquery.Schema{
								{
									Name:     "string_to_string",
									Type:     bigquery.RecordFieldType,
									Repeated: true,
									Schema: bigquery.Schema{
										{Name: "key", Type: bigquery.StringFieldType},
										{Name: "value", Type: bigquery.StringFieldType},
									},
								},
							},
						},
					},
				},
			},
			expected: &examplev1.ExampleMap{
				StringToNested: map[string]*examplev1.ExampleMap_Nested{
					"a": {
						StringToString: map[string]string{
							"a": "b",
						},
					},
				},
			},
		},

		{
			name: "datetime (without offset)",
			row: []bigquery.Value{
				civil.DateTime{
					Date: civil.Date{
						Year:  2021,
						Month: time.February,
						Day:   1,
					},
					Time: civil.Time{
						Hour:       8,
						Minute:     30,
						Second:     1,
						Nanosecond: 2,
					},
				},
			},
			schema: bigquery.Schema{
				{Name: "date_time", Type: bigquery.DateTimeFieldType},
			},
			opt: UnmarshalOptions{
				Schema: SchemaOptions{
					UseDateTimeWithoutOffset: true,
				},
			},
			expected: &examplev1.ExampleDateTime{
				DateTime: &datetime.DateTime{
					Year:    2021,
					Month:   int32(time.February),
					Day:     1,
					Hours:   8,
					Minutes: 30,
					Seconds: 1,
					Nanos:   2,
				},
			},
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := proto.Clone(tt.expected)
			proto.Reset(actual)
			if err := tt.opt.Load(tt.row, tt.schema, actual); tt.errorContains != "" {
				assert.ErrorContains(t, err, tt.errorContains)
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tt.expected, actual, protocmp.Transform())
			}
		})
	}
}
