package count

import (
	"bufio"
	"bytes"
)

type splitFuncCounter struct {
	Count int
	split bufio.SplitFunc
}

func newSplitFuncCounter(split bufio.SplitFunc) splitFuncCounter {
	return splitFuncCounter{
		split: split,
	}
}

func (c *splitFuncCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(c.split)

	// Since we are scanning from a []byte we do not need to worry about io error handling
	for s.Scan() {
		c.Count++
	}

	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	d := newSplitFuncCounter(bufio.ScanWords)
	len, _ := d.Write(p)
	*c = WordCounter(d.Count)
	return len, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	d := newSplitFuncCounter(bufio.ScanLines)
	len, _ := d.Write(p)
	*c = LineCounter(d.Count)
	return len, nil
}
