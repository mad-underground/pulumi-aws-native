// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class DatasetTriggeringDatasetArgs : global::Pulumi.ResourceArgs
    {
        [Input("datasetName", required: true)]
        public Input<string> DatasetName { get; set; } = null!;

        public DatasetTriggeringDatasetArgs()
        {
        }
        public static new DatasetTriggeringDatasetArgs Empty => new DatasetTriggeringDatasetArgs();
    }
}
