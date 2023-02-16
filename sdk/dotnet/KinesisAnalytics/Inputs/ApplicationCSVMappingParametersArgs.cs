// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalytics.Inputs
{

    public sealed class ApplicationCSVMappingParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("recordColumnDelimiter", required: true)]
        public Input<string> RecordColumnDelimiter { get; set; } = null!;

        [Input("recordRowDelimiter", required: true)]
        public Input<string> RecordRowDelimiter { get; set; } = null!;

        public ApplicationCSVMappingParametersArgs()
        {
        }
        public static new ApplicationCSVMappingParametersArgs Empty => new ApplicationCSVMappingParametersArgs();
    }
}
