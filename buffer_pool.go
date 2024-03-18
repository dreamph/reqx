package reqx

import (
	"bytes"
	"sync"
)

var (
	multipartBufferPool = &sync.Pool{
		New: func() any {
			return &bytes.Buffer{}
		},
	}
)

func getMultipartBufferPool() *bytes.Buffer {
	return multipartBufferPool.Get().(*bytes.Buffer)
}

func putMultipartBufferPool(buf *bytes.Buffer) {
	buf.Reset()
	multipartBufferPool.Put(buf)
}
