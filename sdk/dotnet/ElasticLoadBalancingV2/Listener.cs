// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2
{
    /// <summary>
    /// Specifies a listener for an Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.
    /// </summary>
    [AwsNativeResourceType("aws-native:elasticloadbalancingv2:Listener")]
    public partial class Listener : global::Pulumi.CustomResource
    {
        /// <summary>
        /// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
        /// </summary>
        [Output("alpnPolicy")]
        public Output<ImmutableArray<string>> AlpnPolicy { get; private set; } = null!;

        /// <summary>
        /// The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
        ///  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
        /// </summary>
        [Output("certificates")]
        public Output<ImmutableArray<Outputs.ListenerCertificate>> Certificates { get; private set; } = null!;

        /// <summary>
        /// The actions for the default rule. You cannot define a condition for a default rule.
        ///  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
        /// </summary>
        [Output("defaultActions")]
        public Output<ImmutableArray<Outputs.ListenerAction>> DefaultActions { get; private set; } = null!;

        [Output("listenerArn")]
        public Output<string> ListenerArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the load balancer.
        /// </summary>
        [Output("loadBalancerArn")]
        public Output<string> LoadBalancerArn { get; private set; } = null!;

        /// <summary>
        /// The mutual authentication configuration information.
        /// </summary>
        [Output("mutualAuthentication")]
        public Output<Outputs.ListenerMutualAuthentication?> MutualAuthentication { get; private set; } = null!;

        /// <summary>
        /// The port on which the load balancer is listening. You cannot specify a port for a Gateway Load Balancer.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
        ///  For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide*.
        /// </summary>
        [Output("sslPolicy")]
        public Output<string?> SslPolicy { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:elasticloadbalancingv2:Listener", name, args ?? new ListenerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Listener(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:elasticloadbalancingv2:Listener", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "loadBalancerArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Listener resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listener Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Listener(name, id, options);
        }
    }

    public sealed class ListenerArgs : global::Pulumi.ResourceArgs
    {
        [Input("alpnPolicy")]
        private InputList<string>? _alpnPolicy;

        /// <summary>
        /// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
        /// </summary>
        public InputList<string> AlpnPolicy
        {
            get => _alpnPolicy ?? (_alpnPolicy = new InputList<string>());
            set => _alpnPolicy = value;
        }

        [Input("certificates")]
        private InputList<Inputs.ListenerCertificateArgs>? _certificates;

        /// <summary>
        /// The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
        ///  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
        /// </summary>
        public InputList<Inputs.ListenerCertificateArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.ListenerCertificateArgs>());
            set => _certificates = value;
        }

        [Input("defaultActions", required: true)]
        private InputList<Inputs.ListenerActionArgs>? _defaultActions;

        /// <summary>
        /// The actions for the default rule. You cannot define a condition for a default rule.
        ///  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
        /// </summary>
        public InputList<Inputs.ListenerActionArgs> DefaultActions
        {
            get => _defaultActions ?? (_defaultActions = new InputList<Inputs.ListenerActionArgs>());
            set => _defaultActions = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the load balancer.
        /// </summary>
        [Input("loadBalancerArn", required: true)]
        public Input<string> LoadBalancerArn { get; set; } = null!;

        /// <summary>
        /// The mutual authentication configuration information.
        /// </summary>
        [Input("mutualAuthentication")]
        public Input<Inputs.ListenerMutualAuthenticationArgs>? MutualAuthentication { get; set; }

        /// <summary>
        /// The port on which the load balancer is listening. You cannot specify a port for a Gateway Load Balancer.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
        ///  For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide*.
        /// </summary>
        [Input("sslPolicy")]
        public Input<string>? SslPolicy { get; set; }

        public ListenerArgs()
        {
        }
        public static new ListenerArgs Empty => new ListenerArgs();
    }
}
