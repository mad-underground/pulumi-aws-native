// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::VerifiedAccessEndpoint resource creates an AWS EC2 Verified Access Endpoint.
type VerifiedAccessEndpoint struct {
	pulumi.CustomResourceState

	// The DNS name for users to reach your application.
	ApplicationDomain pulumi.StringOutput `pulumi:"applicationDomain"`
	// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
	AttachmentType pulumi.StringOutput `pulumi:"attachmentType"`
	// The creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// A description for the AWS Verified Access endpoint.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Returned if endpoint has a device trust provider attached.
	DeviceValidationDomain pulumi.StringOutput `pulumi:"deviceValidationDomain"`
	// The ARN of a public TLS/SSL certificate imported into or created with ACM.
	DomainCertificateArn pulumi.StringOutput `pulumi:"domainCertificateArn"`
	// A DNS name that is generated for the endpoint.
	EndpointDomain pulumi.StringOutput `pulumi:"endpointDomain"`
	// A custom identifier that gets prepended to a DNS name that is generated for the endpoint.
	EndpointDomainPrefix pulumi.StringOutput `pulumi:"endpointDomainPrefix"`
	// The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
	EndpointType pulumi.StringOutput `pulumi:"endpointType"`
	// The last updated time.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// The load balancer details if creating the AWS Verified Access endpoint as load-balancer type.
	LoadBalancerOptions VerifiedAccessEndpointLoadBalancerOptionsPtrOutput `pulumi:"loadBalancerOptions"`
	// The options for network-interface type endpoint.
	NetworkInterfaceOptions VerifiedAccessEndpointNetworkInterfaceOptionsPtrOutput `pulumi:"networkInterfaceOptions"`
	// The AWS Verified Access policy document.
	PolicyDocument pulumi.StringPtrOutput `pulumi:"policyDocument"`
	// The status of the Verified Access policy.
	PolicyEnabled pulumi.BoolPtrOutput `pulumi:"policyEnabled"`
	// The IDs of the security groups for the endpoint.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The endpoint status.
	Status pulumi.StringOutput `pulumi:"status"`
	// An array of key-value pairs to apply to this resource.
	Tags VerifiedAccessEndpointTagArrayOutput `pulumi:"tags"`
	// The ID of the AWS Verified Access endpoint.
	VerifiedAccessEndpointId pulumi.StringOutput `pulumi:"verifiedAccessEndpointId"`
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId pulumi.StringOutput `pulumi:"verifiedAccessGroupId"`
	// The ID of the AWS Verified Access instance.
	VerifiedAccessInstanceId pulumi.StringOutput `pulumi:"verifiedAccessInstanceId"`
}

// NewVerifiedAccessEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVerifiedAccessEndpoint(ctx *pulumi.Context,
	name string, args *VerifiedAccessEndpointArgs, opts ...pulumi.ResourceOption) (*VerifiedAccessEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationDomain == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationDomain'")
	}
	if args.AttachmentType == nil {
		return nil, errors.New("invalid value for required argument 'AttachmentType'")
	}
	if args.DomainCertificateArn == nil {
		return nil, errors.New("invalid value for required argument 'DomainCertificateArn'")
	}
	if args.EndpointDomainPrefix == nil {
		return nil, errors.New("invalid value for required argument 'EndpointDomainPrefix'")
	}
	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	if args.VerifiedAccessGroupId == nil {
		return nil, errors.New("invalid value for required argument 'VerifiedAccessGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VerifiedAccessEndpoint
	err := ctx.RegisterResource("aws-native:ec2:VerifiedAccessEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVerifiedAccessEndpoint gets an existing VerifiedAccessEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVerifiedAccessEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VerifiedAccessEndpointState, opts ...pulumi.ResourceOption) (*VerifiedAccessEndpoint, error) {
	var resource VerifiedAccessEndpoint
	err := ctx.ReadResource("aws-native:ec2:VerifiedAccessEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VerifiedAccessEndpoint resources.
type verifiedAccessEndpointState struct {
}

type VerifiedAccessEndpointState struct {
}

func (VerifiedAccessEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*verifiedAccessEndpointState)(nil)).Elem()
}

type verifiedAccessEndpointArgs struct {
	// The DNS name for users to reach your application.
	ApplicationDomain string `pulumi:"applicationDomain"`
	// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
	AttachmentType string `pulumi:"attachmentType"`
	// A description for the AWS Verified Access endpoint.
	Description *string `pulumi:"description"`
	// The ARN of a public TLS/SSL certificate imported into or created with ACM.
	DomainCertificateArn string `pulumi:"domainCertificateArn"`
	// A custom identifier that gets prepended to a DNS name that is generated for the endpoint.
	EndpointDomainPrefix string `pulumi:"endpointDomainPrefix"`
	// The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
	EndpointType string `pulumi:"endpointType"`
	// The load balancer details if creating the AWS Verified Access endpoint as load-balancer type.
	LoadBalancerOptions *VerifiedAccessEndpointLoadBalancerOptions `pulumi:"loadBalancerOptions"`
	// The options for network-interface type endpoint.
	NetworkInterfaceOptions *VerifiedAccessEndpointNetworkInterfaceOptions `pulumi:"networkInterfaceOptions"`
	// The AWS Verified Access policy document.
	PolicyDocument *string `pulumi:"policyDocument"`
	// The status of the Verified Access policy.
	PolicyEnabled *bool `pulumi:"policyEnabled"`
	// The IDs of the security groups for the endpoint.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// An array of key-value pairs to apply to this resource.
	Tags []VerifiedAccessEndpointTag `pulumi:"tags"`
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId string `pulumi:"verifiedAccessGroupId"`
}

// The set of arguments for constructing a VerifiedAccessEndpoint resource.
type VerifiedAccessEndpointArgs struct {
	// The DNS name for users to reach your application.
	ApplicationDomain pulumi.StringInput
	// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
	AttachmentType pulumi.StringInput
	// A description for the AWS Verified Access endpoint.
	Description pulumi.StringPtrInput
	// The ARN of a public TLS/SSL certificate imported into or created with ACM.
	DomainCertificateArn pulumi.StringInput
	// A custom identifier that gets prepended to a DNS name that is generated for the endpoint.
	EndpointDomainPrefix pulumi.StringInput
	// The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
	EndpointType pulumi.StringInput
	// The load balancer details if creating the AWS Verified Access endpoint as load-balancer type.
	LoadBalancerOptions VerifiedAccessEndpointLoadBalancerOptionsPtrInput
	// The options for network-interface type endpoint.
	NetworkInterfaceOptions VerifiedAccessEndpointNetworkInterfaceOptionsPtrInput
	// The AWS Verified Access policy document.
	PolicyDocument pulumi.StringPtrInput
	// The status of the Verified Access policy.
	PolicyEnabled pulumi.BoolPtrInput
	// The IDs of the security groups for the endpoint.
	SecurityGroupIds pulumi.StringArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags VerifiedAccessEndpointTagArrayInput
	// The ID of the AWS Verified Access group.
	VerifiedAccessGroupId pulumi.StringInput
}

func (VerifiedAccessEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*verifiedAccessEndpointArgs)(nil)).Elem()
}

type VerifiedAccessEndpointInput interface {
	pulumi.Input

	ToVerifiedAccessEndpointOutput() VerifiedAccessEndpointOutput
	ToVerifiedAccessEndpointOutputWithContext(ctx context.Context) VerifiedAccessEndpointOutput
}

func (*VerifiedAccessEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**VerifiedAccessEndpoint)(nil)).Elem()
}

func (i *VerifiedAccessEndpoint) ToVerifiedAccessEndpointOutput() VerifiedAccessEndpointOutput {
	return i.ToVerifiedAccessEndpointOutputWithContext(context.Background())
}

func (i *VerifiedAccessEndpoint) ToVerifiedAccessEndpointOutputWithContext(ctx context.Context) VerifiedAccessEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VerifiedAccessEndpointOutput)
}

type VerifiedAccessEndpointOutput struct{ *pulumi.OutputState }

func (VerifiedAccessEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VerifiedAccessEndpoint)(nil)).Elem()
}

