// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kms
{
    public static class GetAlias
    {
        /// <summary>
        /// The AWS::KMS::Alias resource specifies a display name for an AWS KMS key in AWS Key Management Service (AWS KMS). You can use an alias to identify an AWS KMS key in cryptographic operations.
        /// </summary>
        public static Task<GetAliasResult> InvokeAsync(GetAliasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAliasResult>("aws-native:kms:getAlias", args ?? new GetAliasArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::KMS::Alias resource specifies a display name for an AWS KMS key in AWS Key Management Service (AWS KMS). You can use an alias to identify an AWS KMS key in cryptographic operations.
        /// </summary>
        public static Output<GetAliasResult> Invoke(GetAliasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAliasResult>("aws-native:kms:getAlias", args ?? new GetAliasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAliasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.
        /// </summary>
        [Input("aliasName", required: true)]
        public string AliasName { get; set; } = null!;

        public GetAliasArgs()
        {
        }
        public static new GetAliasArgs Empty => new GetAliasArgs();
    }

    public sealed class GetAliasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.
        /// </summary>
        [Input("aliasName", required: true)]
        public Input<string> AliasName { get; set; } = null!;

        public GetAliasInvokeArgs()
        {
        }
        public static new GetAliasInvokeArgs Empty => new GetAliasInvokeArgs();
    }


    [OutputType]
    public sealed class GetAliasResult
    {
        /// <summary>
        /// Identifies the AWS KMS key to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.
        /// </summary>
        public readonly string? TargetKeyId;

        [OutputConstructor]
        private GetAliasResult(string? targetKeyId)
        {
            TargetKeyId = targetKeyId;
        }
    }
}