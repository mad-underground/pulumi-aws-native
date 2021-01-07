// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFront.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html
    /// </summary>
    public sealed class CloudFrontOriginAccessIdentityPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
        /// </summary>
        [Input("CloudFrontOriginAccessIdentityConfig", required: true)]
        public Input<Inputs.CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs> CloudFrontOriginAccessIdentityConfig { get; set; } = null!;

        public CloudFrontOriginAccessIdentityPropertiesArgs()
        {
        }
    }
}