func (o VerifiedAccessEndpointOutput) ToVerifiedAccessEndpointOutput() VerifiedAccessEndpointOutput {
	return o
}

func (o VerifiedAccessEndpointOutput) ToVerifiedAccessEndpointOutputWithContext(ctx context.Context) VerifiedAccessEndpointOutput {
	return o
}

// The DNS name for users to reach your application.
func (o VerifiedAccessEndpointOutput) ApplicationDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.ApplicationDomain }).(pulumi.StringOutput)
}

// The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the application.
func (o VerifiedAccessEndpointOutput) AttachmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.AttachmentType }).(pulumi.StringOutput)
}

// The creation time.
func (o VerifiedAccessEndpointOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// A description for the AWS Verified Access endpoint.
func (o VerifiedAccessEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Returned if endpoint has a device trust provider attached.
func (o VerifiedAccessEndpointOutput) DeviceValidationDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.DeviceValidationDomain }).(pulumi.StringOutput)
}

// The ARN of a public TLS/SSL certificate imported into or created with ACM.
func (o VerifiedAccessEndpointOutput) DomainCertificateArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.DomainCertificateArn }).(pulumi.StringOutput)
}

// A DNS name that is generated for the endpoint.
func (o VerifiedAccessEndpointOutput) EndpointDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.EndpointDomain }).(pulumi.StringOutput)
}

// A custom identifier that gets prepended to a DNS name that is generated for the endpoint.
func (o VerifiedAccessEndpointOutput) EndpointDomainPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.EndpointDomainPrefix }).(pulumi.StringOutput)
}

// The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified.
func (o VerifiedAccessEndpointOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.EndpointType }).(pulumi.StringOutput)
}

// The last updated time.
func (o VerifiedAccessEndpointOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// The load balancer details if creating the AWS Verified Access endpoint as load-balancer type.
func (o VerifiedAccessEndpointOutput) LoadBalancerOptions() VerifiedAccessEndpointLoadBalancerOptionsPtrOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) VerifiedAccessEndpointLoadBalancerOptionsPtrOutput {
		return v.LoadBalancerOptions
	}).(VerifiedAccessEndpointLoadBalancerOptionsPtrOutput)
}

// The options for network-interface type endpoint.
func (o VerifiedAccessEndpointOutput) NetworkInterfaceOptions() VerifiedAccessEndpointNetworkInterfaceOptionsPtrOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) VerifiedAccessEndpointNetworkInterfaceOptionsPtrOutput {
		return v.NetworkInterfaceOptions
	}).(VerifiedAccessEndpointNetworkInterfaceOptionsPtrOutput)
}

// The AWS Verified Access policy document.
func (o VerifiedAccessEndpointOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringPtrOutput { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// The status of the Verified Access policy.
func (o VerifiedAccessEndpointOutput) PolicyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.BoolPtrOutput { return v.PolicyEnabled }).(pulumi.BoolPtrOutput)
}

// The IDs of the security groups for the endpoint.
func (o VerifiedAccessEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The endpoint status.
func (o VerifiedAccessEndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// An array of key-value pairs to apply to this resource.
func (o VerifiedAccessEndpointOutput) Tags() VerifiedAccessEndpointTagArrayOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) VerifiedAccessEndpointTagArrayOutput { return v.Tags }).(VerifiedAccessEndpointTagArrayOutput)
}

// The ID of the AWS Verified Access endpoint.
func (o VerifiedAccessEndpointOutput) VerifiedAccessEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.VerifiedAccessEndpointId }).(pulumi.StringOutput)
}

// The ID of the AWS Verified Access group.
func (o VerifiedAccessEndpointOutput) VerifiedAccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.VerifiedAccessGroupId }).(pulumi.StringOutput)
}

// The ID of the AWS Verified Access instance.
func (o VerifiedAccessEndpointOutput) VerifiedAccessInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VerifiedAccessEndpoint) pulumi.StringOutput { return v.VerifiedAccessInstanceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VerifiedAccessEndpointInput)(nil)).Elem(), &VerifiedAccessEndpoint{})
	pulumi.RegisterOutputType(VerifiedAccessEndpointOutput{})
}
