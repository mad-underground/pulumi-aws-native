// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisVisualCustomActionOperationArgs : global::Pulumi.ResourceArgs
    {
        [Input("filterOperation")]
        public Input<Inputs.AnalysisCustomActionFilterOperationArgs>? FilterOperation { get; set; }

        [Input("navigationOperation")]
        public Input<Inputs.AnalysisCustomActionNavigationOperationArgs>? NavigationOperation { get; set; }

        [Input("setParametersOperation")]
        public Input<Inputs.AnalysisCustomActionSetParametersOperationArgs>? SetParametersOperation { get; set; }

        [Input("uRLOperation")]
        public Input<Inputs.AnalysisCustomActionURLOperationArgs>? URLOperation { get; set; }

        public AnalysisVisualCustomActionOperationArgs()
        {
        }
        public static new AnalysisVisualCustomActionOperationArgs Empty => new AnalysisVisualCustomActionOperationArgs();
    }
}
