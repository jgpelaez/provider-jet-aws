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

type CatalogDataObservation struct {
}

type CatalogDataParameters struct {

	// +kubebuilder:validation:Optional
	AboutText *string `json:"aboutText,omitempty" tf:"about_text,omitempty"`

	// +kubebuilder:validation:Optional
	Architectures []*string `json:"architectures,omitempty" tf:"architectures,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LogoImageBlob *string `json:"logoImageBlob,omitempty" tf:"logo_image_blob,omitempty"`

	// +kubebuilder:validation:Optional
	OperatingSystems []*string `json:"operatingSystems,omitempty" tf:"operating_systems,omitempty"`

	// +kubebuilder:validation:Optional
	UsageText *string `json:"usageText,omitempty" tf:"usage_text,omitempty"`
}

type PublicRepositoryObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	RepositoryURI *string `json:"repositoryUri,omitempty" tf:"repository_uri,omitempty"`
}

type PublicRepositoryParameters struct {

	// +kubebuilder:validation:Optional
	CatalogData []CatalogDataParameters `json:"catalogData,omitempty" tf:"catalog_data,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PublicRepositorySpec defines the desired state of PublicRepository
type PublicRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicRepositoryParameters `json:"forProvider"`
}

// PublicRepositoryStatus defines the observed state of PublicRepository.
type PublicRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PublicRepository is the Schema for the PublicRepositorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type PublicRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicRepositorySpec   `json:"spec"`
	Status            PublicRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicRepositoryList contains a list of PublicRepositorys
type PublicRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicRepository `json:"items"`
}

// Repository type metadata.
var (
	PublicRepositoryKind             = "PublicRepository"
	PublicRepositoryGroupKind        = schema.GroupKind{Group: Group, Kind: PublicRepositoryKind}.String()
	PublicRepositoryKindAPIVersion   = PublicRepositoryKind + "." + GroupVersion.String()
	PublicRepositoryGroupVersionKind = GroupVersion.WithKind(PublicRepositoryKind)
)

func init() {
	SchemeBuilder.Register(&PublicRepository{}, &PublicRepositoryList{})
}
