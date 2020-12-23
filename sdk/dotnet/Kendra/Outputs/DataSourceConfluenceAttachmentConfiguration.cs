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
    public sealed class DataSourceConfluenceAttachmentConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceattachmentconfiguration.html#cfn-kendra-datasource-confluenceattachmentconfiguration-attachmentfieldmappings
        /// </summary>
        public readonly Outputs.DataSourceConfluenceAttachmentFieldMappingsList? AttachmentFieldMappings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceattachmentconfiguration.html#cfn-kendra-datasource-confluenceattachmentconfiguration-crawlattachments
        /// </summary>
        public readonly bool? CrawlAttachments;

        [OutputConstructor]
        private DataSourceConfluenceAttachmentConfiguration(
            Outputs.DataSourceConfluenceAttachmentFieldMappingsList? AttachmentFieldMappings,

            bool? CrawlAttachments)
        {
            this.AttachmentFieldMappings = AttachmentFieldMappings;
            this.CrawlAttachments = CrawlAttachments;
        }
    }
}
