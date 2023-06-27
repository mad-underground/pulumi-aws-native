// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KeyspaceRegionListItem string

const (
	KeyspaceRegionListItemApNortheast1 = KeyspaceRegionListItem("ap-northeast-1")
	KeyspaceRegionListItemApNortheast2 = KeyspaceRegionListItem("ap-northeast-2")
	KeyspaceRegionListItemApSouth1     = KeyspaceRegionListItem("ap-south-1")
	KeyspaceRegionListItemApSoutheast1 = KeyspaceRegionListItem("ap-southeast-1")
	KeyspaceRegionListItemApSoutheast2 = KeyspaceRegionListItem("ap-southeast-2")
	KeyspaceRegionListItemCaCentral1   = KeyspaceRegionListItem("ca-central-1")
	KeyspaceRegionListItemEuCentral1   = KeyspaceRegionListItem("eu-central-1")
	KeyspaceRegionListItemEuNorth1     = KeyspaceRegionListItem("eu-north-1")
	KeyspaceRegionListItemEuWest1      = KeyspaceRegionListItem("eu-west-1")
	KeyspaceRegionListItemEuWest2      = KeyspaceRegionListItem("eu-west-2")
	KeyspaceRegionListItemEuWest3      = KeyspaceRegionListItem("eu-west-3")
	KeyspaceRegionListItemSaEast1      = KeyspaceRegionListItem("sa-east-1")
	KeyspaceRegionListItemUsEast1      = KeyspaceRegionListItem("us-east-1")
	KeyspaceRegionListItemUsEast2      = KeyspaceRegionListItem("us-east-2")
	KeyspaceRegionListItemUsWest1      = KeyspaceRegionListItem("us-west-1")
	KeyspaceRegionListItemUsWest2      = KeyspaceRegionListItem("us-west-2")
)

func (KeyspaceRegionListItem) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyspaceRegionListItem)(nil)).Elem()
}

func (e KeyspaceRegionListItem) ToKeyspaceRegionListItemOutput() KeyspaceRegionListItemOutput {
	return pulumi.ToOutput(e).(KeyspaceRegionListItemOutput)
}

func (e KeyspaceRegionListItem) ToKeyspaceRegionListItemOutputWithContext(ctx context.Context) KeyspaceRegionListItemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyspaceRegionListItemOutput)
}

func (e KeyspaceRegionListItem) ToKeyspaceRegionListItemPtrOutput() KeyspaceRegionListItemPtrOutput {
	return e.ToKeyspaceRegionListItemPtrOutputWithContext(context.Background())
}

func (e KeyspaceRegionListItem) ToKeyspaceRegionListItemPtrOutputWithContext(ctx context.Context) KeyspaceRegionListItemPtrOutput {
	return KeyspaceRegionListItem(e).ToKeyspaceRegionListItemOutputWithContext(ctx).ToKeyspaceRegionListItemPtrOutputWithContext(ctx)
}

func (e KeyspaceRegionListItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyspaceRegionListItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyspaceRegionListItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyspaceRegionListItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyspaceRegionListItemOutput struct{ *pulumi.OutputState }

func (KeyspaceRegionListItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyspaceRegionListItem)(nil)).Elem()
}

func (o KeyspaceRegionListItemOutput) ToKeyspaceRegionListItemOutput() KeyspaceRegionListItemOutput {
	return o
}

func (o KeyspaceRegionListItemOutput) ToKeyspaceRegionListItemOutputWithContext(ctx context.Context) KeyspaceRegionListItemOutput {
	return o
}

func (o KeyspaceRegionListItemOutput) ToKeyspaceRegionListItemPtrOutput() KeyspaceRegionListItemPtrOutput {
	return o.ToKeyspaceRegionListItemPtrOutputWithContext(context.Background())
}

func (o KeyspaceRegionListItemOutput) ToKeyspaceRegionListItemPtrOutputWithContext(ctx context.Context) KeyspaceRegionListItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyspaceRegionListItem) *KeyspaceRegionListItem {
		return &v
	}).(KeyspaceRegionListItemPtrOutput)
}

