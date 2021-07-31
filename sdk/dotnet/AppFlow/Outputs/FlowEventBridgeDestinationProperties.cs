// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-eventbridgedestinationproperties.html
    /// </summary>
    [OutputType]
    public sealed class FlowEventBridgeDestinationProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-eventbridgedestinationproperties.html#cfn-appflow-flow-eventbridgedestinationproperties-errorhandlingconfig
        /// </summary>
        public readonly Outputs.FlowErrorHandlingConfig? ErrorHandlingConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-eventbridgedestinationproperties.html#cfn-appflow-flow-eventbridgedestinationproperties-object
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private FlowEventBridgeDestinationProperties(
            Outputs.FlowErrorHandlingConfig? errorHandlingConfig,

            string @object)
        {
            ErrorHandlingConfig = errorHandlingConfig;
            Object = @object;
        }
    }
}
