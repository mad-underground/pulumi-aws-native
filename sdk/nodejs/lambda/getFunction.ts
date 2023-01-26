// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Function
 */
export function getFunction(args: GetFunctionArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:lambda:getFunction", {
        "functionName": args.functionName,
    }, opts);
}

export interface GetFunctionArgs {
    /**
     * The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.
     */
    functionName: string;
}

export interface GetFunctionResult {
    readonly architectures?: enums.lambda.FunctionArchitecturesItem[];
    /**
     * Unique identifier for function resources
     */
    readonly arn?: string;
    /**
     * A unique Arn for CodeSigningConfig resource
     */
    readonly codeSigningConfigArn?: string;
    /**
     * A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.
     */
    readonly deadLetterConfig?: outputs.lambda.FunctionDeadLetterConfig;
    /**
     * A description of the function.
     */
    readonly description?: string;
    /**
     * Environment variables that are accessible from function code during execution.
     */
    readonly environment?: outputs.lambda.FunctionEnvironment;
    /**
     * A function's ephemeral storage settings.
     */
    readonly ephemeralStorage?: outputs.lambda.FunctionEphemeralStorage;
    /**
     * Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.
     */
    readonly fileSystemConfigs?: outputs.lambda.FunctionFileSystemConfig[];
    /**
     * The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime
     */
    readonly handler?: string;
    /**
     * ImageConfig
     */
    readonly imageConfig?: outputs.lambda.FunctionImageConfig;
    /**
     * The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
     */
    readonly kmsKeyArn?: string;
    /**
     * A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.
     */
    readonly layers?: string[];
    /**
     * The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.
     */
    readonly memorySize?: number;
    /**
     * PackageType.
     */
    readonly packageType?: enums.lambda.FunctionPackageType;
    /**
     * The number of simultaneous executions to reserve for the function.
     */
    readonly reservedConcurrentExecutions?: number;
    /**
     * The Amazon Resource Name (ARN) of the function's execution role.
     */
    readonly role?: string;
    /**
     * The identifier of the function's runtime.
     */
    readonly runtime?: string;
    /**
     * RuntimeManagementConfig
     */
    readonly runtimeManagementConfig?: outputs.lambda.FunctionRuntimeManagementConfig;
    /**
     * The SnapStart setting of your function
     */
    readonly snapStart?: outputs.lambda.FunctionSnapStart;
    /**
     * The SnapStart response of your function
     */
    readonly snapStartResponse?: outputs.lambda.FunctionSnapStartResponse;
    /**
     * A list of tags to apply to the function.
     */
    readonly tags?: outputs.lambda.FunctionTag[];
    /**
     * The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.
     */
    readonly timeout?: number;
    /**
     * Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.
     */
    readonly tracingConfig?: outputs.lambda.FunctionTracingConfig;
    /**
     * For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC.
     */
    readonly vpcConfig?: outputs.lambda.FunctionVpcConfig;
}
/**
 * Resource Type definition for AWS::Lambda::Function
 */
export function getFunctionOutput(args: GetFunctionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionResult> {
    return pulumi.output(args).apply((a: any) => getFunction(a, opts))
}

export interface GetFunctionOutputArgs {
    /**
     * The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.
     */
    functionName: pulumi.Input<string>;
}
