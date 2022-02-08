// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::Instance
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("aws-native:ec2:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	Id string `pulumi:"id"`
}

type LookupInstanceResult struct {
	AdditionalInfo                    *string                      `pulumi:"additionalInfo"`
	Affinity                          *string                      `pulumi:"affinity"`
	BlockDeviceMappings               []InstanceBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	CreditSpecification               *InstanceCreditSpecification `pulumi:"creditSpecification"`
	DisableApiTermination             *bool                        `pulumi:"disableApiTermination"`
	EbsOptimized                      *bool                        `pulumi:"ebsOptimized"`
	HostId                            *string                      `pulumi:"hostId"`
	IamInstanceProfile                *string                      `pulumi:"iamInstanceProfile"`
	Id                                *string                      `pulumi:"id"`
	InstanceInitiatedShutdownBehavior *string                      `pulumi:"instanceInitiatedShutdownBehavior"`
	InstanceType                      *string                      `pulumi:"instanceType"`
	KernelId                          *string                      `pulumi:"kernelId"`
	Monitoring                        *bool                        `pulumi:"monitoring"`
	PrivateDnsName                    *string                      `pulumi:"privateDnsName"`
	PrivateIp                         *string                      `pulumi:"privateIp"`
	PropagateTagsToVolumeOnCreation   *bool                        `pulumi:"propagateTagsToVolumeOnCreation"`
	PublicDnsName                     *string                      `pulumi:"publicDnsName"`
	PublicIp                          *string                      `pulumi:"publicIp"`
	RamdiskId                         *string                      `pulumi:"ramdiskId"`
	SecurityGroupIds                  []string                     `pulumi:"securityGroupIds"`
	SourceDestCheck                   *bool                        `pulumi:"sourceDestCheck"`
	SsmAssociations                   []InstanceSsmAssociation     `pulumi:"ssmAssociations"`
	Tags                              []InstanceTag                `pulumi:"tags"`
	Tenancy                           *string                      `pulumi:"tenancy"`
	UserData                          *string                      `pulumi:"userData"`
	Volumes                           []InstanceVolume             `pulumi:"volumes"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) AdditionalInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.AdditionalInfo }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) Affinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Affinity }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) BlockDeviceMappings() InstanceBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []InstanceBlockDeviceMapping { return v.BlockDeviceMappings }).(InstanceBlockDeviceMappingArrayOutput)
}

func (o LookupInstanceResultOutput) CreditSpecification() InstanceCreditSpecificationPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *InstanceCreditSpecification { return v.CreditSpecification }).(InstanceCreditSpecificationPtrOutput)
}

func (o LookupInstanceResultOutput) DisableApiTermination() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.DisableApiTermination }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceResultOutput) EbsOptimized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.EbsOptimized }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceResultOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) IamInstanceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.IamInstanceProfile }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) InstanceInitiatedShutdownBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.InstanceInitiatedShutdownBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) KernelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.KernelId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) Monitoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.Monitoring }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceResultOutput) PrivateDnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.PrivateDnsName }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) PrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.PrivateIp }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) PropagateTagsToVolumeOnCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.PropagateTagsToVolumeOnCreation }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceResultOutput) PublicDnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.PublicDnsName }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) PublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.PublicIp }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) RamdiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.RamdiskId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceResultOutput) SourceDestCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.SourceDestCheck }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceResultOutput) SsmAssociations() InstanceSsmAssociationArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []InstanceSsmAssociation { return v.SsmAssociations }).(InstanceSsmAssociationArrayOutput)
}

func (o LookupInstanceResultOutput) Tags() InstanceTagArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []InstanceTag { return v.Tags }).(InstanceTagArrayOutput)
}

func (o LookupInstanceResultOutput) Tenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.Tenancy }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceResultOutput) Volumes() InstanceVolumeArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []InstanceVolume { return v.Volumes }).(InstanceVolumeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
