/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClusterModeObservation struct {
}

type ClusterModeParameters struct {

	// +kubebuilder:validation:Optional
	NumNodeGroups *int64 `json:"numNodeGroups,omitempty" tf:"num_node_groups,omitempty"`

	// +kubebuilder:validation:Required
	ReplicasPerNodeGroup *int64 `json:"replicasPerNodeGroup" tf:"replicas_per_node_group,omitempty"`
}

type ReplicationGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ClusterEnabled *bool `json:"clusterEnabled,omitempty" tf:"cluster_enabled,omitempty"`

	ConfigurationEndpointAddress *string `json:"configurationEndpointAddress,omitempty" tf:"configuration_endpoint_address,omitempty"`

	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	MemberClusters []*string `json:"memberClusters,omitempty" tf:"member_clusters,omitempty"`

	PrimaryEndpointAddress *string `json:"primaryEndpointAddress,omitempty" tf:"primary_endpoint_address,omitempty"`

	ReaderEndpointAddress *string `json:"readerEndpointAddress,omitempty" tf:"reader_endpoint_address,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReplicationGroupParameters struct {

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// +kubebuilder:validation:Optional
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AuthTokenSecretRef v1.SecretKeySelector `json:"authTokenSecretRef,omitempty" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticFailoverEnabled *bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterMode []ClusterModeParameters `json:"clusterMode,omitempty" tf:"cluster_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalReplicationGroupID *string `json:"globalReplicationGroupId,omitempty" tf:"global_replication_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	MultiAzEnabled *bool `json:"multiAzEnabled,omitempty" tf:"multi_az_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`

	// +kubebuilder:validation:Optional
	NumberCacheClusters *int64 `json:"numberCacheClusters,omitempty" tf:"number_cache_clusters,omitempty"`

	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Required
	ReplicationGroupDescription *string `json:"replicationGroupDescription" tf:"replication_group_description,omitempty"`

	// +kubebuilder:validation:Required
	ReplicationGroupID *string `json:"replicationGroupId" tf:"replication_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
}

// ReplicationGroupSpec defines the desired state of ReplicationGroup
type ReplicationGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationGroupParameters `json:"forProvider"`
}

// ReplicationGroupStatus defines the observed state of ReplicationGroup.
type ReplicationGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationGroup is the Schema for the ReplicationGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationGroupSpec   `json:"spec"`
	Status            ReplicationGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationGroupList contains a list of ReplicationGroups
type ReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationGroup `json:"items"`
}

// Repository type metadata.
var (
	ReplicationGroupKind             = "ReplicationGroup"
	ReplicationGroupGroupKind        = schema.GroupKind{Group: Group, Kind: ReplicationGroupKind}.String()
	ReplicationGroupKindAPIVersion   = ReplicationGroupKind + "." + GroupVersion.String()
	ReplicationGroupGroupVersionKind = GroupVersion.WithKind(ReplicationGroupKind)
)

func init() {
	SchemeBuilder.Register(&ReplicationGroup{}, &ReplicationGroupList{})
}
