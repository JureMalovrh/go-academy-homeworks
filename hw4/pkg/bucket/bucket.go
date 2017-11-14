package bucket

import (
	"errors"
	"fmt"
)

// MockBucket is a holder of the data
type MockBucket struct {
	Data map[int]int
	Name string
}

// NewBucket returns new MockBucket, with map initialized
func NewBucket(name string) MockBucket {
	data := make(map[int]int)
	return MockBucket{data, name}
}

// CreateMockDB function that creates whole db and returns ReadWriter
func CreateMockDB(name string) (MockReadWriter, error) {
	bucket := NewBucket(name)
	reader := MockReader{&bucket}
	writer := MockWriter{&bucket}
	readWriter := MockReadWriter{&reader, &writer}
	return readWriter, nil
}

// MockReader is a reader of the data
type MockReader struct {
	bucket *MockBucket
}

// Get is a function that returns value from the "db" or it return an error
func (r *MockReadWriter) Get(id int) (int, error) {
	val, nonExisting := r.Reader.bucket.Data[id]
	if !nonExisting {
		return -1, errors.New(fmt.Sprintf("Key not found $%d", id))
	}
	return val, nil
}

// MockWriter is a writer of the data
type MockWriter struct {
	bucket *MockBucket
}

// Insert is a function that inserts data into the "db"
func (w *MockReadWriter) Insert(id int, val int) int {
	w.Writer.bucket.Data[id] = val
	return val
}

// MockReadWriter is a holder of Reader and Writer
type MockReadWriter struct {
	Reader *MockReader
	Writer *MockWriter
}
