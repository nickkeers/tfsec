package gke
// 
// // ATTENTION!
// // This rule was autogenerated!
// // Before making changes, consider updating the generator.
// 
// import (
// 	"github.com/aquasecurity/defsec/provider"
// 	"github.com/aquasecurity/defsec/result"
// 	"github.com/aquasecurity/defsec/severity"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
// 	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
// 	"github.com/aquasecurity/tfsec/pkg/rule"
// )
// 
// func init() {
// 	scanner.RegisterCheckRule(rule.Rule{
// 		Provider:  provider.GoogleProvider,
// 		Service:   "gke",
// 		ShortCode: "enable-auto-repair",
// 		Documentation: rule.RuleDocumentation{
// 			Summary:     "Kubernetes should have 'Automatic repair' enabled",
// 			Explanation: `Automatic repair will monitor nodes and attempt repair when a node fails multiple subsequent health checks`,
// 			Impact:      "Failing nodes will require manual repair.",
// 			Resolution:  "Enable automatic repair",
// 			BadExample: []string{`
// resource "google_service_account" "default" {
//   account_id   = "service-account-id"
//   display_name = "Service Account"
// }
// 
// resource "google_container_cluster" "primary" {
//   name     = "my-gke-cluster"
//   location = "us-central1"
// 
//   # We can't create a cluster with no node pool defined, but we want to only use
//   # separately managed node pools. So we create the smallest possible default
//   # node pool and immediately delete it.
//   remove_default_node_pool = true
//   initial_node_count       = 1
// }
// 
// resource "google_container_node_pool" "bad_example" {
//   name       = "my-node-pool"
//   cluster    = google_container_cluster.primary.id
//   node_count = 1
// 
//   node_config {
//     preemptible  = true
//     machine_type = "e2-medium"
// 
//     # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
//     service_account = google_service_account.default.email
//     oauth_scopes = [
//       "https://www.googleapis.com/auth/cloud-platform"
//     ]
//   }
//   management {
//     auto_repair = false
//   }
// }
// `},
// 			GoodExample: []string{`
// resource "google_service_account" "default" {
//   account_id   = "service-account-id"
//   display_name = "Service Account"
// }
// 
// resource "google_container_cluster" "primary" {
//   name     = "my-gke-cluster"
//   location = "us-central1"
// 
//   # We can't create a cluster with no node pool defined, but we want to only use
//   # separately managed node pools. So we create the smallest possible default
//   # node pool and immediately delete it.
//   remove_default_node_pool = true
//   initial_node_count       = 1
// }
// 
// resource "google_container_node_pool" "good_example" {
//   name       = "my-node-pool"
//   cluster    = google_container_cluster.primary.id
//   node_count = 1
// 
//   node_config {
//     preemptible  = true
//     machine_type = "e2-medium"
// 
//     # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
//     service_account = google_service_account.default.email
//     oauth_scopes = [
//       "https://www.googleapis.com/auth/cloud-platform"
//     ]
//   }
//   management {
//     auto_repair = true
//   }
// }
// `},
// 			Links: []string{
// 				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_node_pool#auto_repair",
// 			},
// 		},
// 		RequiredTypes: []string{
// 			"resource",
// 		},
// 		RequiredLabels: []string{
// 			"google_container_node_pool",
// 		},
// 		DefaultSeverity: severity.Low,
// 		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
// 			if autoRepairAttr := resourceBlock.GetBlock("management").GetAttribute("auto_repair"); autoRepairAttr.IsNil() { // alert on use of default value
// 				set.AddResult().
// 					WithDescription("Resource '%s' uses default value for management.auto_repair", resourceBlock.FullName())
// 			} else if autoRepairAttr.IsFalse() {
// 				set.AddResult().
// 					WithDescription("Resource '%s' does not have management.auto_repair set to true", resourceBlock.FullName()).
// 					WithAttribute("")
// 			}
// 		},
// 	})
// }