// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// Sets the runtime management configuration for a function's version. For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html).
    /// </summary>
    public sealed class FunctionRuntimeManagementConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the runtime version you want the function to use.
        ///   This is only required if you're using the *Manual* runtime update mode.
        /// </summary>
        [Input("runtimeVersionArn")]
        public Input<string>? RuntimeVersionArn { get; set; }

        /// <summary>
        /// Specify the runtime update mode.
        ///   + *Auto (default)* - Automatically update to the most recent and secure runtime version using a [Two-phase runtime version rollout](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-two-phase). This is the best choice for most customers to ensure they always benefit from runtime updates.
        ///  + *FunctionUpdate* - LAM updates the runtime of you function to the most recent and secure runtime version when you update your function. This approach synchronizes runtime updates with function deployments, giving you control over when runtime updates are applied and allowing you to detect and mitigate rare runtime update incompatibilities early. When using this setting, you need to regularly update your functions to keep their runtime up-to-date.
        ///  + *Manual* - You specify a runtime version in your function configuration. The function will use this runtime version indefinitely. In the rare case where a new runtime version is incompatible with an existing function, this allows you to roll back your function to an earlier runtime version. For more information, see [Roll back a runtime version](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-rollback).
        ///  
        ///  *Valid Values*: ``Auto`` | ``FunctionUpdate`` | ``Manual``
        /// </summary>
        [Input("updateRuntimeOn", required: true)]
        public Input<Pulumi.AwsNative.Lambda.FunctionRuntimeManagementConfigUpdateRuntimeOn> UpdateRuntimeOn { get; set; } = null!;

        public FunctionRuntimeManagementConfigArgs()
        {
        }
        public static new FunctionRuntimeManagementConfigArgs Empty => new FunctionRuntimeManagementConfigArgs();
    }
}
