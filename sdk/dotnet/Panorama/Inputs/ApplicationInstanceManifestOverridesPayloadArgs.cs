// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Panorama.Inputs
{

    public sealed class ApplicationInstanceManifestOverridesPayloadArgs : global::Pulumi.ResourceArgs
    {
        [Input("payloadData")]
        public Input<string>? PayloadData { get; set; }

        public ApplicationInstanceManifestOverridesPayloadArgs()
        {
        }
        public static new ApplicationInstanceManifestOverridesPayloadArgs Empty => new ApplicationInstanceManifestOverridesPayloadArgs();
    }
}