func (o KeyspaceRegionListItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyspaceRegionListItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyspaceRegionListItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyspaceRegionListItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyspaceRegionListItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyspaceRegionListItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyspaceRegionListItemPtrOutput struct{ *pulumi.OutputState }

func (KeyspaceRegionListItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyspaceRegionListItem)(nil)).Elem()
}

func (o KeyspaceRegionListItemPtrOutput) ToKeyspaceRegionListItemPtrOutput() KeyspaceRegionListItemPtrOutput {
	return o
}

func (o KeyspaceRegionListItemPtrOutput) ToKeyspaceRegionListItemPtrOutputWithContext(ctx context.Context) KeyspaceRegionListItemPtrOutput {
	return o
}

func (o KeyspaceRegionListItemPtrOutput) Elem() KeyspaceRegionListItemOutput {
	return o.ApplyT(func(v *KeyspaceRegionListItem) KeyspaceRegionListItem {
		if v != nil {
			return *v
		}
		var ret KeyspaceRegionListItem
		return ret
	}).(KeyspaceRegionListItemOutput)
}

func (o KeyspaceRegionListItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyspaceRegionListItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyspaceRegionListItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyspaceRegionListItemInput is an input type that accepts KeyspaceRegionListItemArgs and KeyspaceRegionListItemOutput values.
// You can construct a concrete instance of `KeyspaceRegionListItemInput` via:
//
//	KeyspaceRegionListItemArgs{...}
type KeyspaceRegionListItemInput interface {
	pulumi.Input

	ToKeyspaceRegionListItemOutput() KeyspaceRegionListItemOutput
	ToKeyspaceRegionListItemOutputWithContext(context.Context) KeyspaceRegionListItemOutput
}

var keyspaceRegionListItemPtrType = reflect.TypeOf((**KeyspaceRegionListItem)(nil)).Elem()

type KeyspaceRegionListItemPtrInput interface {
	pulumi.Input

	ToKeyspaceRegionListItemPtrOutput() KeyspaceRegionListItemPtrOutput
	ToKeyspaceRegionListItemPtrOutputWithContext(context.Context) KeyspaceRegionListItemPtrOutput
}

type keyspaceRegionListItemPtr string

func KeyspaceRegionListItemPtr(v string) KeyspaceRegionListItemPtrInput {
	return (*keyspaceRegionListItemPtr)(&v)
}

func (*keyspaceRegionListItemPtr) ElementType() reflect.Type {
	return keyspaceRegionListItemPtrType
}

func (in *keyspaceRegionListItemPtr) ToKeyspaceRegionListItemPtrOutput() KeyspaceRegionListItemPtrOutput {
	return pulumi.ToOutput(in).(KeyspaceRegionListItemPtrOutput)
}

func (in *keyspaceRegionListItemPtr) ToKeyspaceRegionListItemPtrOutputWithContext(ctx context.Context) KeyspaceRegionListItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyspaceRegionListItemPtrOutput)
}

// KeyspaceRegionListItemArrayInput is an input type that accepts KeyspaceRegionListItemArray and KeyspaceRegionListItemArrayOutput values.
// You can construct a concrete instance of `KeyspaceRegionListItemArrayInput` via:
//
//	KeyspaceRegionListItemArray{ KeyspaceRegionListItemArgs{...} }
type KeyspaceRegionListItemArrayInput interface {
	pulumi.Input

	ToKeyspaceRegionListItemArrayOutput() KeyspaceRegionListItemArrayOutput
	ToKeyspaceRegionListItemArrayOutputWithContext(context.Context) KeyspaceRegionListItemArrayOutput
}

type KeyspaceRegionListItemArray []KeyspaceRegionListItem

func (KeyspaceRegionListItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyspaceRegionListItem)(nil)).Elem()
}

func (i KeyspaceRegionListItemArray) ToKeyspaceRegionListItemArrayOutput() KeyspaceRegionListItemArrayOutput {
	return i.ToKeyspaceRegionListItemArrayOutputWithContext(context.Background())
}

