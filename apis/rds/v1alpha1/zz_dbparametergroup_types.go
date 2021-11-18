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

type DBParameterGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DBParameterGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Family *string `json:"family" tf:"family,omitempty"`

	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ParameterObservation struct {
}

type ParameterParameters struct {

	// +kubebuilder:validation:Optional
	ApplyMethod *string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// DBParameterGroupSpec defines the desired state of DBParameterGroup
type DBParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DBParameterGroupParameters `json:"forProvider"`
}

// DBParameterGroupStatus defines the observed state of DBParameterGroup.
type DBParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DBParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DBParameterGroup is the Schema for the DBParameterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DBParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBParameterGroupSpec   `json:"spec"`
	Status            DBParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBParameterGroupList contains a list of DBParameterGroups
type DBParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	DBParameterGroup_Kind             = "DBParameterGroup"
	DBParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DBParameterGroup_Kind}.String()
	DBParameterGroup_KindAPIVersion   = DBParameterGroup_Kind + "." + CRDGroupVersion.String()
	DBParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(DBParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&DBParameterGroup{}, &DBParameterGroupList{})
}
