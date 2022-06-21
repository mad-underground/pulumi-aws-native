// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshiftserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NamespaceType struct {
	AdminUsername     *string              `pulumi:"adminUsername"`
	CreationDate      *string              `pulumi:"creationDate"`
	DbName            *string              `pulumi:"dbName"`
	DefaultIamRoleArn *string              `pulumi:"defaultIamRoleArn"`
	IamRoles          []string             `pulumi:"iamRoles"`
	KmsKeyId          *string              `pulumi:"kmsKeyId"`
	LogExports        []NamespaceLogExport `pulumi:"logExports"`
	NamespaceArn      *string              `pulumi:"namespaceArn"`
	NamespaceId       *string              `pulumi:"namespaceId"`
	NamespaceName     *string              `pulumi:"namespaceName"`
	Status            *NamespaceStatus     `pulumi:"status"`
}

type NamespaceTypeOutput struct{ *pulumi.OutputState }

func (NamespaceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceType)(nil)).Elem()
}

func (o NamespaceTypeOutput) ToNamespaceTypeOutput() NamespaceTypeOutput {
	return o
}

func (o NamespaceTypeOutput) ToNamespaceTypeOutputWithContext(ctx context.Context) NamespaceTypeOutput {
	return o
}

func (o NamespaceTypeOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) DbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.DbName }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) DefaultIamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.DefaultIamRoleArn }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) IamRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamespaceType) []string { return v.IamRoles }).(pulumi.StringArrayOutput)
}

func (o NamespaceTypeOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) LogExports() NamespaceLogExportArrayOutput {
	return o.ApplyT(func(v NamespaceType) []NamespaceLogExport { return v.LogExports }).(NamespaceLogExportArrayOutput)
}

func (o NamespaceTypeOutput) NamespaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.NamespaceArn }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.NamespaceId }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespaceType) *string { return v.NamespaceName }).(pulumi.StringPtrOutput)
}

func (o NamespaceTypeOutput) Status() NamespaceStatusPtrOutput {
	return o.ApplyT(func(v NamespaceType) *NamespaceStatus { return v.Status }).(NamespaceStatusPtrOutput)
}

type NamespaceTypePtrOutput struct{ *pulumi.OutputState }

func (NamespaceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceType)(nil)).Elem()
}

func (o NamespaceTypePtrOutput) ToNamespaceTypePtrOutput() NamespaceTypePtrOutput {
	return o
}

func (o NamespaceTypePtrOutput) ToNamespaceTypePtrOutputWithContext(ctx context.Context) NamespaceTypePtrOutput {
	return o
}

func (o NamespaceTypePtrOutput) Elem() NamespaceTypeOutput {
	return o.ApplyT(func(v *NamespaceType) NamespaceType {
		if v != nil {
			return *v
		}
		var ret NamespaceType
		return ret
	}).(NamespaceTypeOutput)
}

func (o NamespaceTypePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.CreationDate
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) DbName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.DbName
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) DefaultIamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.DefaultIamRoleArn
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) IamRoles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceType) []string {
		if v == nil {
			return nil
		}
		return v.IamRoles
	}).(pulumi.StringArrayOutput)
}

func (o NamespaceTypePtrOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.KmsKeyId
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) LogExports() NamespaceLogExportArrayOutput {
	return o.ApplyT(func(v *NamespaceType) []NamespaceLogExport {
		if v == nil {
			return nil
		}
		return v.LogExports
	}).(NamespaceLogExportArrayOutput)
}

func (o NamespaceTypePtrOutput) NamespaceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceArn
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) NamespaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceId
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) NamespaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *string {
		if v == nil {
			return nil
		}
		return v.NamespaceName
	}).(pulumi.StringPtrOutput)
}

func (o NamespaceTypePtrOutput) Status() NamespaceStatusPtrOutput {
	return o.ApplyT(func(v *NamespaceType) *NamespaceStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(NamespaceStatusPtrOutput)
}

type NamespaceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// NamespaceTagInput is an input type that accepts NamespaceTagArgs and NamespaceTagOutput values.
// You can construct a concrete instance of `NamespaceTagInput` via:
//
//          NamespaceTagArgs{...}
type NamespaceTagInput interface {
	pulumi.Input

	ToNamespaceTagOutput() NamespaceTagOutput
	ToNamespaceTagOutputWithContext(context.Context) NamespaceTagOutput
}

type NamespaceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (NamespaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceTag)(nil)).Elem()
}

func (i NamespaceTagArgs) ToNamespaceTagOutput() NamespaceTagOutput {
	return i.ToNamespaceTagOutputWithContext(context.Background())
}

func (i NamespaceTagArgs) ToNamespaceTagOutputWithContext(ctx context.Context) NamespaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceTagOutput)
}

// NamespaceTagArrayInput is an input type that accepts NamespaceTagArray and NamespaceTagArrayOutput values.
// You can construct a concrete instance of `NamespaceTagArrayInput` via:
//
//          NamespaceTagArray{ NamespaceTagArgs{...} }
type NamespaceTagArrayInput interface {
	pulumi.Input

	ToNamespaceTagArrayOutput() NamespaceTagArrayOutput
	ToNamespaceTagArrayOutputWithContext(context.Context) NamespaceTagArrayOutput
}

type NamespaceTagArray []NamespaceTagInput

func (NamespaceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceTag)(nil)).Elem()
}

func (i NamespaceTagArray) ToNamespaceTagArrayOutput() NamespaceTagArrayOutput {
	return i.ToNamespaceTagArrayOutputWithContext(context.Background())
}

func (i NamespaceTagArray) ToNamespaceTagArrayOutputWithContext(ctx context.Context) NamespaceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceTagArrayOutput)
}

type NamespaceTagOutput struct{ *pulumi.OutputState }

func (NamespaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceTag)(nil)).Elem()
}

func (o NamespaceTagOutput) ToNamespaceTagOutput() NamespaceTagOutput {
	return o
}

func (o NamespaceTagOutput) ToNamespaceTagOutputWithContext(ctx context.Context) NamespaceTagOutput {
	return o
}

func (o NamespaceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v NamespaceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o NamespaceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v NamespaceTag) string { return v.Value }).(pulumi.StringOutput)
}

type NamespaceTagArrayOutput struct{ *pulumi.OutputState }

func (NamespaceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceTag)(nil)).Elem()
}

func (o NamespaceTagArrayOutput) ToNamespaceTagArrayOutput() NamespaceTagArrayOutput {
	return o
}

func (o NamespaceTagArrayOutput) ToNamespaceTagArrayOutputWithContext(ctx context.Context) NamespaceTagArrayOutput {
	return o
}

func (o NamespaceTagArrayOutput) Index(i pulumi.IntInput) NamespaceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceTag {
		return vs[0].([]NamespaceTag)[vs[1].(int)]
	}).(NamespaceTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceTagInput)(nil)).Elem(), NamespaceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NamespaceTagArrayInput)(nil)).Elem(), NamespaceTagArray{})
	pulumi.RegisterOutputType(NamespaceTypeOutput{})
	pulumi.RegisterOutputType(NamespaceTypePtrOutput{})
	pulumi.RegisterOutputType(NamespaceTagOutput{})
	pulumi.RegisterOutputType(NamespaceTagArrayOutput{})
}