func (i KeyspaceRegionListItemArray) ToKeyspaceRegionListItemArrayOutputWithContext(ctx context.Context) KeyspaceRegionListItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyspaceRegionListItemArrayOutput)
}

type KeyspaceRegionListItemArrayOutput struct{ *pulumi.OutputState }

func (KeyspaceRegionListItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyspaceRegionListItem)(nil)).Elem()
}

func (o KeyspaceRegionListItemArrayOutput) ToKeyspaceRegionListItemArrayOutput() KeyspaceRegionListItemArrayOutput {
	return o
}

func (o KeyspaceRegionListItemArrayOutput) ToKeyspaceRegionListItemArrayOutputWithContext(ctx context.Context) KeyspaceRegionListItemArrayOutput {
	return o
}

func (o KeyspaceRegionListItemArrayOutput) Index(i pulumi.IntInput) KeyspaceRegionListItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyspaceRegionListItem {
		return vs[0].([]KeyspaceRegionListItem)[vs[1].(int)]
	}).(KeyspaceRegionListItemOutput)
}

type KeyspaceReplicationSpecificationReplicationStrategy string

const (
	KeyspaceReplicationSpecificationReplicationStrategySingleRegion = KeyspaceReplicationSpecificationReplicationStrategy("SINGLE_REGION")
	KeyspaceReplicationSpecificationReplicationStrategyMultiRegion  = KeyspaceReplicationSpecificationReplicationStrategy("MULTI_REGION")
)

func (KeyspaceReplicationSpecificationReplicationStrategy) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyspaceReplicationSpecificationReplicationStrategy)(nil)).Elem()
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToKeyspaceReplicationSpecificationReplicationStrategyOutput() KeyspaceReplicationSpecificationReplicationStrategyOutput {
	return pulumi.ToOutput(e).(KeyspaceReplicationSpecificationReplicationStrategyOutput)
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToKeyspaceReplicationSpecificationReplicationStrategyOutputWithContext(ctx context.Context) KeyspaceReplicationSpecificationReplicationStrategyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyspaceReplicationSpecificationReplicationStrategyOutput)
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutput() KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return e.ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(context.Background())
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(ctx context.Context) KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return KeyspaceReplicationSpecificationReplicationStrategy(e).ToKeyspaceReplicationSpecificationReplicationStrategyOutputWithContext(ctx).ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(ctx)
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyspaceReplicationSpecificationReplicationStrategy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyspaceReplicationSpecificationReplicationStrategyOutput struct{ *pulumi.OutputState }

func (KeyspaceReplicationSpecificationReplicationStrategyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyspaceReplicationSpecificationReplicationStrategy)(nil)).Elem()
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToKeyspaceReplicationSpecificationReplicationStrategyOutput() KeyspaceReplicationSpecificationReplicationStrategyOutput {
	return o
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToKeyspaceReplicationSpecificationReplicationStrategyOutputWithContext(ctx context.Context) KeyspaceReplicationSpecificationReplicationStrategyOutput {
	return o
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutput() KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return o.ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(context.Background())
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(ctx context.Context) KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyspaceReplicationSpecificationReplicationStrategy) *KeyspaceReplicationSpecificationReplicationStrategy {
		return &v
	}).(KeyspaceReplicationSpecificationReplicationStrategyPtrOutput)
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyspaceReplicationSpecificationReplicationStrategy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyspaceReplicationSpecificationReplicationStrategyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyspaceReplicationSpecificationReplicationStrategy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyspaceReplicationSpecificationReplicationStrategyPtrOutput struct{ *pulumi.OutputState }

func (KeyspaceReplicationSpecificationReplicationStrategyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyspaceReplicationSpecificationReplicationStrategy)(nil)).Elem()
}

func (o KeyspaceReplicationSpecificationReplicationStrategyPtrOutput) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutput() KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return o
}

func (o KeyspaceReplicationSpecificationReplicationStrategyPtrOutput) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(ctx context.Context) KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return o
}

