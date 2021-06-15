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

// Indicates whether the user requires a password to authenticate.
type Authentication struct {
	PasswordCount *int64  `json:"passwordCount,omitempty"`
	Type          *string `json:"type_,omitempty"`
}

// Describes an Availability Zone in which the cluster is launched.
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// Contains all of the attributes of a specific cluster.
type CacheCluster struct {
	ARN                       *string      `json:"arn,omitempty"`
	AtRestEncryptionEnabled   *bool        `json:"atRestEncryptionEnabled,omitempty"`
	AuthTokenEnabled          *bool        `json:"authTokenEnabled,omitempty"`
	AuthTokenLastModifiedDate *metav1.Time `json:"authTokenLastModifiedDate,omitempty"`
	AutoMinorVersionUpgrade   *bool        `json:"autoMinorVersionUpgrade,omitempty"`
	CacheClusterCreateTime    *metav1.Time `json:"cacheClusterCreateTime,omitempty"`
	CacheClusterID            *string      `json:"cacheClusterID,omitempty"`
	CacheClusterStatus        *string      `json:"cacheClusterStatus,omitempty"`
	CacheNodeType             *string      `json:"cacheNodeType,omitempty"`
	CacheSubnetGroupName      *string      `json:"cacheSubnetGroupName,omitempty"`
	ClientDownloadLandingPage *string      `json:"clientDownloadLandingPage,omitempty"`
	// Represents the information required for client programs to connect to a cache
	// node.
	ConfigurationEndpoint      *Endpoint `json:"configurationEndpoint,omitempty"`
	Engine                     *string   `json:"engine,omitempty"`
	EngineVersion              *string   `json:"engineVersion,omitempty"`
	NumCacheNodes              *int64    `json:"numCacheNodes,omitempty"`
	PreferredAvailabilityZone  *string   `json:"preferredAvailabilityZone,omitempty"`
	PreferredMaintenanceWindow *string   `json:"preferredMaintenanceWindow,omitempty"`
	PreferredOutpostARN        *string   `json:"preferredOutpostARN,omitempty"`
	ReplicationGroupID         *string   `json:"replicationGroupID,omitempty"`
	SnapshotRetentionLimit     *int64    `json:"snapshotRetentionLimit,omitempty"`
	SnapshotWindow             *string   `json:"snapshotWindow,omitempty"`
	TransitEncryptionEnabled   *bool     `json:"transitEncryptionEnabled,omitempty"`
}

// Provides all of the details about a particular cache engine version.
type CacheEngineVersion struct {
	CacheEngineDescription        *string `json:"cacheEngineDescription,omitempty"`
	CacheEngineVersionDescription *string `json:"cacheEngineVersionDescription,omitempty"`
	CacheParameterGroupFamily     *string `json:"cacheParameterGroupFamily,omitempty"`
	Engine                        *string `json:"engine,omitempty"`
	EngineVersion                 *string `json:"engineVersion,omitempty"`
}

