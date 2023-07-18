// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package memorydb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// A key-value pair to associate with a resource.
type ACLTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value *string `pulumi:"value"`
}

// ACLTagInput is an input type that accepts ACLTagArgs and ACLTagOutput values.
// You can construct a concrete instance of `ACLTagInput` via:
//
//	ACLTagArgs{...}
type ACLTagInput interface {
	pulumi.Input

	ToACLTagOutput() ACLTagOutput
	ToACLTagOutputWithContext(context.Context) ACLTagOutput
}

// A key-value pair to associate with a resource.
type ACLTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ACLTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ACLTag)(nil)).Elem()
}

func (i ACLTagArgs) ToACLTagOutput() ACLTagOutput {
	return i.ToACLTagOutputWithContext(context.Background())
}

func (i ACLTagArgs) ToACLTagOutputWithContext(ctx context.Context) ACLTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACLTagOutput)
}

// ACLTagArrayInput is an input type that accepts ACLTagArray and ACLTagArrayOutput values.
// You can construct a concrete instance of `ACLTagArrayInput` via:
//
//	ACLTagArray{ ACLTagArgs{...} }
type ACLTagArrayInput interface {
	pulumi.Input

	ToACLTagArrayOutput() ACLTagArrayOutput
	ToACLTagArrayOutputWithContext(context.Context) ACLTagArrayOutput
}

type ACLTagArray []ACLTagInput

func (ACLTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ACLTag)(nil)).Elem()
}

func (i ACLTagArray) ToACLTagArrayOutput() ACLTagArrayOutput {
	return i.ToACLTagArrayOutputWithContext(context.Background())
}

func (i ACLTagArray) ToACLTagArrayOutputWithContext(ctx context.Context) ACLTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ACLTagArrayOutput)
}

// A key-value pair to associate with a resource.
type ACLTagOutput struct{ *pulumi.OutputState }

func (ACLTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ACLTag)(nil)).Elem()
}

func (o ACLTagOutput) ToACLTagOutput() ACLTagOutput {
	return o
}

func (o ACLTagOutput) ToACLTagOutputWithContext(ctx context.Context) ACLTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o ACLTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ACLTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o ACLTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ACLTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ACLTagArrayOutput struct{ *pulumi.OutputState }

func (ACLTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ACLTag)(nil)).Elem()
}

func (o ACLTagArrayOutput) ToACLTagArrayOutput() ACLTagArrayOutput {
	return o
}

func (o ACLTagArrayOutput) ToACLTagArrayOutputWithContext(ctx context.Context) ACLTagArrayOutput {
	return o
}

func (o ACLTagArrayOutput) Index(i pulumi.IntInput) ACLTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ACLTag {
		return vs[0].([]ACLTag)[vs[1].(int)]
	}).(ACLTagOutput)
}

type AuthenticationModeProperties struct {
	// Passwords used for this user account. You can create up to two passwords for each user.
	Passwords []string `pulumi:"passwords"`
	// Type of authentication strategy for this user.
	Type *UserAuthenticationModePropertiesType `pulumi:"type"`
}

// AuthenticationModePropertiesInput is an input type that accepts AuthenticationModePropertiesArgs and AuthenticationModePropertiesOutput values.
// You can construct a concrete instance of `AuthenticationModePropertiesInput` via:
//
//	AuthenticationModePropertiesArgs{...}
type AuthenticationModePropertiesInput interface {
	pulumi.Input

	ToAuthenticationModePropertiesOutput() AuthenticationModePropertiesOutput
	ToAuthenticationModePropertiesOutputWithContext(context.Context) AuthenticationModePropertiesOutput
}

type AuthenticationModePropertiesArgs struct {
	// Passwords used for this user account. You can create up to two passwords for each user.
	Passwords pulumi.StringArrayInput `pulumi:"passwords"`
	// Type of authentication strategy for this user.
	Type UserAuthenticationModePropertiesTypePtrInput `pulumi:"type"`
}

func (AuthenticationModePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationModeProperties)(nil)).Elem()
}

