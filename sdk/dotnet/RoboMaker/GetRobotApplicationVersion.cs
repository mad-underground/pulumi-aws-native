// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker
{
    public static class GetRobotApplicationVersion
    {
        /// <summary>
        /// AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.
        /// </summary>
        public static Task<GetRobotApplicationVersionResult> InvokeAsync(GetRobotApplicationVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRobotApplicationVersionResult>("aws-native:robomaker:getRobotApplicationVersion", args ?? new GetRobotApplicationVersionArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.
        /// </summary>
        public static Output<GetRobotApplicationVersionResult> Invoke(GetRobotApplicationVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRobotApplicationVersionResult>("aws-native:robomaker:getRobotApplicationVersion", args ?? new GetRobotApplicationVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRobotApplicationVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetRobotApplicationVersionArgs()
        {
        }
        public static new GetRobotApplicationVersionArgs Empty => new GetRobotApplicationVersionArgs();
    }

    public sealed class GetRobotApplicationVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetRobotApplicationVersionInvokeArgs()
        {
        }
        public static new GetRobotApplicationVersionInvokeArgs Empty => new GetRobotApplicationVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetRobotApplicationVersionResult
    {
        public readonly string? ApplicationVersion;
        public readonly string? Arn;

        [OutputConstructor]
        private GetRobotApplicationVersionResult(
            string? applicationVersion,

            string? arn)
        {
            ApplicationVersion = applicationVersion;
            Arn = arn;
        }
    }
}
