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

type DBClusterParameterGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBClusterParameterGroupTagInput is an input type that accepts DBClusterParameterGroupTagArgs and DBClusterParameterGroupTagOutput values.
// You can construct a concrete instance of `DBClusterParameterGroupTagInput` via:
//
//	DBClusterParameterGroupTagArgs{...}
type DBClusterParameterGroupTagInput interface {
	pulumi.Input

	ToDBClusterParameterGroupTagOutput() DBClusterParameterGroupTagOutput
	ToDBClusterParameterGroupTagOutputWithContext(context.Context) DBClusterParameterGroupTagOutput
}

type DBClusterParameterGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBClusterParameterGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterParameterGroupTag)(nil)).Elem()
}

func (i DBClusterParameterGroupTagArgs) ToDBClusterParameterGroupTagOutput() DBClusterParameterGroupTagOutput {
	return i.ToDBClusterParameterGroupTagOutputWithContext(context.Background())
}

func (i DBClusterParameterGroupTagArgs) ToDBClusterParameterGroupTagOutputWithContext(ctx context.Context) DBClusterParameterGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterParameterGroupTagOutput)
}

// DBClusterParameterGroupTagArrayInput is an input type that accepts DBClusterParameterGroupTagArray and DBClusterParameterGroupTagArrayOutput values.
// You can construct a concrete instance of `DBClusterParameterGroupTagArrayInput` via:
//
//	DBClusterParameterGroupTagArray{ DBClusterParameterGroupTagArgs{...} }
type DBClusterParameterGroupTagArrayInput interface {
	pulumi.Input

	ToDBClusterParameterGroupTagArrayOutput() DBClusterParameterGroupTagArrayOutput
	ToDBClusterParameterGroupTagArrayOutputWithContext(context.Context) DBClusterParameterGroupTagArrayOutput
}

type DBClusterParameterGroupTagArray []DBClusterParameterGroupTagInput

func (DBClusterParameterGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterParameterGroupTag)(nil)).Elem()
}

func (i DBClusterParameterGroupTagArray) ToDBClusterParameterGroupTagArrayOutput() DBClusterParameterGroupTagArrayOutput {
	return i.ToDBClusterParameterGroupTagArrayOutputWithContext(context.Background())
}

func (i DBClusterParameterGroupTagArray) ToDBClusterParameterGroupTagArrayOutputWithContext(ctx context.Context) DBClusterParameterGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterParameterGroupTagArrayOutput)
}

type DBClusterParameterGroupTagOutput struct{ *pulumi.OutputState }

func (DBClusterParameterGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterParameterGroupTag)(nil)).Elem()
}

func (o DBClusterParameterGroupTagOutput) ToDBClusterParameterGroupTagOutput() DBClusterParameterGroupTagOutput {
	return o
}

func (o DBClusterParameterGroupTagOutput) ToDBClusterParameterGroupTagOutputWithContext(ctx context.Context) DBClusterParameterGroupTagOutput {
	return o
}

func (o DBClusterParameterGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterParameterGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBClusterParameterGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterParameterGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBClusterParameterGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DBClusterParameterGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterParameterGroupTag)(nil)).Elem()
}

func (o DBClusterParameterGroupTagArrayOutput) ToDBClusterParameterGroupTagArrayOutput() DBClusterParameterGroupTagArrayOutput {
	return o
}

func (o DBClusterParameterGroupTagArrayOutput) ToDBClusterParameterGroupTagArrayOutputWithContext(ctx context.Context) DBClusterParameterGroupTagArrayOutput {
	return o
}

func (o DBClusterParameterGroupTagArrayOutput) Index(i pulumi.IntInput) DBClusterParameterGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBClusterParameterGroupTag {
		return vs[0].([]DBClusterParameterGroupTag)[vs[1].(int)]
	}).(DBClusterParameterGroupTagOutput)
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
type DBClusterRole struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
	FeatureName *string `pulumi:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn string `pulumi:"roleArn"`
}

// DBClusterRoleInput is an input type that accepts DBClusterRoleArgs and DBClusterRoleOutput values.
// You can construct a concrete instance of `DBClusterRoleInput` via:
//
//	DBClusterRoleArgs{...}
type DBClusterRoleInput interface {
	pulumi.Input

	ToDBClusterRoleOutput() DBClusterRoleOutput
	ToDBClusterRoleOutputWithContext(context.Context) DBClusterRoleOutput
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
type DBClusterRoleArgs struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
	FeatureName pulumi.StringPtrInput `pulumi:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
}