// Represents an individual cache node within a cluster. Each cache node runs
// its own instance of the cluster's protocol-compliant caching software - either
// Memcached or Redis.
//
// The following node types are supported by ElastiCache. Generally speaking,
// the current generation types provide more memory and computational power
// at lower cost when compared to their equivalent previous generation counterparts.
//
//    * General purpose: Current generation: M6g node types (available only
//    for Redis engine version 5.0.6 onward and for Memcached engine version
//    1.5.16 onward). cache.m6g.large, cache.m6g.xlarge, cache.m6g.2xlarge,
//    cache.m6g.4xlarge, cache.m6g.8xlarge, cache.m6g.12xlarge, cache.m6g.16xlarge
//    For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
//    M5 node types: cache.m5.large, cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge,
//    cache.m5.12xlarge, cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge,
//    cache.m4.2xlarge, cache.m4.4xlarge, cache.m4.10xlarge T3 node types: cache.t3.micro,
//    cache.t3.small, cache.t3.medium T2 node types: cache.t2.micro, cache.t2.small,
//    cache.t2.medium Previous generation: (not recommended) T1 node types:
//    cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large,
//    cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge,
//    cache.m3.2xlarge
//
//    * Compute optimized: Previous generation: (not recommended) C1 node types:
//    cache.c1.xlarge
//
//    * Memory optimized: Current generation: R6g node types (available only
//    for Redis engine version 5.0.6 onward and for Memcached engine version
//    1.5.16 onward). cache.r6g.large, cache.r6g.xlarge, cache.r6g.2xlarge,
//    cache.r6g.4xlarge, cache.r6g.8xlarge, cache.r6g.12xlarge, cache.r6g.16xlarge
//    For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
//    R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge,
//    cache.r5.12xlarge, cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge,
//    cache.r4.2xlarge, cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge
//    Previous generation: (not recommended) M2 node types: cache.m2.xlarge,
//    cache.m2.2xlarge, cache.m2.4xlarge R3 node types: cache.r3.large, cache.r3.xlarge,
//    cache.r3.2xlarge, cache.r3.4xlarge, cache.r3.8xlarge
//
// Additional node type info
//
//    * All current generation instance types are created in Amazon VPC by default.
//
//    * Redis append-only files (AOF) are not supported for T1 or T2 instances.
//
//    * Redis Multi-AZ with automatic failover is not supported on T1 instances.
//
//    * Redis configuration variables appendonly and appendfsync are not supported
//    on Redis version 2.8.22 and later.
type CacheNode struct {
	CacheNodeCreateTime      *metav1.Time `json:"cacheNodeCreateTime,omitempty"`
	CacheNodeID              *string      `json:"cacheNodeID,omitempty"`
	CacheNodeStatus          *string      `json:"cacheNodeStatus,omitempty"`
	CustomerAvailabilityZone *string      `json:"customerAvailabilityZone,omitempty"`
	CustomerOutpostARN       *string      `json:"customerOutpostARN,omitempty"`
	// Represents the information required for client programs to connect to a cache
	// node.
	Endpoint             *Endpoint `json:"endpoint,omitempty"`
	ParameterGroupStatus *string   `json:"parameterGroupStatus,omitempty"`
	SourceCacheNodeID    *string   `json:"sourceCacheNodeID,omitempty"`
}

