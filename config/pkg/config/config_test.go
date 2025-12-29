package config_test

import (
	"testing"

	"github.com/pfremaux/golibs/config/pkg/config"
)

func TestLoadYaml(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filePath string
		out      any
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := config.LoadYaml(tt.filePath, tt.out)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadYaml() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadYaml() succeeded unexpectedly")
			}
		})
	}
}

func TestLoadFlagsConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		params []config.Parameter
		want   map[string]*string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.LoadFlagsConfig(tt.params)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("LoadFlagsConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
