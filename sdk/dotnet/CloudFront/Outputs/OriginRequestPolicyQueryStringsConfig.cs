// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class OriginRequestPolicyQueryStringsConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-querystringsconfig.html#cfn-cloudfront-originrequestpolicy-querystringsconfig-querystringbehavior
        /// </summary>
        public readonly string QueryStringBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-querystringsconfig.html#cfn-cloudfront-originrequestpolicy-querystringsconfig-querystrings
        /// </summary>
        public readonly ImmutableArray<string> QueryStrings;

        [OutputConstructor]
        private OriginRequestPolicyQueryStringsConfig(
            string queryStringBehavior,

            ImmutableArray<string> queryStrings)
        {
            QueryStringBehavior = queryStringBehavior;
            QueryStrings = queryStrings;
        }
    }
}
