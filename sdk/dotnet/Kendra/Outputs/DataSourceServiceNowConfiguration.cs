// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceServiceNowConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-hosturl
        /// </summary>
        public readonly string HostUrl;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-knowledgearticleconfiguration
        /// </summary>
        public readonly Outputs.DataSourceServiceNowKnowledgeArticleConfiguration? KnowledgeArticleConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-secretarn
        /// </summary>
        public readonly string SecretArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-servicecatalogconfiguration
        /// </summary>
        public readonly Outputs.DataSourceServiceNowServiceCatalogConfiguration? ServiceCatalogConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-servicenowbuildversion
        /// </summary>
        public readonly string ServiceNowBuildVersion;

        [OutputConstructor]
        private DataSourceServiceNowConfiguration(
            string hostUrl,

            Outputs.DataSourceServiceNowKnowledgeArticleConfiguration? knowledgeArticleConfiguration,

            string secretArn,

            Outputs.DataSourceServiceNowServiceCatalogConfiguration? serviceCatalogConfiguration,

            string serviceNowBuildVersion)
        {
            HostUrl = hostUrl;
            KnowledgeArticleConfiguration = knowledgeArticleConfiguration;
            SecretArn = secretArn;
            ServiceCatalogConfiguration = serviceCatalogConfiguration;
            ServiceNowBuildVersion = serviceNowBuildVersion;
        }
    }
}
