// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::MaintenanceWindowTask
//
// Deprecated: MaintenanceWindowTask is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type MaintenanceWindowTask struct {
	pulumi.CustomResourceState

	Description              pulumi.StringPtrOutput                                 `pulumi:"description"`
	LoggingInfo              MaintenanceWindowTaskLoggingInfoPtrOutput              `pulumi:"loggingInfo"`
	MaxConcurrency           pulumi.StringPtrOutput                                 `pulumi:"maxConcurrency"`
	MaxErrors                pulumi.StringPtrOutput                                 `pulumi:"maxErrors"`
	Name                     pulumi.StringPtrOutput                                 `pulumi:"name"`
	Priority                 pulumi.IntOutput                                       `pulumi:"priority"`
	ServiceRoleArn           pulumi.StringPtrOutput                                 `pulumi:"serviceRoleArn"`
	Targets                  MaintenanceWindowTaskTargetArrayOutput                 `pulumi:"targets"`
	TaskArn                  pulumi.StringOutput                                    `pulumi:"taskArn"`
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrOutput `pulumi:"taskInvocationParameters"`
	TaskParameters           pulumi.AnyOutput                                       `pulumi:"taskParameters"`
	TaskType                 pulumi.StringOutput                                    `pulumi:"taskType"`
	WindowId                 pulumi.StringOutput                                    `pulumi:"windowId"`
}

// NewMaintenanceWindowTask registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindowTask(ctx *pulumi.Context,
	name string, args *MaintenanceWindowTaskArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindowTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.TaskArn == nil {
		return nil, errors.New("invalid value for required argument 'TaskArn'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	if args.WindowId == nil {
		return nil, errors.New("invalid value for required argument 'WindowId'")
	}
	var resource MaintenanceWindowTask
	err := ctx.RegisterResource("aws-native:ssm:MaintenanceWindowTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindowTask gets an existing MaintenanceWindowTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindowTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowTaskState, opts ...pulumi.ResourceOption) (*MaintenanceWindowTask, error) {
	var resource MaintenanceWindowTask
	err := ctx.ReadResource("aws-native:ssm:MaintenanceWindowTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindowTask resources.
type maintenanceWindowTaskState struct {
}

type MaintenanceWindowTaskState struct {
}

func (MaintenanceWindowTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTaskState)(nil)).Elem()
}

type maintenanceWindowTaskArgs struct {
	Description              *string                                        `pulumi:"description"`
	LoggingInfo              *MaintenanceWindowTaskLoggingInfo              `pulumi:"loggingInfo"`
	MaxConcurrency           *string                                        `pulumi:"maxConcurrency"`
	MaxErrors                *string                                        `pulumi:"maxErrors"`
	Name                     *string                                        `pulumi:"name"`
	Priority                 int                                            `pulumi:"priority"`
	ServiceRoleArn           *string                                        `pulumi:"serviceRoleArn"`
	Targets                  []MaintenanceWindowTaskTarget                  `pulumi:"targets"`
	TaskArn                  string                                         `pulumi:"taskArn"`
	TaskInvocationParameters *MaintenanceWindowTaskTaskInvocationParameters `pulumi:"taskInvocationParameters"`
	TaskParameters           interface{}                                    `pulumi:"taskParameters"`
	TaskType                 string                                         `pulumi:"taskType"`
	WindowId                 string                                         `pulumi:"windowId"`
}

// The set of arguments for constructing a MaintenanceWindowTask resource.
type MaintenanceWindowTaskArgs struct {
	Description              pulumi.StringPtrInput
	LoggingInfo              MaintenanceWindowTaskLoggingInfoPtrInput
	MaxConcurrency           pulumi.StringPtrInput
	MaxErrors                pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	Priority                 pulumi.IntInput
	ServiceRoleArn           pulumi.StringPtrInput
	Targets                  MaintenanceWindowTaskTargetArrayInput
	TaskArn                  pulumi.StringInput
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrInput
	TaskParameters           pulumi.Input
	TaskType                 pulumi.StringInput
	WindowId                 pulumi.StringInput
}

func (MaintenanceWindowTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTaskArgs)(nil)).Elem()
}

type MaintenanceWindowTaskInput interface {
	pulumi.Input

	ToMaintenanceWindowTaskOutput() MaintenanceWindowTaskOutput
	ToMaintenanceWindowTaskOutputWithContext(ctx context.Context) MaintenanceWindowTaskOutput
}

func (*MaintenanceWindowTask) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowTask)(nil))
}

func (i *MaintenanceWindowTask) ToMaintenanceWindowTaskOutput() MaintenanceWindowTaskOutput {
	return i.ToMaintenanceWindowTaskOutputWithContext(context.Background())
}

func (i *MaintenanceWindowTask) ToMaintenanceWindowTaskOutputWithContext(ctx context.Context) MaintenanceWindowTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTaskOutput)
}

type MaintenanceWindowTaskOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowTask)(nil))
}

func (o MaintenanceWindowTaskOutput) ToMaintenanceWindowTaskOutput() MaintenanceWindowTaskOutput {
	return o
}

func (o MaintenanceWindowTaskOutput) ToMaintenanceWindowTaskOutputWithContext(ctx context.Context) MaintenanceWindowTaskOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MaintenanceWindowTaskOutput{})
}