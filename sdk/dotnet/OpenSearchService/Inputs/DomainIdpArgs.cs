// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Inputs
{

    public sealed class DomainIdpArgs : global::Pulumi.ResourceArgs
    {
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        [Input("metadataContent", required: true)]
        public Input<string> MetadataContent { get; set; } = null!;

        public DomainIdpArgs()
        {
        }
        public static new DomainIdpArgs Empty => new DomainIdpArgs();
    }
}
