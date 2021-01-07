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
    public sealed class RuleGroupTargetTypes
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-targettypes.html#cfn-networkfirewall-rulegroup-targettypes-targettypes
        /// </summary>
        public readonly ImmutableArray<string> TargetTypes;

        [OutputConstructor]
        private RuleGroupTargetTypes(ImmutableArray<string> TargetTypes)
        {
            this.TargetTypes = TargetTypes;
        }
    }
}
