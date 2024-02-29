// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda
{
    public static class GetFunction
    {
        /// <summary>
        /// The ``AWS::Lambda::Function`` resource creates a Lambda function. To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html). The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
        ///  You set the package type to ``Image`` if the deployment package is a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html). For a container image, the code property must include the URI of a container image in the Amazon ECR registry. You do not need to specify the handler and runtime properties. 
        ///  You set the package type to ``Zip`` if the deployment package is a [.zip file archive](https://docs.aws.amazon.com/lam
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("aws-native:lambda:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::Lambda::Function`` resource creates a Lambda function. To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html). The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
        ///  You set the package type to ``Image`` if the deployment package is a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html). For a container image, the code property must include the URI of a container image in the Amazon ECR registry. You do not need to specify the handler and runtime properties. 
        ///  You set the package type to ``Zip`` if the deployment package is a [.zip file archive](https://docs.aws.amazon.com/lam
        /// </summary>
        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("aws-native:lambda:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Lambda function, up to 64 characters in length. If you don't specify a name, CFN generates one.
        ///  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        public GetFunctionArgs()
        {
        }
        public static new GetFunctionArgs Empty => new GetFunctionArgs();
    }

    public sealed class GetFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Lambda function, up to 64 characters in length. If you don't specify a name, CFN generates one.
        ///  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        public GetFunctionInvokeArgs()
        {
        }
        public static new GetFunctionInvokeArgs Empty => new GetFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is ``x86_64``.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Lambda.FunctionArchitecturesItem> Architectures;
        public readonly string? Arn;
        /// <summary>
        /// To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
        /// </summary>
        public readonly string? CodeSigningConfigArn;
        /// <summary>
        /// A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see [Dead-letter queues](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).
        /// </summary>
        public readonly Outputs.FunctionDeadLetterConfig? DeadLetterConfig;
        /// <summary>
        /// A description of the function.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Environment variables that are accessible from function code during execution.
        /// </summary>
        public readonly Outputs.FunctionEnvironment? Environment;
        /// <summary>
        /// The size of the function's ``/tmp`` directory in MB. The default value is 512, but it can be any whole number between 512 and 10,240 MB.
        /// </summary>
        public readonly Outputs.FunctionEphemeralStorage? EphemeralStorage;
        /// <summary>
        /// Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) resource, you must also specify a ``DependsOn`` attribute to ensure that the mount target is created or updated before the function.
        ///  For more information about using the ``DependsOn`` attribute, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html).
        /// </summary>
        public readonly ImmutableArray<Outputs.FunctionFileSystemConfig> FileSystemConfigs;
        /// <summary>
        /// The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Lambda programming model](https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).
        /// </summary>
        public readonly string? Handler;
        /// <summary>
        /// Configuration values that override the container image Dockerfile settings. For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms).
        /// </summary>
        public readonly Outputs.FunctionImageConfig? ImageConfig;
        /// <summary>
        /// The ARN of the KMSlong (KMS) customer managed key that's used to encrypt your function's [environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption). When [Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html) is activated, Lambda also uses this key is to encrypt your function's snapshot. If you deploy your function using a container image, Lambda also uses this key to encrypt your function when it's deployed. Note that this is not the same key that's used to protect your container image in the Amazon Elastic Container Registry (Amazon ECR). If you don't provide a customer managed key, Lambda uses a default service key.
        /// </summary>
        public readonly string? KmsKeyArn;
        /// <summary>
        /// A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.
        /// </summary>
        public readonly ImmutableArray<string> Layers;
        /// <summary>
        /// The function's Amazon CloudWatch Logs configuration settings.
        /// </summary>
        public readonly Outputs.FunctionLoggingConfig? LoggingConfig;
        /// <summary>
        /// The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB. Note that new AWS accounts have reduced concurrency and memory quotas. AWS raises these quotas automatically based on your usage. You can also request a quota increase.
        /// </summary>
        public readonly int? MemorySize;
        /// <summary>
        /// The type of deployment package. Set to ``Image`` for container image and set ``Zip`` for .zip file archive.
        /// </summary>
        public readonly Pulumi.AwsNative.Lambda.FunctionPackageType? PackageType;
        /// <summary>
        /// The number of simultaneous executions to reserve for the function.
        /// </summary>
        public readonly int? ReservedConcurrentExecutions;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the function's execution role.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.
        ///  The following list includes deprecated runtimes. For more information, see [Runtime deprecation policy](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy).
        /// </summary>
        public readonly string? Runtime;
        /// <summary>
        /// Sets the runtime management configuration for a function's version. For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html).
        /// </summary>
        public readonly Outputs.FunctionRuntimeManagementConfig? RuntimeManagementConfig;
        public readonly Outputs.FunctionSnapStartResponse? SnapStartResponse;
        /// <summary>
        /// A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).
        /// </summary>
        public readonly int? Timeout;
        /// <summary>
        /// Set ``Mode`` to ``Active`` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).
        /// </summary>
        public readonly Outputs.FunctionTracingConfig? TracingConfig;
        /// <summary>
        /// For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see [Configuring a Lambda function to access resources in a VPC](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
        /// </summary>
        public readonly Outputs.FunctionVpcConfig? VpcConfig;

        [OutputConstructor]
        private GetFunctionResult(
            ImmutableArray<Pulumi.AwsNative.Lambda.FunctionArchitecturesItem> architectures,

            string? arn,

            string? codeSigningConfigArn,

            Outputs.FunctionDeadLetterConfig? deadLetterConfig,

            string? description,

            Outputs.FunctionEnvironment? environment,

            Outputs.FunctionEphemeralStorage? ephemeralStorage,

            ImmutableArray<Outputs.FunctionFileSystemConfig> fileSystemConfigs,

            string? handler,

            Outputs.FunctionImageConfig? imageConfig,

            string? kmsKeyArn,

            ImmutableArray<string> layers,

            Outputs.FunctionLoggingConfig? loggingConfig,

            int? memorySize,

            Pulumi.AwsNative.Lambda.FunctionPackageType? packageType,

            int? reservedConcurrentExecutions,

            string? role,

            string? runtime,

            Outputs.FunctionRuntimeManagementConfig? runtimeManagementConfig,

            Outputs.FunctionSnapStartResponse? snapStartResponse,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            int? timeout,

            Outputs.FunctionTracingConfig? tracingConfig,

            Outputs.FunctionVpcConfig? vpcConfig)
        {
            Architectures = architectures;
            Arn = arn;
            CodeSigningConfigArn = codeSigningConfigArn;
            DeadLetterConfig = deadLetterConfig;
            Description = description;
            Environment = environment;
            EphemeralStorage = ephemeralStorage;
            FileSystemConfigs = fileSystemConfigs;
            Handler = handler;
            ImageConfig = imageConfig;
            KmsKeyArn = kmsKeyArn;
            Layers = layers;
            LoggingConfig = loggingConfig;
            MemorySize = memorySize;
            PackageType = packageType;
            ReservedConcurrentExecutions = reservedConcurrentExecutions;
            Role = role;
            Runtime = runtime;
            RuntimeManagementConfig = runtimeManagementConfig;
            SnapStartResponse = snapStartResponse;
            Tags = tags;
            Timeout = timeout;
            TracingConfig = tracingConfig;
            VpcConfig = vpcConfig;
        }
    }
}
