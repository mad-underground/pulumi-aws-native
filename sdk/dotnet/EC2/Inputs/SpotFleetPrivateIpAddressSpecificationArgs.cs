// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-privateipaddressspecification.html
    /// </summary>
    public sealed class SpotFleetPrivateIpAddressSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-privateipaddressspecification.html#cfn-ec2-spotfleet-privateipaddressspecification-primary
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-privateipaddressspecification.html#cfn-ec2-spotfleet-privateipaddressspecification-privateipaddress
        /// </summary>
        [Input("privateIpAddress", required: true)]
        public Input<string> PrivateIpAddress { get; set; } = null!;

        public SpotFleetPrivateIpAddressSpecificationArgs()
        {
        }
    }
}