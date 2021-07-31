// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-carriergateway-tags.html
    /// </summary>
    [OutputType]
    public sealed class CarrierGatewayTags
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-carriergateway-tags.html#cfn-ec2-carriergateway-tags-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private CarrierGatewayTags(ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Tags = tags;
        }
    }
}
