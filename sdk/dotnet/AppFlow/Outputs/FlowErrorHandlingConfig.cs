// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowErrorHandlingConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html#cfn-appflow-flow-errorhandlingconfig-bucketname
        /// </summary>
        public readonly string? BucketName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html#cfn-appflow-flow-errorhandlingconfig-bucketprefix
        /// </summary>
        public readonly string? BucketPrefix;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-errorhandlingconfig.html#cfn-appflow-flow-errorhandlingconfig-failonfirsterror
        /// </summary>
        public readonly bool? FailOnFirstError;

        [OutputConstructor]
        private FlowErrorHandlingConfig(
            string? BucketName,

            string? BucketPrefix,

            bool? FailOnFirstError)
        {
            this.BucketName = BucketName;
            this.BucketPrefix = BucketPrefix;
            this.FailOnFirstError = FailOnFirstError;
        }
    }
}
