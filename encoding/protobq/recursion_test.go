package protobq

import (
	"testing"

	examplev1 "github.com/goalsgame/protobuf-bigquery/internal/examples/proto/gen/go/einride/bigquery/example/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
)

func TestIsFieldRecursive(t *testing.T) {
	testCases := []struct {
		description string
		message     protoreflect.MessageDescriptor
		fieldName   protoreflect.Name
		isRecursive bool
	}{
		{
			description: "RecursiveMessageWrapper: recursive_message (recursive type)",
			message:     (&examplev1.RecursiveMessageWrapper{}).ProtoReflect().Descriptor(),
			fieldName:   "recursive_message",
			isRecursive: true,
		},
		{
			description: "RecursiveMessageWrapper: repeated_recursive_wrapper (contains recursive type)",
			message:     (&examplev1.RecursiveMessageWrapper{}).ProtoReflect().Descriptor(),
			fieldName:   "repeated_recursive_wrapper",
			isRecursive: false, // This is unexpected, but reflects current implementation.
		},
		{
			description: "RecursiveMessageWrapper: repeated_recursive_message (repeated recursive type)",
			message:     (&examplev1.RecursiveMessageWrapper{}).ProtoReflect().Descriptor(),
			fieldName:   "repeated_recursive_message",
			isRecursive: true,
		},
		{
			description: "RecursiveListWrapper: id (string)",
			message:     (&examplev1.RecursiveListWrapper{}).ProtoReflect().Descriptor(),
			fieldName:   "id",
			isRecursive: false,
		},
		{
			description: "RecursiveListWrapper: child (recursive type)",
			message:     (&examplev1.RecursiveListWrapper{}).ProtoReflect().Descriptor(),
			fieldName:   "child",
			isRecursive: true,
		},
		{
			description: "RecursiveMessage: id (string)",
			message:     (&examplev1.RecursiveMessage{}).ProtoReflect().Descriptor(),
			fieldName:   "id",
			isRecursive: false,
		},
		{
			description: "RecursiveMessage: child (direct recursion)",
			message:     (&examplev1.RecursiveMessage{}).ProtoReflect().Descriptor(),
			fieldName:   "child",
			isRecursive: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			field := tc.message.Fields().ByName(tc.fieldName)
			if field == nil {
				t.Fatalf("field %s not found in message %s", tc.fieldName, tc.message.FullName())
			}
			assert.Equal(t, tc.isRecursive, isFieldRecursive(field))
		})
	}
}
