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
    public sealed class KeyGroupProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html#cfn-cloudfront-keygroup-keygroupconfig
        /// </summary>
        public readonly Outputs.KeyGroupKeyGroupConfig KeyGroupConfig;

        [OutputConstructor]
        private KeyGroupProperties(Outputs.KeyGroupKeyGroupConfig KeyGroupConfig)
        {
            this.KeyGroupConfig = KeyGroupConfig;
        }
    }
}
