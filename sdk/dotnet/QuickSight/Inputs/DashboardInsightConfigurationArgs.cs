// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardInsightConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("computations")]
        private InputList<Inputs.DashboardComputationArgs>? _computations;
        public InputList<Inputs.DashboardComputationArgs> Computations
        {
            get => _computations ?? (_computations = new InputList<Inputs.DashboardComputationArgs>());
            set => _computations = value;
        }

        [Input("customNarrative")]
        public Input<Inputs.DashboardCustomNarrativeOptionsArgs>? CustomNarrative { get; set; }

        public DashboardInsightConfigurationArgs()
        {
        }
        public static new DashboardInsightConfigurationArgs Empty => new DashboardInsightConfigurationArgs();
    }
}
