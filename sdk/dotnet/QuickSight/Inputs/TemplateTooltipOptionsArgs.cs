// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTooltipOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldBasedTooltip")]
        public Input<Inputs.TemplateFieldBasedTooltipArgs>? FieldBasedTooltip { get; set; }

        [Input("selectedTooltipType")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateSelectedTooltipType>? SelectedTooltipType { get; set; }

        [Input("tooltipVisibility")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisibility>? TooltipVisibility { get; set; }

        public TemplateTooltipOptionsArgs()
        {
        }
        public static new TemplateTooltipOptionsArgs Empty => new TemplateTooltipOptionsArgs();
    }
}
