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
    public sealed class BranchAttributes
    {
        public readonly string Arn;
        public readonly string BranchName;

        [OutputConstructor]
        private BranchAttributes(
            string Arn,

            string BranchName)
        {
            this.Arn = Arn;
            this.BranchName = BranchName;
        }
    }
}