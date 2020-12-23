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
    public sealed class DataSourceConfluenceSpaceToIndexFieldMapping
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacetoindexfieldmapping.html#cfn-kendra-datasource-confluencespacetoindexfieldmapping-datasourcefieldname
        /// </summary>
        public readonly string DataSourceFieldName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacetoindexfieldmapping.html#cfn-kendra-datasource-confluencespacetoindexfieldmapping-datefieldformat
        /// </summary>
        public readonly string? DateFieldFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacetoindexfieldmapping.html#cfn-kendra-datasource-confluencespacetoindexfieldmapping-indexfieldname
        /// </summary>
        public readonly string IndexFieldName;

        [OutputConstructor]
        private DataSourceConfluenceSpaceToIndexFieldMapping(
            string DataSourceFieldName,

            string? DateFieldFormat,

            string IndexFieldName)
        {
            this.DataSourceFieldName = DataSourceFieldName;
            this.DateFieldFormat = DateFieldFormat;
            this.IndexFieldName = IndexFieldName;
        }
    }
}
