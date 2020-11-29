# Protobuf + BigQuery + Go

Go utilities for saving protocol buffers to BigQuery.

## Examples

### Schema inference

```go
schema := protobq.Schema(&library.Book{})
```

### Value marshaling

```go
value, err := protobq.Marshal(&library.Book{Name: "test"})
```
