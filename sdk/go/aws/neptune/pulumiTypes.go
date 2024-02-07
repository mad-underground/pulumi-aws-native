// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
type DbClusterDbClusterRole struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
	FeatureName *string `pulumi:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn string `pulumi:"roleArn"`
}

// DbClusterDbClusterRoleInput is an input type that accepts DbClusterDbClusterRoleArgs and DbClusterDbClusterRoleOutput values.
// You can construct a concrete instance of `DbClusterDbClusterRoleInput` via:
//
//	DbClusterDbClusterRoleArgs{...}
type DbClusterDbClusterRoleInput interface {
	pulumi.Input

	ToDbClusterDbClusterRoleOutput() DbClusterDbClusterRoleOutput
	ToDbClusterDbClusterRoleOutputWithContext(context.Context) DbClusterDbClusterRoleOutput
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
type DbClusterDbClusterRoleArgs struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
	FeatureName pulumi.StringPtrInput `pulumi:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
}

func (DbClusterDbClusterRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterDbClusterRole)(nil)).Elem()
}

func (i DbClusterDbClusterRoleArgs) ToDbClusterDbClusterRoleOutput() DbClusterDbClusterRoleOutput {
	return i.ToDbClusterDbClusterRoleOutputWithContext(context.Background())
}

func (i DbClusterDbClusterRoleArgs) ToDbClusterDbClusterRoleOutputWithContext(ctx context.Context) DbClusterDbClusterRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterDbClusterRoleOutput)
}

// DbClusterDbClusterRoleArrayInput is an input type that accepts DbClusterDbClusterRoleArray and DbClusterDbClusterRoleArrayOutput values.
// You can construct a concrete instance of `DbClusterDbClusterRoleArrayInput` via:
//
//	DbClusterDbClusterRoleArray{ DbClusterDbClusterRoleArgs{...} }
type DbClusterDbClusterRoleArrayInput interface {
	pulumi.Input

	ToDbClusterDbClusterRoleArrayOutput() DbClusterDbClusterRoleArrayOutput
	ToDbClusterDbClusterRoleArrayOutputWithContext(context.Context) DbClusterDbClusterRoleArrayOutput
}

type DbClusterDbClusterRoleArray []DbClusterDbClusterRoleInput

func (DbClusterDbClusterRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbClusterDbClusterRole)(nil)).Elem()
}

func (i DbClusterDbClusterRoleArray) ToDbClusterDbClusterRoleArrayOutput() DbClusterDbClusterRoleArrayOutput {
	return i.ToDbClusterDbClusterRoleArrayOutputWithContext(context.Background())
}

func (i DbClusterDbClusterRoleArray) ToDbClusterDbClusterRoleArrayOutputWithContext(ctx context.Context) DbClusterDbClusterRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterDbClusterRoleArrayOutput)
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
type DbClusterDbClusterRoleOutput struct{ *pulumi.OutputState }

func (DbClusterDbClusterRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterDbClusterRole)(nil)).Elem()
}

func (o DbClusterDbClusterRoleOutput) ToDbClusterDbClusterRoleOutput() DbClusterDbClusterRoleOutput {
	return o
}

func (o DbClusterDbClusterRoleOutput) ToDbClusterDbClusterRoleOutputWithContext(ctx context.Context) DbClusterDbClusterRoleOutput {
	return o
}

// The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
func (o DbClusterDbClusterRoleOutput) FeatureName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DbClusterDbClusterRole) *string { return v.FeatureName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
func (o DbClusterDbClusterRoleOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v DbClusterDbClusterRole) string { return v.RoleArn }).(pulumi.StringOutput)
}

type DbClusterDbClusterRoleArrayOutput struct{ *pulumi.OutputState }

func (DbClusterDbClusterRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbClusterDbClusterRole)(nil)).Elem()
}

func (o DbClusterDbClusterRoleArrayOutput) ToDbClusterDbClusterRoleArrayOutput() DbClusterDbClusterRoleArrayOutput {
	return o
}

func (o DbClusterDbClusterRoleArrayOutput) ToDbClusterDbClusterRoleArrayOutputWithContext(ctx context.Context) DbClusterDbClusterRoleArrayOutput {
	return o
}

