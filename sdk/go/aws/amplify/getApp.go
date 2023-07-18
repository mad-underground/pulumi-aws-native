// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAppResult
	err := ctx.Invoke("aws-native:amplify:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupAppResult struct {
	AppId                    *string                  `pulumi:"appId"`
	AppName                  *string                  `pulumi:"appName"`
	Arn                      *string                  `pulumi:"arn"`
	BuildSpec                *string                  `pulumi:"buildSpec"`
	CustomHeaders            *string                  `pulumi:"customHeaders"`
	CustomRules              []AppCustomRule          `pulumi:"customRules"`
	DefaultDomain            *string                  `pulumi:"defaultDomain"`
	Description              *string                  `pulumi:"description"`
	EnableBranchAutoDeletion *bool                    `pulumi:"enableBranchAutoDeletion"`
	EnvironmentVariables     []AppEnvironmentVariable `pulumi:"environmentVariables"`
	IAMServiceRole           *string                  `pulumi:"iAMServiceRole"`
	Name                     *string                  `pulumi:"name"`
	Platform                 *AppPlatform             `pulumi:"platform"`
	Repository               *string                  `pulumi:"repository"`
	Tags                     []AppTag                 `pulumi:"tags"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) AppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.AppName }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) BuildSpec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.BuildSpec }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) CustomHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.CustomHeaders }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) CustomRules() AppCustomRuleArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []AppCustomRule { return v.CustomRules }).(AppCustomRuleArrayOutput)
}

func (o LookupAppResultOutput) DefaultDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.DefaultDomain }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) EnableBranchAutoDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *bool { return v.EnableBranchAutoDeletion }).(pulumi.BoolPtrOutput)
}

func (o LookupAppResultOutput) EnvironmentVariables() AppEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []AppEnvironmentVariable { return v.EnvironmentVariables }).(AppEnvironmentVariableArrayOutput)
}

func (o LookupAppResultOutput) IAMServiceRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.IAMServiceRole }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Platform() AppPlatformPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *AppPlatform { return v.Platform }).(AppPlatformPtrOutput)
}

func (o LookupAppResultOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

func (o LookupAppResultOutput) Tags() AppTagArrayOutput {
	return o.ApplyT(func(v LookupAppResult) []AppTag { return v.Tags }).(AppTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
