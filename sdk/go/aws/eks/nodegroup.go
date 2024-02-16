// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EKS::Nodegroup
type Nodegroup struct {
	pulumi.CustomResourceState

	// The AMI type for your node group.
	AmiType pulumi.StringPtrOutput `pulumi:"amiType"`
	Arn     pulumi.StringOutput    `pulumi:"arn"`
	// The capacity type of your managed node group.
	CapacityType pulumi.StringPtrOutput `pulumi:"capacityType"`
	// Name of the cluster to create the node group in.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The root device disk size (in GiB) for your node group instances.
	DiskSize pulumi.IntPtrOutput `pulumi:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateEnabled pulumi.BoolPtrOutput `pulumi:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// An object representing a node group's launch template specification.
	LaunchTemplate NodegroupLaunchTemplateSpecificationPtrOutput `pulumi:"launchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	NodeRole pulumi.StringOutput `pulumi:"nodeRole"`
	// The unique name to give your node group.
	NodegroupName pulumi.StringPtrOutput `pulumi:"nodegroupName"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	ReleaseVersion pulumi.StringPtrOutput `pulumi:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess NodegroupRemoteAccessPtrOutput `pulumi:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig NodegroupScalingConfigPtrOutput `pulumi:"scalingConfig"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	Subnets pulumi.StringArrayOutput `pulumi:"subnets"`
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints NodegroupTaintArrayOutput `pulumi:"taints"`
	// The node group update configuration.
	UpdateConfig NodegroupUpdateConfigPtrOutput `pulumi:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewNodegroup registers a new resource with the given unique name, arguments, and options.
func NewNodegroup(ctx *pulumi.Context,
	name string, args *NodegroupArgs, opts ...pulumi.ResourceOption) (*Nodegroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.NodeRole == nil {
		return nil, errors.New("invalid value for required argument 'NodeRole'")
	}
	if args.Subnets == nil {
		return nil, errors.New("invalid value for required argument 'Subnets'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"amiType",
		"capacityType",
		"clusterName",
		"diskSize",
		"instanceTypes[*]",
		"nodeRole",
		"nodegroupName",
		"remoteAccess",
		"subnets[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nodegroup
	err := ctx.RegisterResource("aws-native:eks:Nodegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodegroup gets an existing Nodegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodegroupState, opts ...pulumi.ResourceOption) (*Nodegroup, error) {
	var resource Nodegroup
	err := ctx.ReadResource("aws-native:eks:Nodegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nodegroup resources.
type nodegroupState struct {
}

type NodegroupState struct {
}

func (NodegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodegroupState)(nil)).Elem()
}

type nodegroupArgs struct {
	// The AMI type for your node group.
	AmiType *string `pulumi:"amiType"`
	// The capacity type of your managed node group.
	CapacityType *string `pulumi:"capacityType"`
	// Name of the cluster to create the node group in.
	ClusterName string `pulumi:"clusterName"`
	// The root device disk size (in GiB) for your node group instances.
	DiskSize *int `pulumi:"diskSize"`
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateEnabled *bool `pulumi:"forceUpdateEnabled"`
	// Specify the instance types for a node group.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels map[string]string `pulumi:"labels"`
	// An object representing a node group's launch template specification.
	LaunchTemplate *NodegroupLaunchTemplateSpecification `pulumi:"launchTemplate"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	NodeRole string `pulumi:"nodeRole"`
	// The unique name to give your node group.
	NodegroupName *string `pulumi:"nodegroupName"`
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess *NodegroupRemoteAccess `pulumi:"remoteAccess"`
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig *NodegroupScalingConfig `pulumi:"scalingConfig"`
	// The subnets to use for the Auto Scaling group that is created for your node group.
	Subnets []string `pulumi:"subnets"`
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
	Tags map[string]string `pulumi:"tags"`
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints []NodegroupTaint `pulumi:"taints"`
	// The node group update configuration.
	UpdateConfig *NodegroupUpdateConfig `pulumi:"updateConfig"`
	// The Kubernetes version to use for your managed nodes.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Nodegroup resource.
type NodegroupArgs struct {
	// The AMI type for your node group.
	AmiType pulumi.StringPtrInput
	// The capacity type of your managed node group.
	CapacityType pulumi.StringPtrInput
	// Name of the cluster to create the node group in.
	ClusterName pulumi.StringInput
	// The root device disk size (in GiB) for your node group instances.
	DiskSize pulumi.IntPtrInput
	// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
	ForceUpdateEnabled pulumi.BoolPtrInput
	// Specify the instance types for a node group.
	InstanceTypes pulumi.StringArrayInput
	// The Kubernetes labels to be applied to the nodes in the node group when they are created.
	Labels pulumi.StringMapInput
	// An object representing a node group's launch template specification.
	LaunchTemplate NodegroupLaunchTemplateSpecificationPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
	NodeRole pulumi.StringInput
	// The unique name to give your node group.
	NodegroupName pulumi.StringPtrInput
	// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
	ReleaseVersion pulumi.StringPtrInput
	// The remote access (SSH) configuration to use with your node group.
	RemoteAccess NodegroupRemoteAccessPtrInput
	// The scaling configuration details for the Auto Scaling group that is created for your node group.
	ScalingConfig NodegroupScalingConfigPtrInput
	// The subnets to use for the Auto Scaling group that is created for your node group.
	Subnets pulumi.StringArrayInput
	// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
	Tags pulumi.StringMapInput
	// The Kubernetes taints to be applied to the nodes in the node group when they are created.
	Taints NodegroupTaintArrayInput
	// The node group update configuration.
	UpdateConfig NodegroupUpdateConfigPtrInput
	// The Kubernetes version to use for your managed nodes.
	Version pulumi.StringPtrInput
}

func (NodegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodegroupArgs)(nil)).Elem()
}

type NodegroupInput interface {
	pulumi.Input

	ToNodegroupOutput() NodegroupOutput
	ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput
}

func (*Nodegroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Nodegroup)(nil)).Elem()
}

func (i *Nodegroup) ToNodegroupOutput() NodegroupOutput {
	return i.ToNodegroupOutputWithContext(context.Background())
}

func (i *Nodegroup) ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodegroupOutput)
}

type NodegroupOutput struct{ *pulumi.OutputState }

func (NodegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nodegroup)(nil)).Elem()
}

func (o NodegroupOutput) ToNodegroupOutput() NodegroupOutput {
	return o
}

func (o NodegroupOutput) ToNodegroupOutputWithContext(ctx context.Context) NodegroupOutput {
	return o
}

// The AMI type for your node group.
func (o NodegroupOutput) AmiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringPtrOutput { return v.AmiType }).(pulumi.StringPtrOutput)
}

func (o NodegroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The capacity type of your managed node group.
func (o NodegroupOutput) CapacityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringPtrOutput { return v.CapacityType }).(pulumi.StringPtrOutput)
}

// Name of the cluster to create the node group in.
func (o NodegroupOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// The root device disk size (in GiB) for your node group instances.
func (o NodegroupOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.IntPtrOutput { return v.DiskSize }).(pulumi.IntPtrOutput)
}

// Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.
func (o NodegroupOutput) ForceUpdateEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.BoolPtrOutput { return v.ForceUpdateEnabled }).(pulumi.BoolPtrOutput)
}

// Specify the instance types for a node group.
func (o NodegroupOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringArrayOutput { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

// The Kubernetes labels to be applied to the nodes in the node group when they are created.
func (o NodegroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// An object representing a node group's launch template specification.
func (o NodegroupOutput) LaunchTemplate() NodegroupLaunchTemplateSpecificationPtrOutput {
	return o.ApplyT(func(v *Nodegroup) NodegroupLaunchTemplateSpecificationPtrOutput { return v.LaunchTemplate }).(NodegroupLaunchTemplateSpecificationPtrOutput)
}

// The Amazon Resource Name (ARN) of the IAM role to associate with your node group.
func (o NodegroupOutput) NodeRole() pulumi.StringOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringOutput { return v.NodeRole }).(pulumi.StringOutput)
}

// The unique name to give your node group.
func (o NodegroupOutput) NodegroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringPtrOutput { return v.NodegroupName }).(pulumi.StringPtrOutput)
}

// The AMI version of the Amazon EKS-optimized AMI to use with your node group.
func (o NodegroupOutput) ReleaseVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringPtrOutput { return v.ReleaseVersion }).(pulumi.StringPtrOutput)
}

// The remote access (SSH) configuration to use with your node group.
func (o NodegroupOutput) RemoteAccess() NodegroupRemoteAccessPtrOutput {
	return o.ApplyT(func(v *Nodegroup) NodegroupRemoteAccessPtrOutput { return v.RemoteAccess }).(NodegroupRemoteAccessPtrOutput)
}

// The scaling configuration details for the Auto Scaling group that is created for your node group.
func (o NodegroupOutput) ScalingConfig() NodegroupScalingConfigPtrOutput {
	return o.ApplyT(func(v *Nodegroup) NodegroupScalingConfigPtrOutput { return v.ScalingConfig }).(NodegroupScalingConfigPtrOutput)
}

// The subnets to use for the Auto Scaling group that is created for your node group.
func (o NodegroupOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringArrayOutput { return v.Subnets }).(pulumi.StringArrayOutput)
}

// The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.
func (o NodegroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The Kubernetes taints to be applied to the nodes in the node group when they are created.
func (o NodegroupOutput) Taints() NodegroupTaintArrayOutput {
	return o.ApplyT(func(v *Nodegroup) NodegroupTaintArrayOutput { return v.Taints }).(NodegroupTaintArrayOutput)
}

// The node group update configuration.
func (o NodegroupOutput) UpdateConfig() NodegroupUpdateConfigPtrOutput {
	return o.ApplyT(func(v *Nodegroup) NodegroupUpdateConfigPtrOutput { return v.UpdateConfig }).(NodegroupUpdateConfigPtrOutput)
}

// The Kubernetes version to use for your managed nodes.
func (o NodegroupOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nodegroup) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodegroupInput)(nil)).Elem(), &Nodegroup{})
	pulumi.RegisterOutputType(NodegroupOutput{})
}
