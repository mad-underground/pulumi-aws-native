// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceSalesforceKnowledgeArticleConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html#cfn-kendra-datasource-salesforceknowledgearticleconfiguration-customknowledgearticletypeconfigurations
        /// </summary>
        public readonly Outputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationList? CustomKnowledgeArticleTypeConfigurations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html#cfn-kendra-datasource-salesforceknowledgearticleconfiguration-includedstates
        /// </summary>
        public readonly Outputs.DataSourceSalesforceKnowledgeArticleStateList IncludedStates;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticleconfiguration.html#cfn-kendra-datasource-salesforceknowledgearticleconfiguration-standardknowledgearticletypeconfiguration
        /// </summary>
        public readonly Outputs.DataSourceSalesforceStandardKnowledgeArticleTypeConfiguration? StandardKnowledgeArticleTypeConfiguration;

        [OutputConstructor]
        private DataSourceSalesforceKnowledgeArticleConfiguration(
            Outputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationList? CustomKnowledgeArticleTypeConfigurations,

            Outputs.DataSourceSalesforceKnowledgeArticleStateList IncludedStates,

            Outputs.DataSourceSalesforceStandardKnowledgeArticleTypeConfiguration? StandardKnowledgeArticleTypeConfiguration)
        {
            this.CustomKnowledgeArticleTypeConfigurations = CustomKnowledgeArticleTypeConfigurations;
            this.IncludedStates = IncludedStates;
            this.StandardKnowledgeArticleTypeConfiguration = StandardKnowledgeArticleTypeConfiguration;
        }
    }
}