func (DBClusterRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterRole)(nil)).Elem()
}

func (i DBClusterRoleArgs) ToDBClusterRoleOutput() DBClusterRoleOutput {
	return i.ToDBClusterRoleOutputWithContext(context.Background())
}

func (i DBClusterRoleArgs) ToDBClusterRoleOutputWithContext(ctx context.Context) DBClusterRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterRoleOutput)
}

// DBClusterRoleArrayInput is an input type that accepts DBClusterRoleArray and DBClusterRoleArrayOutput values.
// You can construct a concrete instance of `DBClusterRoleArrayInput` via:
//
//	DBClusterRoleArray{ DBClusterRoleArgs{...} }
type DBClusterRoleArrayInput interface {
	pulumi.Input

	ToDBClusterRoleArrayOutput() DBClusterRoleArrayOutput
	ToDBClusterRoleArrayOutputWithContext(context.Context) DBClusterRoleArrayOutput
}

type DBClusterRoleArray []DBClusterRoleInput

func (DBClusterRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterRole)(nil)).Elem()
}

func (i DBClusterRoleArray) ToDBClusterRoleArrayOutput() DBClusterRoleArrayOutput {
	return i.ToDBClusterRoleArrayOutputWithContext(context.Background())
}

func (i DBClusterRoleArray) ToDBClusterRoleArrayOutputWithContext(ctx context.Context) DBClusterRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterRoleArrayOutput)
}

// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
type DBClusterRoleOutput struct{ *pulumi.OutputState }

func (DBClusterRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterRole)(nil)).Elem()
}

func (o DBClusterRoleOutput) ToDBClusterRoleOutput() DBClusterRoleOutput {
	return o
}

func (o DBClusterRoleOutput) ToDBClusterRoleOutputWithContext(ctx context.Context) DBClusterRoleOutput {
	return o
}

// The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.
func (o DBClusterRoleOutput) FeatureName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DBClusterRole) *string { return v.FeatureName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
func (o DBClusterRoleOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterRole) string { return v.RoleArn }).(pulumi.StringOutput)
}

type DBClusterRoleArrayOutput struct{ *pulumi.OutputState }

func (DBClusterRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterRole)(nil)).Elem()
}

func (o DBClusterRoleArrayOutput) ToDBClusterRoleArrayOutput() DBClusterRoleArrayOutput {
	return o
}

func (o DBClusterRoleArrayOutput) ToDBClusterRoleArrayOutputWithContext(ctx context.Context) DBClusterRoleArrayOutput {
	return o
}

func (o DBClusterRoleArrayOutput) Index(i pulumi.IntInput) DBClusterRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBClusterRole {
		return vs[0].([]DBClusterRole)[vs[1].(int)]
	}).(DBClusterRoleOutput)
}

// Contains the scaling configuration of an Neptune Serverless DB cluster.
type DBClusterServerlessScalingConfiguration struct {
	// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
	MaxCapacity float64 `pulumi:"maxCapacity"`
	// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
	MinCapacity float64 `pulumi:"minCapacity"`
}

// DBClusterServerlessScalingConfigurationInput is an input type that accepts DBClusterServerlessScalingConfigurationArgs and DBClusterServerlessScalingConfigurationOutput values.
// You can construct a concrete instance of `DBClusterServerlessScalingConfigurationInput` via:
//
//	DBClusterServerlessScalingConfigurationArgs{...}
type DBClusterServerlessScalingConfigurationInput interface {
	pulumi.Input

	ToDBClusterServerlessScalingConfigurationOutput() DBClusterServerlessScalingConfigurationOutput
	ToDBClusterServerlessScalingConfigurationOutputWithContext(context.Context) DBClusterServerlessScalingConfigurationOutput
}