func (i AuthenticationModePropertiesArgs) ToAuthenticationModePropertiesOutput() AuthenticationModePropertiesOutput {
	return i.ToAuthenticationModePropertiesOutputWithContext(context.Background())
}

func (i AuthenticationModePropertiesArgs) ToAuthenticationModePropertiesOutputWithContext(ctx context.Context) AuthenticationModePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationModePropertiesOutput)
}

func (i AuthenticationModePropertiesArgs) ToAuthenticationModePropertiesPtrOutput() AuthenticationModePropertiesPtrOutput {
	return i.ToAuthenticationModePropertiesPtrOutputWithContext(context.Background())
}

func (i AuthenticationModePropertiesArgs) ToAuthenticationModePropertiesPtrOutputWithContext(ctx context.Context) AuthenticationModePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationModePropertiesOutput).ToAuthenticationModePropertiesPtrOutputWithContext(ctx)
}

// AuthenticationModePropertiesPtrInput is an input type that accepts AuthenticationModePropertiesArgs, AuthenticationModePropertiesPtr and AuthenticationModePropertiesPtrOutput values.
// You can construct a concrete instance of `AuthenticationModePropertiesPtrInput` via:
//
//	        AuthenticationModePropertiesArgs{...}
//
//	or:
//
//	        nil
type AuthenticationModePropertiesPtrInput interface {
	pulumi.Input

	ToAuthenticationModePropertiesPtrOutput() AuthenticationModePropertiesPtrOutput
	ToAuthenticationModePropertiesPtrOutputWithContext(context.Context) AuthenticationModePropertiesPtrOutput
}

type authenticationModePropertiesPtrType AuthenticationModePropertiesArgs

func AuthenticationModePropertiesPtr(v *AuthenticationModePropertiesArgs) AuthenticationModePropertiesPtrInput {
	return (*authenticationModePropertiesPtrType)(v)
}

func (*authenticationModePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationModeProperties)(nil)).Elem()
}

func (i *authenticationModePropertiesPtrType) ToAuthenticationModePropertiesPtrOutput() AuthenticationModePropertiesPtrOutput {
	return i.ToAuthenticationModePropertiesPtrOutputWithContext(context.Background())
}

func (i *authenticationModePropertiesPtrType) ToAuthenticationModePropertiesPtrOutputWithContext(ctx context.Context) AuthenticationModePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationModePropertiesPtrOutput)
}

type AuthenticationModePropertiesOutput struct{ *pulumi.OutputState }

func (AuthenticationModePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationModeProperties)(nil)).Elem()
}

func (o AuthenticationModePropertiesOutput) ToAuthenticationModePropertiesOutput() AuthenticationModePropertiesOutput {
	return o
}

func (o AuthenticationModePropertiesOutput) ToAuthenticationModePropertiesOutputWithContext(ctx context.Context) AuthenticationModePropertiesOutput {
	return o
}

func (o AuthenticationModePropertiesOutput) ToAuthenticationModePropertiesPtrOutput() AuthenticationModePropertiesPtrOutput {
	return o.ToAuthenticationModePropertiesPtrOutputWithContext(context.Background())
}

func (o AuthenticationModePropertiesOutput) ToAuthenticationModePropertiesPtrOutputWithContext(ctx context.Context) AuthenticationModePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationModeProperties) *AuthenticationModeProperties {
		return &v
	}).(AuthenticationModePropertiesPtrOutput)
}

// Passwords used for this user account. You can create up to two passwords for each user.
func (o AuthenticationModePropertiesOutput) Passwords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthenticationModeProperties) []string { return v.Passwords }).(pulumi.StringArrayOutput)
}

// Type of authentication strategy for this user.
func (o AuthenticationModePropertiesOutput) Type() UserAuthenticationModePropertiesTypePtrOutput {
	return o.ApplyT(func(v AuthenticationModeProperties) *UserAuthenticationModePropertiesType { return v.Type }).(UserAuthenticationModePropertiesTypePtrOutput)
}

type AuthenticationModePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AuthenticationModePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationModeProperties)(nil)).Elem()
}

func (o AuthenticationModePropertiesPtrOutput) ToAuthenticationModePropertiesPtrOutput() AuthenticationModePropertiesPtrOutput {
	return o
}

func (o AuthenticationModePropertiesPtrOutput) ToAuthenticationModePropertiesPtrOutputWithContext(ctx context.Context) AuthenticationModePropertiesPtrOutput {
	return o
}

func (o AuthenticationModePropertiesPtrOutput) Elem() AuthenticationModePropertiesOutput {
	return o.ApplyT(func(v *AuthenticationModeProperties) AuthenticationModeProperties {
		if v != nil {
			return *v
		}
		var ret AuthenticationModeProperties
		return ret
	}).(AuthenticationModePropertiesOutput)
}

// Passwords used for this user account. You can create up to two passwords for each user.
func (o AuthenticationModePropertiesPtrOutput) Passwords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthenticationModeProperties) []string {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(pulumi.StringArrayOutput)
}

// Type of authentication strategy for this user.
func (o AuthenticationModePropertiesPtrOutput) Type() UserAuthenticationModePropertiesTypePtrOutput {
	return o.ApplyT(func(v *AuthenticationModeProperties) *UserAuthenticationModePropertiesType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(UserAuthenticationModePropertiesTypePtrOutput)
}

type ClusterEndpoint struct {
	// The DNS address of the primary read-write node.
	Address *string `pulumi:"address"`
	// The port number that the engine is listening on.
	Port *int `pulumi:"port"`
}

// ClusterEndpointInput is an input type that accepts ClusterEndpointArgs and ClusterEndpointOutput values.
// You can construct a concrete instance of `ClusterEndpointInput` via:
//
//	ClusterEndpointArgs{...}
type ClusterEndpointInput interface {
	pulumi.Input

	ToClusterEndpointOutput() ClusterEndpointOutput
	ToClusterEndpointOutputWithContext(context.Context) ClusterEndpointOutput
}

type ClusterEndpointArgs struct {
	// The DNS address of the primary read-write node.
	Address pulumi.StringPtrInput `pulumi:"address"`
	// The port number that the engine is listening on.
	Port pulumi.IntPtrInput `pulumi:"port"`
}

func (ClusterEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEndpoint)(nil)).Elem()
}

func (i ClusterEndpointArgs) ToClusterEndpointOutput() ClusterEndpointOutput {
	return i.ToClusterEndpointOutputWithContext(context.Background())
}

func (i ClusterEndpointArgs) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointOutput)
}

func (i ClusterEndpointArgs) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return i.ToClusterEndpointPtrOutputWithContext(context.Background())
}

func (i ClusterEndpointArgs) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointOutput).ToClusterEndpointPtrOutputWithContext(ctx)
}

// ClusterEndpointPtrInput is an input type that accepts ClusterEndpointArgs, ClusterEndpointPtr and ClusterEndpointPtrOutput values.
// You can construct a concrete instance of `ClusterEndpointPtrInput` via:
//
//	        ClusterEndpointArgs{...}
//
//	or:
//
//	        nil
type ClusterEndpointPtrInput interface {
	pulumi.Input

	ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput
	ToClusterEndpointPtrOutputWithContext(context.Context) ClusterEndpointPtrOutput
}

type clusterEndpointPtrType ClusterEndpointArgs

func ClusterEndpointPtr(v *ClusterEndpointArgs) ClusterEndpointPtrInput {
	return (*clusterEndpointPtrType)(v)
}

func (*clusterEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (i *clusterEndpointPtrType) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return i.ToClusterEndpointPtrOutputWithContext(context.Background())
}

func (i *clusterEndpointPtrType) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointPtrOutput)
}

type ClusterEndpointOutput struct{ *pulumi.OutputState }

func (ClusterEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointOutput) ToClusterEndpointOutput() ClusterEndpointOutput {
	return o
}

func (o ClusterEndpointOutput) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return o
}

