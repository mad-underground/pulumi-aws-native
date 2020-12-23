// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcecustomknowledgearticletypeconfigurationlist.html
    /// </summary>
    public sealed class DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationListArgs : Pulumi.ResourceArgs
    {
        [Input("SalesforceCustomKnowledgeArticleTypeConfigurationList")]
        private InputList<Inputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs>? _SalesforceCustomKnowledgeArticleTypeConfigurationList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcecustomknowledgearticletypeconfigurationlist.html#cfn-kendra-datasource-salesforcecustomknowledgearticletypeconfigurationlist-salesforcecustomknowledgearticletypeconfigurationlist
        /// </summary>
        public InputList<Inputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs> SalesforceCustomKnowledgeArticleTypeConfigurationList
        {
            get => _SalesforceCustomKnowledgeArticleTypeConfigurationList ?? (_SalesforceCustomKnowledgeArticleTypeConfigurationList = new InputList<Inputs.DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationArgs>());
            set => _SalesforceCustomKnowledgeArticleTypeConfigurationList = value;
        }

        public DataSourceSalesforceCustomKnowledgeArticleTypeConfigurationListArgs()
        {
        }
    }
}
