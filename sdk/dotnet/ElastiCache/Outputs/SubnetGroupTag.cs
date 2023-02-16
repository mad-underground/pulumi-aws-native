// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElastiCache.Outputs
{

    /// <summary>
    /// A tag that can be added to an ElastiCache subnet group. Tags are composed of a Key/Value pair. You can use tags to categorize and track all your subnet groups. A tag with a null Value is permitted.
    /// </summary>
    [OutputType]
    public sealed class SubnetGroupTag
    {
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private SubnetGroupTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
