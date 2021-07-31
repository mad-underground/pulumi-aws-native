// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancer.html
    /// </summary>
    [OutputType]
    public sealed class ServiceLoadBalancer
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancer.html#cfn-ecs-service-loadbalancer-containername
        /// </summary>
        public readonly string? ContainerName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancer.html#cfn-ecs-service-loadbalancer-containerport
        /// </summary>
        public readonly int? ContainerPort;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancer.html#cfn-ecs-service-loadbalancer-loadbalancername
        /// </summary>
        public readonly string? LoadBalancerName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-loadbalancer.html#cfn-ecs-service-loadbalancer-targetgrouparn
        /// </summary>
        public readonly string? TargetGroupArn;

        [OutputConstructor]
        private ServiceLoadBalancer(
            string? containerName,

            int? containerPort,

            string? loadBalancerName,

            string? targetGroupArn)
        {
            ContainerName = containerName;
            ContainerPort = containerPort;
            LoadBalancerName = loadBalancerName;
            TargetGroupArn = targetGroupArn;
        }
    }
}
