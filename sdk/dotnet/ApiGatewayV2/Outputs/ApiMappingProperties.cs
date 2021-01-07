// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class ApiMappingProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apiid
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-apimappingkey
        /// </summary>
        public readonly string? ApiMappingKey;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-domainname
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html#cfn-apigatewayv2-apimapping-stage
        /// </summary>
        public readonly string Stage;

        [OutputConstructor]
        private ApiMappingProperties(
            string ApiId,

            string? ApiMappingKey,

            string DomainName,

            string Stage)
        {
            this.ApiId = ApiId;
            this.ApiMappingKey = ApiMappingKey;
            this.DomainName = DomainName;
            this.Stage = Stage;
        }
    }
}