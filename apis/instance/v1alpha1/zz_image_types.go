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

type AdditionalVolumesObservation struct {

	// Date of the volume creation.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// The export URI of the volume.
	ExportURI *string `json:"exportUri,omitempty" tf:"export_uri,omitempty"`

	// The ID of the volume.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date of volume latest update.
	ModificationDate *string `json:"modificationDate,omitempty" tf:"modification_date,omitempty"`

	// The name of the volume.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID the volume is associated with.
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// ID of the project the volume is associated with
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Description of the server containing the volume (in case the image is a backup from a server).
	Server map[string]*string `json:"server,omitempty" tf:"server,omitempty"`

	// The size of the volume.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// State of the volume.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// List of tags associated with the volume.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of volume, possible values are l_ssd and b_ssd.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`

	// The zone in which the volume is.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type AdditionalVolumesParameters struct {
}

type ImageObservation struct {

	// The description of the extra volumes attached to the image.
	// Specs of the additional volumes attached to the image
	AdditionalVolumes []AdditionalVolumesObservation `json:"additionalVolumes,omitempty" tf:"additional_volumes,omitempty"`

	// Date of the image creation.
	// The date and time of the creation of the image
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// ID of the server the image is based on (in case it is a backup).
	// The ID of the backed-up server from which the snapshot was taken
	FromServerID *string `json:"fromServerId,omitempty" tf:"from_server_id,omitempty"`

	// The ID of the image.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date of image latest update.
	// The date and time of the last modification of the Redis cluster
	ModificationDate *string `json:"modificationDate,omitempty" tf:"modification_date,omitempty"`

	// The organization ID the image is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// State of the image. Possible values are: available, creating or error.
	// The state of the image [ available | creating | error ]
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type ImageParameters struct {

	// List of IDs of the snapshots of the additional volumes to be attached to the image.
	// The IDs of the additional volumes attached to the image
	// +kubebuilder:validation:Optional
	AdditionalVolumeIds []*string `json:"additionalVolumeIds,omitempty" tf:"additional_volume_ids,omitempty"`

	// The architecture the image is compatible with. Possible values are: x86_64 or arm.
	// Architecture of the image (default = x86_64)
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// The name of the image. If not provided it will be randomly generated.
	// The name of the image
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the image is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Set to true if the image is public.
	// If true, the image will be public
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// The ID of the snapshot of the volume to be used as root in the image.
	// UUID of the snapshot from which the image is to be created
	// +crossplane:generate:reference:type=Snapshot
	// +kubebuilder:validation:Optional
	RootVolumeID *string `json:"rootVolumeId,omitempty" tf:"root_volume_id,omitempty"`

	// Reference to a Snapshot to populate rootVolumeId.
	// +kubebuilder:validation:Optional
	RootVolumeIDRef *v1.Reference `json:"rootVolumeIdRef,omitempty" tf:"-"`

	// Selector for a Snapshot to populate rootVolumeId.
	// +kubebuilder:validation:Optional
	RootVolumeIDSelector *v1.Selector `json:"rootVolumeIdSelector,omitempty" tf:"-"`

	// A list of tags to apply to the image.
	// List of tags ["tag1", "tag2", ...] attached to the image
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the image should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Image is the Schema for the Images API. Manages Scaleway Instance Images.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec"`
	Status            ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
