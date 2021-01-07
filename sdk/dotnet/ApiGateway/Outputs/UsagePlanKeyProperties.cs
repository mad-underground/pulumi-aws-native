// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApiGateway.Outputs
{

    [OutputType]
    public sealed class UsagePlanKeyProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keyid
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keytype
        /// </summary>
        public readonly string KeyType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-usageplanid
        /// </summary>
        public readonly string UsagePlanId;

        [OutputConstructor]
        private UsagePlanKeyProperties(
            string KeyId,

            string KeyType,

            string UsagePlanId)
        {
            this.KeyId = KeyId;
            this.KeyType = KeyType;
            this.UsagePlanId = UsagePlanId;
        }
    }
}