func (o KeyspaceReplicationSpecificationReplicationStrategyPtrOutput) Elem() KeyspaceReplicationSpecificationReplicationStrategyOutput {
	return o.ApplyT(func(v *KeyspaceReplicationSpecificationReplicationStrategy) KeyspaceReplicationSpecificationReplicationStrategy {
		if v != nil {
			return *v
		}
		var ret KeyspaceReplicationSpecificationReplicationStrategy
		return ret
	}).(KeyspaceReplicationSpecificationReplicationStrategyOutput)
}

func (o KeyspaceReplicationSpecificationReplicationStrategyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyspaceReplicationSpecificationReplicationStrategyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyspaceReplicationSpecificationReplicationStrategy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyspaceReplicationSpecificationReplicationStrategyInput is an input type that accepts KeyspaceReplicationSpecificationReplicationStrategyArgs and KeyspaceReplicationSpecificationReplicationStrategyOutput values.
// You can construct a concrete instance of `KeyspaceReplicationSpecificationReplicationStrategyInput` via:
//
//	KeyspaceReplicationSpecificationReplicationStrategyArgs{...}
type KeyspaceReplicationSpecificationReplicationStrategyInput interface {
	pulumi.Input

	ToKeyspaceReplicationSpecificationReplicationStrategyOutput() KeyspaceReplicationSpecificationReplicationStrategyOutput
	ToKeyspaceReplicationSpecificationReplicationStrategyOutputWithContext(context.Context) KeyspaceReplicationSpecificationReplicationStrategyOutput
}

var keyspaceReplicationSpecificationReplicationStrategyPtrType = reflect.TypeOf((**KeyspaceReplicationSpecificationReplicationStrategy)(nil)).Elem()

type KeyspaceReplicationSpecificationReplicationStrategyPtrInput interface {
	pulumi.Input

	ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutput() KeyspaceReplicationSpecificationReplicationStrategyPtrOutput
	ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(context.Context) KeyspaceReplicationSpecificationReplicationStrategyPtrOutput
}

type keyspaceReplicationSpecificationReplicationStrategyPtr string

func KeyspaceReplicationSpecificationReplicationStrategyPtr(v string) KeyspaceReplicationSpecificationReplicationStrategyPtrInput {
	return (*keyspaceReplicationSpecificationReplicationStrategyPtr)(&v)
}

func (*keyspaceReplicationSpecificationReplicationStrategyPtr) ElementType() reflect.Type {
	return keyspaceReplicationSpecificationReplicationStrategyPtrType
}

func (in *keyspaceReplicationSpecificationReplicationStrategyPtr) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutput() KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return pulumi.ToOutput(in).(KeyspaceReplicationSpecificationReplicationStrategyPtrOutput)
}

func (in *keyspaceReplicationSpecificationReplicationStrategyPtr) ToKeyspaceReplicationSpecificationReplicationStrategyPtrOutputWithContext(ctx context.Context) KeyspaceReplicationSpecificationReplicationStrategyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyspaceReplicationSpecificationReplicationStrategyPtrOutput)
}

type TableClusteringKeyColumnOrderBy string

const (
	TableClusteringKeyColumnOrderByAsc  = TableClusteringKeyColumnOrderBy("ASC")
	TableClusteringKeyColumnOrderByDesc = TableClusteringKeyColumnOrderBy("DESC")
)

func (TableClusteringKeyColumnOrderBy) ElementType() reflect.Type {
	return reflect.TypeOf((*TableClusteringKeyColumnOrderBy)(nil)).Elem()
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByOutput() TableClusteringKeyColumnOrderByOutput {
	return pulumi.ToOutput(e).(TableClusteringKeyColumnOrderByOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableClusteringKeyColumnOrderByOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return e.ToTableClusteringKeyColumnOrderByPtrOutputWithContext(context.Background())
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return TableClusteringKeyColumnOrderBy(e).ToTableClusteringKeyColumnOrderByOutputWithContext(ctx).ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx)
}

func (e TableClusteringKeyColumnOrderBy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableClusteringKeyColumnOrderBy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableClusteringKeyColumnOrderByOutput struct{ *pulumi.OutputState }

func (TableClusteringKeyColumnOrderByOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableClusteringKeyColumnOrderBy)(nil)).Elem()
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByOutput() TableClusteringKeyColumnOrderByOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return o.ToTableClusteringKeyColumnOrderByPtrOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableClusteringKeyColumnOrderBy) *TableClusteringKeyColumnOrderBy {
		return &v
	}).(TableClusteringKeyColumnOrderByPtrOutput)
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableClusteringKeyColumnOrderBy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableClusteringKeyColumnOrderBy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableClusteringKeyColumnOrderByPtrOutput struct{ *pulumi.OutputState }

func (TableClusteringKeyColumnOrderByPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableClusteringKeyColumnOrderBy)(nil)).Elem()
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByPtrOutput) Elem() TableClusteringKeyColumnOrderByOutput {
	return o.ApplyT(func(v *TableClusteringKeyColumnOrderBy) TableClusteringKeyColumnOrderBy {
		if v != nil {
			return *v
		}
		var ret TableClusteringKeyColumnOrderBy
		return ret
	}).(TableClusteringKeyColumnOrderByOutput)
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableClusteringKeyColumnOrderBy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableClusteringKeyColumnOrderByInput is an input type that accepts TableClusteringKeyColumnOrderByArgs and TableClusteringKeyColumnOrderByOutput values.
// You can construct a concrete instance of `TableClusteringKeyColumnOrderByInput` via:
//
//	TableClusteringKeyColumnOrderByArgs{...}
type TableClusteringKeyColumnOrderByInput interface {
	pulumi.Input

	ToTableClusteringKeyColumnOrderByOutput() TableClusteringKeyColumnOrderByOutput
	ToTableClusteringKeyColumnOrderByOutputWithContext(context.Context) TableClusteringKeyColumnOrderByOutput
}

var tableClusteringKeyColumnOrderByPtrType = reflect.TypeOf((**TableClusteringKeyColumnOrderBy)(nil)).Elem()

type TableClusteringKeyColumnOrderByPtrInput interface {
	pulumi.Input

	ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput
	ToTableClusteringKeyColumnOrderByPtrOutputWithContext(context.Context) TableClusteringKeyColumnOrderByPtrOutput
}

type tableClusteringKeyColumnOrderByPtr string

func TableClusteringKeyColumnOrderByPtr(v string) TableClusteringKeyColumnOrderByPtrInput {
	return (*tableClusteringKeyColumnOrderByPtr)(&v)
}

func (*tableClusteringKeyColumnOrderByPtr) ElementType() reflect.Type {
	return tableClusteringKeyColumnOrderByPtrType
}

func (in *tableClusteringKeyColumnOrderByPtr) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return pulumi.ToOutput(in).(TableClusteringKeyColumnOrderByPtrOutput)
}

func (in *tableClusteringKeyColumnOrderByPtr) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableClusteringKeyColumnOrderByPtrOutput)
}

// Server-side encryption type
type TableEncryptionType string

const (
	TableEncryptionTypeAwsOwnedKmsKey        = TableEncryptionType("AWS_OWNED_KMS_KEY")
	TableEncryptionTypeCustomerManagedKmsKey = TableEncryptionType("CUSTOMER_MANAGED_KMS_KEY")
)

func (TableEncryptionType) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEncryptionType)(nil)).Elem()
}

func (e TableEncryptionType) ToTableEncryptionTypeOutput() TableEncryptionTypeOutput {
	return pulumi.ToOutput(e).(TableEncryptionTypeOutput)
}

func (e TableEncryptionType) ToTableEncryptionTypeOutputWithContext(ctx context.Context) TableEncryptionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableEncryptionTypeOutput)
}

func (e TableEncryptionType) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return e.ToTableEncryptionTypePtrOutputWithContext(context.Background())
}

func (e TableEncryptionType) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return TableEncryptionType(e).ToTableEncryptionTypeOutputWithContext(ctx).ToTableEncryptionTypePtrOutputWithContext(ctx)
}

