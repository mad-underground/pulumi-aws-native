// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Inputs
{

    public sealed class DomainEbsOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("ebsEnabled")]
        public Input<bool>? EbsEnabled { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("throughput")]
        public Input<int>? Throughput { get; set; }

        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public DomainEbsOptionsArgs()
        {
        }
        public static new DomainEbsOptionsArgs Empty => new DomainEbsOptionsArgs();
    }
}
