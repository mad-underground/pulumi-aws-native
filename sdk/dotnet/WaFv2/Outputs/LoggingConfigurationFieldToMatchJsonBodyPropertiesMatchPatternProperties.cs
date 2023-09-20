// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// The patterns to look for in the JSON body. AWS WAF inspects the results of these pattern matches against the rule inspection criteria. 
    /// </summary>
    [OutputType]
    public sealed class LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchPatternProperties
    {
        /// <summary>
        /// Match all of the elements. See also MatchScope in JsonBody. You must specify either this setting or the IncludedPaths setting, but not both.
        /// </summary>
        public readonly object? All;
        /// <summary>
        /// Match only the specified include paths. See also MatchScope in JsonBody.
        /// </summary>
        public readonly ImmutableArray<string> IncludedPaths;

        [OutputConstructor]
        private LoggingConfigurationFieldToMatchJsonBodyPropertiesMatchPatternProperties(
            object? all,

            ImmutableArray<string> includedPaths)
        {
            All = all;
            IncludedPaths = includedPaths;
        }
    }
}