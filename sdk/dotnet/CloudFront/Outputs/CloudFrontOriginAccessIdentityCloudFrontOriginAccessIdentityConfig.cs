// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFront.Outputs
{

    [OutputType]
    public sealed class CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig-comment
        /// </summary>
        public readonly string Comment;

        [OutputConstructor]
        private CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig(string Comment)
        {
            this.Comment = Comment;
        }
    }
}