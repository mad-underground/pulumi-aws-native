// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.QLDB.Outputs
{

    [OutputType]
    public sealed class StreamKinesisConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html#cfn-qldb-stream-kinesisconfiguration-aggregationenabled
        /// </summary>
        public readonly bool? AggregationEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qldb-stream-kinesisconfiguration.html#cfn-qldb-stream-kinesisconfiguration-streamarn
        /// </summary>
        public readonly string? StreamArn;

        [OutputConstructor]
        private StreamKinesisConfiguration(
            bool? AggregationEnabled,

            string? StreamArn)
        {
            this.AggregationEnabled = AggregationEnabled;
            this.StreamArn = StreamArn;
        }
    }
}
