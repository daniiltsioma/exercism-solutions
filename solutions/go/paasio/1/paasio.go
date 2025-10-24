package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	mu sync.Mutex

	r io.Reader
	bytesRead int64
	readOps int
}

type writeCounter struct {
	mu sync.Mutex

	w io.Writer
	bytesWritten int64
	writeOps int
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		w: writer,
		bytesWritten: 0,
		writeOps: 0,
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		r: reader,
		bytesRead: 0,
		readOps: 0,
	}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		ReadCounter: NewReadCounter(readwriter),
		WriteCounter: NewWriteCounter(readwriter),
	}		
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	n, err := rc.r.Read(p); 

	rc.bytesRead += int64(n)
	rc.readOps++

	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	return rc.bytesRead, rc.readOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	n, err := wc.w.Write(p);

	wc.bytesWritten += int64(n)
	wc.writeOps++

	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	return wc.bytesWritten, wc.writeOps
}
