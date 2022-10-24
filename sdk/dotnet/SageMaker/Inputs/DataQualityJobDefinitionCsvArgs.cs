// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The CSV format
    /// </summary>
    public sealed class DataQualityJobDefinitionCsvArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean flag indicating if given CSV has header
        /// </summary>
        [Input("header")]
        public Input<bool>? Header { get; set; }

        public DataQualityJobDefinitionCsvArgs()
        {
        }
        public static new DataQualityJobDefinitionCsvArgs Empty => new DataQualityJobDefinitionCsvArgs();
    }
}
