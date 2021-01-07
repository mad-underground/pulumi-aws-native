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
    public sealed class StageProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-accesslogsetting
        /// </summary>
        public readonly Outputs.StageAccessLogSetting? AccessLogSetting;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclusterenabled
        /// </summary>
        public readonly bool? CacheClusterEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclustersize
        /// </summary>
        public readonly string? CacheClusterSize;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-canarysetting
        /// </summary>
        public readonly Outputs.StageCanarySetting? CanarySetting;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-clientcertificateid
        /// </summary>
        public readonly string? ClientCertificateId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-deploymentid
        /// </summary>
        public readonly string? DeploymentId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-documentationversion
        /// </summary>
        public readonly string? DocumentationVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-methodsettings
        /// </summary>
        public readonly ImmutableArray<Outputs.StageMethodSetting> MethodSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-restapiid
        /// </summary>
        public readonly string RestApiId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-stagename
        /// </summary>
        public readonly string? StageName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-tracingenabled
        /// </summary>
        public readonly bool? TracingEnabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-variables
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Variables;

        [OutputConstructor]
        private StageProperties(
            Outputs.StageAccessLogSetting? AccessLogSetting,

            bool? CacheClusterEnabled,

            string? CacheClusterSize,

            Outputs.StageCanarySetting? CanarySetting,

            string? ClientCertificateId,

            string? DeploymentId,

            string? Description,

            string? DocumentationVersion,

            ImmutableArray<Outputs.StageMethodSetting> MethodSettings,

            string RestApiId,

            string? StageName,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            bool? TracingEnabled,

            ImmutableDictionary<string, string>? Variables)
        {
            this.AccessLogSetting = AccessLogSetting;
            this.CacheClusterEnabled = CacheClusterEnabled;
            this.CacheClusterSize = CacheClusterSize;
            this.CanarySetting = CanarySetting;
            this.ClientCertificateId = ClientCertificateId;
            this.DeploymentId = DeploymentId;
            this.Description = Description;
            this.DocumentationVersion = DocumentationVersion;
            this.MethodSettings = MethodSettings;
            this.RestApiId = RestApiId;
            this.StageName = StageName;
            this.Tags = Tags;
            this.TracingEnabled = TracingEnabled;
            this.Variables = Variables;
        }
    }
}