func (e TableEncryptionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableEncryptionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableEncryptionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableEncryptionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableEncryptionTypeOutput struct{ *pulumi.OutputState }

func (TableEncryptionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEncryptionType)(nil)).Elem()
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypeOutput() TableEncryptionTypeOutput {
	return o
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypeOutputWithContext(ctx context.Context) TableEncryptionTypeOutput {
	return o
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return o.ToTableEncryptionTypePtrOutputWithContext(context.Background())
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableEncryptionType) *TableEncryptionType {
		return &v
	}).(TableEncryptionTypePtrOutput)
}

func (o TableEncryptionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableEncryptionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableEncryptionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableEncryptionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableEncryptionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableEncryptionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableEncryptionTypePtrOutput struct{ *pulumi.OutputState }

func (TableEncryptionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableEncryptionType)(nil)).Elem()
}

func (o TableEncryptionTypePtrOutput) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return o
}

func (o TableEncryptionTypePtrOutput) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return o
}

func (o TableEncryptionTypePtrOutput) Elem() TableEncryptionTypeOutput {
	return o.ApplyT(func(v *TableEncryptionType) TableEncryptionType {
		if v != nil {
			return *v
		}
		var ret TableEncryptionType
		return ret
	}).(TableEncryptionTypeOutput)
}

func (o TableEncryptionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableEncryptionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableEncryptionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableEncryptionTypeInput is an input type that accepts TableEncryptionTypeArgs and TableEncryptionTypeOutput values.
// You can construct a concrete instance of `TableEncryptionTypeInput` via:
//
//	TableEncryptionTypeArgs{...}
type TableEncryptionTypeInput interface {
	pulumi.Input

	ToTableEncryptionTypeOutput() TableEncryptionTypeOutput
	ToTableEncryptionTypeOutputWithContext(context.Context) TableEncryptionTypeOutput
}

var tableEncryptionTypePtrType = reflect.TypeOf((**TableEncryptionType)(nil)).Elem()

type TableEncryptionTypePtrInput interface {
	pulumi.Input

	ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput
	ToTableEncryptionTypePtrOutputWithContext(context.Context) TableEncryptionTypePtrOutput
}

type tableEncryptionTypePtr string

func TableEncryptionTypePtr(v string) TableEncryptionTypePtrInput {
	return (*tableEncryptionTypePtr)(&v)
}

func (*tableEncryptionTypePtr) ElementType() reflect.Type {
	return tableEncryptionTypePtrType
}

func (in *tableEncryptionTypePtr) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return pulumi.ToOutput(in).(TableEncryptionTypePtrOutput)
}

func (in *tableEncryptionTypePtr) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableEncryptionTypePtrOutput)
}

// Capacity mode for the specified table
type TableMode string

const (
	TableModeProvisioned = TableMode("PROVISIONED")
	TableModeOnDemand    = TableMode("ON_DEMAND")
)

func (TableMode) ElementType() reflect.Type {
	return reflect.TypeOf((*TableMode)(nil)).Elem()
}

func (e TableMode) ToTableModeOutput() TableModeOutput {
	return pulumi.ToOutput(e).(TableModeOutput)
}

func (e TableMode) ToTableModeOutputWithContext(ctx context.Context) TableModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableModeOutput)
}

func (e TableMode) ToTableModePtrOutput() TableModePtrOutput {
	return e.ToTableModePtrOutputWithContext(context.Background())
}

func (e TableMode) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return TableMode(e).ToTableModeOutputWithContext(ctx).ToTableModePtrOutputWithContext(ctx)
}

func (e TableMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableModeOutput struct{ *pulumi.OutputState }

func (TableModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableMode)(nil)).Elem()
}

func (o TableModeOutput) ToTableModeOutput() TableModeOutput {
	return o
}

func (o TableModeOutput) ToTableModeOutputWithContext(ctx context.Context) TableModeOutput {
	return o
}

func (o TableModeOutput) ToTableModePtrOutput() TableModePtrOutput {
	return o.ToTableModePtrOutputWithContext(context.Background())
}

