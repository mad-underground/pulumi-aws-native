// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    /// <summary>
    /// Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.
    /// </summary>
    [OutputType]
    public sealed class FunctionConfigurationAppSyncRuntime
    {
        /// <summary>
        /// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The version of the runtime to use. Currently, the only allowed version is 1.0.0.
        /// </summary>
        public readonly string RuntimeVersion;

        [OutputConstructor]
        private FunctionConfigurationAppSyncRuntime(
            string name,

            string runtimeVersion)
        {
            Name = name;
            RuntimeVersion = runtimeVersion;
        }
    }
}
