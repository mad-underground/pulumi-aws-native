// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Elasticsearch.Inputs
{

    public sealed class DomainLogPublishingOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudWatchLogsLogGroupArn")]
        public Input<string>? CloudWatchLogsLogGroupArn { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public DomainLogPublishingOptionArgs()
        {
        }
        public static new DomainLogPublishingOptionArgs Empty => new DomainLogPublishingOptionArgs();
    }
}