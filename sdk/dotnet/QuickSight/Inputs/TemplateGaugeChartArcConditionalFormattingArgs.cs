// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateGaugeChartArcConditionalFormattingArgs : global::Pulumi.ResourceArgs
    {
        [Input("foregroundColor")]
        public Input<Inputs.TemplateConditionalFormattingColorArgs>? ForegroundColor { get; set; }

        public TemplateGaugeChartArcConditionalFormattingArgs()
        {
        }
        public static new TemplateGaugeChartArcConditionalFormattingArgs Empty => new TemplateGaugeChartArcConditionalFormattingArgs();
    }
}
