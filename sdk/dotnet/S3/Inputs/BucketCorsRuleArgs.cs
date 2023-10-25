// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketCorsRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedHeaders")]
        private InputList<string>? _allowedHeaders;
        public InputList<string> AllowedHeaders
        {
            get => _allowedHeaders ?? (_allowedHeaders = new InputList<string>());
            set => _allowedHeaders = value;
        }

        [Input("allowedMethods", required: true)]
        private InputList<string>? _allowedMethods;
        public InputList<string> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<string>());
            set => _allowedMethods = value;
        }

        [Input("allowedOrigins", required: true)]
        private InputList<string>? _allowedOrigins;
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        [Input("exposedHeaders")]
        private InputList<string>? _exposedHeaders;
        public InputList<string> ExposedHeaders
        {
            get => _exposedHeaders ?? (_exposedHeaders = new InputList<string>());
            set => _exposedHeaders = value;
        }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        public BucketCorsRuleArgs()
        {
        }
        public static new BucketCorsRuleArgs Empty => new BucketCorsRuleArgs();
    }
}
