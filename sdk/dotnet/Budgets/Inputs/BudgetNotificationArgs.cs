// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets.Inputs
{

    public sealed class BudgetNotificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("comparisonOperator", required: true)]
        public Input<string> ComparisonOperator { get; set; } = null!;

        [Input("notificationType", required: true)]
        public Input<string> NotificationType { get; set; } = null!;

        [Input("threshold", required: true)]
        public Input<double> Threshold { get; set; } = null!;

        [Input("thresholdType")]
        public Input<string>? ThresholdType { get; set; }

        public BudgetNotificationArgs()
        {
        }
        public static new BudgetNotificationArgs Empty => new BudgetNotificationArgs();
    }
}
