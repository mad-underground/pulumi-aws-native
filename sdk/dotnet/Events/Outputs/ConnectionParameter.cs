// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Outputs
{

    [OutputType]
    public sealed class ConnectionParameter
    {
        public readonly bool? IsValueSecret;
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private ConnectionParameter(
            bool? isValueSecret,

            string key,

            string value)
        {
            IsValueSecret = isValueSecret;
            Key = key;
            Value = value;
        }
    }
}