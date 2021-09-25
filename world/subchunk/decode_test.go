package subchunk

import (
	"testing"

	"github.com/danhale-git/mine/mock"
)

func TestNew(t *testing.T) {
	_, err := Decode(mock.SubChunkValue)
	if err != nil {
		t.Errorf("unexpected error returned: %s", err)
	}
}

func TestReadStateIndices(t *testing.T) {
	r := mock.SubChunkReader()
	_, _ = r.Read(make([]byte, 2))

	indices, err := readStateIndices(r)
	if err != nil {
		t.Errorf("unexpected error returned: %s", err)
	}

	palette := mock.StatePaletteIDs()

	expected := mock.BlockStateIndices

	for i, stateIndex := range indices {
		if stateIndex >= len(palette) {
			t.Fatalf("block state index %d is out of range of state palette with length %d", stateIndex, len(palette))
		}

		if stateIndex != expected[i] {
			t.Fatalf("expected palette index '%d' but got '%d'", expected[i], stateIndex)
		}
	}

	if len(indices) != BlockCount {
		t.Errorf("expected %d blocks state indices: got %d", BlockCount, len(indices))
	}
}
