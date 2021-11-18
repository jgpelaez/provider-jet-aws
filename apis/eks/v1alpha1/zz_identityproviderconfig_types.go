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

type IdentityProviderConfigObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type IdentityProviderConfigOidcObservation struct {
}

type IdentityProviderConfigOidcParameters struct {

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Optional
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// +kubebuilder:validation:Optional
	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix,omitempty"`

	// +kubebuilder:validation:Required
	IdentityProviderConfigName *string `json:"identityProviderConfigName" tf:"identity_provider_config_name,omitempty"`

	// +kubebuilder:validation:Required
	IssuerURL *string `json:"issuerUrl" tf:"issuer_url,omitempty"`

	// +kubebuilder:validation:Optional
	RequiredClaims map[string]*string `json:"requiredClaims,omitempty" tf:"required_claims,omitempty"`

	// +kubebuilder:validation:Optional
	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim,omitempty"`

	// +kubebuilder:validation:Optional
	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix,omitempty"`
}

type IdentityProviderConfigParameters struct {

	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Oidc []IdentityProviderConfigOidcParameters `json:"oidc" tf:"oidc,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// IdentityProviderConfigSpec defines the desired state of IdentityProviderConfig
type IdentityProviderConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderConfigParameters `json:"forProvider"`
}

// IdentityProviderConfigStatus defines the observed state of IdentityProviderConfig.
type IdentityProviderConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderConfig is the Schema for the IdentityProviderConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IdentityProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityProviderConfigSpec   `json:"spec"`
	Status            IdentityProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderConfigList contains a list of IdentityProviderConfigs
type IdentityProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderConfig `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderConfig_Kind             = "IdentityProviderConfig"
	IdentityProviderConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderConfig_Kind}.String()
	IdentityProviderConfig_KindAPIVersion   = IdentityProviderConfig_Kind + "." + CRDGroupVersion.String()
	IdentityProviderConfig_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderConfig{}, &IdentityProviderConfigList{})
}