// Contains the scaling configuration of an Neptune Serverless DB cluster.
type DBClusterServerlessScalingConfigurationArgs struct {
	// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
	MaxCapacity pulumi.Float64Input `pulumi:"maxCapacity"`
	// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
	MinCapacity pulumi.Float64Input `pulumi:"minCapacity"`
}

func (DBClusterServerlessScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (i DBClusterServerlessScalingConfigurationArgs) ToDBClusterServerlessScalingConfigurationOutput() DBClusterServerlessScalingConfigurationOutput {
	return i.ToDBClusterServerlessScalingConfigurationOutputWithContext(context.Background())
}

func (i DBClusterServerlessScalingConfigurationArgs) ToDBClusterServerlessScalingConfigurationOutputWithContext(ctx context.Context) DBClusterServerlessScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterServerlessScalingConfigurationOutput)
}

func (i DBClusterServerlessScalingConfigurationArgs) ToDBClusterServerlessScalingConfigurationPtrOutput() DBClusterServerlessScalingConfigurationPtrOutput {
	return i.ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i DBClusterServerlessScalingConfigurationArgs) ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DBClusterServerlessScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterServerlessScalingConfigurationOutput).ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(ctx)
}

// DBClusterServerlessScalingConfigurationPtrInput is an input type that accepts DBClusterServerlessScalingConfigurationArgs, DBClusterServerlessScalingConfigurationPtr and DBClusterServerlessScalingConfigurationPtrOutput values.
// You can construct a concrete instance of `DBClusterServerlessScalingConfigurationPtrInput` via:
//
//	        DBClusterServerlessScalingConfigurationArgs{...}
//
//	or:
//
//	        nil
type DBClusterServerlessScalingConfigurationPtrInput interface {
	pulumi.Input

	ToDBClusterServerlessScalingConfigurationPtrOutput() DBClusterServerlessScalingConfigurationPtrOutput
	ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(context.Context) DBClusterServerlessScalingConfigurationPtrOutput
}

type dbclusterServerlessScalingConfigurationPtrType DBClusterServerlessScalingConfigurationArgs

func DBClusterServerlessScalingConfigurationPtr(v *DBClusterServerlessScalingConfigurationArgs) DBClusterServerlessScalingConfigurationPtrInput {
	return (*dbclusterServerlessScalingConfigurationPtrType)(v)
}

func (*dbclusterServerlessScalingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (i *dbclusterServerlessScalingConfigurationPtrType) ToDBClusterServerlessScalingConfigurationPtrOutput() DBClusterServerlessScalingConfigurationPtrOutput {
	return i.ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i *dbclusterServerlessScalingConfigurationPtrType) ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DBClusterServerlessScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterServerlessScalingConfigurationPtrOutput)
}

// Contains the scaling configuration of an Neptune Serverless DB cluster.
type DBClusterServerlessScalingConfigurationOutput struct{ *pulumi.OutputState }

func (DBClusterServerlessScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (o DBClusterServerlessScalingConfigurationOutput) ToDBClusterServerlessScalingConfigurationOutput() DBClusterServerlessScalingConfigurationOutput {
	return o
}

func (o DBClusterServerlessScalingConfigurationOutput) ToDBClusterServerlessScalingConfigurationOutputWithContext(ctx context.Context) DBClusterServerlessScalingConfigurationOutput {
	return o
}

func (o DBClusterServerlessScalingConfigurationOutput) ToDBClusterServerlessScalingConfigurationPtrOutput() DBClusterServerlessScalingConfigurationPtrOutput {
	return o.ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(context.Background())
}

func (o DBClusterServerlessScalingConfigurationOutput) ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DBClusterServerlessScalingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DBClusterServerlessScalingConfiguration) *DBClusterServerlessScalingConfiguration {
		return &v
	}).(DBClusterServerlessScalingConfigurationPtrOutput)
}

// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
func (o DBClusterServerlessScalingConfigurationOutput) MaxCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v DBClusterServerlessScalingConfiguration) float64 { return v.MaxCapacity }).(pulumi.Float64Output)
}

// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
func (o DBClusterServerlessScalingConfigurationOutput) MinCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v DBClusterServerlessScalingConfiguration) float64 { return v.MinCapacity }).(pulumi.Float64Output)
}

type DBClusterServerlessScalingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (DBClusterServerlessScalingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DBClusterServerlessScalingConfiguration)(nil)).Elem()
}

func (o DBClusterServerlessScalingConfigurationPtrOutput) ToDBClusterServerlessScalingConfigurationPtrOutput() DBClusterServerlessScalingConfigurationPtrOutput {
	return o
}

func (o DBClusterServerlessScalingConfigurationPtrOutput) ToDBClusterServerlessScalingConfigurationPtrOutputWithContext(ctx context.Context) DBClusterServerlessScalingConfigurationPtrOutput {
	return o
}

func (o DBClusterServerlessScalingConfigurationPtrOutput) Elem() DBClusterServerlessScalingConfigurationOutput {
	return o.ApplyT(func(v *DBClusterServerlessScalingConfiguration) DBClusterServerlessScalingConfiguration {
		if v != nil {
			return *v
		}
		var ret DBClusterServerlessScalingConfiguration
		return ret
	}).(DBClusterServerlessScalingConfigurationOutput)
}

// The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.
func (o DBClusterServerlessScalingConfigurationPtrOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DBClusterServerlessScalingConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return &v.MaxCapacity
	}).(pulumi.Float64PtrOutput)
}

// The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.
func (o DBClusterServerlessScalingConfigurationPtrOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DBClusterServerlessScalingConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return &v.MinCapacity
	}).(pulumi.Float64PtrOutput)
}

// A key-value pair to associate with a resource.
type DBClusterTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value *string `pulumi:"value"`
}

// DBClusterTagInput is an input type that accepts DBClusterTagArgs and DBClusterTagOutput values.
// You can construct a concrete instance of `DBClusterTagInput` via:
//
//	DBClusterTagArgs{...}
type DBClusterTagInput interface {
	pulumi.Input

	ToDBClusterTagOutput() DBClusterTagOutput
	ToDBClusterTagOutputWithContext(context.Context) DBClusterTagOutput
}

// A key-value pair to associate with a resource.
type DBClusterTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DBClusterTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterTag)(nil)).Elem()
}

func (i DBClusterTagArgs) ToDBClusterTagOutput() DBClusterTagOutput {
	return i.ToDBClusterTagOutputWithContext(context.Background())
}

func (i DBClusterTagArgs) ToDBClusterTagOutputWithContext(ctx context.Context) DBClusterTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterTagOutput)
}

// DBClusterTagArrayInput is an input type that accepts DBClusterTagArray and DBClusterTagArrayOutput values.
// You can construct a concrete instance of `DBClusterTagArrayInput` via:
//
//	DBClusterTagArray{ DBClusterTagArgs{...} }
type DBClusterTagArrayInput interface {
	pulumi.Input

	ToDBClusterTagArrayOutput() DBClusterTagArrayOutput
	ToDBClusterTagArrayOutputWithContext(context.Context) DBClusterTagArrayOutput
}

type DBClusterTagArray []DBClusterTagInput

func (DBClusterTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterTag)(nil)).Elem()
}

func (i DBClusterTagArray) ToDBClusterTagArrayOutput() DBClusterTagArrayOutput {
	return i.ToDBClusterTagArrayOutputWithContext(context.Background())
}

func (i DBClusterTagArray) ToDBClusterTagArrayOutputWithContext(ctx context.Context) DBClusterTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBClusterTagArrayOutput)
}

// A key-value pair to associate with a resource.
type DBClusterTagOutput struct{ *pulumi.OutputState }

func (DBClusterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBClusterTag)(nil)).Elem()
}

func (o DBClusterTagOutput) ToDBClusterTagOutput() DBClusterTagOutput {
	return o
}

func (o DBClusterTagOutput) ToDBClusterTagOutputWithContext(ctx context.Context) DBClusterTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o DBClusterTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBClusterTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o DBClusterTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DBClusterTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DBClusterTagArrayOutput struct{ *pulumi.OutputState }

func (DBClusterTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBClusterTag)(nil)).Elem()
}

func (o DBClusterTagArrayOutput) ToDBClusterTagArrayOutput() DBClusterTagArrayOutput {
	return o
}

func (o DBClusterTagArrayOutput) ToDBClusterTagArrayOutputWithContext(ctx context.Context) DBClusterTagArrayOutput {
	return o
}

