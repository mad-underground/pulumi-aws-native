// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketAccessControlTranslationArgs : global::Pulumi.ResourceArgs
    {
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        public BucketAccessControlTranslationArgs()
        {
        }
        public static new BucketAccessControlTranslationArgs Empty => new BucketAccessControlTranslationArgs();
    }
}