// A parameter that has a different value for each cache node type it is applied
// to. For example, in a Redis cluster, a cache.m1.large cache node type would
// have a larger maxmemory value than a cache.m1.small type.
type CacheNodeTypeSpecificParameter struct {
	AllowedValues        *string `json:"allowedValues,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	Description          *string `json:"description,omitempty"`
	IsModifiable         *bool   `json:"isModifiable,omitempty"`
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
	ParameterName        *string `json:"parameterName,omitempty"`
	Source               *string `json:"source,omitempty"`
}

// A value that applies only to a certain cache node type.
type CacheNodeTypeSpecificValue struct {
	CacheNodeType *string `json:"cacheNodeType,omitempty"`
	Value         *string `json:"value,omitempty"`
}

// The status of the service update on the cache node
type CacheNodeUpdateStatus struct {
	CacheNodeID                  *string      `json:"cacheNodeID,omitempty"`
	NodeDeletionDate             *metav1.Time `json:"nodeDeletionDate,omitempty"`
	NodeUpdateEndDate            *metav1.Time `json:"nodeUpdateEndDate,omitempty"`
	NodeUpdateInitiatedDate      *metav1.Time `json:"nodeUpdateInitiatedDate,omitempty"`
	NodeUpdateStartDate          *metav1.Time `json:"nodeUpdateStartDate,omitempty"`
	NodeUpdateStatusModifiedDate *metav1.Time `json:"nodeUpdateStatusModifiedDate,omitempty"`
}

// Status of the cache parameter group.
type CacheParameterGroupStatus_SDK struct {
	CacheParameterGroupName *string `json:"cacheParameterGroupName,omitempty"`
	ParameterApplyStatus    *string `json:"parameterApplyStatus,omitempty"`
}

// Represents the output of a CreateCacheParameterGroup operation.
type CacheParameterGroup_SDK struct {
	ARN                       *string `json:"arn,omitempty"`
	CacheParameterGroupFamily *string `json:"cacheParameterGroupFamily,omitempty"`
	CacheParameterGroupName   *string `json:"cacheParameterGroupName,omitempty"`
	Description               *string `json:"description,omitempty"`
	IsGlobal                  *bool   `json:"isGlobal,omitempty"`
}

// Represents the output of one of the following operations:
//
//    * AuthorizeCacheSecurityGroupIngress
//
//    * CreateCacheSecurityGroup
//
//    * RevokeCacheSecurityGroupIngress
type CacheSecurityGroup struct {
	ARN                    *string `json:"arn,omitempty"`
	CacheSecurityGroupName *string `json:"cacheSecurityGroupName,omitempty"`
	Description            *string `json:"description,omitempty"`
	OwnerID                *string `json:"ownerID,omitempty"`
}

// Represents a cluster's status within a particular cache security group.
type CacheSecurityGroupMembership struct {
	CacheSecurityGroupName *string `json:"cacheSecurityGroupName,omitempty"`
	Status                 *string `json:"status,omitempty"`
}

// Represents the output of one of the following operations:
//
//    * CreateCacheSubnetGroup
//
//    * ModifyCacheSubnetGroup
type CacheSubnetGroup_SDK struct {
	ARN                         *string   `json:"arn,omitempty"`
	CacheSubnetGroupDescription *string   `json:"cacheSubnetGroupDescription,omitempty"`
	CacheSubnetGroupName        *string   `json:"cacheSubnetGroupName,omitempty"`
	Subnets                     []*Subnet `json:"subnets,omitempty"`
	VPCID                       *string   `json:"vpcID,omitempty"`
}

// Node group (shard) configuration options when adding or removing replicas.
// Each node group (shard) configuration has the following members: NodeGroupId,
// NewReplicaCount, and PreferredAvailabilityZones.
type ConfigureShard struct {
	NewReplicaCount *int64  `json:"newReplicaCount,omitempty"`
	NodeGroupID     *string `json:"nodeGroupID,omitempty"`
}

// The endpoint from which data should be migrated.
type CustomerNodeEndpoint struct {
	Address *string `json:"address,omitempty"`
	Port    *int64  `json:"port,omitempty"`
}

// Provides ownership and status information for an Amazon EC2 security group.
type EC2SecurityGroup struct {
	EC2SecurityGroupName    *string `json:"ec2SecurityGroupName,omitempty"`
	EC2SecurityGroupOwnerID *string `json:"ec2SecurityGroupOwnerID,omitempty"`
	Status                  *string `json:"status,omitempty"`
}

// Represents the information required for client programs to connect to a cache
// node.
type Endpoint struct {
	Address *string `json:"address,omitempty"`
	Port    *int64  `json:"port,omitempty"`
}

// Represents the output of a DescribeEngineDefaultParameters operation.
type EngineDefaults struct {
	CacheParameterGroupFamily *string `json:"cacheParameterGroupFamily,omitempty"`
	Marker                    *string `json:"marker,omitempty"`
}

// Represents a single occurrence of something interesting within the system.
// Some examples of events are creating a cluster, adding or removing a cache
// node, or rebooting a node.
type Event struct {
	Date             *metav1.Time `json:"date,omitempty"`
	Message          *string      `json:"message,omitempty"`
	SourceIdentifier *string      `json:"sourceIdentifier,omitempty"`
}

// Used to streamline results of a search based on the property being filtered.
type Filter struct {
	Name   *string   `json:"name,omitempty"`
	Values []*string `json:"values,omitempty"`
}

// Indicates the slot configuration and global identifier for a slice group.
type GlobalNodeGroup struct {
	GlobalNodeGroupID *string `json:"globalNodeGroupID,omitempty"`
	Slots             *string `json:"slots,omitempty"`
}

// Consists of a primary cluster that accepts writes and an associated secondary
// cluster that resides in a different AWS region. The secondary cluster accepts
// only reads. The primary cluster automatically replicates updates to the secondary
// cluster.
//
//    * The GlobalReplicationGroupIdSuffix represents the name of the Global
//    Datastore, which is what you use to associate a secondary cluster.
type GlobalReplicationGroup struct {
	ARN                               *string `json:"arn,omitempty"`
	AtRestEncryptionEnabled           *bool   `json:"atRestEncryptionEnabled,omitempty"`
	AuthTokenEnabled                  *bool   `json:"authTokenEnabled,omitempty"`
	CacheNodeType                     *string `json:"cacheNodeType,omitempty"`
	ClusterEnabled                    *bool   `json:"clusterEnabled,omitempty"`
	Engine                            *string `json:"engine,omitempty"`
	EngineVersion                     *string `json:"engineVersion,omitempty"`
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty"`
	GlobalReplicationGroupID          *string `json:"globalReplicationGroupID,omitempty"`
	Status                            *string `json:"status,omitempty"`
	TransitEncryptionEnabled          *bool   `json:"transitEncryptionEnabled,omitempty"`
}