func (o DbClusterDbClusterRoleArrayOutput) Index(i pulumi.IntInput) DbClusterDbClusterRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbClusterDbClusterRole {
		return vs[0].([]DbClusterDbClusterRole)[vs[1].(int)]
	}).(DbClusterDbClusterRoleOutput)
}

type DbClusterParameterGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DbClusterParameterGroupTagInput is an input type that accepts DbClusterParameterGroupTagArgs and DbClusterParameterGroupTagOutput values.
// You can construct a concrete instance of `DbClusterParameterGroupTagInput` via:
//
//	DbClusterParameterGroupTagArgs{...}
type DbClusterParameterGroupTagInput interface {
	pulumi.Input

	ToDbClusterParameterGroupTagOutput() DbClusterParameterGroupTagOutput
	ToDbClusterParameterGroupTagOutputWithContext(context.Context) DbClusterParameterGroupTagOutput
}

type DbClusterParameterGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DbClusterParameterGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterParameterGroupTag)(nil)).Elem()
}

func (i DbClusterParameterGroupTagArgs) ToDbClusterParameterGroupTagOutput() DbClusterParameterGroupTagOutput {
	return i.ToDbClusterParameterGroupTagOutputWithContext(context.Background())
}

func (i DbClusterParameterGroupTagArgs) ToDbClusterParameterGroupTagOutputWithContext(ctx context.Context) DbClusterParameterGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterParameterGroupTagOutput)
}

// DbClusterParameterGroupTagArrayInput is an input type that accepts DbClusterParameterGroupTagArray and DbClusterParameterGroupTagArrayOutput values.
// You can construct a concrete instance of `DbClusterParameterGroupTagArrayInput` via:
//
//	DbClusterParameterGroupTagArray{ DbClusterParameterGroupTagArgs{...} }
type DbClusterParameterGroupTagArrayInput interface {
	pulumi.Input

	ToDbClusterParameterGroupTagArrayOutput() DbClusterParameterGroupTagArrayOutput
	ToDbClusterParameterGroupTagArrayOutputWithContext(context.Context) DbClusterParameterGroupTagArrayOutput
}

type DbClusterParameterGroupTagArray []DbClusterParameterGroupTagInput

func (DbClusterParameterGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbClusterParameterGroupTag)(nil)).Elem()
}

func (i DbClusterParameterGroupTagArray) ToDbClusterParameterGroupTagArrayOutput() DbClusterParameterGroupTagArrayOutput {
	return i.ToDbClusterParameterGroupTagArrayOutputWithContext(context.Background())
}

func (i DbClusterParameterGroupTagArray) ToDbClusterParameterGroupTagArrayOutputWithContext(ctx context.Context) DbClusterParameterGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterParameterGroupTagArrayOutput)
}

type DbClusterParameterGroupTagOutput struct{ *pulumi.OutputState }

func (DbClusterParameterGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterParameterGroupTag)(nil)).Elem()
}

func (o DbClusterParameterGroupTagOutput) ToDbClusterParameterGroupTagOutput() DbClusterParameterGroupTagOutput {
	return o
}

func (o DbClusterParameterGroupTagOutput) ToDbClusterParameterGroupTagOutputWithContext(ctx context.Context) DbClusterParameterGroupTagOutput {
	return o
}

func (o DbClusterParameterGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DbClusterParameterGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DbClusterParameterGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DbClusterParameterGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DbClusterParameterGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DbClusterParameterGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbClusterParameterGroupTag)(nil)).Elem()
}

func (o DbClusterParameterGroupTagArrayOutput) ToDbClusterParameterGroupTagArrayOutput() DbClusterParameterGroupTagArrayOutput {
	return o
}

func (o DbClusterParameterGroupTagArrayOutput) ToDbClusterParameterGroupTagArrayOutputWithContext(ctx context.Context) DbClusterParameterGroupTagArrayOutput {
	return o
}

func (o DbClusterParameterGroupTagArrayOutput) Index(i pulumi.IntInput) DbClusterParameterGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbClusterParameterGroupTag {
		return vs[0].([]DbClusterParameterGroupTag)[vs[1].(int)]
	}).(DbClusterParameterGroupTagOutput)
}

// Contains the scaling configuration of an Neptune Serverless DB cluster.
type DbClusterServerlessScalingConfiguration struct {
	// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
	MaxCapacity float64 `pulumi:"maxCapacity"`
	// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
	MinCapacity float64 `pulumi:"minCapacity"`
}

