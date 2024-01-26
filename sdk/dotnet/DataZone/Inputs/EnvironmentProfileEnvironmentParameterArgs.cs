// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// The parameter details of an environment profile.
    /// </summary>
    public sealed class EnvironmentProfileEnvironmentParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of an environment profile parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The value of an environment profile parameter.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public EnvironmentProfileEnvironmentParameterArgs()
        {
        }
        public static new EnvironmentProfileEnvironmentParameterArgs Empty => new EnvironmentProfileEnvironmentParameterArgs();
    }
}