// The name of the Global Datastore and role of this replication group in the
// Global Datastore.
type GlobalReplicationGroupInfo struct {
	GlobalReplicationGroupID         *string `json:"globalReplicationGroupID,omitempty"`
	GlobalReplicationGroupMemberRole *string `json:"globalReplicationGroupMemberRole,omitempty"`
}

// A member of a Global Datastore. It contains the Replication Group Id, the
// AWS region and the role of the replication group.
type GlobalReplicationGroupMember struct {
	AutomaticFailover      *string `json:"automaticFailover,omitempty"`
	ReplicationGroupID     *string `json:"replicationGroupID,omitempty"`
	ReplicationGroupRegion *string `json:"replicationGroupRegion,omitempty"`
	Role                   *string `json:"role,omitempty"`
	Status                 *string `json:"status,omitempty"`
}

// Represents a collection of cache nodes in a replication group. One node in
// the node group is the read/write primary node. All the other nodes are read-only
// Replica nodes.
type NodeGroup struct {
	NodeGroupID      *string            `json:"nodeGroupID,omitempty"`
	NodeGroupMembers []*NodeGroupMember `json:"nodeGroupMembers,omitempty"`
	// Represents the information required for client programs to connect to a cache
	// node.
	PrimaryEndpoint *Endpoint `json:"primaryEndpoint,omitempty"`
	// Represents the information required for client programs to connect to a cache
	// node.
	ReaderEndpoint *Endpoint `json:"readerEndpoint,omitempty"`
	Slots          *string   `json:"slots,omitempty"`
	Status         *string   `json:"status,omitempty"`
}

// Node group (shard) configuration options. Each node group (shard) configuration
// has the following: Slots, PrimaryAvailabilityZone, ReplicaAvailabilityZones,
// ReplicaCount.
type NodeGroupConfiguration struct {
	NodeGroupID              *string   `json:"nodeGroupID,omitempty"`
	PrimaryAvailabilityZone  *string   `json:"primaryAvailabilityZone,omitempty"`
	PrimaryOutpostARN        *string   `json:"primaryOutpostARN,omitempty"`
	ReplicaAvailabilityZones []*string `json:"replicaAvailabilityZones,omitempty"`
	ReplicaCount             *int64    `json:"replicaCount,omitempty"`
	ReplicaOutpostARNs       []*string `json:"replicaOutpostARNs,omitempty"`
	Slots                    *string   `json:"slots,omitempty"`
}

// Represents a single node within a node group (shard).
type NodeGroupMember struct {
	CacheClusterID            *string `json:"cacheClusterID,omitempty"`
	CacheNodeID               *string `json:"cacheNodeID,omitempty"`
	CurrentRole               *string `json:"currentRole,omitempty"`
	PreferredAvailabilityZone *string `json:"preferredAvailabilityZone,omitempty"`
	PreferredOutpostARN       *string `json:"preferredOutpostARN,omitempty"`
	// Represents the information required for client programs to connect to a cache
	// node.
	ReadEndpoint *Endpoint `json:"readEndpoint,omitempty"`
}

// The status of the service update on the node group member
type NodeGroupMemberUpdateStatus struct {
	CacheClusterID               *string      `json:"cacheClusterID,omitempty"`
	CacheNodeID                  *string      `json:"cacheNodeID,omitempty"`
	NodeDeletionDate             *metav1.Time `json:"nodeDeletionDate,omitempty"`
	NodeUpdateEndDate            *metav1.Time `json:"nodeUpdateEndDate,omitempty"`
	NodeUpdateInitiatedDate      *metav1.Time `json:"nodeUpdateInitiatedDate,omitempty"`
	NodeUpdateStartDate          *metav1.Time `json:"nodeUpdateStartDate,omitempty"`
	NodeUpdateStatusModifiedDate *metav1.Time `json:"nodeUpdateStatusModifiedDate,omitempty"`
}