// DbClusterServerlessScalingConfigurationInput is an input type that accepts DbClusterServerlessScalingConfigurationArgs and DbClusterServerlessScalingConfigurationOutput values.
// You can construct a concrete instance of `DbClusterServerlessScalingConfigurationInput` via:
//
//	DbClusterServerlessScalingConfigurationArgs{...}
type DbClusterServerlessScalingConfigurationInput interface {
	pulumi.Input

	ToDbClusterServerlessScalingConfigurationOutput() DbClusterServerlessScalingConfigurationOutput
	ToDbClusterServerlessScalingConfigurationOutputWithContext(context.Context) DbClusterServerlessScalingConfigurationOutput
}

// Contains the scaling configuration of an Neptune Serverless DB cluster.
type DbClusterServerlessScalingConfigurationArgs struct {
	// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
	MaxCapacity pulumi.Float64Input `pulumi:"maxCapacity"`
	// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
	MinCapacity pulumi.Float64Input `pulumi:"minCapacity"`
}

func (DbClusterServerlessScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (i DbClusterServerlessScalingConfigurationArgs) ToDbClusterServerlessScalingConfigurationOutput() DbClusterServerlessScalingConfigurationOutput {
	return i.ToDbClusterServerlessScalingConfigurationOutputWithContext(context.Background())
}

func (i DbClusterServerlessScalingConfigurationArgs) ToDbClusterServerlessScalingConfigurationOutputWithContext(ctx context.Context) DbClusterServerlessScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterServerlessScalingConfigurationOutput)
}

func (i DbClusterServerlessScalingConfigurationArgs) ToDbClusterServerlessScalingConfigurationPtrOutput() DbClusterServerlessScalingConfigurationPtrOutput {
	return i.ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i DbClusterServerlessScalingConfigurationArgs) ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DbClusterServerlessScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterServerlessScalingConfigurationOutput).ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(ctx)
}

// DbClusterServerlessScalingConfigurationPtrInput is an input type that accepts DbClusterServerlessScalingConfigurationArgs, DbClusterServerlessScalingConfigurationPtr and DbClusterServerlessScalingConfigurationPtrOutput values.
// You can construct a concrete instance of `DbClusterServerlessScalingConfigurationPtrInput` via:
//
//	        DbClusterServerlessScalingConfigurationArgs{...}
//
//	or:
//
//	        nil
type DbClusterServerlessScalingConfigurationPtrInput interface {
	pulumi.Input

	ToDbClusterServerlessScalingConfigurationPtrOutput() DbClusterServerlessScalingConfigurationPtrOutput
	ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(context.Context) DbClusterServerlessScalingConfigurationPtrOutput
}

type dbClusterServerlessScalingConfigurationPtrType DbClusterServerlessScalingConfigurationArgs

func DbClusterServerlessScalingConfigurationPtr(v *DbClusterServerlessScalingConfigurationArgs) DbClusterServerlessScalingConfigurationPtrInput {
	return (*dbClusterServerlessScalingConfigurationPtrType)(v)
}

func (*dbClusterServerlessScalingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DbClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (i *dbClusterServerlessScalingConfigurationPtrType) ToDbClusterServerlessScalingConfigurationPtrOutput() DbClusterServerlessScalingConfigurationPtrOutput {
	return i.ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i *dbClusterServerlessScalingConfigurationPtrType) ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DbClusterServerlessScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterServerlessScalingConfigurationPtrOutput)
}

// Contains the scaling configuration of an Neptune Serverless DB cluster.
type DbClusterServerlessScalingConfigurationOutput struct{ *pulumi.OutputState }

func (DbClusterServerlessScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (o DbClusterServerlessScalingConfigurationOutput) ToDbClusterServerlessScalingConfigurationOutput() DbClusterServerlessScalingConfigurationOutput {
	return o
}

func (o DbClusterServerlessScalingConfigurationOutput) ToDbClusterServerlessScalingConfigurationOutputWithContext(ctx context.Context) DbClusterServerlessScalingConfigurationOutput {
	return o
}

func (o DbClusterServerlessScalingConfigurationOutput) ToDbClusterServerlessScalingConfigurationPtrOutput() DbClusterServerlessScalingConfigurationPtrOutput {
	return o.ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(context.Background())
}

func (o DbClusterServerlessScalingConfigurationOutput) ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DbClusterServerlessScalingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DbClusterServerlessScalingConfiguration) *DbClusterServerlessScalingConfiguration {
		return &v
	}).(DbClusterServerlessScalingConfigurationPtrOutput)
}

// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
func (o DbClusterServerlessScalingConfigurationOutput) MaxCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v DbClusterServerlessScalingConfiguration) float64 { return v.MaxCapacity }).(pulumi.Float64Output)
}

// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
func (o DbClusterServerlessScalingConfigurationOutput) MinCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v DbClusterServerlessScalingConfiguration) float64 { return v.MinCapacity }).(pulumi.Float64Output)
}

type DbClusterServerlessScalingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (DbClusterServerlessScalingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (o DbClusterServerlessScalingConfigurationPtrOutput) ToDbClusterServerlessScalingConfigurationPtrOutput() DbClusterServerlessScalingConfigurationPtrOutput {
	return o
}

func (o DbClusterServerlessScalingConfigurationPtrOutput) ToDbClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DbClusterServerlessScalingConfigurationPtrOutput {
	return o
}

func (o DbClusterServerlessScalingConfigurationPtrOutput) Elem() DbClusterServerlessScalingConfigurationOutput {
	return o.ApplyT(func(v *DbClusterServerlessScalingConfiguration) DbClusterServerlessScalingConfiguration {
		if v != nil {
			return *v
		}
		var ret DbClusterServerlessScalingConfiguration
		return ret
	}).(DbClusterServerlessScalingConfigurationOutput)
}

// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
func (o DbClusterServerlessScalingConfigurationPtrOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DbClusterServerlessScalingConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return &v.MaxCapacity
	}).(pulumi.Float64PtrOutput)
}

// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
func (o DbClusterServerlessScalingConfigurationPtrOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DbClusterServerlessScalingConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return &v.MinCapacity
	}).(pulumi.Float64PtrOutput)
}

// A key-value pair to associate with a resource.
type DbClusterTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value *string `pulumi:"value"`
}

// DbClusterTagInput is an input type that accepts DbClusterTagArgs and DbClusterTagOutput values.
// You can construct a concrete instance of `DbClusterTagInput` via:
//
//	DbClusterTagArgs{...}
type DbClusterTagInput interface {
	pulumi.Input

	ToDbClusterTagOutput() DbClusterTagOutput
	ToDbClusterTagOutputWithContext(context.Context) DbClusterTagOutput
}

// A key-value pair to associate with a resource.
type DbClusterTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DbClusterTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterTag)(nil)).Elem()
}

func (i DbClusterTagArgs) ToDbClusterTagOutput() DbClusterTagOutput {
	return i.ToDbClusterTagOutputWithContext(context.Background())
}

func (i DbClusterTagArgs) ToDbClusterTagOutputWithContext(ctx context.Context) DbClusterTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterTagOutput)
}

// DbClusterTagArrayInput is an input type that accepts DbClusterTagArray and DbClusterTagArrayOutput values.
// You can construct a concrete instance of `DbClusterTagArrayInput` via:
//
//	DbClusterTagArray{ DbClusterTagArgs{...} }
type DbClusterTagArrayInput interface {
	pulumi.Input

	ToDbClusterTagArrayOutput() DbClusterTagArrayOutput
	ToDbClusterTagArrayOutputWithContext(context.Context) DbClusterTagArrayOutput
}

type DbClusterTagArray []DbClusterTagInput

func (DbClusterTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbClusterTag)(nil)).Elem()
}

func (i DbClusterTagArray) ToDbClusterTagArrayOutput() DbClusterTagArrayOutput {
	return i.ToDbClusterTagArrayOutputWithContext(context.Background())
}

func (i DbClusterTagArray) ToDbClusterTagArrayOutputWithContext(ctx context.Context) DbClusterTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbClusterTagArrayOutput)
}

// A key-value pair to associate with a resource.
type DbClusterTagOutput struct{ *pulumi.OutputState }

func (DbClusterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbClusterTag)(nil)).Elem()
}

func (o DbClusterTagOutput) ToDbClusterTagOutput() DbClusterTagOutput {
	return o
}

func (o DbClusterTagOutput) ToDbClusterTagOutputWithContext(ctx context.Context) DbClusterTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o DbClusterTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DbClusterTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o DbClusterTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DbClusterTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DbClusterTagArrayOutput struct{ *pulumi.OutputState }

