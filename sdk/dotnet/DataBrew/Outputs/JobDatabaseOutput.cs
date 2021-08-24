// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html
    /// </summary>
    [OutputType]
    public sealed class JobDatabaseOutput
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html#cfn-databrew-job-databaseoutput-databaseoptions
        /// </summary>
        public readonly Outputs.JobDatabaseTableOutputOptions DatabaseOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html#cfn-databrew-job-databaseoutput-databaseoutputmode
        /// </summary>
        public readonly string? DatabaseOutputMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-databaseoutput.html#cfn-databrew-job-databaseoutput-glueconnectionname
        /// </summary>
        public readonly string GlueConnectionName;

        [OutputConstructor]
        private JobDatabaseOutput(
            Outputs.JobDatabaseTableOutputOptions databaseOptions,

            string? databaseOutputMode,

            string glueConnectionName)
        {
            DatabaseOptions = databaseOptions;
            DatabaseOutputMode = databaseOutputMode;
            GlueConnectionName = glueConnectionName;
        }
    }
}