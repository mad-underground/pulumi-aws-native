// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppSync.Outputs
{

    [OutputType]
    public sealed class ApiKeyAttributes
    {
        public readonly string ApiKey;
        public readonly string Arn;

        [OutputConstructor]
        private ApiKeyAttributes(
            string ApiKey,

            string Arn)
        {
            this.ApiKey = ApiKey;
            this.Arn = Arn;
        }
    }
}