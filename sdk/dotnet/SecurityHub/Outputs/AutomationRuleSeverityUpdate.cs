// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Outputs
{

    [OutputType]
    public sealed class AutomationRuleSeverityUpdate
    {
        public readonly Pulumi.AwsNative.SecurityHub.AutomationRuleSeverityUpdateLabel? Label;
        public readonly int? Normalized;
        public readonly double? Product;

        [OutputConstructor]
        private AutomationRuleSeverityUpdate(
            Pulumi.AwsNative.SecurityHub.AutomationRuleSeverityUpdateLabel? label,

            int? normalized,

            double? product)
        {
            Label = label;
            Normalized = normalized;
            Product = product;
        }
    }
}