func (o ClusterEndpointOutput) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return o.ToClusterEndpointPtrOutputWithContext(context.Background())
}

func (o ClusterEndpointOutput) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterEndpoint) *ClusterEndpoint {
		return &v
	}).(ClusterEndpointPtrOutput)
}

// The DNS address of the primary read-write node.
func (o ClusterEndpointOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterEndpoint) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The port number that the engine is listening on.
func (o ClusterEndpointOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterEndpoint) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ClusterEndpointPtrOutput struct{ *pulumi.OutputState }

func (ClusterEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointPtrOutput) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return o
}

func (o ClusterEndpointPtrOutput) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return o
}

func (o ClusterEndpointPtrOutput) Elem() ClusterEndpointOutput {
	return o.ApplyT(func(v *ClusterEndpoint) ClusterEndpoint {
		if v != nil {
			return *v
		}
		var ret ClusterEndpoint
		return ret
	}).(ClusterEndpointOutput)
}

// The DNS address of the primary read-write node.
func (o ClusterEndpointPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

// The port number that the engine is listening on.
func (o ClusterEndpointPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

// A key-value pair to associate with a resource.
type ClusterTag struct {
	// The key for the tag. May not be null.
	Key string `pulumi:"key"`
	// The tag's value. May be null.
	Value string `pulumi:"value"`
}

// ClusterTagInput is an input type that accepts ClusterTagArgs and ClusterTagOutput values.
// You can construct a concrete instance of `ClusterTagInput` via:
//
//	ClusterTagArgs{...}
type ClusterTagInput interface {
	pulumi.Input

	ToClusterTagOutput() ClusterTagOutput
	ToClusterTagOutputWithContext(context.Context) ClusterTagOutput
}

// A key-value pair to associate with a resource.
type ClusterTagArgs struct {
	// The key for the tag. May not be null.
	Key pulumi.StringInput `pulumi:"key"`
	// The tag's value. May be null.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterTag)(nil)).Elem()
}

func (i ClusterTagArgs) ToClusterTagOutput() ClusterTagOutput {
	return i.ToClusterTagOutputWithContext(context.Background())
}

func (i ClusterTagArgs) ToClusterTagOutputWithContext(ctx context.Context) ClusterTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTagOutput)
}

// ClusterTagArrayInput is an input type that accepts ClusterTagArray and ClusterTagArrayOutput values.
// You can construct a concrete instance of `ClusterTagArrayInput` via:
//
//	ClusterTagArray{ ClusterTagArgs{...} }
type ClusterTagArrayInput interface {
	pulumi.Input

	ToClusterTagArrayOutput() ClusterTagArrayOutput
	ToClusterTagArrayOutputWithContext(context.Context) ClusterTagArrayOutput
}

type ClusterTagArray []ClusterTagInput

func (ClusterTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterTag)(nil)).Elem()
}

func (i ClusterTagArray) ToClusterTagArrayOutput() ClusterTagArrayOutput {
	return i.ToClusterTagArrayOutputWithContext(context.Background())
}

func (i ClusterTagArray) ToClusterTagArrayOutputWithContext(ctx context.Context) ClusterTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterTagArrayOutput)
}

// A key-value pair to associate with a resource.
type ClusterTagOutput struct{ *pulumi.OutputState }

func (ClusterTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterTag)(nil)).Elem()
}

func (o ClusterTagOutput) ToClusterTagOutput() ClusterTagOutput {
	return o
}

func (o ClusterTagOutput) ToClusterTagOutputWithContext(ctx context.Context) ClusterTagOutput {
	return o
}

// The key for the tag. May not be null.
func (o ClusterTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterTag) string { return v.Key }).(pulumi.StringOutput)
}

// The tag's value. May be null.
func (o ClusterTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterTag) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterTagArrayOutput struct{ *pulumi.OutputState }

func (ClusterTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterTag)(nil)).Elem()
}

func (o ClusterTagArrayOutput) ToClusterTagArrayOutput() ClusterTagArrayOutput {
	return o
}

func (o ClusterTagArrayOutput) ToClusterTagArrayOutputWithContext(ctx context.Context) ClusterTagArrayOutput {
	return o
}

