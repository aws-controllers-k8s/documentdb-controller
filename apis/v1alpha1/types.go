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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Information about an Availability Zone.
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// A certificate authority (CA) certificate for an Amazon Web Services account.
type Certificate struct {
	CertificateARN        *string      `json:"certificateARN,omitempty"`
	CertificateIdentifier *string      `json:"certificateIdentifier,omitempty"`
	CertificateType       *string      `json:"certificateType,omitempty"`
	Thumbprint            *string      `json:"thumbprint,omitempty"`
	ValidFrom             *metav1.Time `json:"validFrom,omitempty"`
	ValidTill             *metav1.Time `json:"validTill,omitempty"`
}

// Returns the details of the DB instance’s server certificate.
//
// For more information, see Updating Your Amazon DocumentDB TLS Certificates
// (https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)
// and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)
// in the Amazon DocumentDB Developer Guide.
type CertificateDetails struct {
	CAIdentifier *string      `json:"cAIdentifier,omitempty"`
	ValidTill    *metav1.Time `json:"validTill,omitempty"`
}

// The configuration setting for the log types to be enabled for export to Amazon
// CloudWatch Logs for a specific instance or cluster.
//
// The EnableLogTypes and DisableLogTypes arrays determine which logs are exported
// (or not exported) to CloudWatch Logs. The values within these arrays depend
// on the engine that is being used.
type CloudwatchLogsExportConfiguration struct {
	DisableLogTypes []*string `json:"disableLogTypes,omitempty"`
	EnableLogTypes  []*string `json:"enableLogTypes,omitempty"`
}

// Contains information about an instance that is part of a cluster.
type DBClusterMember struct {
	DBClusterParameterGroupStatus *string `json:"dbClusterParameterGroupStatus,omitempty"`
	DBInstanceIdentifier          *string `json:"dbInstanceIdentifier,omitempty"`
	IsClusterWriter               *bool   `json:"isClusterWriter,omitempty"`
	PromotionTier                 *int64  `json:"promotionTier,omitempty"`
}

// Detailed information about a cluster parameter group.
type DBClusterParameterGroup struct {
	DBClusterParameterGroupARN  *string `json:"dbClusterParameterGroupARN,omitempty"`
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty"`
	DBParameterGroupFamily      *string `json:"dbParameterGroupFamily,omitempty"`
	Description                 *string `json:"description,omitempty"`
}

// Describes an Identity and Access Management (IAM) role that is associated
// with a cluster.
type DBClusterRole struct {
	RoleARN *string `json:"roleARN,omitempty"`
	Status  *string `json:"status,omitempty"`
}

// Detailed information about a cluster snapshot.
type DBClusterSnapshot struct {
	AvailabilityZones           []*string    `json:"availabilityZones,omitempty"`
	ClusterCreateTime           *metav1.Time `json:"clusterCreateTime,omitempty"`
	DBClusterIdentifier         *string      `json:"dbClusterIdentifier,omitempty"`
	DBClusterSnapshotARN        *string      `json:"dbClusterSnapshotARN,omitempty"`
	DBClusterSnapshotIdentifier *string      `json:"dbClusterSnapshotIdentifier,omitempty"`
	Engine                      *string      `json:"engine,omitempty"`
	EngineVersion               *string      `json:"engineVersion,omitempty"`
	KMSKeyID                    *string      `json:"kmsKeyID,omitempty"`
	MasterUsername              *string      `json:"masterUsername,omitempty"`
	PercentProgress             *int64       `json:"percentProgress,omitempty"`
	Port                        *int64       `json:"port,omitempty"`
	SnapshotCreateTime          *metav1.Time `json:"snapshotCreateTime,omitempty"`
	SnapshotType                *string      `json:"snapshotType,omitempty"`
	SourceDBClusterSnapshotARN  *string      `json:"sourceDBClusterSnapshotARN,omitempty"`
	Status                      *string      `json:"status,omitempty"`
	StorageEncrypted            *bool        `json:"storageEncrypted,omitempty"`
	StorageType                 *string      `json:"storageType,omitempty"`
	VPCID                       *string      `json:"vpcID,omitempty"`
}

