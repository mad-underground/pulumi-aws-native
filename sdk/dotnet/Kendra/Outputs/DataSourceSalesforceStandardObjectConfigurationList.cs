// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceSalesforceStandardObjectConfigurationList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardobjectconfigurationlist.html#cfn-kendra-datasource-salesforcestandardobjectconfigurationlist-salesforcestandardobjectconfigurationlist
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceSalesforceStandardObjectConfiguration> SalesforceStandardObjectConfigurationList;

        [OutputConstructor]
        private DataSourceSalesforceStandardObjectConfigurationList(ImmutableArray<Outputs.DataSourceSalesforceStandardObjectConfiguration> salesforceStandardObjectConfigurationList)
        {
            SalesforceStandardObjectConfigurationList = salesforceStandardObjectConfigurationList;
        }
    }
}
