// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppIntegrations::DataIntegration
type DataIntegration struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the data integration.
	DataIntegrationArn pulumi.StringOutput `pulumi:"dataIntegrationArn"`
	// The data integration description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The KMS key of the data integration.
	KmsKey pulumi.StringOutput `pulumi:"kmsKey"`
	// The name of the data integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the data and how often it should be pulled from the source.
	ScheduleConfig DataIntegrationScheduleConfigOutput `pulumi:"scheduleConfig"`
	// The URI of the data source.
	SourceURI pulumi.StringOutput `pulumi:"sourceURI"`
	// The tags (keys and values) associated with the data integration.
	Tags DataIntegrationTagArrayOutput `pulumi:"tags"`
}

// NewDataIntegration registers a new resource with the given unique name, arguments, and options.
func NewDataIntegration(ctx *pulumi.Context,
	name string, args *DataIntegrationArgs, opts ...pulumi.ResourceOption) (*DataIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KmsKey == nil {
		return nil, errors.New("invalid value for required argument 'KmsKey'")
	}
	if args.ScheduleConfig == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleConfig'")
	}
	if args.SourceURI == nil {
		return nil, errors.New("invalid value for required argument 'SourceURI'")
	}
	var resource DataIntegration
	err := ctx.RegisterResource("aws-native:appintegrations:DataIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataIntegration gets an existing DataIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataIntegrationState, opts ...pulumi.ResourceOption) (*DataIntegration, error) {
	var resource DataIntegration
	err := ctx.ReadResource("aws-native:appintegrations:DataIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataIntegration resources.
type dataIntegrationState struct {
}

type DataIntegrationState struct {
}

func (DataIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataIntegrationState)(nil)).Elem()
}

type dataIntegrationArgs struct {
	// The data integration description.
	Description *string `pulumi:"description"`
	// The KMS key of the data integration.
	KmsKey string `pulumi:"kmsKey"`
	// The name of the data integration.
	Name *string `pulumi:"name"`
	// The name of the data and how often it should be pulled from the source.
	ScheduleConfig DataIntegrationScheduleConfig `pulumi:"scheduleConfig"`
	// The URI of the data source.
	SourceURI string `pulumi:"sourceURI"`
	// The tags (keys and values) associated with the data integration.
	Tags []DataIntegrationTag `pulumi:"tags"`
}

// The set of arguments for constructing a DataIntegration resource.
type DataIntegrationArgs struct {
	// The data integration description.
	Description pulumi.StringPtrInput
	// The KMS key of the data integration.
	KmsKey pulumi.StringInput
	// The name of the data integration.
	Name pulumi.StringPtrInput
	// The name of the data and how often it should be pulled from the source.
	ScheduleConfig DataIntegrationScheduleConfigInput
	// The URI of the data source.
	SourceURI pulumi.StringInput
	// The tags (keys and values) associated with the data integration.
	Tags DataIntegrationTagArrayInput
}

func (DataIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataIntegrationArgs)(nil)).Elem()
}

type DataIntegrationInput interface {
	pulumi.Input

	ToDataIntegrationOutput() DataIntegrationOutput
	ToDataIntegrationOutputWithContext(ctx context.Context) DataIntegrationOutput
}

func (*DataIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**DataIntegration)(nil)).Elem()
}

func (i *DataIntegration) ToDataIntegrationOutput() DataIntegrationOutput {
	return i.ToDataIntegrationOutputWithContext(context.Background())
}

func (i *DataIntegration) ToDataIntegrationOutputWithContext(ctx context.Context) DataIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationOutput)
}

type DataIntegrationOutput struct{ *pulumi.OutputState }

func (DataIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataIntegration)(nil)).Elem()
}

func (o DataIntegrationOutput) ToDataIntegrationOutput() DataIntegrationOutput {
	return o
}

func (o DataIntegrationOutput) ToDataIntegrationOutputWithContext(ctx context.Context) DataIntegrationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationInput)(nil)).Elem(), &DataIntegration{})
	pulumi.RegisterOutputType(DataIntegrationOutput{})
}