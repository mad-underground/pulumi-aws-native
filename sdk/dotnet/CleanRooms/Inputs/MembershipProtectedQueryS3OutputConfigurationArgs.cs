// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class MembershipProtectedQueryS3OutputConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("keyPrefix")]
        public Input<string>? KeyPrefix { get; set; }

        [Input("resultFormat", required: true)]
        public Input<Pulumi.AwsNative.CleanRooms.MembershipResultFormat> ResultFormat { get; set; } = null!;

        public MembershipProtectedQueryS3OutputConfigurationArgs()
        {
        }
        public static new MembershipProtectedQueryS3OutputConfigurationArgs Empty => new MembershipProtectedQueryS3OutputConfigurationArgs();
    }
}