func (DbClusterTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbClusterTag)(nil)).Elem()
}

func (o DbClusterTagArrayOutput) ToDbClusterTagArrayOutput() DbClusterTagArrayOutput {
	return o
}

func (o DbClusterTagArrayOutput) ToDbClusterTagArrayOutputWithContext(ctx context.Context) DbClusterTagArrayOutput {
	return o
}

func (o DbClusterTagArrayOutput) Index(i pulumi.IntInput) DbClusterTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbClusterTag {
		return vs[0].([]DbClusterTag)[vs[1].(int)]
	}).(DbClusterTagOutput)
}

type DbInstanceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DbInstanceTagInput is an input type that accepts DbInstanceTagArgs and DbInstanceTagOutput values.
// You can construct a concrete instance of `DbInstanceTagInput` via:
//
//	DbInstanceTagArgs{...}
type DbInstanceTagInput interface {
	pulumi.Input

	ToDbInstanceTagOutput() DbInstanceTagOutput
	ToDbInstanceTagOutputWithContext(context.Context) DbInstanceTagOutput
}

type DbInstanceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DbInstanceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbInstanceTag)(nil)).Elem()
}

func (i DbInstanceTagArgs) ToDbInstanceTagOutput() DbInstanceTagOutput {
	return i.ToDbInstanceTagOutputWithContext(context.Background())
}

func (i DbInstanceTagArgs) ToDbInstanceTagOutputWithContext(ctx context.Context) DbInstanceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstanceTagOutput)
}

// DbInstanceTagArrayInput is an input type that accepts DbInstanceTagArray and DbInstanceTagArrayOutput values.
// You can construct a concrete instance of `DbInstanceTagArrayInput` via:
//
//	DbInstanceTagArray{ DbInstanceTagArgs{...} }
type DbInstanceTagArrayInput interface {
	pulumi.Input

	ToDbInstanceTagArrayOutput() DbInstanceTagArrayOutput
	ToDbInstanceTagArrayOutputWithContext(context.Context) DbInstanceTagArrayOutput
}

type DbInstanceTagArray []DbInstanceTagInput

func (DbInstanceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbInstanceTag)(nil)).Elem()
}

func (i DbInstanceTagArray) ToDbInstanceTagArrayOutput() DbInstanceTagArrayOutput {
	return i.ToDbInstanceTagArrayOutputWithContext(context.Background())
}

func (i DbInstanceTagArray) ToDbInstanceTagArrayOutputWithContext(ctx context.Context) DbInstanceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbInstanceTagArrayOutput)
}

type DbInstanceTagOutput struct{ *pulumi.OutputState }

func (DbInstanceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbInstanceTag)(nil)).Elem()
}

func (o DbInstanceTagOutput) ToDbInstanceTagOutput() DbInstanceTagOutput {
	return o
}

func (o DbInstanceTagOutput) ToDbInstanceTagOutputWithContext(ctx context.Context) DbInstanceTagOutput {
	return o
}

func (o DbInstanceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DbInstanceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DbInstanceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DbInstanceTag) string { return v.Value }).(pulumi.StringOutput)
}

type DbInstanceTagArrayOutput struct{ *pulumi.OutputState }

func (DbInstanceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbInstanceTag)(nil)).Elem()
}

func (o DbInstanceTagArrayOutput) ToDbInstanceTagArrayOutput() DbInstanceTagArrayOutput {
	return o
}

func (o DbInstanceTagArrayOutput) ToDbInstanceTagArrayOutputWithContext(ctx context.Context) DbInstanceTagArrayOutput {
	return o
}

func (o DbInstanceTagArrayOutput) Index(i pulumi.IntInput) DbInstanceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbInstanceTag {
		return vs[0].([]DbInstanceTag)[vs[1].(int)]
	}).(DbInstanceTagOutput)
}

type DbParameterGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DbParameterGroupTagInput is an input type that accepts DbParameterGroupTagArgs and DbParameterGroupTagOutput values.
// You can construct a concrete instance of `DbParameterGroupTagInput` via:
//
//	DbParameterGroupTagArgs{...}
type DbParameterGroupTagInput interface {
	pulumi.Input

	ToDbParameterGroupTagOutput() DbParameterGroupTagOutput
	ToDbParameterGroupTagOutputWithContext(context.Context) DbParameterGroupTagOutput
}

type DbParameterGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DbParameterGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbParameterGroupTag)(nil)).Elem()
}

func (i DbParameterGroupTagArgs) ToDbParameterGroupTagOutput() DbParameterGroupTagOutput {
	return i.ToDbParameterGroupTagOutputWithContext(context.Background())
}

func (i DbParameterGroupTagArgs) ToDbParameterGroupTagOutputWithContext(ctx context.Context) DbParameterGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbParameterGroupTagOutput)
}

// DbParameterGroupTagArrayInput is an input type that accepts DbParameterGroupTagArray and DbParameterGroupTagArrayOutput values.
// You can construct a concrete instance of `DbParameterGroupTagArrayInput` via:
//
//	DbParameterGroupTagArray{ DbParameterGroupTagArgs{...} }
type DbParameterGroupTagArrayInput interface {
	pulumi.Input

	ToDbParameterGroupTagArrayOutput() DbParameterGroupTagArrayOutput
	ToDbParameterGroupTagArrayOutputWithContext(context.Context) DbParameterGroupTagArrayOutput
}

type DbParameterGroupTagArray []DbParameterGroupTagInput

func (DbParameterGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbParameterGroupTag)(nil)).Elem()
}

func (i DbParameterGroupTagArray) ToDbParameterGroupTagArrayOutput() DbParameterGroupTagArrayOutput {
	return i.ToDbParameterGroupTagArrayOutputWithContext(context.Background())
}

func (i DbParameterGroupTagArray) ToDbParameterGroupTagArrayOutputWithContext(ctx context.Context) DbParameterGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbParameterGroupTagArrayOutput)
}

type DbParameterGroupTagOutput struct{ *pulumi.OutputState }

func (DbParameterGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbParameterGroupTag)(nil)).Elem()
}

func (o DbParameterGroupTagOutput) ToDbParameterGroupTagOutput() DbParameterGroupTagOutput {
	return o
}

func (o DbParameterGroupTagOutput) ToDbParameterGroupTagOutputWithContext(ctx context.Context) DbParameterGroupTagOutput {
	return o
}

func (o DbParameterGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DbParameterGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DbParameterGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DbParameterGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DbParameterGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DbParameterGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbParameterGroupTag)(nil)).Elem()
}

func (o DbParameterGroupTagArrayOutput) ToDbParameterGroupTagArrayOutput() DbParameterGroupTagArrayOutput {
	return o
}

func (o DbParameterGroupTagArrayOutput) ToDbParameterGroupTagArrayOutputWithContext(ctx context.Context) DbParameterGroupTagArrayOutput {
	return o
}

func (o DbParameterGroupTagArrayOutput) Index(i pulumi.IntInput) DbParameterGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbParameterGroupTag {
		return vs[0].([]DbParameterGroupTag)[vs[1].(int)]
	}).(DbParameterGroupTagOutput)
}

type DbSubnetGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DbSubnetGroupTagInput is an input type that accepts DbSubnetGroupTagArgs and DbSubnetGroupTagOutput values.
// You can construct a concrete instance of `DbSubnetGroupTagInput` via:
//
//	DbSubnetGroupTagArgs{...}
type DbSubnetGroupTagInput interface {
	pulumi.Input

	ToDbSubnetGroupTagOutput() DbSubnetGroupTagOutput
	ToDbSubnetGroupTagOutputWithContext(context.Context) DbSubnetGroupTagOutput
}

type DbSubnetGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DbSubnetGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DbSubnetGroupTag)(nil)).Elem()
}

func (i DbSubnetGroupTagArgs) ToDbSubnetGroupTagOutput() DbSubnetGroupTagOutput {
	return i.ToDbSubnetGroupTagOutputWithContext(context.Background())
}

func (i DbSubnetGroupTagArgs) ToDbSubnetGroupTagOutputWithContext(ctx context.Context) DbSubnetGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbSubnetGroupTagOutput)
}

// DbSubnetGroupTagArrayInput is an input type that accepts DbSubnetGroupTagArray and DbSubnetGroupTagArrayOutput values.
// You can construct a concrete instance of `DbSubnetGroupTagArrayInput` via:
//
//	DbSubnetGroupTagArray{ DbSubnetGroupTagArgs{...} }
type DbSubnetGroupTagArrayInput interface {
	pulumi.Input

	ToDbSubnetGroupTagArrayOutput() DbSubnetGroupTagArrayOutput
	ToDbSubnetGroupTagArrayOutputWithContext(context.Context) DbSubnetGroupTagArrayOutput
}

