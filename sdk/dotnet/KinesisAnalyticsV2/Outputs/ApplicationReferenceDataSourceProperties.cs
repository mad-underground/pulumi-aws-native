// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationReferenceDataSourceProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationreferencedatasource.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-applicationname
        /// </summary>
        public readonly string ApplicationName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationreferencedatasource.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-referencedatasource
        /// </summary>
        public readonly Outputs.ApplicationReferenceDataSourceReferenceDataSource ReferenceDataSource;

        [OutputConstructor]
        private ApplicationReferenceDataSourceProperties(
            string ApplicationName,

            Outputs.ApplicationReferenceDataSourceReferenceDataSource ReferenceDataSource)
        {
            this.ApplicationName = ApplicationName;
            this.ReferenceDataSource = ReferenceDataSource;
        }
    }
}