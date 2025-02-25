// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This schema is for testing purpose only.
type SimulationApplication struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The current revision id.
	CurrentRevisionId pulumi.StringPtrOutput `pulumi:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// The name of the simulation application.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The rendering engine for the simulation application.
	RenderingEngine SimulationApplicationRenderingEnginePtrOutput `pulumi:"renderingEngine"`
	// The robot software suite used by the simulation application.
	RobotSoftwareSuite SimulationApplicationRobotSoftwareSuiteOutput `pulumi:"robotSoftwareSuite"`
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite SimulationApplicationSimulationSoftwareSuiteOutput `pulumi:"simulationSoftwareSuite"`
	// The sources of the simulation application.
	Sources SimulationApplicationSourceConfigArrayOutput `pulumi:"sources"`
	Tags    pulumi.StringMapOutput                       `pulumi:"tags"`
}

// NewSimulationApplication registers a new resource with the given unique name, arguments, and options.
func NewSimulationApplication(ctx *pulumi.Context,
	name string, args *SimulationApplicationArgs, opts ...pulumi.ResourceOption) (*SimulationApplication, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RobotSoftwareSuite == nil {
		return nil, errors.New("invalid value for required argument 'RobotSoftwareSuite'")
	}
	if args.SimulationSoftwareSuite == nil {
		return nil, errors.New("invalid value for required argument 'SimulationSoftwareSuite'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SimulationApplication
	err := ctx.RegisterResource("aws-native:robomaker:SimulationApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSimulationApplication gets an existing SimulationApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSimulationApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimulationApplicationState, opts ...pulumi.ResourceOption) (*SimulationApplication, error) {
	var resource SimulationApplication
	err := ctx.ReadResource("aws-native:robomaker:SimulationApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SimulationApplication resources.
type simulationApplicationState struct {
}

type SimulationApplicationState struct {
}

func (SimulationApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*simulationApplicationState)(nil)).Elem()
}

type simulationApplicationArgs struct {
	// The current revision id.
	CurrentRevisionId *string `pulumi:"currentRevisionId"`
	// The URI of the Docker image for the robot application.
	Environment *string `pulumi:"environment"`
	// The name of the simulation application.
	Name *string `pulumi:"name"`
	// The rendering engine for the simulation application.
	RenderingEngine *SimulationApplicationRenderingEngine `pulumi:"renderingEngine"`
	// The robot software suite used by the simulation application.
	RobotSoftwareSuite SimulationApplicationRobotSoftwareSuite `pulumi:"robotSoftwareSuite"`
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite SimulationApplicationSimulationSoftwareSuite `pulumi:"simulationSoftwareSuite"`
	// The sources of the simulation application.
	Sources []SimulationApplicationSourceConfig `pulumi:"sources"`
	Tags    map[string]string                   `pulumi:"tags"`
}

// The set of arguments for constructing a SimulationApplication resource.
type SimulationApplicationArgs struct {
	// The current revision id.
	CurrentRevisionId pulumi.StringPtrInput
	// The URI of the Docker image for the robot application.
	Environment pulumi.StringPtrInput
	// The name of the simulation application.
	Name pulumi.StringPtrInput
	// The rendering engine for the simulation application.
	RenderingEngine SimulationApplicationRenderingEnginePtrInput
	// The robot software suite used by the simulation application.
	RobotSoftwareSuite SimulationApplicationRobotSoftwareSuiteInput
	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite SimulationApplicationSimulationSoftwareSuiteInput
	// The sources of the simulation application.
	Sources SimulationApplicationSourceConfigArrayInput
	Tags    pulumi.StringMapInput
}

func (SimulationApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simulationApplicationArgs)(nil)).Elem()
}

type SimulationApplicationInput interface {
	pulumi.Input

	ToSimulationApplicationOutput() SimulationApplicationOutput
	ToSimulationApplicationOutputWithContext(ctx context.Context) SimulationApplicationOutput
}

func (*SimulationApplication) ElementType() reflect.Type {
	return reflect.TypeOf((**SimulationApplication)(nil)).Elem()
}

func (i *SimulationApplication) ToSimulationApplicationOutput() SimulationApplicationOutput {
	return i.ToSimulationApplicationOutputWithContext(context.Background())
}

func (i *SimulationApplication) ToSimulationApplicationOutputWithContext(ctx context.Context) SimulationApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimulationApplicationOutput)
}

type SimulationApplicationOutput struct{ *pulumi.OutputState }

func (SimulationApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimulationApplication)(nil)).Elem()
}

func (o SimulationApplicationOutput) ToSimulationApplicationOutput() SimulationApplicationOutput {
	return o
}

func (o SimulationApplicationOutput) ToSimulationApplicationOutputWithContext(ctx context.Context) SimulationApplicationOutput {
	return o
}

func (o SimulationApplicationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *SimulationApplication) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The current revision id.
func (o SimulationApplicationOutput) CurrentRevisionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimulationApplication) pulumi.StringPtrOutput { return v.CurrentRevisionId }).(pulumi.StringPtrOutput)
}

// The URI of the Docker image for the robot application.
func (o SimulationApplicationOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimulationApplication) pulumi.StringPtrOutput { return v.Environment }).(pulumi.StringPtrOutput)
}

// The name of the simulation application.
func (o SimulationApplicationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SimulationApplication) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The rendering engine for the simulation application.
func (o SimulationApplicationOutput) RenderingEngine() SimulationApplicationRenderingEnginePtrOutput {
	return o.ApplyT(func(v *SimulationApplication) SimulationApplicationRenderingEnginePtrOutput { return v.RenderingEngine }).(SimulationApplicationRenderingEnginePtrOutput)
}

// The robot software suite used by the simulation application.
func (o SimulationApplicationOutput) RobotSoftwareSuite() SimulationApplicationRobotSoftwareSuiteOutput {
	return o.ApplyT(func(v *SimulationApplication) SimulationApplicationRobotSoftwareSuiteOutput {
		return v.RobotSoftwareSuite
	}).(SimulationApplicationRobotSoftwareSuiteOutput)
}

// The simulation software suite used by the simulation application.
func (o SimulationApplicationOutput) SimulationSoftwareSuite() SimulationApplicationSimulationSoftwareSuiteOutput {
	return o.ApplyT(func(v *SimulationApplication) SimulationApplicationSimulationSoftwareSuiteOutput {
		return v.SimulationSoftwareSuite
	}).(SimulationApplicationSimulationSoftwareSuiteOutput)
}

// The sources of the simulation application.
func (o SimulationApplicationOutput) Sources() SimulationApplicationSourceConfigArrayOutput {
	return o.ApplyT(func(v *SimulationApplication) SimulationApplicationSourceConfigArrayOutput { return v.Sources }).(SimulationApplicationSourceConfigArrayOutput)
}

func (o SimulationApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimulationApplication) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SimulationApplicationInput)(nil)).Elem(), &SimulationApplication{})
	pulumi.RegisterOutputType(SimulationApplicationOutput{})
}
