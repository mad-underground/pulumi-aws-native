// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaStore.Inputs
{

    public sealed class ContainerMetricPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("containerLevelMetrics", required: true)]
        public Input<string> ContainerLevelMetrics { get; set; } = null!;

        [Input("metricPolicyRules")]
        private InputList<Inputs.ContainerMetricPolicyRuleArgs>? _metricPolicyRules;
        public InputList<Inputs.ContainerMetricPolicyRuleArgs> MetricPolicyRules
        {
            get => _metricPolicyRules ?? (_metricPolicyRules = new InputList<Inputs.ContainerMetricPolicyRuleArgs>());
            set => _metricPolicyRules = value;
        }

        public ContainerMetricPolicyArgs()
        {
        }
        public static new ContainerMetricPolicyArgs Empty => new ContainerMetricPolicyArgs();
    }
}
