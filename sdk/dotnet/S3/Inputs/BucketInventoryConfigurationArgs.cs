// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketInventoryConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<Inputs.BucketDestinationArgs> Destination { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("includedObjectVersions", required: true)]
        public Input<string> IncludedObjectVersions { get; set; } = null!;

        [Input("optionalFields")]
        private InputList<string>? _optionalFields;
        public InputList<string> OptionalFields
        {
            get => _optionalFields ?? (_optionalFields = new InputList<string>());
            set => _optionalFields = value;
        }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("scheduleFrequency", required: true)]
        public Input<string> ScheduleFrequency { get; set; } = null!;

        public BucketInventoryConfigurationArgs()
        {
        }
        public static new BucketInventoryConfigurationArgs Empty => new BucketInventoryConfigurationArgs();
    }
}