func (o ClusterTagArrayOutput) Index(i pulumi.IntInput) ClusterTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterTag {
		return vs[0].([]ClusterTag)[vs[1].(int)]
	}).(ClusterTagOutput)
}

// A key-value pair to associate with a resource.
type ParameterGroupTag struct {
	// The key for the tag. May not be null.
	Key string `pulumi:"key"`
	// The tag's value. May be null.
	Value string `pulumi:"value"`
}

// ParameterGroupTagInput is an input type that accepts ParameterGroupTagArgs and ParameterGroupTagOutput values.
// You can construct a concrete instance of `ParameterGroupTagInput` via:
//
//	ParameterGroupTagArgs{...}
type ParameterGroupTagInput interface {
	pulumi.Input

	ToParameterGroupTagOutput() ParameterGroupTagOutput
	ToParameterGroupTagOutputWithContext(context.Context) ParameterGroupTagOutput
}

// A key-value pair to associate with a resource.
type ParameterGroupTagArgs struct {
	// The key for the tag. May not be null.
	Key pulumi.StringInput `pulumi:"key"`
	// The tag's value. May be null.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupTag)(nil)).Elem()
}

func (i ParameterGroupTagArgs) ToParameterGroupTagOutput() ParameterGroupTagOutput {
	return i.ToParameterGroupTagOutputWithContext(context.Background())
}

func (i ParameterGroupTagArgs) ToParameterGroupTagOutputWithContext(ctx context.Context) ParameterGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupTagOutput)
}

// ParameterGroupTagArrayInput is an input type that accepts ParameterGroupTagArray and ParameterGroupTagArrayOutput values.
// You can construct a concrete instance of `ParameterGroupTagArrayInput` via:
//
//	ParameterGroupTagArray{ ParameterGroupTagArgs{...} }
type ParameterGroupTagArrayInput interface {
	pulumi.Input

	ToParameterGroupTagArrayOutput() ParameterGroupTagArrayOutput
	ToParameterGroupTagArrayOutputWithContext(context.Context) ParameterGroupTagArrayOutput
}

type ParameterGroupTagArray []ParameterGroupTagInput

func (ParameterGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupTag)(nil)).Elem()
}

func (i ParameterGroupTagArray) ToParameterGroupTagArrayOutput() ParameterGroupTagArrayOutput {
	return i.ToParameterGroupTagArrayOutputWithContext(context.Background())
}

func (i ParameterGroupTagArray) ToParameterGroupTagArrayOutputWithContext(ctx context.Context) ParameterGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupTagArrayOutput)
}

// A key-value pair to associate with a resource.
type ParameterGroupTagOutput struct{ *pulumi.OutputState }

func (ParameterGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupTag)(nil)).Elem()
}

func (o ParameterGroupTagOutput) ToParameterGroupTagOutput() ParameterGroupTagOutput {
	return o
}

func (o ParameterGroupTagOutput) ToParameterGroupTagOutputWithContext(ctx context.Context) ParameterGroupTagOutput {
	return o
}

// The key for the tag. May not be null.
func (o ParameterGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

// The tag's value. May be null.
func (o ParameterGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupTagArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupTag)(nil)).Elem()
}

func (o ParameterGroupTagArrayOutput) ToParameterGroupTagArrayOutput() ParameterGroupTagArrayOutput {
	return o
}

func (o ParameterGroupTagArrayOutput) ToParameterGroupTagArrayOutputWithContext(ctx context.Context) ParameterGroupTagArrayOutput {
	return o
}

func (o ParameterGroupTagArrayOutput) Index(i pulumi.IntInput) ParameterGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupTag {
		return vs[0].([]ParameterGroupTag)[vs[1].(int)]
	}).(ParameterGroupTagOutput)
}

// A key-value pair to associate with a resource.
type SubnetGroupTag struct {
	// The key for the tag. May not be null.
	Key string `pulumi:"key"`
	// The tag's value. May be null.
	Value string `pulumi:"value"`
}

