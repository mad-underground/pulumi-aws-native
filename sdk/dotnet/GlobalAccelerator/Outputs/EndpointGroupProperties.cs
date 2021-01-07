// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.GlobalAccelerator.Outputs
{

    [OutputType]
    public sealed class EndpointGroupProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointconfigurations
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointGroupEndpointConfiguration> EndpointConfigurations;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-endpointgroupregion
        /// </summary>
        public readonly string EndpointGroupRegion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckintervalseconds
        /// </summary>
        public readonly int? HealthCheckIntervalSeconds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckpath
        /// </summary>
        public readonly string? HealthCheckPath;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckport
        /// </summary>
        public readonly int? HealthCheckPort;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-healthcheckprotocol
        /// </summary>
        public readonly string? HealthCheckProtocol;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-listenerarn
        /// </summary>
        public readonly string ListenerArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-portoverrides
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointGroupPortOverride> PortOverrides;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-thresholdcount
        /// </summary>
        public readonly int? ThresholdCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-endpointgroup.html#cfn-globalaccelerator-endpointgroup-trafficdialpercentage
        /// </summary>
        public readonly double? TrafficDialPercentage;

        [OutputConstructor]
        private EndpointGroupProperties(
            ImmutableArray<Outputs.EndpointGroupEndpointConfiguration> EndpointConfigurations,

            string EndpointGroupRegion,

            int? HealthCheckIntervalSeconds,

            string? HealthCheckPath,

            int? HealthCheckPort,

            string? HealthCheckProtocol,

            string ListenerArn,

            ImmutableArray<Outputs.EndpointGroupPortOverride> PortOverrides,

            int? ThresholdCount,

            double? TrafficDialPercentage)
        {
            this.EndpointConfigurations = EndpointConfigurations;
            this.EndpointGroupRegion = EndpointGroupRegion;
            this.HealthCheckIntervalSeconds = HealthCheckIntervalSeconds;
            this.HealthCheckPath = HealthCheckPath;
            this.HealthCheckPort = HealthCheckPort;
            this.HealthCheckProtocol = HealthCheckProtocol;
            this.ListenerArn = ListenerArn;
            this.PortOverrides = PortOverrides;
            this.ThresholdCount = ThresholdCount;
            this.TrafficDialPercentage = TrafficDialPercentage;
        }
    }
}