// Contains the name and values of a manual cluster snapshot attribute.
//
// Manual cluster snapshot attributes are used to authorize other Amazon Web
// Services accounts to restore a manual cluster snapshot.
type DBClusterSnapshotAttribute struct {
	AttributeName *string `json:"attributeName,omitempty"`
}

// Detailed information about the attributes that are associated with a cluster
// snapshot.
type DBClusterSnapshotAttributesResult struct {
	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty"`
}

// Detailed information about a cluster.
type DBCluster_SDK struct {
	AssociatedRoles              []*DBClusterRole              `json:"associatedRoles,omitempty"`
	AvailabilityZones            []*string                     `json:"availabilityZones,omitempty"`
	BackupRetentionPeriod        *int64                        `json:"backupRetentionPeriod,omitempty"`
	CloneGroupID                 *string                       `json:"cloneGroupID,omitempty"`
	ClusterCreateTime            *metav1.Time                  `json:"clusterCreateTime,omitempty"`
	DBClusterARN                 *string                       `json:"dbClusterARN,omitempty"`
	DBClusterIdentifier          *string                       `json:"dbClusterIdentifier,omitempty"`
	DBClusterMembers             []*DBClusterMember            `json:"dbClusterMembers,omitempty"`
	DBClusterParameterGroup      *string                       `json:"dbClusterParameterGroup,omitempty"`
	DBSubnetGroup                *string                       `json:"dbSubnetGroup,omitempty"`
	DBClusterResourceID          *string                       `json:"dbClusterResourceID,omitempty"`
	DeletionProtection           *bool                         `json:"deletionProtection,omitempty"`
	EarliestRestorableTime       *metav1.Time                  `json:"earliestRestorableTime,omitempty"`
	EnabledCloudwatchLogsExports []*string                     `json:"enabledCloudwatchLogsExports,omitempty"`
	Endpoint                     *string                       `json:"endpoint,omitempty"`
	Engine                       *string                       `json:"engine,omitempty"`
	EngineVersion                *string                       `json:"engineVersion,omitempty"`
	HostedZoneID                 *string                       `json:"hostedZoneID,omitempty"`
	KMSKeyID                     *string                       `json:"kmsKeyID,omitempty"`
	LatestRestorableTime         *metav1.Time                  `json:"latestRestorableTime,omitempty"`
	MasterUsername               *string                       `json:"masterUsername,omitempty"`
	MultiAZ                      *bool                         `json:"multiAZ,omitempty"`
	PercentProgress              *string                       `json:"percentProgress,omitempty"`
	Port                         *int64                        `json:"port,omitempty"`
	PreferredBackupWindow        *string                       `json:"preferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow   *string                       `json:"preferredMaintenanceWindow,omitempty"`
	ReadReplicaIdentifiers       []*string                     `json:"readReplicaIdentifiers,omitempty"`
	ReaderEndpoint               *string                       `json:"readerEndpoint,omitempty"`
	ReplicationSourceIdentifier  *string                       `json:"replicationSourceIdentifier,omitempty"`
	Status                       *string                       `json:"status,omitempty"`
	StorageEncrypted             *bool                         `json:"storageEncrypted,omitempty"`
	StorageType                  *string                       `json:"storageType,omitempty"`
	VPCSecurityGroups            []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

// Detailed information about an engine version.
type DBEngineVersion struct {
	DBEngineDescription                       *string   `json:"dbEngineDescription,omitempty"`
	DBEngineVersionDescription                *string   `json:"dbEngineVersionDescription,omitempty"`
	DBParameterGroupFamily                    *string   `json:"dbParameterGroupFamily,omitempty"`
	Engine                                    *string   `json:"engine,omitempty"`
	EngineVersion                             *string   `json:"engineVersion,omitempty"`
	ExportableLogTypes                        []*string `json:"exportableLogTypes,omitempty"`
	SupportsCertificateRotationWithoutRestart *bool     `json:"supportsCertificateRotationWithoutRestart,omitempty"`
	SupportsLogExportsToCloudwatchLogs        *bool     `json:"supportsLogExportsToCloudwatchLogs,omitempty"`
}

// Provides a list of status information for an instance.
type DBInstanceStatusInfo struct {
	Message    *string `json:"message,omitempty"`
	Normal     *bool   `json:"normal,omitempty"`
	Status     *string `json:"status,omitempty"`
	StatusType *string `json:"statusType,omitempty"`
}

// Detailed information about an instance.
type DBInstance_SDK struct {
	AutoMinorVersionUpgrade *bool   `json:"autoMinorVersionUpgrade,omitempty"`
	AvailabilityZone        *string `json:"availabilityZone,omitempty"`
	BackupRetentionPeriod   *int64  `json:"backupRetentionPeriod,omitempty"`
	CACertificateIdentifier *string `json:"caCertificateIdentifier,omitempty"`
	// Returns the details of the DB instance’s server certificate.
	//
	// For more information, see Updating Your Amazon DocumentDB TLS Certificates
	// (https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)
	// and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)
	// in the Amazon DocumentDB Developer Guide.
	CertificateDetails   *CertificateDetails `json:"certificateDetails,omitempty"`
	CopyTagsToSnapshot   *bool               `json:"copyTagsToSnapshot,omitempty"`
	DBClusterIdentifier  *string             `json:"dbClusterIdentifier,omitempty"`
	DBInstanceARN        *string             `json:"dbInstanceARN,omitempty"`
	DBInstanceClass      *string             `json:"dbInstanceClass,omitempty"`
	DBInstanceIdentifier *string             `json:"dbInstanceIdentifier,omitempty"`
	DBInstanceStatus     *string             `json:"dbInstanceStatus,omitempty"`
	// Detailed information about a subnet group.
	DBSubnetGroup                *DBSubnetGroup_SDK `json:"dbSubnetGroup,omitempty"`
	DBIResourceID                *string            `json:"dbiResourceID,omitempty"`
	EnabledCloudwatchLogsExports []*string          `json:"enabledCloudwatchLogsExports,omitempty"`
	// Network information for accessing a cluster or instance. Client programs
	// must specify a valid endpoint to access these Amazon DocumentDB resources.
	Endpoint             *Endpoint    `json:"endpoint,omitempty"`
	Engine               *string      `json:"engine,omitempty"`
	EngineVersion        *string      `json:"engineVersion,omitempty"`
	InstanceCreateTime   *metav1.Time `json:"instanceCreateTime,omitempty"`
	KMSKeyID             *string      `json:"kmsKeyID,omitempty"`
	LatestRestorableTime *metav1.Time `json:"latestRestorableTime,omitempty"`
	// One or more modified settings for an instance. These modified settings have
	// been requested, but haven't been applied yet.
	PendingModifiedValues       *PendingModifiedValues        `json:"pendingModifiedValues,omitempty"`
	PerformanceInsightsEnabled  *bool                         `json:"performanceInsightsEnabled,omitempty"`
	PerformanceInsightsKMSKeyID *string                       `json:"performanceInsightsKMSKeyID,omitempty"`
	PreferredBackupWindow       *string                       `json:"preferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow  *string                       `json:"preferredMaintenanceWindow,omitempty"`
	PromotionTier               *int64                        `json:"promotionTier,omitempty"`
	PubliclyAccessible          *bool                         `json:"publiclyAccessible,omitempty"`
	StatusInfos                 []*DBInstanceStatusInfo       `json:"statusInfos,omitempty"`
	StorageEncrypted            *bool                         `json:"storageEncrypted,omitempty"`
	VPCSecurityGroups           []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

// Detailed information about a subnet group.
type DBSubnetGroup_SDK struct {
	DBSubnetGroupARN         *string   `json:"dbSubnetGroupARN,omitempty"`
	DBSubnetGroupDescription *string   `json:"dbSubnetGroupDescription,omitempty"`
	DBSubnetGroupName        *string   `json:"dbSubnetGroupName,omitempty"`
	SubnetGroupStatus        *string   `json:"subnetGroupStatus,omitempty"`
	Subnets                  []*Subnet `json:"subnets,omitempty"`
	VPCID                    *string   `json:"vpcID,omitempty"`
}

// Network information for accessing a cluster or instance. Client programs
// must specify a valid endpoint to access these Amazon DocumentDB resources.
type Endpoint struct {
	Address      *string `json:"address,omitempty"`
	HostedZoneID *string `json:"hostedZoneID,omitempty"`
	Port         *int64  `json:"port,omitempty"`
}

// Contains the result of a successful invocation of the DescribeEngineDefaultClusterParameters
// operation.
type EngineDefaults struct {
	DBParameterGroupFamily *string `json:"dbParameterGroupFamily,omitempty"`
	Marker                 *string `json:"marker,omitempty"`
}

// Detailed information about an event.
type Event struct {
	Date             *metav1.Time `json:"date,omitempty"`
	Message          *string      `json:"message,omitempty"`
	SourceARN        *string      `json:"sourceARN,omitempty"`
	SourceIdentifier *string      `json:"sourceIdentifier,omitempty"`
}

// An event source type, accompanied by one or more event category names.
type EventCategoriesMap struct {
	SourceType *string `json:"sourceType,omitempty"`
}

// Detailed information about an event to which you have subscribed.
type EventSubscription struct {
	CustSubscriptionID       *string `json:"custSubscriptionID,omitempty"`
	CustomerAWSID            *string `json:"customerAWSID,omitempty"`
	Enabled                  *bool   `json:"enabled,omitempty"`
	EventSubscriptionARN     *string `json:"eventSubscriptionARN,omitempty"`
	SNSTopicARN              *string `json:"snsTopicARN,omitempty"`
	SourceType               *string `json:"sourceType,omitempty"`
	Status                   *string `json:"status,omitempty"`
	SubscriptionCreationTime *string `json:"subscriptionCreationTime,omitempty"`
}

// A named set of filter values, used to return a more specific list of results.
// You can use a filter to match a set of resources by specific criteria, such
// as IDs.
//
// Wildcards are not supported in filters.
type Filter struct {
	Name   *string   `json:"name,omitempty"`
	Values []*string `json:"values,omitempty"`
}

// A data type representing an Amazon DocumentDB global cluster.
type GlobalCluster struct {
	DatabaseName            *string `json:"databaseName,omitempty"`
	DeletionProtection      *bool   `json:"deletionProtection,omitempty"`
	Engine                  *string `json:"engine,omitempty"`
	EngineVersion           *string `json:"engineVersion,omitempty"`
	GlobalClusterARN        *string `json:"globalClusterARN,omitempty"`
	GlobalClusterIdentifier *string `json:"globalClusterIdentifier,omitempty"`
	GlobalClusterResourceID *string `json:"globalClusterResourceID,omitempty"`
	Status                  *string `json:"status,omitempty"`
	StorageEncrypted        *bool   `json:"storageEncrypted,omitempty"`
}

// A data structure with information about any primary and secondary clusters
// associated with an Amazon DocumentDB global clusters.
type GlobalClusterMember struct {
	DBClusterARN *string `json:"dbClusterARN,omitempty"`
	IsWriter     *bool   `json:"isWriter,omitempty"`
}

// The options that are available for an instance.
type OrderableDBInstanceOption struct {
	DBInstanceClass *string `json:"dbInstanceClass,omitempty"`
	Engine          *string `json:"engine,omitempty"`
	EngineVersion   *string `json:"engineVersion,omitempty"`
	LicenseModel    *string `json:"licenseModel,omitempty"`
	StorageType     *string `json:"storageType,omitempty"`
	VPC             *bool   `json:"vpc,omitempty"`
}

// Detailed information about an individual parameter.
type Parameter struct {
	AllowedValues        *string `json:"allowedValues,omitempty"`
	ApplyType            *string `json:"applyType,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	Description          *string `json:"description,omitempty"`
	IsModifiable         *bool   `json:"isModifiable,omitempty"`
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
	ParameterName        *string `json:"parameterName,omitempty"`
	ParameterValue       *string `json:"parameterValue,omitempty"`
	Source               *string `json:"source,omitempty"`
}

// A list of the log types whose configuration is still pending. These log types
// are in the process of being activated or deactivated.
type PendingCloudwatchLogsExports struct {
	LogTypesToDisable []*string `json:"logTypesToDisable,omitempty"`
	LogTypesToEnable  []*string `json:"logTypesToEnable,omitempty"`
}

// Provides information about a pending maintenance action for a resource.
type PendingMaintenanceAction struct {
	Action               *string      `json:"action,omitempty"`
	AutoAppliedAfterDate *metav1.Time `json:"autoAppliedAfterDate,omitempty"`
	CurrentApplyDate     *metav1.Time `json:"currentApplyDate,omitempty"`
	Description          *string      `json:"description,omitempty"`
	ForcedApplyDate      *metav1.Time `json:"forcedApplyDate,omitempty"`
	OptInStatus          *string      `json:"optInStatus,omitempty"`
}

// One or more modified settings for an instance. These modified settings have
// been requested, but haven't been applied yet.
type PendingModifiedValues struct {
	AllocatedStorage        *int64  `json:"allocatedStorage,omitempty"`
	BackupRetentionPeriod   *int64  `json:"backupRetentionPeriod,omitempty"`
	CACertificateIdentifier *string `json:"caCertificateIdentifier,omitempty"`
	DBInstanceClass         *string `json:"dbInstanceClass,omitempty"`
	DBInstanceIdentifier    *string `json:"dbInstanceIdentifier,omitempty"`
	DBSubnetGroupName       *string `json:"dbSubnetGroupName,omitempty"`
	EngineVersion           *string `json:"engineVersion,omitempty"`
	IOPS                    *int64  `json:"iops,omitempty"`
	LicenseModel            *string `json:"licenseModel,omitempty"`
	MasterUserPassword      *string `json:"masterUserPassword,omitempty"`
	MultiAZ                 *bool   `json:"multiAZ,omitempty"`
	// A list of the log types whose configuration is still pending. These log types
	// are in the process of being activated or deactivated.
	PendingCloudwatchLogsExports *PendingCloudwatchLogsExports `json:"pendingCloudwatchLogsExports,omitempty"`
	Port                         *int64                        `json:"port,omitempty"`
	StorageType                  *string                       `json:"storageType,omitempty"`
}

// Represents the output of ApplyPendingMaintenanceAction.
type ResourcePendingMaintenanceActions struct {
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty"`
}

// Detailed information about a subnet.
type Subnet struct {
	// Information about an Availability Zone.
	SubnetAvailabilityZone *AvailabilityZone `json:"subnetAvailabilityZone,omitempty"`
	SubnetIdentifier       *string           `json:"subnetIdentifier,omitempty"`
	SubnetStatus           *string           `json:"subnetStatus,omitempty"`
}

// Metadata assigned to an Amazon DocumentDB resource consisting of a key-value
// pair.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// The version of the database engine that an instance can be upgraded to.
type UpgradeTarget struct {
	AutoUpgrade           *bool   `json:"autoUpgrade,omitempty"`
	Description           *string `json:"description,omitempty"`
	Engine                *string `json:"engine,omitempty"`
	EngineVersion         *string `json:"engineVersion,omitempty"`
	IsMajorVersionUpgrade *bool   `json:"isMajorVersionUpgrade,omitempty"`
}

// Used as a response element for queries on virtual private cloud (VPC) security
// group membership.
type VPCSecurityGroupMembership struct {
	Status             *string `json:"status,omitempty"`
	VPCSecurityGroupID *string `json:"vpcSecurityGroupID,omitempty"`
}
