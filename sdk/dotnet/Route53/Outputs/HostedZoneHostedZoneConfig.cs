// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53.Outputs
{

    [OutputType]
    public sealed class HostedZoneHostedZoneConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
        /// </summary>
        public readonly string? Comment;

        [OutputConstructor]
        private HostedZoneHostedZoneConfig(string? comment)
        {
            Comment = comment;
        }
    }
}
