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

package db_cluster

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

	if len(a.ko.Spec.AvailabilityZones) != len(b.ko.Spec.AvailabilityZones) {
		delta.Add("Spec.AvailabilityZones", a.ko.Spec.AvailabilityZones, b.ko.Spec.AvailabilityZones)
	} else if len(a.ko.Spec.AvailabilityZones) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.AvailabilityZones, b.ko.Spec.AvailabilityZones) {
			delta.Add("Spec.AvailabilityZones", a.ko.Spec.AvailabilityZones, b.ko.Spec.AvailabilityZones)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod) {
		delta.Add("Spec.BackupRetentionPeriod", a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod)
	} else if a.ko.Spec.BackupRetentionPeriod != nil && b.ko.Spec.BackupRetentionPeriod != nil {
		if *a.ko.Spec.BackupRetentionPeriod != *b.ko.Spec.BackupRetentionPeriod {
			delta.Add("Spec.BackupRetentionPeriod", a.ko.Spec.BackupRetentionPeriod, b.ko.Spec.BackupRetentionPeriod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier) {
		delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
	} else if a.ko.Spec.DBClusterIdentifier != nil && b.ko.Spec.DBClusterIdentifier != nil {
		if *a.ko.Spec.DBClusterIdentifier != *b.ko.Spec.DBClusterIdentifier {
			delta.Add("Spec.DBClusterIdentifier", a.ko.Spec.DBClusterIdentifier, b.ko.Spec.DBClusterIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBClusterParameterGroupName, b.ko.Spec.DBClusterParameterGroupName) {
		delta.Add("Spec.DBClusterParameterGroupName", a.ko.Spec.DBClusterParameterGroupName, b.ko.Spec.DBClusterParameterGroupName)
	} else if a.ko.Spec.DBClusterParameterGroupName != nil && b.ko.Spec.DBClusterParameterGroupName != nil {
		if *a.ko.Spec.DBClusterParameterGroupName != *b.ko.Spec.DBClusterParameterGroupName {
			delta.Add("Spec.DBClusterParameterGroupName", a.ko.Spec.DBClusterParameterGroupName, b.ko.Spec.DBClusterParameterGroupName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName) {
		delta.Add("Spec.DBSubnetGroupName", a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName)
	} else if a.ko.Spec.DBSubnetGroupName != nil && b.ko.Spec.DBSubnetGroupName != nil {
		if *a.ko.Spec.DBSubnetGroupName != *b.ko.Spec.DBSubnetGroupName {
			delta.Add("Spec.DBSubnetGroupName", a.ko.Spec.DBSubnetGroupName, b.ko.Spec.DBSubnetGroupName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.DBSubnetGroupRef, b.ko.Spec.DBSubnetGroupRef) {
		delta.Add("Spec.DBSubnetGroupRef", a.ko.Spec.DBSubnetGroupRef, b.ko.Spec.DBSubnetGroupRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection) {
		delta.Add("Spec.DeletionProtection", a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection)
	} else if a.ko.Spec.DeletionProtection != nil && b.ko.Spec.DeletionProtection != nil {
		if *a.ko.Spec.DeletionProtection != *b.ko.Spec.DeletionProtection {
			delta.Add("Spec.DeletionProtection", a.ko.Spec.DeletionProtection, b.ko.Spec.DeletionProtection)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DestinationRegion, b.ko.Spec.DestinationRegion) {
		delta.Add("Spec.DestinationRegion", a.ko.Spec.DestinationRegion, b.ko.Spec.DestinationRegion)
	} else if a.ko.Spec.DestinationRegion != nil && b.ko.Spec.DestinationRegion != nil {
		if *a.ko.Spec.DestinationRegion != *b.ko.Spec.DestinationRegion {
			delta.Add("Spec.DestinationRegion", a.ko.Spec.DestinationRegion, b.ko.Spec.DestinationRegion)
		}
	}
	if len(a.ko.Spec.EnableCloudwatchLogsExports) != len(b.ko.Spec.EnableCloudwatchLogsExports) {
		delta.Add("Spec.EnableCloudwatchLogsExports", a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports)
	} else if len(a.ko.Spec.EnableCloudwatchLogsExports) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports) {
			delta.Add("Spec.EnableCloudwatchLogsExports", a.ko.Spec.EnableCloudwatchLogsExports, b.ko.Spec.EnableCloudwatchLogsExports)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Engine, b.ko.Spec.Engine) {
		delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
	} else if a.ko.Spec.Engine != nil && b.ko.Spec.Engine != nil {
		if *a.ko.Spec.Engine != *b.ko.Spec.Engine {
			delta.Add("Spec.Engine", a.ko.Spec.Engine, b.ko.Spec.Engine)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion) {
		delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
	} else if a.ko.Spec.EngineVersion != nil && b.ko.Spec.EngineVersion != nil {
		if *a.ko.Spec.EngineVersion != *b.ko.Spec.EngineVersion {
			delta.Add("Spec.EngineVersion", a.ko.Spec.EngineVersion, b.ko.Spec.EngineVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GlobalClusterIdentifier, b.ko.Spec.GlobalClusterIdentifier) {
		delta.Add("Spec.GlobalClusterIdentifier", a.ko.Spec.GlobalClusterIdentifier, b.ko.Spec.GlobalClusterIdentifier)
	} else if a.ko.Spec.GlobalClusterIdentifier != nil && b.ko.Spec.GlobalClusterIdentifier != nil {
		if *a.ko.Spec.GlobalClusterIdentifier != *b.ko.Spec.GlobalClusterIdentifier {
			delta.Add("Spec.GlobalClusterIdentifier", a.ko.Spec.GlobalClusterIdentifier, b.ko.Spec.GlobalClusterIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID) {
		delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
	} else if a.ko.Spec.KMSKeyID != nil && b.ko.Spec.KMSKeyID != nil {
		if *a.ko.Spec.KMSKeyID != *b.ko.Spec.KMSKeyID {
			delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.KMSKeyRef, b.ko.Spec.KMSKeyRef) {
		delta.Add("Spec.KMSKeyRef", a.ko.Spec.KMSKeyRef, b.ko.Spec.KMSKeyRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword) {
		delta.Add("Spec.MasterUserPassword", a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword)
	} else if a.ko.Spec.MasterUserPassword != nil && b.ko.Spec.MasterUserPassword != nil {
		if *a.ko.Spec.MasterUserPassword != *b.ko.Spec.MasterUserPassword {
			delta.Add("Spec.MasterUserPassword", a.ko.Spec.MasterUserPassword, b.ko.Spec.MasterUserPassword)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername) {
		delta.Add("Spec.MasterUsername", a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername)
	} else if a.ko.Spec.MasterUsername != nil && b.ko.Spec.MasterUsername != nil {
		if *a.ko.Spec.MasterUsername != *b.ko.Spec.MasterUsername {
			delta.Add("Spec.MasterUsername", a.ko.Spec.MasterUsername, b.ko.Spec.MasterUsername)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Port, b.ko.Spec.Port) {
		delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
	} else if a.ko.Spec.Port != nil && b.ko.Spec.Port != nil {
		if *a.ko.Spec.Port != *b.ko.Spec.Port {
			delta.Add("Spec.Port", a.ko.Spec.Port, b.ko.Spec.Port)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreSignedURL, b.ko.Spec.PreSignedURL) {
		delta.Add("Spec.PreSignedURL", a.ko.Spec.PreSignedURL, b.ko.Spec.PreSignedURL)
	} else if a.ko.Spec.PreSignedURL != nil && b.ko.Spec.PreSignedURL != nil {
		if *a.ko.Spec.PreSignedURL != *b.ko.Spec.PreSignedURL {
			delta.Add("Spec.PreSignedURL", a.ko.Spec.PreSignedURL, b.ko.Spec.PreSignedURL)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow) {
		delta.Add("Spec.PreferredBackupWindow", a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow)
	} else if a.ko.Spec.PreferredBackupWindow != nil && b.ko.Spec.PreferredBackupWindow != nil {
		if *a.ko.Spec.PreferredBackupWindow != *b.ko.Spec.PreferredBackupWindow {
			delta.Add("Spec.PreferredBackupWindow", a.ko.Spec.PreferredBackupWindow, b.ko.Spec.PreferredBackupWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow) {
		delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
	} else if a.ko.Spec.PreferredMaintenanceWindow != nil && b.ko.Spec.PreferredMaintenanceWindow != nil {
		if *a.ko.Spec.PreferredMaintenanceWindow != *b.ko.Spec.PreferredMaintenanceWindow {
			delta.Add("Spec.PreferredMaintenanceWindow", a.ko.Spec.PreferredMaintenanceWindow, b.ko.Spec.PreferredMaintenanceWindow)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapshotIdentifier, b.ko.Spec.SnapshotIdentifier) {
		delta.Add("Spec.SnapshotIdentifier", a.ko.Spec.SnapshotIdentifier, b.ko.Spec.SnapshotIdentifier)
	} else if a.ko.Spec.SnapshotIdentifier != nil && b.ko.Spec.SnapshotIdentifier != nil {
		if *a.ko.Spec.SnapshotIdentifier != *b.ko.Spec.SnapshotIdentifier {
			delta.Add("Spec.SnapshotIdentifier", a.ko.Spec.SnapshotIdentifier, b.ko.Spec.SnapshotIdentifier)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SourceRegion, b.ko.Spec.SourceRegion) {
		delta.Add("Spec.SourceRegion", a.ko.Spec.SourceRegion, b.ko.Spec.SourceRegion)
	} else if a.ko.Spec.SourceRegion != nil && b.ko.Spec.SourceRegion != nil {
		if *a.ko.Spec.SourceRegion != *b.ko.Spec.SourceRegion {
			delta.Add("Spec.SourceRegion", a.ko.Spec.SourceRegion, b.ko.Spec.SourceRegion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted) {
		delta.Add("Spec.StorageEncrypted", a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted)
	} else if a.ko.Spec.StorageEncrypted != nil && b.ko.Spec.StorageEncrypted != nil {
		if *a.ko.Spec.StorageEncrypted != *b.ko.Spec.StorageEncrypted {
			delta.Add("Spec.StorageEncrypted", a.ko.Spec.StorageEncrypted, b.ko.Spec.StorageEncrypted)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StorageType, b.ko.Spec.StorageType) {
		delta.Add("Spec.StorageType", a.ko.Spec.StorageType, b.ko.Spec.StorageType)
	} else if a.ko.Spec.StorageType != nil && b.ko.Spec.StorageType != nil {
		if *a.ko.Spec.StorageType != *b.ko.Spec.StorageType {
			delta.Add("Spec.StorageType", a.ko.Spec.StorageType, b.ko.Spec.StorageType)
		}
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if len(a.ko.Spec.VPCSecurityGroupIDs) != len(b.ko.Spec.VPCSecurityGroupIDs) {
		delta.Add("Spec.VPCSecurityGroupIDs", a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs)
	} else if len(a.ko.Spec.VPCSecurityGroupIDs) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs) {
			delta.Add("Spec.VPCSecurityGroupIDs", a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.VPCSecurityGroupRefs, b.ko.Spec.VPCSecurityGroupRefs) {
		delta.Add("Spec.VPCSecurityGroupRefs", a.ko.Spec.VPCSecurityGroupRefs, b.ko.Spec.VPCSecurityGroupRefs)
	}

	return delta
}
