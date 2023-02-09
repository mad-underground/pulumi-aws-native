// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DataSync::Agent.
type Agent struct {
	pulumi.CustomResourceState

	// Activation key of the Agent.
	ActivationKey pulumi.StringPtrOutput `pulumi:"activationKey"`
	// The DataSync Agent ARN.
	AgentArn pulumi.StringOutput `pulumi:"agentArn"`
	// The name configured for the agent. Text reference used to identify the agent in the console.
	AgentName pulumi.StringPtrOutput `pulumi:"agentName"`
	// The service endpoints that the agent will connect to.
	EndpointType AgentEndpointTypeOutput `pulumi:"endpointType"`
	// The ARNs of the security group used to protect your data transfer task subnets.
	SecurityGroupArns pulumi.StringArrayOutput `pulumi:"securityGroupArns"`
	// The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	SubnetArns pulumi.StringArrayOutput `pulumi:"subnetArns"`
	// An array of key-value pairs to apply to this resource.
	Tags AgentTagArrayOutput `pulumi:"tags"`
	// The ID of the VPC endpoint that the agent has access to.
	VpcEndpointId pulumi.StringPtrOutput `pulumi:"vpcEndpointId"`
}

// NewAgent registers a new resource with the given unique name, arguments, and options.
func NewAgent(ctx *pulumi.Context,
	name string, args *AgentArgs, opts ...pulumi.ResourceOption) (*Agent, error) {
	if args == nil {
		args = &AgentArgs{}
	}

	var resource Agent
	err := ctx.RegisterResource("aws-native:datasync:Agent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgent gets an existing Agent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentState, opts ...pulumi.ResourceOption) (*Agent, error) {
	var resource Agent
	err := ctx.ReadResource("aws-native:datasync:Agent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agent resources.
type agentState struct {
}

type AgentState struct {
}

func (AgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentState)(nil)).Elem()
}

type agentArgs struct {
	// Activation key of the Agent.
	ActivationKey *string `pulumi:"activationKey"`
	// The name configured for the agent. Text reference used to identify the agent in the console.
	AgentName *string `pulumi:"agentName"`
	// The ARNs of the security group used to protect your data transfer task subnets.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	SubnetArns []string `pulumi:"subnetArns"`
	// An array of key-value pairs to apply to this resource.
	Tags []AgentTag `pulumi:"tags"`
	// The ID of the VPC endpoint that the agent has access to.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
}

// The set of arguments for constructing a Agent resource.
type AgentArgs struct {
	// Activation key of the Agent.
	ActivationKey pulumi.StringPtrInput
	// The name configured for the agent. Text reference used to identify the agent in the console.
	AgentName pulumi.StringPtrInput
	// The ARNs of the security group used to protect your data transfer task subnets.
	SecurityGroupArns pulumi.StringArrayInput
	// The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
	SubnetArns pulumi.StringArrayInput
	// An array of key-value pairs to apply to this resource.
	Tags AgentTagArrayInput
	// The ID of the VPC endpoint that the agent has access to.
	VpcEndpointId pulumi.StringPtrInput
}

func (AgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentArgs)(nil)).Elem()
}

type AgentInput interface {
	pulumi.Input

	ToAgentOutput() AgentOutput
	ToAgentOutputWithContext(ctx context.Context) AgentOutput
}

func (*Agent) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (i *Agent) ToAgentOutput() AgentOutput {
	return i.ToAgentOutputWithContext(context.Background())
}

func (i *Agent) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentOutput)
}

type AgentOutput struct{ *pulumi.OutputState }

func (AgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Agent)(nil)).Elem()
}

func (o AgentOutput) ToAgentOutput() AgentOutput {
	return o
}

func (o AgentOutput) ToAgentOutputWithContext(ctx context.Context) AgentOutput {
	return o
}

// Activation key of the Agent.
func (o AgentOutput) ActivationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringPtrOutput { return v.ActivationKey }).(pulumi.StringPtrOutput)
}

// The DataSync Agent ARN.
func (o AgentOutput) AgentArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringOutput { return v.AgentArn }).(pulumi.StringOutput)
}

// The name configured for the agent. Text reference used to identify the agent in the console.
func (o AgentOutput) AgentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringPtrOutput { return v.AgentName }).(pulumi.StringPtrOutput)
}

// The service endpoints that the agent will connect to.
func (o AgentOutput) EndpointType() AgentEndpointTypeOutput {
	return o.ApplyT(func(v *Agent) AgentEndpointTypeOutput { return v.EndpointType }).(AgentEndpointTypeOutput)
}

// The ARNs of the security group used to protect your data transfer task subnets.
func (o AgentOutput) SecurityGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringArrayOutput { return v.SecurityGroupArns }).(pulumi.StringArrayOutput)
}

// The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
func (o AgentOutput) SubnetArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringArrayOutput { return v.SubnetArns }).(pulumi.StringArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o AgentOutput) Tags() AgentTagArrayOutput {
	return o.ApplyT(func(v *Agent) AgentTagArrayOutput { return v.Tags }).(AgentTagArrayOutput)
}

// The ID of the VPC endpoint that the agent has access to.
func (o AgentOutput) VpcEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Agent) pulumi.StringPtrOutput { return v.VpcEndpointId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AgentInput)(nil)).Elem(), &Agent{})
	pulumi.RegisterOutputType(AgentOutput{})
}
