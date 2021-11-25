// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// A label for tagging Lex resources
    /// </summary>
    [OutputType]
    public sealed class BotAliasTag
    {
        /// <summary>
        /// A string used to identify this tag
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// A string containing the value for the tag
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private BotAliasTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}