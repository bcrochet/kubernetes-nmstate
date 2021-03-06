package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nmstate/kubernetes-nmstate/pkg/apis/nmstate/shared"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkState is the Schema for the nodenetworkstates API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodenetworkstates,shortName=nns,scope=Cluster
type NodeNetworkState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status shared.NodeNetworkStateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkStateList contains a list of NodeNetworkState
type NodeNetworkStateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetworkState `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeNetworkState{}, &NodeNetworkStateList{})
}
