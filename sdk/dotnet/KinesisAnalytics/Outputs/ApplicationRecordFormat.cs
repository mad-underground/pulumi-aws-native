// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisAnalytics.Outputs
{

    [OutputType]
    public sealed class ApplicationRecordFormat
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-mappingparameters
        /// </summary>
        public readonly Outputs.ApplicationMappingParameters? MappingParameters;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-recordformattype
        /// </summary>
        public readonly string RecordFormatType;

        [OutputConstructor]
        private ApplicationRecordFormat(
            Outputs.ApplicationMappingParameters? MappingParameters,

            string RecordFormatType)
        {
            this.MappingParameters = MappingParameters;
            this.RecordFormatType = RecordFormatType;
        }
    }
}