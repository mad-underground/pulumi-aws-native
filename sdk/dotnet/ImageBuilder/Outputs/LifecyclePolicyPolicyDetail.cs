// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// The policy detail of the lifecycle policy.
    /// </summary>
    [OutputType]
    public sealed class LifecyclePolicyPolicyDetail
    {
        public readonly Outputs.LifecyclePolicyAction Action;
        public readonly Outputs.LifecyclePolicyExclusionRules? ExclusionRules;
        public readonly Outputs.LifecyclePolicyFilter Filter;

        [OutputConstructor]
        private LifecyclePolicyPolicyDetail(
            Outputs.LifecyclePolicyAction action,

            Outputs.LifecyclePolicyExclusionRules? exclusionRules,

            Outputs.LifecyclePolicyFilter filter)
        {
            Action = action;
            ExclusionRules = exclusionRules;
            Filter = filter;
        }
    }
}
