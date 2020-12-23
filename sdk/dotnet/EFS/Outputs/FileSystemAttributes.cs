// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EFS.Outputs
{

    [OutputType]
    public sealed class FileSystemAttributes
    {
        public readonly string Arn;
        public readonly string FileSystemId;

        [OutputConstructor]
        private FileSystemAttributes(
            string Arn,

            string FileSystemId)
        {
            this.Arn = Arn;
            this.FileSystemId = FileSystemId;
        }
    }
}
