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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html
    /// </summary>
    public sealed class KeyGroupPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html#cfn-cloudfront-keygroup-keygroupconfig
        /// </summary>
        [Input("KeyGroupConfig", required: true)]
        public Input<Inputs.KeyGroupKeyGroupConfigArgs> KeyGroupConfig { get; set; } = null!;

        public KeyGroupPropertiesArgs()
        {
        }
    }
}
