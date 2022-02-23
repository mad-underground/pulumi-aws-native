// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Signer
{
    public static class GetProfilePermission
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetProfilePermissionResult> InvokeAsync(GetProfilePermissionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProfilePermissionResult>("aws-native:signer:getProfilePermission", args ?? new GetProfilePermissionArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetProfilePermissionResult> Invoke(GetProfilePermissionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProfilePermissionResult>("aws-native:signer:getProfilePermission", args ?? new GetProfilePermissionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProfilePermissionArgs : Pulumi.InvokeArgs
    {
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        [Input("statementId", required: true)]
        public string StatementId { get; set; } = null!;

        public GetProfilePermissionArgs()
        {
        }
    }

    public sealed class GetProfilePermissionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        [Input("statementId", required: true)]
        public Input<string> StatementId { get; set; } = null!;

        public GetProfilePermissionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProfilePermissionResult
    {
        [OutputConstructor]
        private GetProfilePermissionResult()
        {
        }
    }
}