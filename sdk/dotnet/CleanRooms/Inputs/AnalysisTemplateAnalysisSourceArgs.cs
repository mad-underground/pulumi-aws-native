// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class AnalysisTemplateAnalysisSourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        public AnalysisTemplateAnalysisSourceArgs()
        {
        }
        public static new AnalysisTemplateAnalysisSourceArgs Empty => new AnalysisTemplateAnalysisSourceArgs();
    }
}
