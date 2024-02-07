// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Amplify::Branch resource creates a new branch within an app.
type Branch struct {
	pulumi.CustomResourceState

	AppId                      pulumi.StringOutput                  `pulumi:"appId"`
	Arn                        pulumi.StringOutput                  `pulumi:"arn"`
	Backend                    BranchBackendPtrOutput               `pulumi:"backend"`
	BasicAuthConfig            BranchBasicAuthConfigPtrOutput       `pulumi:"basicAuthConfig"`
	BranchName                 pulumi.StringOutput                  `pulumi:"branchName"`
	BuildSpec                  pulumi.StringPtrOutput               `pulumi:"buildSpec"`
	Description                pulumi.StringPtrOutput               `pulumi:"description"`
	EnableAutoBuild            pulumi.BoolPtrOutput                 `pulumi:"enableAutoBuild"`
	EnablePerformanceMode      pulumi.BoolPtrOutput                 `pulumi:"enablePerformanceMode"`
	EnablePullRequestPreview   pulumi.BoolPtrOutput                 `pulumi:"enablePullRequestPreview"`
	EnvironmentVariables       BranchEnvironmentVariableArrayOutput `pulumi:"environmentVariables"`
	Framework                  pulumi.StringPtrOutput               `pulumi:"framework"`
	PullRequestEnvironmentName pulumi.StringPtrOutput               `pulumi:"pullRequestEnvironmentName"`
	Stage                      BranchStagePtrOutput                 `pulumi:"stage"`
	Tags                       BranchTagArrayOutput                 `pulumi:"tags"`
}

// NewBranch registers a new resource with the given unique name, arguments, and options.
func NewBranch(ctx *pulumi.Context,
	name string, args *BranchArgs, opts ...pulumi.ResourceOption) (*Branch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"appId",
		"branchName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Branch
	err := ctx.RegisterResource("aws-native:amplify:Branch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBranch gets an existing Branch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBranch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BranchState, opts ...pulumi.ResourceOption) (*Branch, error) {
	var resource Branch
	err := ctx.ReadResource("aws-native:amplify:Branch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Branch resources.
type branchState struct {
}

type BranchState struct {
}

func (BranchState) ElementType() reflect.Type {
	return reflect.TypeOf((*branchState)(nil)).Elem()
}

type branchArgs struct {
	AppId                      string                      `pulumi:"appId"`
	Backend                    *BranchBackend              `pulumi:"backend"`
	BasicAuthConfig            *BranchBasicAuthConfig      `pulumi:"basicAuthConfig"`
	BranchName                 *string                     `pulumi:"branchName"`
	BuildSpec                  *string                     `pulumi:"buildSpec"`
	Description                *string                     `pulumi:"description"`
	EnableAutoBuild            *bool                       `pulumi:"enableAutoBuild"`
	EnablePerformanceMode      *bool                       `pulumi:"enablePerformanceMode"`
	EnablePullRequestPreview   *bool                       `pulumi:"enablePullRequestPreview"`
	EnvironmentVariables       []BranchEnvironmentVariable `pulumi:"environmentVariables"`
	Framework                  *string                     `pulumi:"framework"`
	PullRequestEnvironmentName *string                     `pulumi:"pullRequestEnvironmentName"`
	Stage                      *BranchStage                `pulumi:"stage"`
	Tags                       []BranchTag                 `pulumi:"tags"`
}

// The set of arguments for constructing a Branch resource.
type BranchArgs struct {
	AppId                      pulumi.StringInput
	Backend                    BranchBackendPtrInput
	BasicAuthConfig            BranchBasicAuthConfigPtrInput
	BranchName                 pulumi.StringPtrInput
	BuildSpec                  pulumi.StringPtrInput
	Description                pulumi.StringPtrInput
	EnableAutoBuild            pulumi.BoolPtrInput
	EnablePerformanceMode      pulumi.BoolPtrInput
	EnablePullRequestPreview   pulumi.BoolPtrInput
	EnvironmentVariables       BranchEnvironmentVariableArrayInput
	Framework                  pulumi.StringPtrInput
	PullRequestEnvironmentName pulumi.StringPtrInput
	Stage                      BranchStagePtrInput
	Tags                       BranchTagArrayInput
}

func (BranchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*branchArgs)(nil)).Elem()
}

type BranchInput interface {
	pulumi.Input

	ToBranchOutput() BranchOutput
	ToBranchOutputWithContext(ctx context.Context) BranchOutput
}

func (*Branch) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil)).Elem()
}

func (i *Branch) ToBranchOutput() BranchOutput {
	return i.ToBranchOutputWithContext(context.Background())
}

func (i *Branch) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BranchOutput)
}

type BranchOutput struct{ *pulumi.OutputState }

func (BranchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Branch)(nil)).Elem()
}

func (o BranchOutput) ToBranchOutput() BranchOutput {
	return o
}

func (o BranchOutput) ToBranchOutputWithContext(ctx context.Context) BranchOutput {
	return o
}

func (o BranchOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o BranchOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o BranchOutput) Backend() BranchBackendPtrOutput {
	return o.ApplyT(func(v *Branch) BranchBackendPtrOutput { return v.Backend }).(BranchBackendPtrOutput)
}

func (o BranchOutput) BasicAuthConfig() BranchBasicAuthConfigPtrOutput {
	return o.ApplyT(func(v *Branch) BranchBasicAuthConfigPtrOutput { return v.BasicAuthConfig }).(BranchBasicAuthConfigPtrOutput)
}

func (o BranchOutput) BranchName() pulumi.StringOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringOutput { return v.BranchName }).(pulumi.StringOutput)
}

func (o BranchOutput) BuildSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringPtrOutput { return v.BuildSpec }).(pulumi.StringPtrOutput)
}

func (o BranchOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BranchOutput) EnableAutoBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolPtrOutput { return v.EnableAutoBuild }).(pulumi.BoolPtrOutput)
}

func (o BranchOutput) EnablePerformanceMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolPtrOutput { return v.EnablePerformanceMode }).(pulumi.BoolPtrOutput)
}

func (o BranchOutput) EnablePullRequestPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.BoolPtrOutput { return v.EnablePullRequestPreview }).(pulumi.BoolPtrOutput)
}

func (o BranchOutput) EnvironmentVariables() BranchEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v *Branch) BranchEnvironmentVariableArrayOutput { return v.EnvironmentVariables }).(BranchEnvironmentVariableArrayOutput)
}

func (o BranchOutput) Framework() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringPtrOutput { return v.Framework }).(pulumi.StringPtrOutput)
}

func (o BranchOutput) PullRequestEnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Branch) pulumi.StringPtrOutput { return v.PullRequestEnvironmentName }).(pulumi.StringPtrOutput)
}

func (o BranchOutput) Stage() BranchStagePtrOutput {
	return o.ApplyT(func(v *Branch) BranchStagePtrOutput { return v.Stage }).(BranchStagePtrOutput)
}

func (o BranchOutput) Tags() BranchTagArrayOutput {
	return o.ApplyT(func(v *Branch) BranchTagArrayOutput { return v.Tags }).(BranchTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BranchInput)(nil)).Elem(), &Branch{})
	pulumi.RegisterOutputType(BranchOutput{})
}
