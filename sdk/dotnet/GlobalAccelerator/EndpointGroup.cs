// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GlobalAccelerator
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html
    /// </summary>
    [AwsNativeResourceType("aws-native:GlobalAccelerator:EndpointGroup")]
    public partial class EndpointGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
        /// </summary>
        [Output("EndpointConfigurations")]
        public Output<ImmutableArray<Outputs.EndpointGroupEndpointConfiguration>> EndpointConfigurations { get; private set; } = null!;

        [Output("EndpointGroupArn")]
        public Output<string> EndpointGroupArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
        /// </summary>
        [Output("EndpointGroupRegion")]
        public Output<string> EndpointGroupRegion { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
        /// </summary>
        [Output("HealthCheckIntervalSeconds")]
        public Output<int?> HealthCheckIntervalSeconds { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
        /// </summary>
        [Output("HealthCheckPath")]
        public Output<string?> HealthCheckPath { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
        /// </summary>
        [Output("HealthCheckPort")]
        public Output<int?> HealthCheckPort { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
        /// </summary>
        [Output("HealthCheckProtocol")]
        public Output<string?> HealthCheckProtocol { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
        /// </summary>
        [Output("ListenerArn")]
        public Output<string> ListenerArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-portoverrides
        /// </summary>
        [Output("PortOverrides")]
        public Output<ImmutableArray<Outputs.EndpointGroupPortOverride>> PortOverrides { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
        /// </summary>
        [Output("ThresholdCount")]
        public Output<int?> ThresholdCount { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
        /// </summary>
        [Output("TrafficDialPercentage")]
        public Output<double?> TrafficDialPercentage { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointGroup(string name, EndpointGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:GlobalAccelerator:EndpointGroup", name, args ?? new EndpointGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:GlobalAccelerator:EndpointGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EndpointGroup(name, id, options);
        }
    }

    public sealed class EndpointGroupArgs : Pulumi.ResourceArgs
    {
        [Input("EndpointConfigurations")]
        private InputList<Inputs.EndpointGroupEndpointConfigurationArgs>? _EndpointConfigurations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
        /// </summary>
        public InputList<Inputs.EndpointGroupEndpointConfigurationArgs> EndpointConfigurations
        {
            get => _EndpointConfigurations ?? (_EndpointConfigurations = new InputList<Inputs.EndpointGroupEndpointConfigurationArgs>());
            set => _EndpointConfigurations = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
        /// </summary>
        [Input("EndpointGroupRegion", required: true)]
        public Input<string> EndpointGroupRegion { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
        /// </summary>
        [Input("HealthCheckIntervalSeconds")]
        public Input<int>? HealthCheckIntervalSeconds { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
        /// </summary>
        [Input("HealthCheckPath")]
        public Input<string>? HealthCheckPath { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
        /// </summary>
        [Input("HealthCheckPort")]
        public Input<int>? HealthCheckPort { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
        /// </summary>
        [Input("HealthCheckProtocol")]
        public Input<string>? HealthCheckProtocol { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
        /// </summary>
        [Input("ListenerArn", required: true)]
        public Input<string> ListenerArn { get; set; } = null!;

        [Input("PortOverrides")]
        private InputList<Inputs.EndpointGroupPortOverrideArgs>? _PortOverrides;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-portoverrides
        /// </summary>
        public InputList<Inputs.EndpointGroupPortOverrideArgs> PortOverrides
        {
            get => _PortOverrides ?? (_PortOverrides = new InputList<Inputs.EndpointGroupPortOverrideArgs>());
            set => _PortOverrides = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
        /// </summary>
        [Input("ThresholdCount")]
        public Input<int>? ThresholdCount { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
        /// </summary>
        [Input("TrafficDialPercentage")]
        public Input<double>? TrafficDialPercentage { get; set; }

        public EndpointGroupArgs()
        {
        }
    }
}
