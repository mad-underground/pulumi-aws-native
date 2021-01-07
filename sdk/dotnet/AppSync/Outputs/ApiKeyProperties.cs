// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppSync.Outputs
{

    [OutputType]
    public sealed class ApiKeyProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-apiid
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-apikeyid
        /// </summary>
        public readonly string? ApiKeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apikey.html#cfn-appsync-apikey-expires
        /// </summary>
        public readonly double? Expires;

        [OutputConstructor]
        private ApiKeyProperties(
            string ApiId,

            string? ApiKeyId,

            string? Description,

            double? Expires)
        {
            this.ApiId = ApiId;
            this.ApiKeyId = ApiKeyId;
            this.Description = Description;
            this.Expires = Expires;
        }
    }
}