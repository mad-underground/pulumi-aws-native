// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkSpacesWeb.Outputs
{

    [OutputType]
    public sealed class UserSettingsCookieSpecification
    {
        public readonly string Domain;
        public readonly string? Name;
        public readonly string? Path;

        [OutputConstructor]
        private UserSettingsCookieSpecification(
            string domain,

            string? name,

            string? path)
        {
            Domain = domain;
            Name = name;
            Path = path;
        }
    }
}
