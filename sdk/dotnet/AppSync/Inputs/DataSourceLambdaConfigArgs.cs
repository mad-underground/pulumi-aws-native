// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Inputs
{

    public sealed class DataSourceLambdaConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("lambdaFunctionArn", required: true)]
        public Input<string> LambdaFunctionArn { get; set; } = null!;

        public DataSourceLambdaConfigArgs()
        {
        }
        public static new DataSourceLambdaConfigArgs Empty => new DataSourceLambdaConfigArgs();
    }
}
