// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class FeatureGroupOnlineStoreSecurityConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        public FeatureGroupOnlineStoreSecurityConfigArgs()
        {
        }
        public static new FeatureGroupOnlineStoreSecurityConfigArgs Empty => new FeatureGroupOnlineStoreSecurityConfigArgs();
    }
}
