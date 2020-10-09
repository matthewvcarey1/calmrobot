package mapland

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		size uint
	}
	tests := []struct {
		name string
		args args
		want *MapLand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapLand_GetMinMax(t *testing.T) {
	tests := []struct {
		name  string
		m     *MapLand
		want  int
		want1 int
		want2 int
		want3 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := tt.m.GetMinMax()
			if got != tt.want {
				t.Errorf("MapLand.GetMinMax() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MapLand.GetMinMax() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("MapLand.GetMinMax() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("MapLand.GetMinMax() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
