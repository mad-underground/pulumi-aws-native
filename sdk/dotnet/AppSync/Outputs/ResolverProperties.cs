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
    public sealed class ResolverProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-apiid
        /// </summary>
        public readonly string ApiId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-cachingconfig
        /// </summary>
        public readonly Outputs.ResolverCachingConfig? CachingConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-datasourcename
        /// </summary>
        public readonly string? DataSourceName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-fieldname
        /// </summary>
        public readonly string FieldName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-kind
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-pipelineconfig
        /// </summary>
        public readonly Outputs.ResolverPipelineConfig? PipelineConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplate
        /// </summary>
        public readonly string? RequestMappingTemplate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-requestmappingtemplates3location
        /// </summary>
        public readonly string? RequestMappingTemplateS3Location;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplate
        /// </summary>
        public readonly string? ResponseMappingTemplate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-responsemappingtemplates3location
        /// </summary>
        public readonly string? ResponseMappingTemplateS3Location;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-syncconfig
        /// </summary>
        public readonly Outputs.ResolverSyncConfig? SyncConfig;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html#cfn-appsync-resolver-typename
        /// </summary>
        public readonly string TypeName;

        [OutputConstructor]
        private ResolverProperties(
            string ApiId,

            Outputs.ResolverCachingConfig? CachingConfig,

            string? DataSourceName,

            string FieldName,

            string? Kind,

            Outputs.ResolverPipelineConfig? PipelineConfig,

            string? RequestMappingTemplate,

            string? RequestMappingTemplateS3Location,

            string? ResponseMappingTemplate,

            string? ResponseMappingTemplateS3Location,

            Outputs.ResolverSyncConfig? SyncConfig,

            string TypeName)
        {
            this.ApiId = ApiId;
            this.CachingConfig = CachingConfig;
            this.DataSourceName = DataSourceName;
            this.FieldName = FieldName;
            this.Kind = Kind;
            this.PipelineConfig = PipelineConfig;
            this.RequestMappingTemplate = RequestMappingTemplate;
            this.RequestMappingTemplateS3Location = RequestMappingTemplateS3Location;
            this.ResponseMappingTemplate = ResponseMappingTemplate;
            this.ResponseMappingTemplateS3Location = ResponseMappingTemplateS3Location;
            this.SyncConfig = SyncConfig;
            this.TypeName = TypeName;
        }
    }
}