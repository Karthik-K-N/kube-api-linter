/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package conditions_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conditions"
)

func TestDefaultConfiguration(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := conditions.Initializer().Init(&conditions.ConditionsConfig{})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "a")
}

func TestNotFieldFirst(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := conditions.Initializer().Init(&conditions.ConditionsConfig{
		IsFirstField: conditions.ConditionsFirstFieldIgnore,
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "b")
}

func TestIgnoreProtobuf(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := conditions.Initializer().Init(&conditions.ConditionsConfig{
		UseProtobuf: conditions.ConditionsUseProtobufIgnore,
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "c")
}

func TestForbidProtobuf(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := conditions.Initializer().Init(&conditions.ConditionsConfig{
		UseProtobuf: conditions.ConditionsUseProtobufForbid,
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "d")
}

func TestIgnorePatchStrategy(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := conditions.Initializer().Init(&conditions.ConditionsConfig{
		UsePatchStrategy: conditions.ConditionsUsePatchStrategyIgnore,
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "e")
}

func TestForbidPatchStrategy(t *testing.T) {
	testdata := analysistest.TestData()

	a, err := conditions.Initializer().Init(&conditions.ConditionsConfig{
		UsePatchStrategy: conditions.ConditionsUsePatchStrategyForbid,
	})
	if err != nil {
		t.Fatal(err)
	}

	analysistest.RunWithSuggestedFixes(t, testdata, a, "f")
}
