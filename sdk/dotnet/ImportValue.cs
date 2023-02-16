// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class ImportValue
    {
        public static Task<ImportValueResult> InvokeAsync(ImportValueArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ImportValueResult>("aws-native:index:importValue", args ?? new ImportValueArgs(), options.WithDefaults());

        public static Output<ImportValueResult> Invoke(ImportValueInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ImportValueResult>("aws-native:index:importValue", args ?? new ImportValueInvokeArgs(), options.WithDefaults());
    }


    public sealed class ImportValueArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public ImportValueArgs()
        {
        }
        public static new ImportValueArgs Empty => new ImportValueArgs();
    }

    public sealed class ImportValueInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ImportValueInvokeArgs()
        {
        }
        public static new ImportValueInvokeArgs Empty => new ImportValueInvokeArgs();
    }


    [OutputType]
    public sealed class ImportValueResult
    {
        public readonly object? Value;

        [OutputConstructor]
        private ImportValueResult(object? value)
        {
            Value = value;
        }
    }
}
