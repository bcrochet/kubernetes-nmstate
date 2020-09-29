/*
Copyright 2020 The Kubernetes Authors.

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

package regex

import (
	"fmt"
	"regexp"
)

const (
	releaseRegexStr = `v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-[a-zA-Z0-9]+)*\.*(0|[1-9][0-9]*)?`
	buildRegexStr   = `([0-9]{1,})\+([0-9a-f]{5,40})`
	branchRegexStr  = `master|release-([0-9]{1,})\.([0-9]{1,})(\.([0-9]{1,}))*$`
)

var (
	BranchRegex          = regexp.MustCompile(branchRegexStr)
	ReleaseRegex         = regexp.MustCompile(releaseRegexStr)
	ReleaseAndBuildRegex = regexp.MustCompile(
		fmt.Sprintf("%s.%s", releaseRegexStr, buildRegexStr),
	)
)