// The status of the service update on the node group
type NodeGroupUpdateStatus struct {
	NodeGroupID *string `json:"nodeGroupID,omitempty"`
}

// Represents an individual cache node in a snapshot of a cluster.
type NodeSnapshot struct {
	CacheClusterID      *string      `json:"cacheClusterID,omitempty"`
	CacheNodeCreateTime *metav1.Time `json:"cacheNodeCreateTime,omitempty"`
	CacheNodeID         *string      `json:"cacheNodeID,omitempty"`
	CacheSize           *string      `json:"cacheSize,omitempty"`
	// Node group (shard) configuration options. Each node group (shard) configuration
	// has the following: Slots, PrimaryAvailabilityZone, ReplicaAvailabilityZones,
	// ReplicaCount.
	NodeGroupConfiguration *NodeGroupConfiguration `json:"nodeGroupConfiguration,omitempty"`
	NodeGroupID            *string                 `json:"nodeGroupID,omitempty"`
	SnapshotCreateTime     *metav1.Time            `json:"snapshotCreateTime,omitempty"`
}

// Describes a notification topic and its status. Notification topics are used
// for publishing ElastiCache events to subscribers using Amazon Simple Notification
// Service (SNS).
type NotificationConfiguration struct {
	TopicARN    *string `json:"topicARN,omitempty"`
	TopicStatus *string `json:"topicStatus,omitempty"`
}

// Describes an individual setting that controls some aspect of ElastiCache
// behavior.
type Parameter struct {
	AllowedValues        *string `json:"allowedValues,omitempty"`
	DataType             *string `json:"dataType,omitempty"`
	Description          *string `json:"description,omitempty"`
	IsModifiable         *bool   `json:"isModifiable,omitempty"`
	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`
	ParameterName        *string `json:"parameterName,omitempty"`
	ParameterValue       *string `json:"parameterValue,omitempty"`
	Source               *string `json:"source,omitempty"`
}

// Describes a name-value pair that is used to update the value of a parameter.
type ParameterNameValue struct {
	ParameterName  *string `json:"parameterName,omitempty"`
	ParameterValue *string `json:"parameterValue,omitempty"`
}

// A group of settings that are applied to the cluster in the future, or that
// are currently being applied.
type PendingModifiedValues struct {
	AuthTokenStatus *string `json:"authTokenStatus,omitempty"`
	CacheNodeType   *string `json:"cacheNodeType,omitempty"`
	EngineVersion   *string `json:"engineVersion,omitempty"`
	NumCacheNodes   *int64  `json:"numCacheNodes,omitempty"`
}

// Update action that has been processed for the corresponding apply/stop request
type ProcessedUpdateAction struct {
	CacheClusterID     *string `json:"cacheClusterID,omitempty"`
	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`
	ServiceUpdateName  *string `json:"serviceUpdateName,omitempty"`
}

// Contains the specific price and frequency of a recurring charges for a reserved
// cache node, or for a reserved cache node offering.
type RecurringCharge struct {
	RecurringChargeAmount    *float64 `json:"recurringChargeAmount,omitempty"`
	RecurringChargeFrequency *string  `json:"recurringChargeFrequency,omitempty"`
}

// A list of the replication groups
type RegionalConfiguration struct {
	ReplicationGroupID     *string `json:"replicationGroupID,omitempty"`
	ReplicationGroupRegion *string `json:"replicationGroupRegion,omitempty"`
}

// The settings to be applied to the Redis replication group, either immediately
// or during the next maintenance window.
type ReplicationGroupPendingModifiedValues struct {
	AuthTokenStatus         *string `json:"authTokenStatus,omitempty"`
	AutomaticFailoverStatus *string `json:"automaticFailoverStatus,omitempty"`
	PrimaryClusterID        *string `json:"primaryClusterID,omitempty"`
	// The status of an online resharding operation.
	Resharding *ReshardingStatus `json:"resharding,omitempty"`
	// The status of the user group update.
	UserGroups *UserGroupsUpdateStatus `json:"userGroups,omitempty"`
}

