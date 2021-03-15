// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-mssencryption.html
    /// </summary>
    public sealed class OriginEndpointMssEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-mssencryption.html#cfn-mediapackage-originendpoint-mssencryption-spekekeyprovider
        /// </summary>
        [Input("spekeKeyProvider", required: true)]
        public Input<Inputs.OriginEndpointSpekeKeyProviderArgs> SpekeKeyProvider { get; set; } = null!;

        public OriginEndpointMssEncryptionArgs()
        {
        }
    }
}
