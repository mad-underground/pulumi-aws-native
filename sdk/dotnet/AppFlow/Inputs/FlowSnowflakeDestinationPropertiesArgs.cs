// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-snowflakedestinationproperties.html
    /// </summary>
    public sealed class FlowSnowflakeDestinationPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-snowflakedestinationproperties.html#cfn-appflow-flow-snowflakedestinationproperties-bucketprefix
        /// </summary>
        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-snowflakedestinationproperties.html#cfn-appflow-flow-snowflakedestinationproperties-errorhandlingconfig
        /// </summary>
        [Input("errorHandlingConfig")]
        public Input<Inputs.FlowErrorHandlingConfigArgs>? ErrorHandlingConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-snowflakedestinationproperties.html#cfn-appflow-flow-snowflakedestinationproperties-intermediatebucketname
        /// </summary>
        [Input("intermediateBucketName", required: true)]
        public Input<string> IntermediateBucketName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-snowflakedestinationproperties.html#cfn-appflow-flow-snowflakedestinationproperties-object
        /// </summary>
        [Input("object", required: true)]
        public Input<string> Object { get; set; } = null!;

        public FlowSnowflakeDestinationPropertiesArgs()
        {
        }
    }
}
