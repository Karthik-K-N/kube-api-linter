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

package nodurations

import (
	"fmt"
	"go/ast"
	"go/types"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
)

const name = "nodurations"

// Analyzer is the analyzer for the nodurations package.
// It checks that no struct fields are `time.Duration`.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Durations should be avoided and should instead use fooSeconds, avoids having to implement go duration parsing.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspector.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers) {
		checkFieldType(pass, field, jsonTagInfo)
	})

	return nil, nil //nolint:nilnil
}

func checkFieldType(pass *analysis.Pass, field *ast.Field, tagInfo extractjsontags.FieldTagInfo) {

	fieldName := utils.FieldName(field)
	if fieldName == "" {
		return
	}

	fmt.Println("filedName", fieldName)
	fmt.Println("field type", types.ExprString(field.Type))
	fmt.Println(types.ExprString(field.Type) == "time.Duration")

}
