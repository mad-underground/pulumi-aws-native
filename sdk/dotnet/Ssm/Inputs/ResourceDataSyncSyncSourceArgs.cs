// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ssm.Inputs
{

    public sealed class ResourceDataSyncSyncSourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("awsOrganizationsSource")]
        public Input<Inputs.ResourceDataSyncAwsOrganizationsSourceArgs>? AwsOrganizationsSource { get; set; }

        [Input("includeFutureRegions")]
        public Input<bool>? IncludeFutureRegions { get; set; }

        [Input("sourceRegions", required: true)]
        private InputList<string>? _sourceRegions;
        public InputList<string> SourceRegions
        {
            get => _sourceRegions ?? (_sourceRegions = new InputList<string>());
            set => _sourceRegions = value;
        }

        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        public ResourceDataSyncSyncSourceArgs()
        {
        }
        public static new ResourceDataSyncSyncSourceArgs Empty => new ResourceDataSyncSyncSourceArgs();
    }
}
