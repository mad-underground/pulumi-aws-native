// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFront.Outputs
{

    [OutputType]
    public sealed class RealtimeLogConfigEndPoint
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html#cfn-cloudfront-realtimelogconfig-endpoint-kinesisstreamconfig
        /// </summary>
        public readonly Outputs.RealtimeLogConfigKinesisStreamConfig KinesisStreamConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html#cfn-cloudfront-realtimelogconfig-endpoint-streamtype
        /// </summary>
        public readonly string StreamType;

        [OutputConstructor]
        private RealtimeLogConfigEndPoint(
            Outputs.RealtimeLogConfigKinesisStreamConfig KinesisStreamConfig,

            string StreamType)
        {
            this.KinesisStreamConfig = KinesisStreamConfig;
            this.StreamType = StreamType;
        }
    }
}
