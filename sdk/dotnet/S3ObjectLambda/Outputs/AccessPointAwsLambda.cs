// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda.Outputs
{

    [OutputType]
    public sealed class AccessPointAwsLambda
    {
        public readonly string FunctionArn;
        public readonly string? FunctionPayload;

        [OutputConstructor]
        private AccessPointAwsLambda(
            string functionArn,

            string? functionPayload)
        {
            FunctionArn = functionArn;
            FunctionPayload = functionPayload;
        }
    }
}
