// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig.Outputs
{

    /// <summary>
    /// A parameter for the extension to send to a specific action.
    /// </summary>
    [OutputType]
    public sealed class ExtensionParameter
    {
        /// <summary>
        /// The description of the extension Parameter.
        /// </summary>
        public readonly string? Description;
        public readonly bool Required;

        [OutputConstructor]
        private ExtensionParameter(
            string? description,

            bool required)
        {
            Description = description;
            Required = required;
        }
    }
}