func (o DBClusterTagArrayOutput) Index(i pulumi.IntInput) DBClusterTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBClusterTag {
		return vs[0].([]DBClusterTag)[vs[1].(int)]
	}).(DBClusterTagOutput)
}

type DBInstanceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBInstanceTagInput is an input type that accepts DBInstanceTagArgs and DBInstanceTagOutput values.
// You can construct a concrete instance of `DBInstanceTagInput` via:
//
//	DBInstanceTagArgs{...}
type DBInstanceTagInput interface {
	pulumi.Input

	ToDBInstanceTagOutput() DBInstanceTagOutput
	ToDBInstanceTagOutputWithContext(context.Context) DBInstanceTagOutput
}

type DBInstanceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBInstanceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBInstanceTag)(nil)).Elem()
}

func (i DBInstanceTagArgs) ToDBInstanceTagOutput() DBInstanceTagOutput {
	return i.ToDBInstanceTagOutputWithContext(context.Background())
}

func (i DBInstanceTagArgs) ToDBInstanceTagOutputWithContext(ctx context.Context) DBInstanceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceTagOutput)
}

// DBInstanceTagArrayInput is an input type that accepts DBInstanceTagArray and DBInstanceTagArrayOutput values.
// You can construct a concrete instance of `DBInstanceTagArrayInput` via:
//
//	DBInstanceTagArray{ DBInstanceTagArgs{...} }
type DBInstanceTagArrayInput interface {
	pulumi.Input

	ToDBInstanceTagArrayOutput() DBInstanceTagArrayOutput
	ToDBInstanceTagArrayOutputWithContext(context.Context) DBInstanceTagArrayOutput
}

type DBInstanceTagArray []DBInstanceTagInput

func (DBInstanceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBInstanceTag)(nil)).Elem()
}

func (i DBInstanceTagArray) ToDBInstanceTagArrayOutput() DBInstanceTagArrayOutput {
	return i.ToDBInstanceTagArrayOutputWithContext(context.Background())
}

func (i DBInstanceTagArray) ToDBInstanceTagArrayOutputWithContext(ctx context.Context) DBInstanceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBInstanceTagArrayOutput)
}

type DBInstanceTagOutput struct{ *pulumi.OutputState }

func (DBInstanceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBInstanceTag)(nil)).Elem()
}

func (o DBInstanceTagOutput) ToDBInstanceTagOutput() DBInstanceTagOutput {
	return o
}

func (o DBInstanceTagOutput) ToDBInstanceTagOutputWithContext(ctx context.Context) DBInstanceTagOutput {
	return o
}

func (o DBInstanceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBInstanceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBInstanceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBInstanceTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBInstanceTagArrayOutput struct{ *pulumi.OutputState }

func (DBInstanceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBInstanceTag)(nil)).Elem()
}

func (o DBInstanceTagArrayOutput) ToDBInstanceTagArrayOutput() DBInstanceTagArrayOutput {
	return o
}

func (o DBInstanceTagArrayOutput) ToDBInstanceTagArrayOutputWithContext(ctx context.Context) DBInstanceTagArrayOutput {
	return o
}

func (o DBInstanceTagArrayOutput) Index(i pulumi.IntInput) DBInstanceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBInstanceTag {
		return vs[0].([]DBInstanceTag)[vs[1].(int)]
	}).(DBInstanceTagOutput)
}

type DBParameterGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBParameterGroupTagInput is an input type that accepts DBParameterGroupTagArgs and DBParameterGroupTagOutput values.
// You can construct a concrete instance of `DBParameterGroupTagInput` via:
//
//	DBParameterGroupTagArgs{...}
type DBParameterGroupTagInput interface {
	pulumi.Input

	ToDBParameterGroupTagOutput() DBParameterGroupTagOutput
	ToDBParameterGroupTagOutputWithContext(context.Context) DBParameterGroupTagOutput
}

type DBParameterGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBParameterGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBParameterGroupTag)(nil)).Elem()
}

func (i DBParameterGroupTagArgs) ToDBParameterGroupTagOutput() DBParameterGroupTagOutput {
	return i.ToDBParameterGroupTagOutputWithContext(context.Background())
}

