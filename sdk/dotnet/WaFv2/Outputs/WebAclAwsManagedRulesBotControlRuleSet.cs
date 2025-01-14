// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Configures how to use the Bot Control managed rule group in the web ACL
    /// </summary>
    [OutputType]
    public sealed class WebAclAwsManagedRulesBotControlRuleSet
    {
        public readonly bool? EnableMachineLearning;
        public readonly Pulumi.AwsNative.WaFv2.WebAclAwsManagedRulesBotControlRuleSetInspectionLevel InspectionLevel;

        [OutputConstructor]
        private WebAclAwsManagedRulesBotControlRuleSet(
            bool? enableMachineLearning,

            Pulumi.AwsNative.WaFv2.WebAclAwsManagedRulesBotControlRuleSetInspectionLevel inspectionLevel)
        {
            EnableMachineLearning = enableMachineLearning;
            InspectionLevel = inspectionLevel;
        }
    }
}
