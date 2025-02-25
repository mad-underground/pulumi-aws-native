// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisVisualCustomAction
    {
        public readonly ImmutableArray<Outputs.AnalysisVisualCustomActionOperation> ActionOperations;
        public readonly string CustomActionId;
        public readonly string Name;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisWidgetStatus? Status;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisualCustomActionTrigger Trigger;

        [OutputConstructor]
        private AnalysisVisualCustomAction(
            ImmutableArray<Outputs.AnalysisVisualCustomActionOperation> actionOperations,

            string customActionId,

            string name,

            Pulumi.AwsNative.QuickSight.AnalysisWidgetStatus? status,

            Pulumi.AwsNative.QuickSight.AnalysisVisualCustomActionTrigger trigger)
        {
            ActionOperations = actionOperations;
            CustomActionId = customActionId;
            Name = name;
            Status = status;
            Trigger = trigger;
        }
    }
}
