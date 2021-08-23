package dns
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
// 		Service:   "dns",
// 		ShortCode: "enable-dnssec",
// 		Documentation: rule.RuleDocumentation{
// 			Summary:     "Cloud DNS should use DNSSEC",
// 			Explanation: `DNSSEC authenticates DNS responses, preventing MITM attacks and impersonation.`,
// 			Impact:      "Unverified DNS responses could lead to man-in-the-middle attacks",
// 			Resolution:  "Enable DNSSEC",
// 			BadExample: []string{`
// resource "google_dns_managed_zone" "bad_example" {
//   name        = "example-zone"
//   dns_name    = "example-${random_id.rnd.hex}.com."
//   description = "Example DNS zone"
//   labels = {
//     foo = "bar"
//   }
//   dnssec_config {
//     state = "off"
//   }
// }
// 
// resource "random_id" "rnd" {
//   byte_length = 4
// }
// `},
// 			GoodExample: []string{`
// resource "google_dns_managed_zone" "good_example" {
//   name        = "example-zone"
//   dns_name    = "example-${random_id.rnd.hex}.com."
//   description = "Example DNS zone"
//   labels = {
//     foo = "bar"
//   }
//   dnssec_config {
//     state = "on"
//   }
// }
// 
// resource "random_id" "rnd" {
//   byte_length = 4
// }
// `},
// 			Links: []string{
// 				"https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/dns_managed_zone#state",
// 			},
// 		},
// 		RequiredTypes: []string{
// 			"resource",
// 		},
// 		RequiredLabels: []string{
// 			"google_dns_managed_zone",
// 		},
// 		DefaultSeverity: severity.Medium,
// 		CheckTerraform: func(set result.Set, resourceBlock block.Block, _ block.Module) {
// 			if stateAttr := resourceBlock.GetBlock("dnssec_config").GetAttribute("state"); stateAttr.IsNil() { // alert on use of default value
// 				set.AddResult().
// 					WithDescription("Resource '%s' uses default value for dnssec_config.state", resourceBlock.FullName())
// 			} else if stateAttr.NotEqual("on") {
// 				set.AddResult().
// 					WithDescription("Resource '%s' does not have dnssec_config.state set to on", resourceBlock.FullName()).
// 					WithAttribute("")
// 			}
// 		},
// 	})
// }