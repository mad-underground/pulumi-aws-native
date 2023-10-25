// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Outputs
{

    [OutputType]
    public sealed class BucketRoutingRule
    {
        public readonly Outputs.BucketRedirectRule RedirectRule;
        public readonly Outputs.BucketRoutingRuleCondition? RoutingRuleCondition;

        [OutputConstructor]
        private BucketRoutingRule(
            Outputs.BucketRedirectRule redirectRule,

            Outputs.BucketRoutingRuleCondition? routingRuleCondition)
        {
            RedirectRule = redirectRule;
            RoutingRuleCondition = routingRuleCondition;
        }
    }
}
