// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class ModelCardBarChartMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.ModelCardBarChartMetricType> Type { get; set; } = null!;

        [Input("value", required: true)]
        private InputList<double>? _value;
        public InputList<double> Value
        {
            get => _value ?? (_value = new InputList<double>());
            set => _value = value;
        }

        [Input("xAxisName")]
        private InputList<string>? _xAxisName;
        public InputList<string> XAxisName
        {
            get => _xAxisName ?? (_xAxisName = new InputList<string>());
            set => _xAxisName = value;
        }

        [Input("yAxisName")]
        public Input<string>? YAxisName { get; set; }

        public ModelCardBarChartMetricArgs()
        {
        }
        public static new ModelCardBarChartMetricArgs Empty => new ModelCardBarChartMetricArgs();
    }
}