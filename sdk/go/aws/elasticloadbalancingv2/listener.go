// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a listener for an Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.
type Listener struct {
	pulumi.CustomResourceState

	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	AlpnPolicy pulumi.StringArrayOutput `pulumi:"alpnPolicy"`
	// The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
	Certificates ListenerCertificateTypeArrayOutput `pulumi:"certificates"`
	// The actions for the default rule. You cannot define a condition for a default rule.
	//  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
	DefaultActions ListenerActionArrayOutput `pulumi:"defaultActions"`
	ListenerArn    pulumi.StringOutput       `pulumi:"listenerArn"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn pulumi.StringOutput `pulumi:"loadBalancerArn"`
	// The mutual authentication configuration information.
	MutualAuthentication ListenerMutualAuthenticationPtrOutput `pulumi:"mutualAuthentication"`
	// The port on which the load balancer is listening. You cannot specify a port for a Gateway Load Balancer.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//  For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide*.
	SslPolicy pulumi.StringPtrOutput `pulumi:"sslPolicy"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultActions == nil {
		return nil, errors.New("invalid value for required argument 'DefaultActions'")
	}
	if args.LoadBalancerArn == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"loadBalancerArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Listener
	err := ctx.RegisterResource("aws-native:elasticloadbalancingv2:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("aws-native:elasticloadbalancingv2:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
}

type ListenerState struct {
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	AlpnPolicy []string `pulumi:"alpnPolicy"`
	// The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
	Certificates []ListenerCertificateType `pulumi:"certificates"`
	// The actions for the default rule. You cannot define a condition for a default rule.
	//  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
	DefaultActions []ListenerAction `pulumi:"defaultActions"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn string `pulumi:"loadBalancerArn"`
	// The mutual authentication configuration information.
	MutualAuthentication *ListenerMutualAuthentication `pulumi:"mutualAuthentication"`
	// The port on which the load balancer is listening. You cannot specify a port for a Gateway Load Balancer.
	Port *int `pulumi:"port"`
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	Protocol *string `pulumi:"protocol"`
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//  For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide*.
	SslPolicy *string `pulumi:"sslPolicy"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
	AlpnPolicy pulumi.StringArrayInput
	// The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	//  To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
	Certificates ListenerCertificateTypeArrayInput
	// The actions for the default rule. You cannot define a condition for a default rule.
	//  To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
	DefaultActions ListenerActionArrayInput
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn pulumi.StringInput
	// The mutual authentication configuration information.
	MutualAuthentication ListenerMutualAuthenticationPtrInput
	// The port on which the load balancer is listening. You cannot specify a port for a Gateway Load Balancer.
	Port pulumi.IntPtrInput
	// The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
	Protocol pulumi.StringPtrInput
	// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
	//  For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide*.
	SslPolicy pulumi.StringPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

// [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
func (o ListenerOutput) AlpnPolicy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringArrayOutput { return v.AlpnPolicy }).(pulumi.StringArrayOutput)
}

// The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
//
//	To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html).
func (o ListenerOutput) Certificates() ListenerCertificateTypeArrayOutput {
	return o.ApplyT(func(v *Listener) ListenerCertificateTypeArrayOutput { return v.Certificates }).(ListenerCertificateTypeArrayOutput)
}

// The actions for the default rule. You cannot define a condition for a default rule.
//
//	To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html).
func (o ListenerOutput) DefaultActions() ListenerActionArrayOutput {
	return o.ApplyT(func(v *Listener) ListenerActionArrayOutput { return v.DefaultActions }).(ListenerActionArrayOutput)
}

func (o ListenerOutput) ListenerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ListenerArn }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the load balancer.
func (o ListenerOutput) LoadBalancerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.LoadBalancerArn }).(pulumi.StringOutput)
}

// The mutual authentication configuration information.
func (o ListenerOutput) MutualAuthentication() ListenerMutualAuthenticationPtrOutput {
	return o.ApplyT(func(v *Listener) ListenerMutualAuthenticationPtrOutput { return v.MutualAuthentication }).(ListenerMutualAuthenticationPtrOutput)
}

// The port on which the load balancer is listening. You cannot specify a port for a Gateway Load Balancer.
func (o ListenerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot specify a protocol for a Gateway Load Balancer.
func (o ListenerOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.
//
//	For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies) in the *Network Load Balancers Guide*.
func (o ListenerOutput) SslPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.SslPolicy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterOutputType(ListenerOutput{})
}
