// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateVisualCustomActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionOperations", required: true)]
        private InputList<Inputs.TemplateVisualCustomActionOperationArgs>? _actionOperations;
        public InputList<Inputs.TemplateVisualCustomActionOperationArgs> ActionOperations
        {
            get => _actionOperations ?? (_actionOperations = new InputList<Inputs.TemplateVisualCustomActionOperationArgs>());
            set => _actionOperations = value;
        }

        [Input("customActionId", required: true)]
        public Input<string> CustomActionId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("status")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateWidgetStatus>? Status { get; set; }

        [Input("trigger", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.TemplateVisualCustomActionTrigger> Trigger { get; set; } = null!;

        public TemplateVisualCustomActionArgs()
        {
        }
        public static new TemplateVisualCustomActionArgs Empty => new TemplateVisualCustomActionArgs();
    }
}
