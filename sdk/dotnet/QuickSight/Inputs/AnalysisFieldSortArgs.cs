// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisFieldSortArgs : global::Pulumi.ResourceArgs
    {
        [Input("direction", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisSortDirection> Direction { get; set; } = null!;

        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        public AnalysisFieldSortArgs()
        {
        }
        public static new AnalysisFieldSortArgs Empty => new AnalysisFieldSortArgs();
    }
}
