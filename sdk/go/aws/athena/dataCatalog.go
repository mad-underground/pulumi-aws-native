// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Athena::DataCatalog
type DataCatalog struct {
	pulumi.CustomResourceState

	// A description of the data catalog to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type.
	Parameters pulumi.AnyOutput `pulumi:"parameters"`
	// A list of comma separated tags to add to the data catalog that is created.
	Tags DataCatalogTagArrayOutput `pulumi:"tags"`
	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	Type DataCatalogTypeOutput `pulumi:"type"`
}

// NewDataCatalog registers a new resource with the given unique name, arguments, and options.
func NewDataCatalog(ctx *pulumi.Context,
	name string, args *DataCatalogArgs, opts ...pulumi.ResourceOption) (*DataCatalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataCatalog
	err := ctx.RegisterResource("aws-native:athena:DataCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalog gets an existing DataCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogState, opts ...pulumi.ResourceOption) (*DataCatalog, error) {
	var resource DataCatalog
	err := ctx.ReadResource("aws-native:athena:DataCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalog resources.
type dataCatalogState struct {
}

type DataCatalogState struct {
}

func (DataCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogState)(nil)).Elem()
}

type dataCatalogArgs struct {
	// A description of the data catalog to be created.
	Description *string `pulumi:"description"`
	// The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	Name *string `pulumi:"name"`
	// Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type.
	Parameters interface{} `pulumi:"parameters"`
	// A list of comma separated tags to add to the data catalog that is created.
	Tags []DataCatalogTag `pulumi:"tags"`
	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	Type DataCatalogType `pulumi:"type"`
}

// The set of arguments for constructing a DataCatalog resource.
type DataCatalogArgs struct {
	// A description of the data catalog to be created.
	Description pulumi.StringPtrInput
	// The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
	Name pulumi.StringPtrInput
	// Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type.
	Parameters pulumi.Input
	// A list of comma separated tags to add to the data catalog that is created.
	Tags DataCatalogTagArrayInput
	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
	Type DataCatalogTypeInput
}

func (DataCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogArgs)(nil)).Elem()
}

type DataCatalogInput interface {
	pulumi.Input

	ToDataCatalogOutput() DataCatalogOutput
	ToDataCatalogOutputWithContext(ctx context.Context) DataCatalogOutput
}

func (*DataCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalog)(nil)).Elem()
}

func (i *DataCatalog) ToDataCatalogOutput() DataCatalogOutput {
	return i.ToDataCatalogOutputWithContext(context.Background())
}

func (i *DataCatalog) ToDataCatalogOutputWithContext(ctx context.Context) DataCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogOutput)
}

type DataCatalogOutput struct{ *pulumi.OutputState }

func (DataCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalog)(nil)).Elem()
}

func (o DataCatalogOutput) ToDataCatalogOutput() DataCatalogOutput {
	return o
}

func (o DataCatalogOutput) ToDataCatalogOutputWithContext(ctx context.Context) DataCatalogOutput {
	return o
}

// A description of the data catalog to be created.
func (o DataCatalogOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCatalog) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters.
func (o DataCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type.
func (o DataCatalogOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataCatalog) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

// A list of comma separated tags to add to the data catalog that is created.
func (o DataCatalogOutput) Tags() DataCatalogTagArrayOutput {
	return o.ApplyT(func(v *DataCatalog) DataCatalogTagArrayOutput { return v.Tags }).(DataCatalogTagArrayOutput)
}

// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore.
func (o DataCatalogOutput) Type() DataCatalogTypeOutput {
	return o.ApplyT(func(v *DataCatalog) DataCatalogTypeOutput { return v.Type }).(DataCatalogTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCatalogInput)(nil)).Elem(), &DataCatalog{})
	pulumi.RegisterOutputType(DataCatalogOutput{})
}
