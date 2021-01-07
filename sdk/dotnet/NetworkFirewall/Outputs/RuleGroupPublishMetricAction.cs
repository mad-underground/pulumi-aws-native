// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupPublishMetricAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-publishmetricaction.html#cfn-networkfirewall-rulegroup-publishmetricaction-dimensions
        /// </summary>
        public readonly Outputs.RuleGroupDimensions Dimensions;

        [OutputConstructor]
        private RuleGroupPublishMetricAction(Outputs.RuleGroupDimensions Dimensions)
        {
            this.Dimensions = Dimensions;
        }
    }
}