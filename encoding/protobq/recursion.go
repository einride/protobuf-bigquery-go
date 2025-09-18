package protobq

import (
	"github.com/goalsgame/protobuf-bigquery/internal/wkt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// recursionDetector is an unexported struct used to track visited message types
// during the recursion check to avoid infinite loops.
type recursionDetector struct {
	visited map[protoreflect.MessageDescriptor]bool
}

// isFieldRecursive checks if a specific field introduces recursion.
// A field is considered recursive if its type either:
// 1. Can traverse back to the message containing the field.
// 2. Is an inherently recursive message type.
func isFieldRecursive(field protoreflect.FieldDescriptor) bool {
	// Recursion check is only meaningful for fields that are messages.
	if field.Kind() != protoreflect.MessageKind {
		return false
	}
	if wkt.IsWellKnownType(string(field.Message().FullName())) {
		return false
	}

	// Get the descriptor for the message type that contains this field.
	parentDescriptor, ok := field.Parent().(protoreflect.MessageDescriptor)
	if !ok {
		// This can happen if the parent is a FileDescriptor (for extensions).
		// We'll consider top-level extension fields as non-recursive in this context.
		return false
	}

	// Determine the message type descriptor of the field's value.
	var fieldMessageDescriptor protoreflect.MessageDescriptor
	switch {
	case field.IsList():
		fieldMessageDescriptor = field.Message()
	case field.IsMap():
		// Keys in maps can't be messages, so we only check the value's type.
		if field.MapValue().Kind() != protoreflect.MessageKind {
			return false
		}
		fieldMessageDescriptor = field.MapValue().Message()
	default: // It's a singular message field.
		fieldMessageDescriptor = field.Message()
	}

	if fieldMessageDescriptor == nil {
		return false // Should not happen for a field of MessageKind.
	}

	// Check 1: Does the field's type loop back to the parent message?
	detector := &recursionDetector{
		visited: make(map[protoreflect.MessageDescriptor]bool),
	}
	if detector.detectPathToTarget(fieldMessageDescriptor, parentDescriptor) {
		return true
	}

	// Check 2: Is the field's type an inherently recursive message type?
	return isMessageTypeRecursive(fieldMessageDescriptor)
}

// isMessageTypeRecursive checks if a message type contains a field that refers
// to itself, either directly or indirectly.
func isMessageTypeRecursive(msgDesc protoreflect.MessageDescriptor) bool {
	// A message type is recursive if any of its fields' types can traverse
	// back to the message type itself.
	fields := msgDesc.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if field.Kind() != protoreflect.MessageKind {
			continue
		}

		var fieldMsgDesc protoreflect.MessageDescriptor
		switch {
		case field.IsList():
			fieldMsgDesc = field.Message()
		case field.IsMap():
			if field.MapValue().Kind() == protoreflect.MessageKind {
				fieldMsgDesc = field.MapValue().Message()
			}
		default: // Singular message.
			fieldMsgDesc = field.Message()
		}

		if fieldMsgDesc != nil {
			detector := &recursionDetector{visited: make(map[protoreflect.MessageDescriptor]bool)}
			if detector.detectPathToTarget(fieldMsgDesc, msgDesc) {
				return true
			}
		}
	}
	return false
}

// detectPathToTarget is a helper that performs a depth-first search to check
// if a traversal from `startDesc` can reach `targetDesc`.
func (d *recursionDetector) detectPathToTarget(startDesc, targetDesc protoreflect.MessageDescriptor) bool {
	// If we've reached the target, we've found the recursive path.
	if startDesc == targetDesc {
		return true
	}

	// If we've already visited this descriptor in the current path, it means
	// we've found a cycle, but not one that leads back to our original target.
	// We can stop exploring this path to avoid infinite loops.
	if d.visited[startDesc] {
		return false
	}

	// Mark the current descriptor as visited for this path.
	d.visited[startDesc] = true

	// Iterate over the fields of the current message descriptor.
	fields := startDesc.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		// We only care about fields that are themselves messages.
		if field.Kind() != protoreflect.MessageKind {
			continue
		}

		// Determine the next message descriptor to traverse into.
		var nextDesc protoreflect.MessageDescriptor
		switch {
		case field.IsList():
			nextDesc = field.Message()
		case field.IsMap():
			if field.MapValue().Kind() == protoreflect.MessageKind {
				nextDesc = field.MapValue().Message()
			}
		default: // Singular message.
			nextDesc = field.Message()
		}

		// If there is a next message type to check, recurse.
		if nextDesc != nil {
			if d.detectPathToTarget(nextDesc, targetDesc) {
				return true
			}
		}
	}

	// Backtrack: Unmark the current descriptor as we are leaving this path.
	delete(d.visited, startDesc)

	return false
}