func (i DBParameterGroupTagArgs) ToDBParameterGroupTagOutputWithContext(ctx context.Context) DBParameterGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBParameterGroupTagOutput)
}

// DBParameterGroupTagArrayInput is an input type that accepts DBParameterGroupTagArray and DBParameterGroupTagArrayOutput values.
// You can construct a concrete instance of `DBParameterGroupTagArrayInput` via:
//
//	DBParameterGroupTagArray{ DBParameterGroupTagArgs{...} }
type DBParameterGroupTagArrayInput interface {
	pulumi.Input

	ToDBParameterGroupTagArrayOutput() DBParameterGroupTagArrayOutput
	ToDBParameterGroupTagArrayOutputWithContext(context.Context) DBParameterGroupTagArrayOutput
}

type DBParameterGroupTagArray []DBParameterGroupTagInput

func (DBParameterGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBParameterGroupTag)(nil)).Elem()
}

func (i DBParameterGroupTagArray) ToDBParameterGroupTagArrayOutput() DBParameterGroupTagArrayOutput {
	return i.ToDBParameterGroupTagArrayOutputWithContext(context.Background())
}

func (i DBParameterGroupTagArray) ToDBParameterGroupTagArrayOutputWithContext(ctx context.Context) DBParameterGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBParameterGroupTagArrayOutput)
}

type DBParameterGroupTagOutput struct{ *pulumi.OutputState }

func (DBParameterGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBParameterGroupTag)(nil)).Elem()
}

func (o DBParameterGroupTagOutput) ToDBParameterGroupTagOutput() DBParameterGroupTagOutput {
	return o
}

func (o DBParameterGroupTagOutput) ToDBParameterGroupTagOutputWithContext(ctx context.Context) DBParameterGroupTagOutput {
	return o
}

func (o DBParameterGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBParameterGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBParameterGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBParameterGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBParameterGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DBParameterGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBParameterGroupTag)(nil)).Elem()
}

func (o DBParameterGroupTagArrayOutput) ToDBParameterGroupTagArrayOutput() DBParameterGroupTagArrayOutput {
	return o
}

func (o DBParameterGroupTagArrayOutput) ToDBParameterGroupTagArrayOutputWithContext(ctx context.Context) DBParameterGroupTagArrayOutput {
	return o
}

func (o DBParameterGroupTagArrayOutput) Index(i pulumi.IntInput) DBParameterGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBParameterGroupTag {
		return vs[0].([]DBParameterGroupTag)[vs[1].(int)]
	}).(DBParameterGroupTagOutput)
}

type DBSubnetGroupTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// DBSubnetGroupTagInput is an input type that accepts DBSubnetGroupTagArgs and DBSubnetGroupTagOutput values.
// You can construct a concrete instance of `DBSubnetGroupTagInput` via:
//
//	DBSubnetGroupTagArgs{...}
type DBSubnetGroupTagInput interface {
	pulumi.Input

	ToDBSubnetGroupTagOutput() DBSubnetGroupTagOutput
	ToDBSubnetGroupTagOutputWithContext(context.Context) DBSubnetGroupTagOutput
}

type DBSubnetGroupTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (DBSubnetGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSubnetGroupTag)(nil)).Elem()
}

func (i DBSubnetGroupTagArgs) ToDBSubnetGroupTagOutput() DBSubnetGroupTagOutput {
	return i.ToDBSubnetGroupTagOutputWithContext(context.Background())
}

func (i DBSubnetGroupTagArgs) ToDBSubnetGroupTagOutputWithContext(ctx context.Context) DBSubnetGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBSubnetGroupTagOutput)
}

// DBSubnetGroupTagArrayInput is an input type that accepts DBSubnetGroupTagArray and DBSubnetGroupTagArrayOutput values.
// You can construct a concrete instance of `DBSubnetGroupTagArrayInput` via:
//
//	DBSubnetGroupTagArray{ DBSubnetGroupTagArgs{...} }
type DBSubnetGroupTagArrayInput interface {
	pulumi.Input

	ToDBSubnetGroupTagArrayOutput() DBSubnetGroupTagArrayOutput
	ToDBSubnetGroupTagArrayOutputWithContext(context.Context) DBSubnetGroupTagArrayOutput
}

