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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html
    /// </summary>
    public sealed class FlowErrorHandlingConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html#cfn-appflow-flow-errorhandlingconfig-bucketname
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html#cfn-appflow-flow-errorhandlingconfig-bucketprefix
        /// </summary>
        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html#cfn-appflow-flow-errorhandlingconfig-failonfirsterror
        /// </summary>
        [Input("failOnFirstError")]
        public Input<bool>? FailOnFirstError { get; set; }

        public FlowErrorHandlingConfigArgs()
        {
        }
    }
}