// SubnetGroupTagInput is an input type that accepts SubnetGroupTagArgs and SubnetGroupTagOutput values.
// You can construct a concrete instance of `SubnetGroupTagInput` via:
//
//	SubnetGroupTagArgs{...}
type SubnetGroupTagInput interface {
	pulumi.Input

	ToSubnetGroupTagOutput() SubnetGroupTagOutput
	ToSubnetGroupTagOutputWithContext(context.Context) SubnetGroupTagOutput
}

// A key-value pair to associate with a resource.
type SubnetGroupTagArgs struct {
	// The key for the tag. May not be null.
	Key pulumi.StringInput `pulumi:"key"`
	// The tag's value. May be null.
	Value pulumi.StringInput `pulumi:"value"`
}

func (SubnetGroupTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroupTag)(nil)).Elem()
}

func (i SubnetGroupTagArgs) ToSubnetGroupTagOutput() SubnetGroupTagOutput {
	return i.ToSubnetGroupTagOutputWithContext(context.Background())
}

func (i SubnetGroupTagArgs) ToSubnetGroupTagOutputWithContext(ctx context.Context) SubnetGroupTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupTagOutput)
}

// SubnetGroupTagArrayInput is an input type that accepts SubnetGroupTagArray and SubnetGroupTagArrayOutput values.
// You can construct a concrete instance of `SubnetGroupTagArrayInput` via:
//
//	SubnetGroupTagArray{ SubnetGroupTagArgs{...} }
type SubnetGroupTagArrayInput interface {
	pulumi.Input

	ToSubnetGroupTagArrayOutput() SubnetGroupTagArrayOutput
	ToSubnetGroupTagArrayOutputWithContext(context.Context) SubnetGroupTagArrayOutput
}

type SubnetGroupTagArray []SubnetGroupTagInput

func (SubnetGroupTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetGroupTag)(nil)).Elem()
}

func (i SubnetGroupTagArray) ToSubnetGroupTagArrayOutput() SubnetGroupTagArrayOutput {
	return i.ToSubnetGroupTagArrayOutputWithContext(context.Background())
}

func (i SubnetGroupTagArray) ToSubnetGroupTagArrayOutputWithContext(ctx context.Context) SubnetGroupTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupTagArrayOutput)
}

// A key-value pair to associate with a resource.
type SubnetGroupTagOutput struct{ *pulumi.OutputState }

func (SubnetGroupTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroupTag)(nil)).Elem()
}

func (o SubnetGroupTagOutput) ToSubnetGroupTagOutput() SubnetGroupTagOutput {
	return o
}

func (o SubnetGroupTagOutput) ToSubnetGroupTagOutputWithContext(ctx context.Context) SubnetGroupTagOutput {
	return o
}

// The key for the tag. May not be null.
func (o SubnetGroupTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetGroupTag) string { return v.Key }).(pulumi.StringOutput)
}

// The tag's value. May be null.
func (o SubnetGroupTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetGroupTag) string { return v.Value }).(pulumi.StringOutput)
}

type SubnetGroupTagArrayOutput struct{ *pulumi.OutputState }

func (SubnetGroupTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetGroupTag)(nil)).Elem()
}

func (o SubnetGroupTagArrayOutput) ToSubnetGroupTagArrayOutput() SubnetGroupTagArrayOutput {
	return o
}

func (o SubnetGroupTagArrayOutput) ToSubnetGroupTagArrayOutputWithContext(ctx context.Context) SubnetGroupTagArrayOutput {
	return o
}

func (o SubnetGroupTagArrayOutput) Index(i pulumi.IntInput) SubnetGroupTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetGroupTag {
		return vs[0].([]SubnetGroupTag)[vs[1].(int)]
	}).(SubnetGroupTagOutput)
}

// A key-value pair to associate with a resource.
type UserTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value *string `pulumi:"value"`
}

// UserTagInput is an input type that accepts UserTagArgs and UserTagOutput values.
// You can construct a concrete instance of `UserTagInput` via:
//
//	UserTagArgs{...}
type UserTagInput interface {
	pulumi.Input

	ToUserTagOutput() UserTagOutput
	ToUserTagOutputWithContext(context.Context) UserTagOutput
}

