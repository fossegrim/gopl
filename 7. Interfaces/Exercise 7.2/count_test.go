package count

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var wantedWrittenBytes = int64(3)

	var buf bytes.Buffer
	w, writtenBytesP := CountingWriter(&buf)
	fmt.Fprint(w, "123")

	if *writtenBytesP != wantedWrittenBytes {
		t.Errorf("CountingWriter incorrectly counts %d bytes written, want %d", *writtenBytesP, wantedWrittenBytes)
	}

	wantedWrittenBytes = 9
	fmt.Fprint(w, "456789")
	if *writtenBytesP != wantedWrittenBytes {
		t.Errorf("CountingWriter incorrectly counts %d bytes written, want %d", *writtenBytesP, wantedWrittenBytes)
	}
}
