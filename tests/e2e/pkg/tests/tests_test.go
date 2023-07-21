// Copyright © 2023 Cisco Systems, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"sort"
	"testing"
)

func Test_Classifier_minimal(t *testing.T) {
	type fields struct {
		k8sClusterPool K8sClusterPool
		testCases      []TestCase
	}
	tests := []struct {
		name   string
		fields fields
		want   []TestType
	}{
		{
			name: "simpleCase",
			fields: fields{
				k8sClusterPool: K8sClusterPool{
					k8sClusters: []K8sCluster{
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local1",
								version:   "1.24",
								provider:  "provider1",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local2",
								version:   "1.25",
								provider:  "provider1",
							},
						},
					},
				},
				testCases: []TestCase{
					{
						TestName: "testCase1",
					},
					{
						TestName: "testCase2",
					},
				},
			},
			want: []TestType{
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider1",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Classifier{
				k8sClusterPool: tt.fields.k8sClusterPool,
				testCases:      tt.fields.testCases,
			}

			got := tr.Minimal()
			if len(tt.want) != len(got) {
				t.Fatalf("not equal want: %v got: %v", tt.want, got)
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].Less(got[j])
			})
			sort.SliceStable(tt.want, func(i, j int) bool {
				return tt.want[i].Less(tt.want[j])
			})

			for i := range got {
				if !got[i].Equal(tt.want[i]) {
					t.Errorf("not equal want: %v got: %v", tt.want[i], got[i])
				}
			}
		})
	}
}

func Test_Classifier_providerComplete(t *testing.T) {
	type fields struct {
		k8sClusterPool K8sClusterPool
		testCases      []TestCase
	}
	tests := []struct {
		name   string
		fields fields
		want   []TestType
	}{
		{
			name: "simpleCase",
			fields: fields{
				k8sClusterPool: K8sClusterPool{
					k8sClusters: []K8sCluster{
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local1",
								version:   "1.24",
								provider:  "provider1",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local2",
								version:   "1.25",
								provider:  "provider2",
							},
						},
					},
				},
				testCases: []TestCase{
					{
						TestName: "testCase1",
					},
					{
						TestName: "testCase2",
					},
				},
			},
			want: []TestType{
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider2",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider2",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Classifier{
				k8sClusterPool: tt.fields.k8sClusterPool,
				testCases:      tt.fields.testCases,
			}

			got := tr.ProviderComplete()
			if len(tt.want) != len(got) {
				t.Fatalf("not equal want: %v got: %v", tt.want, got)
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].Less(got[j])
			})
			sort.SliceStable(tt.want, func(i, j int) bool {
				return tt.want[i].Less(tt.want[j])
			})

			for i := range got {
				if !got[i].Equal(tt.want[i]) {
					t.Errorf("not equal want: %v got: %v", tt.want[i], got[i])
				}
			}
		})
	}
}

func Test_Classifier_versionComplete(t *testing.T) {
	type fields struct {
		k8sClusterPool K8sClusterPool
		testCases      []TestCase
	}
	tests := []struct {
		name   string
		fields fields
		want   []TestType
	}{
		{
			name: "simpleCase",
			fields: fields{
				k8sClusterPool: K8sClusterPool{
					k8sClusters: []K8sCluster{
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local1",
								version:   "1.24",
								provider:  "provider1",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local1",
								version:   "1.25",
								provider:  "provider1",
							},
						},
					},
				},
				testCases: []TestCase{
					{
						TestName: "testCase1",
					},
					{
						TestName: "testCase2",
					},
				},
			},
			want: []TestType{
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Classifier{
				k8sClusterPool: tt.fields.k8sClusterPool,
				testCases:      tt.fields.testCases,
			}

			got := tr.VersionComplete()
			if len(tt.want) != len(got) {
				t.Fatalf("not equal want: %v got: %v", tt.want, got)
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].Less(got[j])
			})
			sort.SliceStable(tt.want, func(i, j int) bool {
				return tt.want[i].Less(tt.want[j])
			})

			for i := range got {
				if !got[i].Equal(tt.want[i]) {
					t.Errorf("not equal want: %v got: %v", tt.want[i], got[i])
				}
			}
		})
	}
}

func Test_Classifier_complete(t *testing.T) {
	type fields struct {
		k8sClusterPool K8sClusterPool
		testCases      []TestCase
	}
	tests := []struct {
		name   string
		fields fields
		want   []TestType
	}{
		{
			name: "simpleCase",
			fields: fields{
				k8sClusterPool: K8sClusterPool{
					k8sClusters: []K8sCluster{
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local1",
								version:   "1.24",
								provider:  "provider1",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local2",
								version:   "1.25",
								provider:  "provider2",
							},
						},
					},
				},
				testCases: []TestCase{
					{
						TestName: "testCase1",
					},
					{
						TestName: "testCase2",
					},
				},
			},
			want: []TestType{
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider2",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider2",
						},
					},
				},
			},
		},
		{
			name: "complexCase",
			fields: fields{
				k8sClusterPool: K8sClusterPool{
					k8sClusters: []K8sCluster{
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local1",
								version:   "1.24",
								provider:  "provider1",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local2",
								version:   "1.25",
								provider:  "provider2",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local3",
								version:   "1.25",
								provider:  "provider3",
							},
						},
						{
							reusable: true,
							clusterInfo: K8sClusterInfo{
								clusterID: "local4",
								version:   "1.25",
								provider:  "provider3",
							},
						},
					},
				},
				testCases: []TestCase{
					{
						TestName: "testCase1",
					},
					{
						TestName: "testCase2",
					},
				},
			},
			want: []TestType{
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local1",
							version:   "1.24",
							provider:  "provider1",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider2",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local2",
							version:   "1.25",
							provider:  "provider2",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase1",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local3",
							version:   "1.25",
							provider:  "provider3",
						},
					},
				},
				{
					testCase: TestCase{
						TestName: "testCase2",
					},
					k8sCluster: K8sCluster{
						reusable: true,
						clusterInfo: K8sClusterInfo{
							clusterID: "local3",
							version:   "1.25",
							provider:  "provider3",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Classifier{
				k8sClusterPool: tt.fields.k8sClusterPool,
				testCases:      tt.fields.testCases,
			}

			got := tr.Complete()
			if len(tt.want) != len(got) {
				t.Fatalf("not equal want: %v got: %v", tt.want, got)
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].Less(got[j])
			})
			sort.SliceStable(tt.want, func(i, j int) bool {
				return tt.want[i].Less(tt.want[j])
			})

			for i := range got {
				if !got[i].Equal(tt.want[i]) {
					t.Errorf("not equal want: %v got: %v", tt.want[i].TestID(), got[i].TestID())
				}
			}
		})
	}
}