type DbSubnetGroupTagArray []DbSubnetGroupTagInput

func (DbSubnetGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbSubnetGroupTag)(nil)).Elem()
}

func (i DbSubnetGroupTagArray) ToDbSubnetGroupTagArrayOutput() DbSubnetGroupTagArrayOutput {
	return i.ToDbSubnetGroupTagArrayOutputWithContext(context.Background())
}

func (i DbSubnetGroupTagArray) ToDbSubnetGroupTagArrayOutputWithContext(ctx context.Context) DbSubnetGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbSubnetGroupTagArrayOutput)
}

type DbSubnetGroupTagOutput struct{ *pulumi.OutputState }

func (DbSubnetGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DbSubnetGroupTag)(nil)).Elem()
}

func (o DbSubnetGroupTagOutput) ToDbSubnetGroupTagOutput() DbSubnetGroupTagOutput {
	return o
}

func (o DbSubnetGroupTagOutput) ToDbSubnetGroupTagOutputWithContext(ctx context.Context) DbSubnetGroupTagOutput {
	return o
}

func (o DbSubnetGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DbSubnetGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DbSubnetGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DbSubnetGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DbSubnetGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DbSubnetGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DbSubnetGroupTag)(nil)).Elem()
}

func (o DbSubnetGroupTagArrayOutput) ToDbSubnetGroupTagArrayOutput() DbSubnetGroupTagArrayOutput {
	return o
}

func (o DbSubnetGroupTagArrayOutput) ToDbSubnetGroupTagArrayOutputWithContext(ctx context.Context) DbSubnetGroupTagArrayOutput {
	return o
}

func (o DbSubnetGroupTagArrayOutput) Index(i pulumi.IntInput) DbSubnetGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DbSubnetGroupTag {
		return vs[0].([]DbSubnetGroupTag)[vs[1].(int)]
	}).(DbSubnetGroupTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterDbClusterRoleInput)(nil)).Elem(), DbClusterDbClusterRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterDbClusterRoleArrayInput)(nil)).Elem(), DbClusterDbClusterRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterParameterGroupTagInput)(nil)).Elem(), DbClusterParameterGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterParameterGroupTagArrayInput)(nil)).Elem(), DbClusterParameterGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterServerlessScalingConfigurationInput)(nil)).Elem(), DbClusterServerlessScalingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterServerlessScalingConfigurationPtrInput)(nil)).Elem(), DbClusterServerlessScalingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterTagInput)(nil)).Elem(), DbClusterTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbClusterTagArrayInput)(nil)).Elem(), DbClusterTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstanceTagInput)(nil)).Elem(), DbInstanceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbInstanceTagArrayInput)(nil)).Elem(), DbInstanceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbParameterGroupTagInput)(nil)).Elem(), DbParameterGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbParameterGroupTagArrayInput)(nil)).Elem(), DbParameterGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbSubnetGroupTagInput)(nil)).Elem(), DbSubnetGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbSubnetGroupTagArrayInput)(nil)).Elem(), DbSubnetGroupTagArray{})
	pulumi.RegisterOutputType(DbClusterDbClusterRoleOutput{})
	pulumi.RegisterOutputType(DbClusterDbClusterRoleArrayOutput{})
	pulumi.RegisterOutputType(DbClusterParameterGroupTagOutput{})
	pulumi.RegisterOutputType(DbClusterParameterGroupTagArrayOutput{})
	pulumi.RegisterOutputType(DbClusterServerlessScalingConfigurationOutput{})
	pulumi.RegisterOutputType(DbClusterServerlessScalingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(DbClusterTagOutput{})
	pulumi.RegisterOutputType(DbClusterTagArrayOutput{})
	pulumi.RegisterOutputType(DbInstanceTagOutput{})
	pulumi.RegisterOutputType(DbInstanceTagArrayOutput{})
	pulumi.RegisterOutputType(DbParameterGroupTagOutput{})
	pulumi.RegisterOutputType(DbParameterGroupTagArrayOutput{})
	pulumi.RegisterOutputType(DbSubnetGroupTagOutput{})
	pulumi.RegisterOutputType(DbSubnetGroupTagArrayOutput{})
}