func (o TableModeOutput) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableMode) *TableMode {
		return &v
	}).(TableModePtrOutput)
}

func (o TableModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableModePtrOutput struct{ *pulumi.OutputState }

func (TableModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableMode)(nil)).Elem()
}

func (o TableModePtrOutput) ToTableModePtrOutput() TableModePtrOutput {
	return o
}

func (o TableModePtrOutput) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return o
}

func (o TableModePtrOutput) Elem() TableModeOutput {
	return o.ApplyT(func(v *TableMode) TableMode {
		if v != nil {
			return *v
		}
		var ret TableMode
		return ret
	}).(TableModeOutput)
}

func (o TableModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableModeInput is an input type that accepts TableModeArgs and TableModeOutput values.
// You can construct a concrete instance of `TableModeInput` via:
//
//	TableModeArgs{...}
type TableModeInput interface {
	pulumi.Input

	ToTableModeOutput() TableModeOutput
	ToTableModeOutputWithContext(context.Context) TableModeOutput
}

var tableModePtrType = reflect.TypeOf((**TableMode)(nil)).Elem()

type TableModePtrInput interface {
	pulumi.Input

	ToTableModePtrOutput() TableModePtrOutput
	ToTableModePtrOutputWithContext(context.Context) TableModePtrOutput
}

type tableModePtr string

func TableModePtr(v string) TableModePtrInput {
	return (*tableModePtr)(&v)
}

func (*tableModePtr) ElementType() reflect.Type {
	return tableModePtrType
}

func (in *tableModePtr) ToTableModePtrOutput() TableModePtrOutput {
	return pulumi.ToOutput(in).(TableModePtrOutput)
}

func (in *tableModePtr) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableModePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceRegionListItemInput)(nil)).Elem(), KeyspaceRegionListItem("ap-northeast-1"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceRegionListItemPtrInput)(nil)).Elem(), KeyspaceRegionListItem("ap-northeast-1"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceRegionListItemArrayInput)(nil)).Elem(), KeyspaceRegionListItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceReplicationSpecificationReplicationStrategyInput)(nil)).Elem(), KeyspaceReplicationSpecificationReplicationStrategy("SINGLE_REGION"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceReplicationSpecificationReplicationStrategyPtrInput)(nil)).Elem(), KeyspaceReplicationSpecificationReplicationStrategy("SINGLE_REGION"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableClusteringKeyColumnOrderByInput)(nil)).Elem(), TableClusteringKeyColumnOrderBy("ASC"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableClusteringKeyColumnOrderByPtrInput)(nil)).Elem(), TableClusteringKeyColumnOrderBy("ASC"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableEncryptionTypeInput)(nil)).Elem(), TableEncryptionType("AWS_OWNED_KMS_KEY"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableEncryptionTypePtrInput)(nil)).Elem(), TableEncryptionType("AWS_OWNED_KMS_KEY"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableModeInput)(nil)).Elem(), TableMode("PROVISIONED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableModePtrInput)(nil)).Elem(), TableMode("PROVISIONED"))
	pulumi.RegisterOutputType(KeyspaceRegionListItemOutput{})
	pulumi.RegisterOutputType(KeyspaceRegionListItemPtrOutput{})
	pulumi.RegisterOutputType(KeyspaceRegionListItemArrayOutput{})
	pulumi.RegisterOutputType(KeyspaceReplicationSpecificationReplicationStrategyOutput{})
	pulumi.RegisterOutputType(KeyspaceReplicationSpecificationReplicationStrategyPtrOutput{})
	pulumi.RegisterOutputType(TableClusteringKeyColumnOrderByOutput{})
	pulumi.RegisterOutputType(TableClusteringKeyColumnOrderByPtrOutput{})
	pulumi.RegisterOutputType(TableEncryptionTypeOutput{})
	pulumi.RegisterOutputType(TableEncryptionTypePtrOutput{})
	pulumi.RegisterOutputType(TableModeOutput{})
	pulumi.RegisterOutputType(TableModePtrOutput{})
}
