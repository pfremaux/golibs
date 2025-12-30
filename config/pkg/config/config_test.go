package config_test

import (
	"testing"

	"github.com/pfremaux/golibs/config/pkg/config"
)

type TestConfig struct {
	Field1 string `yaml:"field1"`
	Field2 int    `yaml:"field2"`
}

func TestLoadYaml(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		filePath       string
		out            any
		wantErr        bool
		expectedField1 string
		expectedField2 int
	}{
		// TODO: Add test cases.
		{
			name:           "Valid YAML file",
			filePath:       "../../test/test.yaml",
			out:            &TestConfig{},
			wantErr:        false,
			expectedField1: "testString",
			expectedField2: 42,
		},
		{
			name:     "Non existing YAML file",
			filePath: "../../test/nonExistingTest.yaml",
			out:      &TestConfig{},
			wantErr:  true,
		},
		{
			name:     "Wrong format YAML file",
			filePath: "../../test/invalid.yaml",
			out:      &TestConfig{},
			wantErr:  true,
		},
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
			result := tt.out.(*TestConfig)
			if result.Field1 != tt.expectedField1 {
				t.Errorf("LoadYaml() Field1 = %v, want %v", result.Field1, tt.expectedField1)
			}
			if result.Field2 != tt.expectedField2 {
				t.Errorf("LoadYaml() Field2 = %v, want %v", result.Field2, tt.expectedField2)
			}
		})
	}
}
