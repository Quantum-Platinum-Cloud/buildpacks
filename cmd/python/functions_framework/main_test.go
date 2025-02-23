// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"

	buildpacktest "github.com/GoogleCloudPlatform/buildpacks/internal/buildpacktest"
)

func TestContainsFF(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "ff_present",
			str:  "functions-framework==19.9.0\nflask\n",
			want: true,
		},
		{
			name: "ff_present_with_comment",
			str:  "functions-framework #my-comment\nflask\n",
			want: true,
		},
		{
			name: "ff_present_second_line",
			str:  "flask\nfunctions-framework==19.9.0",
			want: true,
		},
		{
			name: "no_ff_present",
			str:  "functions-framework-example==0.1.0\nflask\n",
			want: false,
		},
		{
			name: "ff_egg_present",
			str:  "git+git://github.com/functions-framework@master#egg=functions-framework\nflask\n",
			want: true,
		},
		{
			name: "ff_egg_not_present",
			str:  "git+git://github.com/functions-framework-example@master#egg=functions-framework-example\nflask\n",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := containsFF(tc.str)
			if got != tc.want {
				t.Errorf("containsFF() got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestDetect(t *testing.T) {
	testCases := []struct {
		name  string
		files map[string]string
		env   []string
		stack string
		want  int
	}{
		{
			name: "with target",
			env:  []string{"GOOGLE_FUNCTION_TARGET=helloWorld"},
			want: 0,
		},
		{
			name: "without target",
			want: 100,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buildpacktest.TestDetectWithStack(t, detectFn, tc.name, tc.files, tc.env, tc.stack, tc.want)
		})
	}
}
