package optional

import (
	"encoding/json"
	"testing"
)

type sample struct {
	Name        Optional[string]  `json:"name"`
	Description Optional[string]  `json:"description"`
	Price       Optional[float64] `json:"price"`
}

func TestOptional_OmittedNullAndValue(t *testing.T) {
	var omitted sample
	if err := json.Unmarshal([]byte(`{}`), &omitted); err != nil {
		t.Fatal(err)
	}
	if omitted.Description.IsSet() {
		t.Fatal("omitted description should not be set")
	}

	var cleared sample
	if err := json.Unmarshal([]byte(`{"description":null}`), &cleared); err != nil {
		t.Fatal(err)
	}
	if !cleared.Description.IsSet() || !cleared.Description.IsNull() {
		t.Fatal("null description should be set and null")
	}

	var valued sample
	if err := json.Unmarshal([]byte(`{"name":"Tea","price":10}`), &valued); err != nil {
		t.Fatal(err)
	}
	if !valued.Name.IsSet() || valued.Name.IsNull() || valued.Name.Value() != "Tea" {
		t.Fatalf("got %+v", valued.Name)
	}
	if !valued.Price.IsSet() || valued.Price.Value() != 10 {
		t.Fatalf("got %+v", valued.Price)
	}
}
