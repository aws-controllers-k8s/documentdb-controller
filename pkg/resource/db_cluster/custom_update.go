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

package db_cluster

import (
	"context"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/docdb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/documentdb-controller/apis/v1alpha1"
)

// customUpdate is required to fix
// https://github.com/aws-controllers-k8s/community/issues/917. The Input shape
// sent to ModifyDBCluster MUST have fields that are unchanged between desired
// and observed set to `nil`.
func (rm *resourceManager) customUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.customUpdate")
	defer exit(err)

	if clusterDeleting(latest) {
		msg := "DB cluster is currently being deleted"
		ackcondition.SetSynced(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitWhileDeleting
	}
	if clusterCreating(latest) {
		msg := "DB cluster is currently being created"
		ackcondition.SetSynced(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitUntilCanModify(latest)
	}
	if !clusterAvailable(latest) {
		msg := "DB cluster is not available for modification in '" +
			*latest.ko.Status.Status + "' status"
		ackcondition.SetSynced(desired, corev1.ConditionFalse, &msg, nil)
		return desired, requeueWaitUntilCanModify(latest)
	}
	if clusterHasTerminalStatus(latest) {
		msg := "DB cluster is in '" + *latest.ko.Status.Status + "' status"
		ackcondition.SetTerminal(desired, corev1.ConditionTrue, &msg, nil)
		ackcondition.SetSynced(desired, corev1.ConditionTrue, nil, nil)
		return desired, nil
	}

	if delta.DifferentAt("Spec.Tags") {
		if err = rm.syncTags(ctx, desired, latest); err != nil {
			return nil, err
		}
	} else if !delta.DifferentExcept("Spec.Tags") {
		// If the only difference between the desired and latest is in the tags,
		// we should not requeue the resource.
		return desired, nil
	}

	input, err := rm.newCustomUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.ModifyDBClusterOutput
	_ = resp
	resp, err = rm.sdkapi.ModifyDBCluster(ctx, input)

	rm.metrics.RecordAPICall("UPDATE", "ModifyDBCluster", err)
	if err != nil {
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.DBCluster.AssociatedRoles != nil {
		f5 := []*svcapitypes.DBClusterRole{}
		for _, f5iter := range resp.DBCluster.AssociatedRoles {
			f5elem := &svcapitypes.DBClusterRole{}
			if f5iter.RoleArn != nil {
				f5elem.RoleARN = f5iter.RoleArn
			}
			if f5iter.Status != nil {
				f5elem.Status = f5iter.Status
			}
			f5 = append(f5, f5elem)
		}
		ko.Status.AssociatedRoles = f5
	} else {
		ko.Status.AssociatedRoles = nil
	}
	if resp.DBCluster.AvailabilityZones != nil {
		ko.Spec.AvailabilityZones = aws.StringSlice(resp.DBCluster.AvailabilityZones)
	} else {
		ko.Spec.AvailabilityZones = nil
	}
	if resp.DBCluster.BackupRetentionPeriod != nil {
		brpcopy := aws.Int64(int64(*resp.DBCluster.BackupRetentionPeriod))
		ko.Spec.BackupRetentionPeriod = brpcopy
	} else {
		ko.Spec.BackupRetentionPeriod = nil
	}
	if resp.DBCluster.CloneGroupId != nil {
		ko.Status.CloneGroupID = resp.DBCluster.CloneGroupId
	} else {
		ko.Status.CloneGroupID = nil
	}
	if resp.DBCluster.ClusterCreateTime != nil {
		ko.Status.ClusterCreateTime = &metav1.Time{*resp.DBCluster.ClusterCreateTime}
	} else {
		ko.Status.ClusterCreateTime = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.DBCluster.DBClusterArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.DBCluster.DBClusterArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.DBCluster.DBClusterIdentifier != nil {
		ko.Spec.DBClusterIdentifier = resp.DBCluster.DBClusterIdentifier
	} else {
		ko.Spec.DBClusterIdentifier = nil
	}
	if resp.DBCluster.DBClusterMembers != nil {
		f19 := []*svcapitypes.DBClusterMember{}
		for _, f19iter := range resp.DBCluster.DBClusterMembers {
			f19elem := &svcapitypes.DBClusterMember{}
			if f19iter.DBClusterParameterGroupStatus != nil {
				f19elem.DBClusterParameterGroupStatus = f19iter.DBClusterParameterGroupStatus
			}
			if f19iter.DBInstanceIdentifier != nil {
				f19elem.DBInstanceIdentifier = f19iter.DBInstanceIdentifier
			}
			if f19iter.IsClusterWriter != nil {
				f19elem.IsClusterWriter = f19iter.IsClusterWriter
			}
			if f19iter.PromotionTier != nil {
				ptcopy := aws.Int64(int64(*f19iter.PromotionTier))
				f19elem.PromotionTier = ptcopy
			}
			f19 = append(f19, f19elem)
		}
		ko.Status.DBClusterMembers = f19
	} else {
		ko.Status.DBClusterMembers = nil
	}
	if resp.DBCluster.DBClusterParameterGroup != nil {
		ko.Status.DBClusterParameterGroup = resp.DBCluster.DBClusterParameterGroup
	} else {
		ko.Status.DBClusterParameterGroup = nil
	}
	if resp.DBCluster.DBSubnetGroup != nil {
		ko.Status.DBSubnetGroup = resp.DBCluster.DBSubnetGroup
	} else {
		ko.Status.DBSubnetGroup = nil
	}
	if resp.DBCluster.DbClusterResourceId != nil {
		ko.Status.DBClusterResourceID = resp.DBCluster.DbClusterResourceId
	} else {
		ko.Status.DBClusterResourceID = nil
	}
	if resp.DBCluster.DeletionProtection != nil {
		ko.Spec.DeletionProtection = resp.DBCluster.DeletionProtection
	} else {
		ko.Spec.DeletionProtection = nil
	}
	if resp.DBCluster.EarliestRestorableTime != nil {
		ko.Status.EarliestRestorableTime = &metav1.Time{*resp.DBCluster.EarliestRestorableTime}
	} else {
		ko.Status.EarliestRestorableTime = nil
	}
	if resp.DBCluster.EnabledCloudwatchLogsExports != nil {
		ko.Status.EnabledCloudwatchLogsExports = aws.StringSlice(resp.DBCluster.EnabledCloudwatchLogsExports)
	} else {
		ko.Status.EnabledCloudwatchLogsExports = nil
	}
	if resp.DBCluster.Endpoint != nil {
		ko.Status.Endpoint = resp.DBCluster.Endpoint
	} else {
		ko.Status.Endpoint = nil
	}
	if resp.DBCluster.Engine != nil {
		ko.Spec.Engine = resp.DBCluster.Engine
	} else {
		ko.Spec.Engine = nil
	}
	if resp.DBCluster.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.DBCluster.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.DBCluster.HostedZoneId != nil {
		ko.Status.HostedZoneID = resp.DBCluster.HostedZoneId
	} else {
		ko.Status.HostedZoneID = nil
	}
	if resp.DBCluster.KmsKeyId != nil {
		ko.Spec.KMSKeyID = resp.DBCluster.KmsKeyId
	} else {
		ko.Spec.KMSKeyID = nil
	}
	if resp.DBCluster.LatestRestorableTime != nil {
		ko.Status.LatestRestorableTime = &metav1.Time{*resp.DBCluster.LatestRestorableTime}
	} else {
		ko.Status.LatestRestorableTime = nil
	}
	if resp.DBCluster.MasterUsername != nil {
		ko.Spec.MasterUsername = resp.DBCluster.MasterUsername
	} else {
		ko.Spec.MasterUsername = nil
	}
	if resp.DBCluster.MultiAZ != nil {
		ko.Status.MultiAZ = resp.DBCluster.MultiAZ
	} else {
		ko.Status.MultiAZ = nil
	}
	if resp.DBCluster.PercentProgress != nil {
		ko.Status.PercentProgress = resp.DBCluster.PercentProgress
	} else {
		ko.Status.PercentProgress = nil
	}
	if resp.DBCluster.Port != nil {
		pcopy := aws.Int64(int64(*resp.DBCluster.Port))
		ko.Spec.Port = pcopy
	} else {
		ko.Spec.Port = nil
	}
	if resp.DBCluster.PreferredBackupWindow != nil {
		ko.Spec.PreferredBackupWindow = resp.DBCluster.PreferredBackupWindow
	} else {
		ko.Spec.PreferredBackupWindow = nil
	}
	if resp.DBCluster.PreferredMaintenanceWindow != nil {
		ko.Spec.PreferredMaintenanceWindow = resp.DBCluster.PreferredMaintenanceWindow
	} else {
		ko.Spec.PreferredMaintenanceWindow = nil
	}
	if resp.DBCluster.ReadReplicaIdentifiers != nil {
		ko.Status.ReadReplicaIdentifiers = aws.StringSlice(resp.DBCluster.ReadReplicaIdentifiers)
	} else {
		ko.Status.ReadReplicaIdentifiers = nil
	}
	if resp.DBCluster.ReaderEndpoint != nil {
		ko.Status.ReaderEndpoint = resp.DBCluster.ReaderEndpoint
	} else {
		ko.Status.ReaderEndpoint = nil
	}
	if resp.DBCluster.Status != nil {
		ko.Status.Status = resp.DBCluster.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.DBCluster.StorageEncrypted != nil {
		ko.Spec.StorageEncrypted = resp.DBCluster.StorageEncrypted
	} else {
		ko.Spec.StorageEncrypted = nil
	}
	if resp.DBCluster.VpcSecurityGroups != nil {
		f55 := []*svcapitypes.VPCSecurityGroupMembership{}
		for _, f55iter := range resp.DBCluster.VpcSecurityGroups {
			f55elem := &svcapitypes.VPCSecurityGroupMembership{}
			if f55iter.Status != nil {
				f55elem.Status = f55iter.Status
			}
			if f55iter.VpcSecurityGroupId != nil {
				f55elem.VPCSecurityGroupID = f55iter.VpcSecurityGroupId
			}
			f55 = append(f55, f55elem)
		}
		ko.Status.VPCSecurityGroups = f55
	} else {
		ko.Status.VPCSecurityGroups = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCustomUpdateRequestPayload returns an SDK-specific struct for the HTTP
// request payload of the Update API call for the resource. It is different
// from the normal newUpdateRequestsPayload in that in addition to checking for
// nil-ness of the Spec fields, it also checks to see if the delta between
// desired and observed contains a diff for the specific field. This is
// required in order to fix
// https://github.com/aws-controllers-k8s/community/issues/917
func (rm *resourceManager) newCustomUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.ModifyDBClusterInput, error) {
	res := &svcsdk.ModifyDBClusterInput{}

	res.ApplyImmediately = aws.Bool(true)
	res.AllowMajorVersionUpgrade = aws.Bool(true)
	if r.ko.Spec.BackupRetentionPeriod != nil && delta.DifferentAt("Spec.BackupRetentionPeriod") {
		res.BackupRetentionPeriod = aws.Int32(int32(*r.ko.Spec.BackupRetentionPeriod))
	}
	// NOTE(jaypipes): This is a required field in the input shape. If not set,
	// we get back a cryptic error message "1 Validation error(s) found."
	if r.ko.Spec.DBClusterIdentifier != nil {
		res.DBClusterIdentifier = r.ko.Spec.DBClusterIdentifier
	}
	if r.ko.Spec.DBClusterParameterGroupName != nil && delta.DifferentAt("Spec.DBClusterParameterGroupName") {
		res.DBClusterParameterGroupName = r.ko.Spec.DBClusterParameterGroupName
	}
	if r.ko.Spec.DeletionProtection != nil && delta.DifferentAt("Spec.DeletionProtection") {
		res.DeletionProtection = r.ko.Spec.DeletionProtection
	}
	if r.ko.Spec.EngineVersion != nil && delta.DifferentAt("Spec.EngineVersion") {
		res.EngineVersion = r.ko.Spec.EngineVersion
	}
	if r.ko.Spec.MasterUserPassword != nil && delta.DifferentAt("Spec.MasterUserPassword") {
		tmpSecret, err := rm.rr.SecretValueFromReference(ctx, r.ko.Spec.MasterUserPassword)
		if err != nil {
			return nil, err
		}
		if tmpSecret != "" {
			res.MasterUserPassword = &tmpSecret
		}
	}
	if r.ko.Spec.Port != nil && delta.DifferentAt("Spec.Port") {
		res.Port = aws.Int32(int32(*r.ko.Spec.Port))
	}
	if r.ko.Spec.PreferredBackupWindow != nil && delta.DifferentAt("Spec.PreferredBackupWindow") {
		res.PreferredBackupWindow = r.ko.Spec.PreferredBackupWindow
	}
	if r.ko.Spec.PreferredMaintenanceWindow != nil && delta.DifferentAt("Spec.PreferredMaintenanceWindow") {
		res.PreferredMaintenanceWindow = r.ko.Spec.PreferredMaintenanceWindow
	}
	if r.ko.Spec.VPCSecurityGroupIDs != nil && delta.DifferentAt("Spec.VPCSecurityGroupIDs") {
		res.VpcSecurityGroupIds = aws.ToStringSlice(r.ko.Spec.VPCSecurityGroupIDs)
	}
	return res, nil
}
