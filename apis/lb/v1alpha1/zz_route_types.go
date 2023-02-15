/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteParameters struct {

	// The ID of the backend to which the route is associated.
	// The backend ID destination of redirection
	// +crossplane:generate:reference:type=Backend
	// +kubebuilder:validation:Optional
	BackendID *string `json:"backendId,omitempty" tf:"backend_id,omitempty"`

	// Reference to a Backend to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDRef *v1.Reference `json:"backendIdRef,omitempty" tf:"-"`

	// Selector for a Backend to populate backendId.
	// +kubebuilder:validation:Optional
	BackendIDSelector *v1.Selector `json:"backendIdSelector,omitempty" tf:"-"`

	// :  The ID of the frontend to which the route is associated.
	// The frontend ID origin of redirection
	// +crossplane:generate:reference:type=Frontend
	// +kubebuilder:validation:Optional
	FrontendID *string `json:"frontendId,omitempty" tf:"frontend_id,omitempty"`

	// Reference to a Frontend to populate frontendId.
	// +kubebuilder:validation:Optional
	FrontendIDRef *v1.Reference `json:"frontendIdRef,omitempty" tf:"-"`

	// Selector for a Frontend to populate frontendId.
	// +kubebuilder:validation:Optional
	FrontendIDSelector *v1.Selector `json:"frontendIdSelector,omitempty" tf:"-"`

	// The SNI to match.
	// The domain to match against
	// +kubebuilder:validation:Optional
	MatchSni *string `json:"matchSni,omitempty" tf:"match_sni,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. Manages Scaleway Load-Balancer Route.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
