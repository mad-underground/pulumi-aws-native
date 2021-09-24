// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentGroupTriggerConfig
    {
        public readonly ImmutableArray<string> TriggerEvents;
        public readonly string? TriggerName;
        public readonly string? TriggerTargetArn;

        [OutputConstructor]
        private DeploymentGroupTriggerConfig(
            ImmutableArray<string> triggerEvents,

            string? triggerName,

            string? triggerTargetArn)
        {
            TriggerEvents = triggerEvents;
            TriggerName = triggerName;
            TriggerTargetArn = triggerTargetArn;
        }
    }
}