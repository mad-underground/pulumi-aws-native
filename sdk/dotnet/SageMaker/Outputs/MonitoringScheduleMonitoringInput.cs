// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Outputs
{

    [OutputType]
    public sealed class MonitoringScheduleMonitoringInput
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringinput.html#cfn-sagemaker-monitoringschedule-monitoringinput-endpointinput
        /// </summary>
        public readonly Outputs.MonitoringScheduleEndpointInput EndpointInput;

        [OutputConstructor]
        private MonitoringScheduleMonitoringInput(Outputs.MonitoringScheduleEndpointInput EndpointInput)
        {
            this.EndpointInput = EndpointInput;
        }
    }
}
