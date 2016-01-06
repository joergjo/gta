package gta

import (
	"reflect"
	"testing"
)

func TestGraphTraversal(t *testing.T) {
	tests := []struct {
		graph   *Graph
		start   string
		got     map[string]bool
		want    map[string]bool
		comment string
	}{
		{
			comment: "A depends on B depends on C, C is dirty, so we expect all of them to be marked",
			graph: &Graph{
				graph: map[string]map[string]bool{
					"C": map[string]bool{
						"B": true,
					},
					"B": map[string]bool{
						"A": true,
					},
				},
			},
			start: "C",
			got:   map[string]bool{},
			want: map[string]bool{
				"A": true,
				"B": true,
				"C": true,
			},
		},
		{
			comment: "A depends on B depends on C, B is dirty, so we expect just A and B, and NOT C to be marked",
			graph: &Graph{
				graph: map[string]map[string]bool{
					"C": map[string]bool{
						"B": true,
					},
					"B": map[string]bool{
						"A": true,
					},
				},
			},
			start: "B",
			got:   map[string]bool{},
			want: map[string]bool{
				"A": true,
				"B": true,
			},
		},
	}

	for _, tt := range tests {
		t.Log(tt.comment)
		tt.graph.Traverse(tt.start, &tt.got)
		if !reflect.DeepEqual(tt.want, tt.got) {
			t.Error("expected want and got to be equal")
			t.Errorf("want: %v", tt.want)
			t.Errorf(" got: %v", tt.got)
		}
	}
}
