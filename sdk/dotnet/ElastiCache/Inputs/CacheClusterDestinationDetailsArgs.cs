// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Inputs
{

    public sealed class CacheClusterDestinationDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudWatchLogsDetails")]
        public Input<Inputs.CacheClusterCloudWatchLogsDestinationDetailsArgs>? CloudWatchLogsDetails { get; set; }

        [Input("kinesisFirehoseDetails")]
        public Input<Inputs.CacheClusterKinesisFirehoseDestinationDetailsArgs>? KinesisFirehoseDetails { get; set; }

        public CacheClusterDestinationDetailsArgs()
        {
        }
        public static new CacheClusterDestinationDetailsArgs Empty => new CacheClusterDestinationDetailsArgs();
    }
}
