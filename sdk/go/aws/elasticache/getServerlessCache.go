// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::ElastiCache::ServerlessCache resource creates an Amazon ElastiCache Serverless Cache.
func LookupServerlessCache(ctx *pulumi.Context, args *LookupServerlessCacheArgs, opts ...pulumi.InvokeOption) (*LookupServerlessCacheResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerlessCacheResult
	err := ctx.Invoke("aws-native:elasticache:getServerlessCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerlessCacheArgs struct {
	// The name of the Serverless Cache. This value must be unique.
	ServerlessCacheName string `pulumi:"serverlessCacheName"`
}

type LookupServerlessCacheResult struct {
	// The ARN of the Serverless Cache.
	Arn              *string                          `pulumi:"arn"`
	CacheUsageLimits *ServerlessCacheCacheUsageLimits `pulumi:"cacheUsageLimits"`
	// The creation time of the Serverless Cache.
	CreateTime *string `pulumi:"createTime"`
	// The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.
	DailySnapshotTime *string `pulumi:"dailySnapshotTime"`
	// The description of the Serverless Cache.
	Description *string                  `pulumi:"description"`
	Endpoint    *ServerlessCacheEndpoint `pulumi:"endpoint"`
	// The full engine version of the Serverless Cache.
	FullEngineVersion *string                  `pulumi:"fullEngineVersion"`
	ReaderEndpoint    *ServerlessCacheEndpoint `pulumi:"readerEndpoint"`
	// One or more Amazon VPC security groups associated with this Serverless Cache.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The snapshot retention limit of the Serverless Cache.
	SnapshotRetentionLimit *int `pulumi:"snapshotRetentionLimit"`
	// The status of the Serverless Cache.
	Status *string `pulumi:"status"`
	// An array of key-value pairs to apply to this Serverless Cache.
	Tags []ServerlessCacheTag `pulumi:"tags"`
	// The ID of the user group.
	UserGroupId *string `pulumi:"userGroupId"`
}

func LookupServerlessCacheOutput(ctx *pulumi.Context, args LookupServerlessCacheOutputArgs, opts ...pulumi.InvokeOption) LookupServerlessCacheResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerlessCacheResult, error) {
			args := v.(LookupServerlessCacheArgs)
			r, err := LookupServerlessCache(ctx, &args, opts...)
			var s LookupServerlessCacheResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerlessCacheResultOutput)
}

type LookupServerlessCacheOutputArgs struct {
	// The name of the Serverless Cache. This value must be unique.
	ServerlessCacheName pulumi.StringInput `pulumi:"serverlessCacheName"`
}

func (LookupServerlessCacheOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerlessCacheArgs)(nil)).Elem()
}

type LookupServerlessCacheResultOutput struct{ *pulumi.OutputState }

func (LookupServerlessCacheResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerlessCacheResult)(nil)).Elem()
}

func (o LookupServerlessCacheResultOutput) ToLookupServerlessCacheResultOutput() LookupServerlessCacheResultOutput {
	return o
}

func (o LookupServerlessCacheResultOutput) ToLookupServerlessCacheResultOutputWithContext(ctx context.Context) LookupServerlessCacheResultOutput {
	return o
}

// The ARN of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupServerlessCacheResultOutput) CacheUsageLimits() ServerlessCacheCacheUsageLimitsPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *ServerlessCacheCacheUsageLimits { return v.CacheUsageLimits }).(ServerlessCacheCacheUsageLimitsPtrOutput)
}

// The creation time of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

// The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) DailySnapshotTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.DailySnapshotTime }).(pulumi.StringPtrOutput)
}

// The description of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupServerlessCacheResultOutput) Endpoint() ServerlessCacheEndpointPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *ServerlessCacheEndpoint { return v.Endpoint }).(ServerlessCacheEndpointPtrOutput)
}

// The full engine version of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) FullEngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.FullEngineVersion }).(pulumi.StringPtrOutput)
}

func (o LookupServerlessCacheResultOutput) ReaderEndpoint() ServerlessCacheEndpointPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *ServerlessCacheEndpoint { return v.ReaderEndpoint }).(ServerlessCacheEndpointPtrOutput)
}

// One or more Amazon VPC security groups associated with this Serverless Cache.
func (o LookupServerlessCacheResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The snapshot retention limit of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) SnapshotRetentionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *int { return v.SnapshotRetentionLimit }).(pulumi.IntPtrOutput)
}

// The status of the Serverless Cache.
func (o LookupServerlessCacheResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this Serverless Cache.
func (o LookupServerlessCacheResultOutput) Tags() ServerlessCacheTagArrayOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) []ServerlessCacheTag { return v.Tags }).(ServerlessCacheTagArrayOutput)
}

// The ID of the user group.
func (o LookupServerlessCacheResultOutput) UserGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerlessCacheResult) *string { return v.UserGroupId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerlessCacheResultOutput{})
}
