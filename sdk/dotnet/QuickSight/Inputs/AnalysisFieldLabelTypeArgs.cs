// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisFieldLabelTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldId")]
        public Input<string>? FieldId { get; set; }

        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        public AnalysisFieldLabelTypeArgs()
        {
        }
        public static new AnalysisFieldLabelTypeArgs Empty => new AnalysisFieldLabelTypeArgs();
    }
}
