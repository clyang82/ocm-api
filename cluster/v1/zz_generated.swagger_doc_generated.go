package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ClientConfig = map[string]string{
	"":         "ClientConfig represents the apiserver address of the managed cluster.",
	"url":      "URL is the url of apiserver endpoint of the managed cluster.",
	"caBundle": "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
}

func (ClientConfig) SwaggerDoc() map[string]string {
	return map_ClientConfig
}

var map_ManagedCluster = map[string]string{
	"":       "ManagedCluster represents the desired state and current status of managed cluster. ManagedCluster is a cluster scoped resource. The name is the cluster UID.\n\nThe cluster join process follows a double opt-in process:\n\n1. agent on managed cluster creates CSR on hub with cluster UID and agent name. 2. agent on managed cluster creates ManagedCluster on hub. 3. cluster admin on hub approves the CSR for the ManagedCluster's UID and agent name. 4. cluster admin sets spec.acceptClient of ManagedCluster to true. 5. cluster admin on managed cluster creates credential of kubeconfig to hub.\n\nOnce the hub creates the cluster namespace, the Klusterlet agent on the Managed Cluster pushes the credential to the hub to use against the managed cluster's kube-apiserver.",
	"spec":   "Spec represents a desired configuration for the agent on the managed cluster.",
	"status": "Status represents the current status of joined managed cluster",
}

func (ManagedCluster) SwaggerDoc() map[string]string {
	return map_ManagedCluster
}

var map_ManagedClusterList = map[string]string{
	"":         "ManagedClusterList is a collection of managed cluster.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of managed cluster.",
}

func (ManagedClusterList) SwaggerDoc() map[string]string {
	return map_ManagedClusterList
}

var map_ManagedClusterSpec = map[string]string{
	"":                            "ManagedClusterSpec provides the information to securely connect to a remote server and verify its identity.",
	"managedClusterClientConfigs": "ManagedClusterClientConfigs represents a list of the apiserver address of the managed cluster. If it is empty, managed cluster has no accessible address to be visited from hub.",
	"hubAcceptsClient":            "hubAcceptsClient represents that hub accepts the join of Klusterlet agent on the managed cluster to the hub. The default value is false, and can only be set true when the user on hub has an RBAC rule to UPDATE on the virtual subresource of managedclusters/accept. When the value is set true, a namespace whose name is same as the name of ManagedCluster is created on hub representing the managed cluster, also role/rolebinding is created on the namespace to grant the permision of access from agent on managed cluster. When the value is set false, the namespace representing the managed cluster is deleted.",
	"leaseDurationSeconds":        "LeaseDurationSeconds is used to coordinate the lease update time of Klusterlet agents on the managed cluster. If its value is zero, the Klusterlet agent will update its lease every 60s by default",
}

func (ManagedClusterSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterSpec
}

var map_ManagedClusterStatus = map[string]string{
	"":            "ManagedClusterStatus represents the current status of joined managed cluster.",
	"conditions":  "Conditions contains the different condition statuses for this managed cluster.",
	"capacity":    "Capacity represents the total resource capacity from all nodeStatuses on the managed cluster.",
	"allocatable": "Allocatable represents the total allocatable resources on the managed cluster.",
	"version":     "Version represents the kubernetes version of the managed cluster.",
}

func (ManagedClusterStatus) SwaggerDoc() map[string]string {
	return map_ManagedClusterStatus
}

var map_ManagedClusterVersion = map[string]string{
	"":           "ManagedClusterVersion represents version information about the managed cluster.",
	"kubernetes": "Kubernetes is the kubernetes version of managed cluster.",
}

func (ManagedClusterVersion) SwaggerDoc() map[string]string {
	return map_ManagedClusterVersion
}

var map_StatusCondition = map[string]string{
	"":                   "StatusCondition contains condition information for a managed cluster.",
	"type":               "Type is the type of the cluster condition.",
	"status":             "Status is the status of the condition. One of True, False, Unknown.",
	"lastTransitionTime": "LastTransitionTime is the last time the condition changed from one status to another.",
	"reason":             "Reason is a (brief) reason for the condition's last status change.",
	"message":            "Message is a human-readable message indicating details about the last status change.",
}

func (StatusCondition) SwaggerDoc() map[string]string {
	return map_StatusCondition
}

// AUTO-GENERATED FUNCTIONS END HERE
