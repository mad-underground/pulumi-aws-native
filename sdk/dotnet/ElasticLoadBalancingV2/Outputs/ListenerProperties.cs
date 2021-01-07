// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class ListenerProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-alpnpolicy
        /// </summary>
        public readonly ImmutableArray<string> AlpnPolicy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-certificates
        /// </summary>
        public readonly ImmutableArray<Outputs.ListenerCertificate> Certificates;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-defaultactions
        /// </summary>
        public readonly ImmutableArray<Outputs.ListenerAction> DefaultActions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-loadbalancerarn
        /// </summary>
        public readonly string LoadBalancerArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-port
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-protocol
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html#cfn-elasticloadbalancingv2-listener-sslpolicy
        /// </summary>
        public readonly string? SslPolicy;

        [OutputConstructor]
        private ListenerProperties(
            ImmutableArray<string> AlpnPolicy,

            ImmutableArray<Outputs.ListenerCertificate> Certificates,

            ImmutableArray<Outputs.ListenerAction> DefaultActions,

            string LoadBalancerArn,

            int? Port,

            string? Protocol,

            string? SslPolicy)
        {
            this.AlpnPolicy = AlpnPolicy;
            this.Certificates = Certificates;
            this.DefaultActions = DefaultActions;
            this.LoadBalancerArn = LoadBalancerArn;
            this.Port = Port;
            this.Protocol = Protocol;
            this.SslPolicy = SslPolicy;
        }
    }
}
