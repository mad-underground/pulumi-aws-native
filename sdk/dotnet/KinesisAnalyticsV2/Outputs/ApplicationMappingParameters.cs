// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// When you configure a SQL-based Kinesis Data Analytics application's input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
    /// </summary>
    [OutputType]
    public sealed class ApplicationMappingParameters
    {
        /// <summary>
        /// Provides additional mapping information when the record format uses delimiters (for example, CSV).
        /// </summary>
        public readonly Outputs.ApplicationCSVMappingParameters? CSVMappingParameters;
        /// <summary>
        /// Provides additional mapping information when JSON is the record format on the streaming source.
        /// </summary>
        public readonly Outputs.ApplicationJSONMappingParameters? JSONMappingParameters;

        [OutputConstructor]
        private ApplicationMappingParameters(
            Outputs.ApplicationCSVMappingParameters? cSVMappingParameters,

            Outputs.ApplicationJSONMappingParameters? jSONMappingParameters)
        {
            CSVMappingParameters = cSVMappingParameters;
            JSONMappingParameters = jSONMappingParameters;
        }
    }
}
