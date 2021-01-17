package count

import "io"

type countingWriter struct {
	wrappedWriter io.Writer
	writtenBytes  int64
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.wrappedWriter.Write(p)
	cw.writtenBytes += int64(n)
	return
}

// CountingWriter returns a wrapped w and a pointer to an int64, such that the pointee always contains the number of bytes written to the writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := countingWriter{
		wrappedWriter: w,
	}
	return &cw, &cw.writtenBytes
}