// A key-value pair to associate with a resource.
type UserTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (UserTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserTag)(nil)).Elem()
}

func (i UserTagArgs) ToUserTagOutput() UserTagOutput {
	return i.ToUserTagOutputWithContext(context.Background())
}

func (i UserTagArgs) ToUserTagOutputWithContext(ctx context.Context) UserTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserTagOutput)
}

// UserTagArrayInput is an input type that accepts UserTagArray and UserTagArrayOutput values.
// You can construct a concrete instance of `UserTagArrayInput` via:
//
//	UserTagArray{ UserTagArgs{...} }
type UserTagArrayInput interface {
	pulumi.Input

	ToUserTagArrayOutput() UserTagArrayOutput
	ToUserTagArrayOutputWithContext(context.Context) UserTagArrayOutput
}

type UserTagArray []UserTagInput

func (UserTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserTag)(nil)).Elem()
}

func (i UserTagArray) ToUserTagArrayOutput() UserTagArrayOutput {
	return i.ToUserTagArrayOutputWithContext(context.Background())
}

func (i UserTagArray) ToUserTagArrayOutputWithContext(ctx context.Context) UserTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserTagArrayOutput)
}

// A key-value pair to associate with a resource.
type UserTagOutput struct{ *pulumi.OutputState }

func (UserTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserTag)(nil)).Elem()
}

func (o UserTagOutput) ToUserTagOutput() UserTagOutput {
	return o
}

func (o UserTagOutput) ToUserTagOutputWithContext(ctx context.Context) UserTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o UserTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v UserTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o UserTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type UserTagArrayOutput struct{ *pulumi.OutputState }

func (UserTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserTag)(nil)).Elem()
}

func (o UserTagArrayOutput) ToUserTagArrayOutput() UserTagArrayOutput {
	return o
}

func (o UserTagArrayOutput) ToUserTagArrayOutputWithContext(ctx context.Context) UserTagArrayOutput {
	return o
}

func (o UserTagArrayOutput) Index(i pulumi.IntInput) UserTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserTag {
		return vs[0].([]UserTag)[vs[1].(int)]
	}).(UserTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ACLTagInput)(nil)).Elem(), ACLTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ACLTagArrayInput)(nil)).Elem(), ACLTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationModePropertiesInput)(nil)).Elem(), AuthenticationModePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationModePropertiesPtrInput)(nil)).Elem(), AuthenticationModePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointInput)(nil)).Elem(), ClusterEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointPtrInput)(nil)).Elem(), ClusterEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTagInput)(nil)).Elem(), ClusterTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterTagArrayInput)(nil)).Elem(), ClusterTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupTagInput)(nil)).Elem(), ParameterGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupTagArrayInput)(nil)).Elem(), ParameterGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetGroupTagInput)(nil)).Elem(), SubnetGroupTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetGroupTagArrayInput)(nil)).Elem(), SubnetGroupTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserTagInput)(nil)).Elem(), UserTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserTagArrayInput)(nil)).Elem(), UserTagArray{})
	pulumi.RegisterOutputType(ACLTagOutput{})
	pulumi.RegisterOutputType(ACLTagArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationModePropertiesOutput{})
	pulumi.RegisterOutputType(AuthenticationModePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterEndpointOutput{})
	pulumi.RegisterOutputType(ClusterEndpointPtrOutput{})
	pulumi.RegisterOutputType(ClusterTagOutput{})
	pulumi.RegisterOutputType(ClusterTagArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupTagOutput{})
	pulumi.RegisterOutputType(ParameterGroupTagArrayOutput{})
	pulumi.RegisterOutputType(SubnetGroupTagOutput{})
	pulumi.RegisterOutputType(SubnetGroupTagArrayOutput{})
	pulumi.RegisterOutputType(UserTagOutput{})
	pulumi.RegisterOutputType(UserTagArrayOutput{})
}
