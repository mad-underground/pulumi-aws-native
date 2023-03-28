// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisForecastScenarioArgs : global::Pulumi.ResourceArgs
    {
        [Input("whatIfPointScenario")]
        public Input<Inputs.AnalysisWhatIfPointScenarioArgs>? WhatIfPointScenario { get; set; }

        [Input("whatIfRangeScenario")]
        public Input<Inputs.AnalysisWhatIfRangeScenarioArgs>? WhatIfRangeScenario { get; set; }

        public AnalysisForecastScenarioArgs()
        {
        }
        public static new AnalysisForecastScenarioArgs Empty => new AnalysisForecastScenarioArgs();
    }
}