// Contains all of the attributes of a specific Redis replication group.
type ReplicationGroup_SDK struct {
	ARN                       *string      `json:"arn,omitempty"`
	AtRestEncryptionEnabled   *bool        `json:"atRestEncryptionEnabled,omitempty"`
	AuthTokenEnabled          *bool        `json:"authTokenEnabled,omitempty"`
	AuthTokenLastModifiedDate *metav1.Time `json:"authTokenLastModifiedDate,omitempty"`
	AutomaticFailover         *string      `json:"automaticFailover,omitempty"`
	CacheNodeType             *string      `json:"cacheNodeType,omitempty"`
	ClusterEnabled            *bool        `json:"clusterEnabled,omitempty"`
	// Represents the information required for client programs to connect to a cache
	// node.
	ConfigurationEndpoint *Endpoint `json:"configurationEndpoint,omitempty"`
	Description           *string   `json:"description,omitempty"`
	// The name of the Global Datastore and role of this replication group in the
	// Global Datastore.
	GlobalReplicationGroupInfo *GlobalReplicationGroupInfo `json:"globalReplicationGroupInfo,omitempty"`
	KMSKeyID                   *string                     `json:"kmsKeyID,omitempty"`
	MemberClusters             []*string                   `json:"memberClusters,omitempty"`
	MemberClustersOutpostARNs  []*string                   `json:"memberClustersOutpostARNs,omitempty"`
	MultiAZ                    *string                     `json:"multiAZ,omitempty"`
	NodeGroups                 []*NodeGroup                `json:"nodeGroups,omitempty"`
	// The settings to be applied to the Redis replication group, either immediately
	// or during the next maintenance window.
	PendingModifiedValues    *ReplicationGroupPendingModifiedValues `json:"pendingModifiedValues,omitempty"`
	ReplicationGroupID       *string                                `json:"replicationGroupID,omitempty"`
	SnapshotRetentionLimit   *int64                                 `json:"snapshotRetentionLimit,omitempty"`
	SnapshotWindow           *string                                `json:"snapshotWindow,omitempty"`
	SnapshottingClusterID    *string                                `json:"snapshottingClusterID,omitempty"`
	Status                   *string                                `json:"status,omitempty"`
	TransitEncryptionEnabled *bool                                  `json:"transitEncryptionEnabled,omitempty"`
	UserGroupIDs             []*string                              `json:"userGroupIDs,omitempty"`
}

// Represents the output of a PurchaseReservedCacheNodesOffering operation.
type ReservedCacheNode struct {
	CacheNodeCount               *int64       `json:"cacheNodeCount,omitempty"`
	CacheNodeType                *string      `json:"cacheNodeType,omitempty"`
	Duration                     *int64       `json:"duration,omitempty"`
	FixedPrice                   *float64     `json:"fixedPrice,omitempty"`
	OfferingType                 *string      `json:"offeringType,omitempty"`
	ProductDescription           *string      `json:"productDescription,omitempty"`
	ReservationARN               *string      `json:"reservationARN,omitempty"`
	ReservedCacheNodeID          *string      `json:"reservedCacheNodeID,omitempty"`
	ReservedCacheNodesOfferingID *string      `json:"reservedCacheNodesOfferingID,omitempty"`
	StartTime                    *metav1.Time `json:"startTime,omitempty"`
	State                        *string      `json:"state,omitempty"`
	UsagePrice                   *float64     `json:"usagePrice,omitempty"`
}

// Describes all of the attributes of a reserved cache node offering.
type ReservedCacheNodesOffering struct {
	CacheNodeType                *string  `json:"cacheNodeType,omitempty"`
	Duration                     *int64   `json:"duration,omitempty"`
	FixedPrice                   *float64 `json:"fixedPrice,omitempty"`
	OfferingType                 *string  `json:"offeringType,omitempty"`
	ProductDescription           *string  `json:"productDescription,omitempty"`
	ReservedCacheNodesOfferingID *string  `json:"reservedCacheNodesOfferingID,omitempty"`
	UsagePrice                   *float64 `json:"usagePrice,omitempty"`
}

