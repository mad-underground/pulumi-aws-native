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
    public sealed class DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcecustomknowledgearticletypeconfigurationlist.html#cfn-kendra-datasource-salesforcecustomknowledgearticletypeconfigurationlist-salesforcecustomknowledgearticletypeconfigurationlist
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfiguration> SalesforceCustomKnowledgeArticleTypeConfigurationList;

        [OutputConstructor]
        private DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationList(ImmutableArray<Outputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfiguration> SalesforceCustomKnowledgeArticleTypeConfigurationList)
        {
            this.SalesforceCustomKnowledgeArticleTypeConfigurationList = SalesforceCustomKnowledgeArticleTypeConfigurationList;
        }
    }
}