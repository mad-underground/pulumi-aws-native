// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Amplify.Outputs
{

    [OutputType]
    public sealed class AppAttributes
    {
        public readonly string AppId;
        public readonly string AppName;
        public readonly string Arn;
        public readonly string DefaultDomain;

        [OutputConstructor]
        private AppAttributes(
            string AppId,

            string AppName,

            string Arn,

            string DefaultDomain)
        {
            this.AppId = AppId;
            this.AppName = AppName;
            this.Arn = Arn;
            this.DefaultDomain = DefaultDomain;
        }
    }
}