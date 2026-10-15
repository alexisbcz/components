package marker

import (
	"strings"
	"testing"

	"github.com/alexisbcz/components/lib"
	h "github.com/alexisbcz/libhtml"
)

func TestMarker(t *testing.T) {
	var b strings.Builder
	if err := Marker(VariantSeparator, lib.AsChild, h.A(h.Href("#x"), Content("t"))).Render(&b); err != nil {
		t.Fatal(err)
	}
	out := b.String()
	if !strings.HasPrefix(out, "<a") || !strings.Contains(out, `data-variant="separator"`) || !strings.Contains(out, "before:h-px") {
		t.Errorf("marker: %s", out)
	}
}
