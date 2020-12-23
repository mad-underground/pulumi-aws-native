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
    public sealed class OriginRequestPolicyOriginRequestPolicyConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-comment
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-cookiesconfig
        /// </summary>
        public readonly Outputs.OriginRequestPolicyCookiesConfig CookiesConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-headersconfig
        /// </summary>
        public readonly Outputs.OriginRequestPolicyHeadersConfig HeadersConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-originrequestpolicyconfig.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig-querystringsconfig
        /// </summary>
        public readonly Outputs.OriginRequestPolicyQueryStringsConfig QueryStringsConfig;

        [OutputConstructor]
        private OriginRequestPolicyOriginRequestPolicyConfig(
            string? Comment,

            Outputs.OriginRequestPolicyCookiesConfig CookiesConfig,

            Outputs.OriginRequestPolicyHeadersConfig HeadersConfig,

            string Name,

            Outputs.OriginRequestPolicyQueryStringsConfig QueryStringsConfig)
        {
            this.Comment = Comment;
            this.CookiesConfig = CookiesConfig;
            this.HeadersConfig = HeadersConfig;
            this.Name = Name;
            this.QueryStringsConfig = QueryStringsConfig;
        }
    }
}
