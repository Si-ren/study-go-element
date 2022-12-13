/*
Copyright 2022.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RouteSpec struct {
	// Domain for the route
	// +kubebuilder:validation:MinLength=1
	Domain string `json:"domain"`
	// The path for the route. Defaults to /.
	// +optional
	Path string `json:"path"`
}

type ServiceSpec struct {
	//+kubebuilder:validation:default=80;optional
	SourcePort int `json:"sourcePort"`
	//+kubebuilder:validation:default=80;optional
	TargetPort int `json:"targetPort"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type (
	// AppServiceSpec defines the desired state of AppService
	AppServiceSpec struct {
		// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
		// Important: Run "make" to regenerate code after modifying this file

		// Foo is an example field of AppService. Edit appservice_types.go to remove/update
		//Foo string `json:"foo,omitempty"`
		//Image runtime image to use.
		// +kubebuilder:validation:Required
		Image string `json:"image"`
		//ImagePullPolicy overrides AppServiceRuntime spec.imagePullPolicy
		//+kubebuilder:validation:Enum=Always;IfNotPresent;Never;Optional
		ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
		// Service if Pod type is Cronjob,this is unnecessary,
		//default service port is 80,target port is 80
		Service ServiceSpec `json:"service"`
		// Ingress Default port is
		Ingress bool `json:"ingress"`
		// Routes for which the ingress is created
		// The first item is set the WP_HOME and WP_SITEURL constants.
		// If no routes are specified, ingress syncing is disabled and WP_HOME de defaults to NAME.NAMESPACE.svc.
		// +optional
		Routes []RouteSpec `json:"routes"`
	}
)

type AppServiceConditionType string

type AppServiceCondition struct {
	// Type of AppService condition.
	Type AppServiceConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// The reason for the condition's last transition.
	Reason string `json:"reason"`
	// A human readable message indicating details about the transition.
	Message string `json:"message"`
}

const (
	// DeploymentReason signals that health of deployment.
	DeploymentReason AppServiceConditionType = "DeploymentError"

	// ServiceErrorReason is the generic reason for Service failures.
	ServiceErrorReason = "ServiceError"

	// IngressReason is the reason for successfully Ingress err.
	IngressReason = "IngressError"
)

// AppServiceStatus defines the observed state of AppService
type AppServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Conditions represents the AppService resource conditions list.
	// +optional
	Conditions []AppServiceCondition `json:"conditions,omitempty"`
}

// AppService is the Schema for the appservices API
//+kubebuilder:object:root=true
// +kubebuilder:resource:shortName=as
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="image",type="string",JSONPath=".spec.image",description="AppService image"
// +kubebuilder:printcolumn:name="statuss",type="string",JSONPath=".status.conditions[?(@.type == 'DeploymentError')].status",description="app status"

type AppService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppServiceSpec   `json:"spec,omitempty"`
	Status AppServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppServiceList contains a list of AppService
type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppService{}, &AppServiceList{})
}
