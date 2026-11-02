// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package db_subnet_group

import (
	"reflect"
	"slices"
	"testing"

	svcapitypes "github.com/aws-controllers-k8s/documentdb-controller/apis/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// diffTags mirrors the tag diff performed by syncTags. It returns the tags to
// add or update (keyed by key, with their values) and the keys to remove.
func diffTags(
	latest []*svcapitypes.Tag,
	desired []*svcapitypes.Tag,
) (map[string]string, []string) {
	from, _ := convertToOrderedACKTags(latest)
	to, toOrder := convertToOrderedACKTags(desired)
	added, _, removed := ackcompare.GetTagsDifference(from, to)
	toAdd := fromACKTags(added, toOrder)

	toAddMap := make(map[string]string, len(toAdd))
	for _, t := range toAdd {
		toAddMap[*t.Key] = *t.Value
	}
	toDeleteKeys := make([]string, 0, len(removed))
	for k := range removed {
		toDeleteKeys = append(toDeleteKeys, k)
	}
	slices.Sort(toDeleteKeys)
	return toAddMap, toDeleteKeys
}

func tag(key, value string) *svcapitypes.Tag {
	return &svcapitypes.Tag{Key: &key, Value: &value}
}

func TestSyncTagsDiff(t *testing.T) {
	tests := []struct {
		name       string
		latest     []*svcapitypes.Tag
		desired    []*svcapitypes.Tag
		wantAdd    map[string]string
		wantDelete []string
	}{
		{
			name:       "no-op when tags are identical",
			latest:     []*svcapitypes.Tag{tag("env", "prod"), tag("team", "db")},
			desired:    []*svcapitypes.Tag{tag("env", "prod"), tag("team", "db")},
			wantAdd:    map[string]string{},
			wantDelete: []string{},
		},
		{
			name:       "add a new tag",
			latest:     []*svcapitypes.Tag{tag("env", "prod")},
			desired:    []*svcapitypes.Tag{tag("env", "prod"), tag("team", "db")},
			wantAdd:    map[string]string{"team": "db"},
			wantDelete: []string{},
		},
		{
			name:       "remove a tag missing from desired",
			latest:     []*svcapitypes.Tag{tag("env", "prod"), tag("team", "db")},
			desired:    []*svcapitypes.Tag{tag("env", "prod")},
			wantAdd:    map[string]string{},
			wantDelete: []string{"team"},
		},
		{
			name:       "value change is add only, not remove",
			latest:     []*svcapitypes.Tag{tag("env", "prod")},
			desired:    []*svcapitypes.Tag{tag("env", "staging")},
			wantAdd:    map[string]string{"env": "staging"},
			wantDelete: []string{},
		},
		{
			name:       "nil keys are skipped without panicking",
			latest:     []*svcapitypes.Tag{{Key: nil, Value: nil}, tag("env", "prod")},
			desired:    []*svcapitypes.Tag{{Key: nil, Value: nil}, tag("env", "prod")},
			wantAdd:    map[string]string{},
			wantDelete: []string{},
		},
		{
			name:       "simultaneous add and remove",
			latest:     []*svcapitypes.Tag{tag("env", "prod"), tag("team", "db")},
			desired:    []*svcapitypes.Tag{tag("env", "prod"), tag("region", "us-east-1")},
			wantAdd:    map[string]string{"region": "us-east-1"},
			wantDelete: []string{"team"},
		},
		{
			name:       "desired nil removes all latest tags",
			latest:     []*svcapitypes.Tag{tag("env", "prod")},
			desired:    nil,
			wantAdd:    map[string]string{},
			wantDelete: []string{"env"},
		},
		{
			name:       "latest nil adds all desired tags",
			latest:     nil,
			desired:    []*svcapitypes.Tag{tag("env", "prod")},
			wantAdd:    map[string]string{"env": "prod"},
			wantDelete: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdd, gotDelete := diffTags(tt.latest, tt.desired)
			if !reflect.DeepEqual(gotAdd, tt.wantAdd) {
				t.Errorf("toAdd = %v, want %v", gotAdd, tt.wantAdd)
			}
			if !reflect.DeepEqual(gotDelete, tt.wantDelete) {
				t.Errorf("toDelete = %v, want %v", gotDelete, tt.wantDelete)
			}
		})
	}
}
