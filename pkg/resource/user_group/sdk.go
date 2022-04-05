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

package user_group

import (
	"context"
	"errors"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/elasticache-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ElastiCache{}
	_ = &svcapitypes.UserGroup{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeUserGroupsOutput
	resp, err = rm.sdkapi.DescribeUserGroupsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeUserGroups", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UserGroupNotFound" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.UserGroups {
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.Engine != nil {
			ko.Spec.Engine = elem.Engine
		} else {
			ko.Spec.Engine = nil
		}
		if elem.PendingChanges != nil {
			f2 := &svcapitypes.UserGroupPendingChanges{}
			if elem.PendingChanges.UserIdsToAdd != nil {
				f2f0 := []*string{}
				for _, f2f0iter := range elem.PendingChanges.UserIdsToAdd {
					var f2f0elem string
					f2f0elem = *f2f0iter
					f2f0 = append(f2f0, &f2f0elem)
				}
				f2.UserIDsToAdd = f2f0
			}
			if elem.PendingChanges.UserIdsToRemove != nil {
				f2f1 := []*string{}
				for _, f2f1iter := range elem.PendingChanges.UserIdsToRemove {
					var f2f1elem string
					f2f1elem = *f2f1iter
					f2f1 = append(f2f1, &f2f1elem)
				}
				f2.UserIDsToRemove = f2f1
			}
			ko.Status.PendingChanges = f2
		} else {
			ko.Status.PendingChanges = nil
		}
		if elem.ReplicationGroups != nil {
			f3 := []*string{}
			for _, f3iter := range elem.ReplicationGroups {
				var f3elem string
				f3elem = *f3iter
				f3 = append(f3, &f3elem)
			}
			ko.Status.ReplicationGroups = f3
		} else {
			ko.Status.ReplicationGroups = nil
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		} else {
			ko.Status.Status = nil
		}
		if elem.UserGroupId != nil {
			ko.Spec.UserGroupID = elem.UserGroupId
		} else {
			ko.Spec.UserGroupID = nil
		}
		if elem.UserIds != nil {
			f6 := []*string{}
			for _, f6iter := range elem.UserIds {
				var f6elem string
				f6elem = *f6iter
				f6 = append(f6, &f6elem)
			}
			ko.Spec.UserIDs = f6
		} else {
			ko.Spec.UserIDs = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomDescribeUserGroupsSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return r.ko.Spec.UserGroupID == nil

}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeUserGroupsInput, error) {
	res := &svcsdk.DescribeUserGroupsInput{}

	if r.ko.Spec.UserGroupID != nil {
		res.SetUserGroupId(*r.ko.Spec.UserGroupID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateUserGroupOutput
	_ = resp
	resp, err = rm.sdkapi.CreateUserGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateUserGroup", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Engine != nil {
		ko.Spec.Engine = resp.Engine
	} else {
		ko.Spec.Engine = nil
	}
	if resp.PendingChanges != nil {
		f2 := &svcapitypes.UserGroupPendingChanges{}
		if resp.PendingChanges.UserIdsToAdd != nil {
			f2f0 := []*string{}
			for _, f2f0iter := range resp.PendingChanges.UserIdsToAdd {
				var f2f0elem string
				f2f0elem = *f2f0iter
				f2f0 = append(f2f0, &f2f0elem)
			}
			f2.UserIDsToAdd = f2f0
		}
		if resp.PendingChanges.UserIdsToRemove != nil {
			f2f1 := []*string{}
			for _, f2f1iter := range resp.PendingChanges.UserIdsToRemove {
				var f2f1elem string
				f2f1elem = *f2f1iter
				f2f1 = append(f2f1, &f2f1elem)
			}
			f2.UserIDsToRemove = f2f1
		}
		ko.Status.PendingChanges = f2
	} else {
		ko.Status.PendingChanges = nil
	}
	if resp.ReplicationGroups != nil {
		f3 := []*string{}
		for _, f3iter := range resp.ReplicationGroups {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		ko.Status.ReplicationGroups = f3
	} else {
		ko.Status.ReplicationGroups = nil
	}
	if resp.Status != nil {
		ko.Status.Status = resp.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.UserGroupId != nil {
		ko.Spec.UserGroupID = resp.UserGroupId
	} else {
		ko.Spec.UserGroupID = nil
	}
	if resp.UserIds != nil {
		f6 := []*string{}
		for _, f6iter := range resp.UserIds {
			var f6elem string
			f6elem = *f6iter
			f6 = append(f6, &f6elem)
		}
		ko.Spec.UserIDs = f6
	} else {
		ko.Spec.UserIDs = nil
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomCreateUserGroupSetOutput(ctx, desired, resp, ko)
	if err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateUserGroupInput, error) {
	res := &svcsdk.CreateUserGroupInput{}

	if r.ko.Spec.Engine != nil {
		res.SetEngine(*r.ko.Spec.Engine)
	}
	if r.ko.Spec.Tags != nil {
		f1 := []*svcsdk.Tag{}
		for _, f1iter := range r.ko.Spec.Tags {
			f1elem := &svcsdk.Tag{}
			if f1iter.Key != nil {
				f1elem.SetKey(*f1iter.Key)
			}
			if f1iter.Value != nil {
				f1elem.SetValue(*f1iter.Value)
			}
			f1 = append(f1, f1elem)
		}
		res.SetTags(f1)
	}
	if r.ko.Spec.UserGroupID != nil {
		res.SetUserGroupId(*r.ko.Spec.UserGroupID)
	}
	if r.ko.Spec.UserIDs != nil {
		f3 := []*string{}
		for _, f3iter := range r.ko.Spec.UserIDs {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		res.SetUserIds(f3)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateUserGroup(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteUserGroupOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteUserGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteUserGroup", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteUserGroupInput, error) {
	res := &svcsdk.DeleteUserGroupInput{}

	if r.ko.Spec.UserGroupID != nil {
		res.SetUserGroupId(*r.ko.Spec.UserGroupID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.UserGroup,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "DuplicateUserNameFault",
		"UserGroupAlreadyExistsFault",
		"InvalidParameterCombination",
		"InvalidParameterValueException",
		"DefaultUserRequired",
		"UserGroupQuotaExceededFault",
		"TagQuotaPerResourceExceeded":
		return true
	default:
		return false
	}
}