type DBSubnetGroupTagArray []DBSubnetGroupTagInput

func (DBSubnetGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBSubnetGroupTag)(nil)).Elem()
}

func (i DBSubnetGroupTagArray) ToDBSubnetGroupTagArrayOutput() DBSubnetGroupTagArrayOutput {
	return i.ToDBSubnetGroupTagArrayOutputWithContext(context.Background())
}

func (i DBSubnetGroupTagArray) ToDBSubnetGroupTagArrayOutputWithContext(ctx context.Context) DBSubnetGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBSubnetGroupTagArrayOutput)
}

type DBSubnetGroupTagOutput struct{ *pulumi.OutputState }

func (DBSubnetGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSubnetGroupTag)(nil)).Elem()
}

func (o DBSubnetGroupTagOutput) ToDBSubnetGroupTagOutput() DBSubnetGroupTagOutput {
	return o
}

func (o DBSubnetGroupTagOutput) ToDBSubnetGroupTagOutputWithContext(ctx context.Context) DBSubnetGroupTagOutput {
	return o
}

func (o DBSubnetGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DBSubnetGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o DBSubnetGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DBSubnetGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type DBSubnetGroupTagArrayOutput struct{ *pulumi.OutputState }

func (DBSubnetGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DBSubnetGroupTag)(nil)).Elem()
}

func (o DBSubnetGroupTagArrayOutput) ToDBSubnetGroupTagArrayOutput() DBSubnetGroupTagArrayOutput {
	return o
}

func (o DBSubnetGroupTagArrayOutput) ToDBSubnetGroupTagArrayOutputWithContext(ctx context.Context) DBSubnetGroupTagArrayOutput {
	return o
}

func (o DBSubnetGroupTagArrayOutput) Index(i pulumi.IntInput) DBSubnetGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DBSubnetGroupTag {
		return vs[0].([]DBSubnetGroupTag)[vs[1].(int)]
	}).(DBSubnetGroupTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterParameterGroupTagInput)(nil)).Elem(), DBClusterParameterGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterParameterGroupTagArrayInput)(nil)).Elem(), DBClusterParameterGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterRoleInput)(nil)).Elem(), DBClusterRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterRoleArrayInput)(nil)).Elem(), DBClusterRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterServerlessScalingConfigurationInput)(nil)).Elem(), DBClusterServerlessScalingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterServerlessScalingConfigurationPtrInput)(nil)).Elem(), DBClusterServerlessScalingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterTagInput)(nil)).Elem(), DBClusterTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBClusterTagArrayInput)(nil)).Elem(), DBClusterTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBInstanceTagInput)(nil)).Elem(), DBInstanceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBInstanceTagArrayInput)(nil)).Elem(), DBInstanceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBParameterGroupTagInput)(nil)).Elem(), DBParameterGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBParameterGroupTagArrayInput)(nil)).Elem(), DBParameterGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBSubnetGroupTagInput)(nil)).Elem(), DBSubnetGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DBSubnetGroupTagArrayInput)(nil)).Elem(), DBSubnetGroupTagArray{})
	pulumi.RegisterOutputType(DBClusterParameterGroupTagOutput{})
	pulumi.RegisterOutputType(DBClusterParameterGroupTagArrayOutput{})
	pulumi.RegisterOutputType(DBClusterRoleOutput{})
	pulumi.RegisterOutputType(DBClusterRoleArrayOutput{})
	pulumi.RegisterOutputType(DBClusterServerlessScalingConfigurationOutput{})
	pulumi.RegisterOutputType(DBClusterServerlessScalingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(DBClusterTagOutput{})
	pulumi.RegisterOutputType(DBClusterTagArrayOutput{})
	pulumi.RegisterOutputType(DBInstanceTagOutput{})
	pulumi.RegisterOutputType(DBInstanceTagArrayOutput{})
	pulumi.RegisterOutputType(DBParameterGroupTagOutput{})
	pulumi.RegisterOutputType(DBParameterGroupTagArrayOutput{})
	pulumi.RegisterOutputType(DBSubnetGroupTagOutput{})
	pulumi.RegisterOutputType(DBSubnetGroupTagArrayOutput{})
}
