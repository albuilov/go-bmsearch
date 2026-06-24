package horspool

import (
	"testing"

	"github.com/albuilov/go-bmsearch/internal/searchtesting"
)

func TestSearch(t *testing.T) {
	searchtesting.Run(t, New())
}
