// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Athena.Outputs
{

    [OutputType]
    public sealed class DataCatalogProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-parameters
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Parameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-tags
        /// </summary>
        public readonly Outputs.DataCatalogTags? Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DataCatalogProperties(
            string? Description,

            string Name,

            ImmutableDictionary<string, string>? Parameters,

            Outputs.DataCatalogTags? Tags,

            string Type)
        {
            this.Description = Description;
            this.Name = Name;
            this.Parameters = Parameters;
            this.Tags = Tags;
            this.Type = Type;
        }
    }
}