// A list of PreferredAvailabilityZones objects that specifies the configuration
// of a node group in the resharded cluster.
type ReshardingConfiguration struct {
	NodeGroupID                *string   `json:"nodeGroupID,omitempty"`
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones,omitempty"`
}

// The status of an online resharding operation.
type ReshardingStatus struct {
	// Represents the progress of an online resharding operation.
	SlotMigration *SlotMigration `json:"slotMigration,omitempty"`
}

// Represents a single cache security group and its status.
type SecurityGroupMembership struct {
	SecurityGroupID *string `json:"securityGroupID,omitempty"`
	Status          *string `json:"status,omitempty"`
}

// An update that you can apply to your Redis clusters.
type ServiceUpdate struct {
	AutoUpdateAfterRecommendedApplyByDate *bool        `json:"autoUpdateAfterRecommendedApplyByDate,omitempty"`
	Engine                                *string      `json:"engine,omitempty"`
	EngineVersion                         *string      `json:"engineVersion,omitempty"`
	EstimatedUpdateTime                   *string      `json:"estimatedUpdateTime,omitempty"`
	ServiceUpdateDescription              *string      `json:"serviceUpdateDescription,omitempty"`
	ServiceUpdateEndDate                  *metav1.Time `json:"serviceUpdateEndDate,omitempty"`
	ServiceUpdateName                     *string      `json:"serviceUpdateName,omitempty"`
	ServiceUpdateRecommendedApplyByDate   *metav1.Time `json:"serviceUpdateRecommendedApplyByDate,omitempty"`
	ServiceUpdateReleaseDate              *metav1.Time `json:"serviceUpdateReleaseDate,omitempty"`
}

// Represents the progress of an online resharding operation.
type SlotMigration struct {
	ProgressPercentage *float64 `json:"progressPercentage,omitempty"`
}

// Represents a copy of an entire Redis cluster as of the time when the snapshot
// was taken.
type Snapshot_SDK struct {
	ARN                         *string         `json:"arn,omitempty"`
	AutoMinorVersionUpgrade     *bool           `json:"autoMinorVersionUpgrade,omitempty"`
	AutomaticFailover           *string         `json:"automaticFailover,omitempty"`
	CacheClusterCreateTime      *metav1.Time    `json:"cacheClusterCreateTime,omitempty"`
	CacheClusterID              *string         `json:"cacheClusterID,omitempty"`
	CacheNodeType               *string         `json:"cacheNodeType,omitempty"`
	CacheParameterGroupName     *string         `json:"cacheParameterGroupName,omitempty"`
	CacheSubnetGroupName        *string         `json:"cacheSubnetGroupName,omitempty"`
	Engine                      *string         `json:"engine,omitempty"`
	EngineVersion               *string         `json:"engineVersion,omitempty"`
	KMSKeyID                    *string         `json:"kmsKeyID,omitempty"`
	NodeSnapshots               []*NodeSnapshot `json:"nodeSnapshots,omitempty"`
	NumCacheNodes               *int64          `json:"numCacheNodes,omitempty"`
	NumNodeGroups               *int64          `json:"numNodeGroups,omitempty"`
	Port                        *int64          `json:"port,omitempty"`
	PreferredAvailabilityZone   *string         `json:"preferredAvailabilityZone,omitempty"`
	PreferredMaintenanceWindow  *string         `json:"preferredMaintenanceWindow,omitempty"`
	PreferredOutpostARN         *string         `json:"preferredOutpostARN,omitempty"`
	ReplicationGroupDescription *string         `json:"replicationGroupDescription,omitempty"`
	ReplicationGroupID          *string         `json:"replicationGroupID,omitempty"`
	SnapshotName                *string         `json:"snapshotName,omitempty"`
	SnapshotRetentionLimit      *int64          `json:"snapshotRetentionLimit,omitempty"`
	SnapshotSource              *string         `json:"snapshotSource,omitempty"`
	SnapshotStatus              *string         `json:"snapshotStatus,omitempty"`
	SnapshotWindow              *string         `json:"snapshotWindow,omitempty"`
	TopicARN                    *string         `json:"topicARN,omitempty"`
	VPCID                       *string         `json:"vpcID,omitempty"`
}

