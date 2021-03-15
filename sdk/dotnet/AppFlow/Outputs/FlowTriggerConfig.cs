// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowTriggerConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-triggerconfig.html#cfn-appflow-flow-triggerconfig-triggerproperties
        /// </summary>
        public readonly Outputs.FlowScheduledTriggerProperties? TriggerProperties;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-triggerconfig.html#cfn-appflow-flow-triggerconfig-triggertype
        /// </summary>
        public readonly string TriggerType;

        [OutputConstructor]
        private FlowTriggerConfig(
            Outputs.FlowScheduledTriggerProperties? triggerProperties,

            string triggerType)
        {
            TriggerProperties = triggerProperties;
            TriggerType = triggerType;
        }
    }
}
