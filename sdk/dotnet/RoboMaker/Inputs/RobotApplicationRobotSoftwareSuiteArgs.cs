// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker.Inputs
{

    public sealed class RobotApplicationRobotSoftwareSuiteArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public RobotApplicationRobotSoftwareSuiteArgs()
        {
        }
    }
}