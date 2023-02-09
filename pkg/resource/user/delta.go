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

// Code generated by ack-generate. DO NOT EDIT.

package user

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AccessString, b.ko.Spec.AccessString) {
		delta.Add("Spec.AccessString", a.ko.Spec.AccessString, b.ko.Spec.AccessString)
	} else if a.ko.Spec.AccessString != nil && b.ko.Spec.AccessString != nil {
		if *a.ko.Spec.AccessString != *b.ko.Spec.AccessString {
			delta.Add("Spec.AccessString", a.ko.Spec.AccessString, b.ko.Spec.AccessString)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NoPasswordRequired, b.ko.Spec.NoPasswordRequired) {
		delta.Add("Spec.NoPasswordRequired", a.ko.Spec.NoPasswordRequired, b.ko.Spec.NoPasswordRequired)
	} else if a.ko.Spec.NoPasswordRequired != nil && b.ko.Spec.NoPasswordRequired != nil {
		if *a.ko.Spec.NoPasswordRequired != *b.ko.Spec.NoPasswordRequired {
			delta.Add("Spec.NoPasswordRequired", a.ko.Spec.NoPasswordRequired, b.ko.Spec.NoPasswordRequired)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.UserID, b.ko.Spec.UserID) {
		delta.Add("Spec.UserID", a.ko.Spec.UserID, b.ko.Spec.UserID)
	} else if a.ko.Spec.UserID != nil && b.ko.Spec.UserID != nil {
		if *a.ko.Spec.UserID != *b.ko.Spec.UserID {
			delta.Add("Spec.UserID", a.ko.Spec.UserID, b.ko.Spec.UserID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.UserName, b.ko.Spec.UserName) {
		delta.Add("Spec.UserName", a.ko.Spec.UserName, b.ko.Spec.UserName)
	} else if a.ko.Spec.UserName != nil && b.ko.Spec.UserName != nil {
		if *a.ko.Spec.UserName != *b.ko.Spec.UserName {
			delta.Add("Spec.UserName", a.ko.Spec.UserName, b.ko.Spec.UserName)
		}
	}

	filterDelta(delta, a, b)
	return delta
}
