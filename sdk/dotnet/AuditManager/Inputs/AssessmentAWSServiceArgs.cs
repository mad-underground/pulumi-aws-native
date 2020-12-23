// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AuditManager.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsservice.html
    /// </summary>
    public sealed class AssessmentAWSServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-auditmanager-assessment-awsservice.html#cfn-auditmanager-assessment-awsservice-servicename
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public AssessmentAWSServiceArgs()
        {
        }
    }
}
