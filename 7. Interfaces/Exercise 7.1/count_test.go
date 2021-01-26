package count

import (
	"fmt"
	"testing"
)

func TestWordCounterAndLineCounterEmpty(t *testing.T) {
	wantedWordCounter := WordCounter(0)
	wantedLineCounter := LineCounter(0)

	bs := []byte("")

	var wordCounter WordCounter
	var lineCounter LineCounter

	fmt.Fprintf(&wordCounter, "%s", bs)
	fmt.Fprintf(&lineCounter, "%s", bs)

	if wordCounter != wantedWordCounter {
		t.Errorf("WordCounter counts %s as %d words, want %d", bs, wordCounter, wantedWordCounter)
	}
	if lineCounter != wantedLineCounter {
		t.Errorf("LineCounter counts %s as %d lines, want %d", bs, lineCounter, wantedLineCounter)
	}
}

func TestWordCounterAndLineCounterNonEmpty(t *testing.T) {
	wantedWordCounter := WordCounter(10)
	wantedLineCounter := LineCounter(8)

	bs := []byte(`there are many words

and lines

in

this
text block`)

	var wordCounter WordCounter
	var lineCounter LineCounter

	fmt.Fprintf(&wordCounter, "%s", bs)
	fmt.Fprintf(&lineCounter, "%s", bs)

	if wordCounter != wantedWordCounter {
		t.Errorf("WordCounter counts %s as %d words, want %d", bs, wordCounter, wantedWordCounter)
	}
	if lineCounter != wantedLineCounter {
		t.Errorf("LineCounter counts %s as %d lines, want %d", bs, lineCounter, wantedLineCounter)
	}
}
