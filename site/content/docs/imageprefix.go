package test_data

// This is based on https://github.com/jlewi/hydros/blob/main/pkg/kustomize/fns/images/imageprefix.go

import (
	"github.com/jlewi/hydros/api/v1alpha1"
	"sigs.k8s.io/kustomize/api/types"
)

const (
	// Kind is the kind for the kustomize function.
	Kind = "ImagePrefix"
)

// ImagePrefixFn is a filter that changes the prefix of all docker images.
// Spec.ImageMappings is a list of mappings from source to destination repo. Each mapping
// specifies a prefix for the source repo that should be changed to the value of ImageMappings.Dest.
type ImagePrefixFn struct {
	// Kind is the API name.
	Kind string `yaml:"kind"`

	// APIVersion is the API version.  Must be examples.kpt.dev/v1alpha1
	APIVersion string `yaml:"apiVersion"`

	// Metadata defines instance metadata.
	Metadata v1alpha1.Metadata `yaml:"metadata"`

	// Spec defines the desired declarative configuration.
	Spec Spec `yaml:"spec"`
}

// Spec is the spec for the kustomize function.
type Spec struct {
	ImageMappings []ImageMapping `yaml:"imageMappings"`

	// FsSlice contains the FieldSpecs to locate an image field,
	// e.g. Path: "spec/myContainers[]/image"
	// This should be set to a list of FieldSpecs that match all the image fields you want to change.
	FsSlice types.FsSlice `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

// ImageMapping represents the mapping from a source
// to a destination repo.
type ImageMapping struct {
	// Src is the source repo to match
	Src string `yaml:"src"`
	// Dest is the new value for the image URI.
	Dest string `yaml:"dest"`
}
