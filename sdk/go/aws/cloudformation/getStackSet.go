// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances
func LookupStackSet(ctx *pulumi.Context, args *LookupStackSetArgs, opts ...pulumi.InvokeOption) (*LookupStackSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStackSetResult
	err := ctx.Invoke("aws-native:cloudformation:getStackSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStackSetArgs struct {
	// The ID of the stack set that you're creating.
	StackSetId string `pulumi:"stackSetId"`
}

type LookupStackSetResult struct {
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
	AdministrationRoleArn *string `pulumi:"administrationRoleArn"`
	// Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.
	AutoDeployment *StackSetAutoDeployment `pulumi:"autoDeployment"`
	// In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.
	Capabilities []StackSetCapability `pulumi:"capabilities"`
	// A description of the stack set. You can use the description to identify the stack set's purpose or other important information.
	Description *string `pulumi:"description"`
	// The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.
	ExecutionRoleName *string `pulumi:"executionRoleName"`
	// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	ManagedExecution *ManagedExecutionProperties `pulumi:"managedExecution"`
	// The input parameters for the stack set template.
	Parameters []StackSetParameter `pulumi:"parameters"`
	// The ID of the stack set that you're creating.
	StackSetId *string `pulumi:"stackSetId"`
	// The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
	Tags []StackSetTag `pulumi:"tags"`
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	TemplateBody *string `pulumi:"templateBody"`
}

func LookupStackSetOutput(ctx *pulumi.Context, args LookupStackSetOutputArgs, opts ...pulumi.InvokeOption) LookupStackSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackSetResult, error) {
			args := v.(LookupStackSetArgs)
			r, err := LookupStackSet(ctx, &args, opts...)
			var s LookupStackSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackSetResultOutput)
}

type LookupStackSetOutputArgs struct {
	// The ID of the stack set that you're creating.
	StackSetId pulumi.StringInput `pulumi:"stackSetId"`
}

func (LookupStackSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackSetArgs)(nil)).Elem()
}

type LookupStackSetResultOutput struct{ *pulumi.OutputState }

func (LookupStackSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackSetResult)(nil)).Elem()
}

func (o LookupStackSetResultOutput) ToLookupStackSetResultOutput() LookupStackSetResultOutput {
	return o
}

func (o LookupStackSetResultOutput) ToLookupStackSetResultOutputWithContext(ctx context.Context) LookupStackSetResultOutput {
	return o
}

// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
func (o LookupStackSetResultOutput) AdministrationRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *string { return v.AdministrationRoleArn }).(pulumi.StringPtrOutput)
}

// Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.
func (o LookupStackSetResultOutput) AutoDeployment() StackSetAutoDeploymentPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *StackSetAutoDeployment { return v.AutoDeployment }).(StackSetAutoDeploymentPtrOutput)
}

// In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.
func (o LookupStackSetResultOutput) Capabilities() StackSetCapabilityArrayOutput {
	return o.ApplyT(func(v LookupStackSetResult) []StackSetCapability { return v.Capabilities }).(StackSetCapabilityArrayOutput)
}

// A description of the stack set. You can use the description to identify the stack set's purpose or other important information.
func (o LookupStackSetResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.
func (o LookupStackSetResultOutput) ExecutionRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *string { return v.ExecutionRoleName }).(pulumi.StringPtrOutput)
}

// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
func (o LookupStackSetResultOutput) ManagedExecution() ManagedExecutionPropertiesPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *ManagedExecutionProperties { return v.ManagedExecution }).(ManagedExecutionPropertiesPtrOutput)
}

// The input parameters for the stack set template.
func (o LookupStackSetResultOutput) Parameters() StackSetParameterArrayOutput {
	return o.ApplyT(func(v LookupStackSetResult) []StackSetParameter { return v.Parameters }).(StackSetParameterArrayOutput)
}

// The ID of the stack set that you're creating.
func (o LookupStackSetResultOutput) StackSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *string { return v.StackSetId }).(pulumi.StringPtrOutput)
}

// The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
func (o LookupStackSetResultOutput) Tags() StackSetTagArrayOutput {
	return o.ApplyT(func(v LookupStackSetResult) []StackSetTag { return v.Tags }).(StackSetTagArrayOutput)
}

// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
func (o LookupStackSetResultOutput) TemplateBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackSetResult) *string { return v.TemplateBody }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackSetResultOutput{})
}