// Represents the subnet associated with a cluster. This parameter refers to
// subnets defined in Amazon Virtual Private Cloud (Amazon VPC) and used with
// ElastiCache.
type Subnet struct {
	// Describes an Availability Zone in which the cluster is launched.
	SubnetAvailabilityZone *AvailabilityZone `json:"subnetAvailabilityZone,omitempty"`
	SubnetIdentifier       *string           `json:"subnetIdentifier,omitempty"`
	// The ID of the outpost subnet.
	SubnetOutpost *SubnetOutpost `json:"subnetOutpost,omitempty"`
}

// The ID of the outpost subnet.
type SubnetOutpost struct {
	SubnetOutpostARN *string `json:"subnetOutpostARN,omitempty"`
}

// A cost allocation Tag that can be added to an ElastiCache cluster or replication
// group. Tags are composed of a Key/Value pair. A tag with a null Value is
// permitted.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Filters update actions from the service updates that are in available status
// during the time range.
type TimeRangeFilter struct {
	EndTime   *metav1.Time `json:"endTime,omitempty"`
	StartTime *metav1.Time `json:"startTime,omitempty"`
}

// Update action that has failed to be processed for the corresponding apply/stop
// request
type UnprocessedUpdateAction struct {
	CacheClusterID     *string `json:"cacheClusterID,omitempty"`
	ErrorMessage       *string `json:"errorMessage,omitempty"`
	ErrorType          *string `json:"errorType,omitempty"`
	ReplicationGroupID *string `json:"replicationGroupID,omitempty"`
	ServiceUpdateName  *string `json:"serviceUpdateName,omitempty"`
}

// The status of the service update for a specific replication group
type UpdateAction struct {
	CacheClusterID                      *string      `json:"cacheClusterID,omitempty"`
	Engine                              *string      `json:"engine,omitempty"`
	EstimatedUpdateTime                 *string      `json:"estimatedUpdateTime,omitempty"`
	NodesUpdated                        *string      `json:"nodesUpdated,omitempty"`
	ReplicationGroupID                  *string      `json:"replicationGroupID,omitempty"`
	ServiceUpdateName                   *string      `json:"serviceUpdateName,omitempty"`
	ServiceUpdateRecommendedApplyByDate *metav1.Time `json:"serviceUpdateRecommendedApplyByDate,omitempty"`
	ServiceUpdateReleaseDate            *metav1.Time `json:"serviceUpdateReleaseDate,omitempty"`
	UpdateActionAvailableDate           *metav1.Time `json:"updateActionAvailableDate,omitempty"`
	UpdateActionStatusModifiedDate      *metav1.Time `json:"updateActionStatusModifiedDate,omitempty"`
}

// Returns the updates being applied to the user group.
type UserGroupPendingChanges struct {
	UserIDsToAdd    []*string `json:"userIDsToAdd,omitempty"`
	UserIDsToRemove []*string `json:"userIDsToRemove,omitempty"`
}

type UserGroup_SDK struct {
	ARN    *string `json:"arn,omitempty"`
	Engine *string `json:"engine,omitempty"`
	// Returns the updates being applied to the user group.
	PendingChanges    *UserGroupPendingChanges `json:"pendingChanges,omitempty"`
	ReplicationGroups []*string                `json:"replicationGroups,omitempty"`
	Status            *string                  `json:"status,omitempty"`
	UserGroupID       *string                  `json:"userGroupID,omitempty"`
	UserIDs           []*string                `json:"userIDs,omitempty"`
}

// The status of the user group update.
type UserGroupsUpdateStatus struct {
	UserGroupIDsToAdd    []*string `json:"userGroupIDsToAdd,omitempty"`
	UserGroupIDsToRemove []*string `json:"userGroupIDsToRemove,omitempty"`
}

type User_SDK struct {
	ARN          *string `json:"arn,omitempty"`
	AccessString *string `json:"accessString,omitempty"`
	// Indicates whether the user requires a password to authenticate.
	Authentication *Authentication `json:"authentication,omitempty"`
	Engine         *string         `json:"engine,omitempty"`
	Status         *string         `json:"status,omitempty"`
	UserGroupIDs   []*string       `json:"userGroupIDs,omitempty"`
	UserID         *string         `json:"userID,omitempty"`
	UserName       *string         `json:"userName,omitempty"`
}
