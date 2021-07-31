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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticlestatelist.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceSalesforceKnowledgeArticleStateList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticlestatelist.html#cfn-kendra-datasource-salesforceknowledgearticlestatelist-salesforceknowledgearticlestatelist
        /// </summary>
        public readonly ImmutableArray<string> SalesforceKnowledgeArticleStateList;

        [OutputConstructor]
        private DataSourceSalesforceKnowledgeArticleStateList(ImmutableArray<string> salesforceKnowledgeArticleStateList)
        {
            SalesforceKnowledgeArticleStateList = salesforceKnowledgeArticleStateList;
        }
    }
}
