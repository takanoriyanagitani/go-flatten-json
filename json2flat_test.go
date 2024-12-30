package json2flat_test

import (
	"testing"

	"encoding/json"

	fl "github.com/takanoriyanagitani/go-flatten-json"
)

func TestJsonToFlat(t *testing.T) {
	t.Parallel()

	t.Run("MapToFlatAny", func(t *testing.T) {
		t.Parallel()

		t.Run("MapToFlatMapDefault", func(t *testing.T) {

			t.Parallel()

			t.Run("empty", func(t *testing.T) {
				t.Parallel()

				var m2f fl.MapToFlatMap = fl.MapToFlatMapDefault()
				var m2a fl.MapToFlatAny = m2f.ToMapToFlatAny()

				var flat map[string]any = m2a(map[string]any{})
				if 0 != len(flat) {
					t.Fatalf("unexpected map: %v\n", flat)
				}
			})

			t.Run("flat-bool", func(t *testing.T) {
				t.Parallel()

				var m2f fl.MapToFlatMap = fl.MapToFlatMapDefault()
				var m2a fl.MapToFlatAny = m2f.ToMapToFlatAny()

				buf := map[string]any{}
				e := json.Unmarshal(
					[]byte(`{
						"active": true
					}`),
					&buf,
				)
				if nil != e {
					t.Fatalf("unexpected error: %v\n", e)
				}

				var flat map[string]any = m2a(buf)
				if 1 != len(flat) {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var active any = flat["active"]
				if true != active.(bool) {
					t.Fatalf("unexpected value: %v\n", active)
				}
			})

			t.Run("flat-double", func(t *testing.T) {
				t.Parallel()

				var m2f fl.MapToFlatMap = fl.MapToFlatMapDefault()
				var m2a fl.MapToFlatAny = m2f.ToMapToFlatAny()

				buf := map[string]any{}
				e := json.Unmarshal(
					[]byte(`{
						"height": 3.776
					}`),
					&buf,
				)
				if nil != e {
					t.Fatalf("unexpected error: %v\n", e)
				}

				var flat map[string]any = m2a(buf)
				if 1 != len(flat) {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var height any = flat["height"]
				if 3.776 != height.(float64) {
					t.Fatalf("unexpected value: %v\n", height)
				}
			})

			t.Run("flat-string", func(t *testing.T) {
				t.Parallel()

				var m2f fl.MapToFlatMap = fl.MapToFlatMapDefault()
				var m2a fl.MapToFlatAny = m2f.ToMapToFlatAny()

				buf := map[string]any{}
				e := json.Unmarshal(
					[]byte(`{
						"name": "fuji"
					}`),
					&buf,
				)
				if nil != e {
					t.Fatalf("unexpected error: %v\n", e)
				}

				var flat map[string]any = m2a(buf)
				if 1 != len(flat) {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var name any = flat["name"]
				if "fuji" != name.(string) {
					t.Fatalf("unexpected value: %v\n", name)
				}
			})

			t.Run("flat-nil", func(t *testing.T) {
				t.Parallel()

				var m2f fl.MapToFlatMap = fl.MapToFlatMapDefault()
				var m2a fl.MapToFlatAny = m2f.ToMapToFlatAny()

				buf := map[string]any{}
				e := json.Unmarshal(
					[]byte(`{
						"updated": null
					}`),
					&buf,
				)
				if nil != e {
					t.Fatalf("unexpected error: %v\n", e)
				}

				var flat map[string]any = m2a(buf)
				if 1 != len(flat) {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var updated any = flat["updated"]
				if nil != updated {
					t.Fatalf("unexpected value: %v\n", updated)
				}
			})

			t.Run("user", func(t *testing.T) {
				t.Parallel()

				var m2f fl.MapToFlatMap = fl.MapToFlatMapDefault()
				var m2a fl.MapToFlatAny = m2f.ToMapToFlatAny()

				buf := map[string]any{}
				e := json.Unmarshal(
					[]byte(`{
						"user": {
							"name": "jd",
							"mails": [
							  {"mail":"dummy.i@dummy.local", "primary":false},
							  {"mail":"dummy.ii@dummy.local", "primary":true}
							]
						}
					}`),
					&buf,
				)
				if nil != e {
					t.Fatalf("unexpected error: %v\n", e)
				}

				var flat map[string]any = m2a(buf)
				if 5 != len(flat) {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var name string = flat["user_name"].(string)
				if "jd" != name {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var mail0 string = flat["user_mails_0_mail"].(string)
				if "dummy.i@dummy.local" != mail0 {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var mail1 string = flat["user_mails_1_mail"].(string)
				if "dummy.ii@dummy.local" != mail1 {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var prim0 bool = flat["user_mails_0_primary"].(bool)
				if false != prim0 {
					t.Fatalf("unexpected map: %v\n", flat)
				}

				var prim1 bool = flat["user_mails_1_primary"].(bool)
				if true != prim1 {
					t.Fatalf("unexpected map: %v\n", flat)
				}

			})

		})
	})
}
