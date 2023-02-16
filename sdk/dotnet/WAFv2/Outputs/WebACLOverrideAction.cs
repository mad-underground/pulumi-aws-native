// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Override a RuleGroup or ManagedRuleGroup behavior. This can only be applied to Rule that has RuleGroupReferenceStatement or ManagedRuleGroupReferenceStatement.
    /// </summary>
    [OutputType]
    public sealed class WebACLOverrideAction
    {
        /// <summary>
        /// Count traffic towards application.
        /// </summary>
        public readonly object? Count;
        /// <summary>
        /// Keep the RuleGroup or ManagedRuleGroup behavior as is.
        /// </summary>
        public readonly object? None;

        [OutputConstructor]
        private WebACLOverrideAction(
            object? count,

            object? none)
        {
            Count = count;
            None = none;
        }
    }
}
