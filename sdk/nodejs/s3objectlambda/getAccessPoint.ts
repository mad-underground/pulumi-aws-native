// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
 */
export function getAccessPoint(args: GetAccessPointArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPointResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:s3objectlambda:getAccessPoint", {
        "name": args.name,
    }, opts);
}

export interface GetAccessPointArgs {
    /**
     * The name you want to assign to this Object lambda Access Point.
     */
    name: string;
}

export interface GetAccessPointResult {
    readonly arn?: string;
    /**
     * The date and time when the Object lambda Access Point was created.
     */
    readonly creationDate?: string;
    /**
     * The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
     */
    readonly objectLambdaConfiguration?: outputs.s3objectlambda.AccessPointObjectLambdaConfiguration;
    readonly policyStatus?: outputs.s3objectlambda.PolicyStatusProperties;
    /**
     * The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
     */
    readonly publicAccessBlockConfiguration?: outputs.s3objectlambda.AccessPointPublicAccessBlockConfiguration;
}
/**
 * The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
 */
export function getAccessPointOutput(args: GetAccessPointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessPointResult> {
    return pulumi.output(args).apply((a: any) => getAccessPoint(a, opts))
}

export interface GetAccessPointOutputArgs {
    /**
     * The name you want to assign to this Object lambda Access Point.
     */
    name: pulumi.Input<string>;
}
