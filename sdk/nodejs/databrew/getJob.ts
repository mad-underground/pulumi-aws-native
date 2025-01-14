// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataBrew::Job.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:databrew:getJob", {
        "name": args.name,
    }, opts);
}

export interface GetJobArgs {
    /**
     * Job name
     */
    name: string;
}

export interface GetJobResult {
    readonly dataCatalogOutputs?: outputs.databrew.JobDataCatalogOutput[];
    readonly databaseOutputs?: outputs.databrew.JobDatabaseOutput[];
    /**
     * Dataset name
     */
    readonly datasetName?: string;
    /**
     * Encryption Key Arn
     */
    readonly encryptionKeyArn?: string;
    /**
     * Encryption mode
     */
    readonly encryptionMode?: enums.databrew.JobEncryptionMode;
    /**
     * Job Sample
     */
    readonly jobSample?: outputs.databrew.JobSample;
    /**
     * Log subscription
     */
    readonly logSubscription?: enums.databrew.JobLogSubscription;
    /**
     * Max capacity
     */
    readonly maxCapacity?: number;
    /**
     * Max retries
     */
    readonly maxRetries?: number;
    /**
     * Output location
     */
    readonly outputLocation?: outputs.databrew.JobOutputLocation;
    readonly outputs?: outputs.databrew.JobOutput[];
    /**
     * Profile Job configuration
     */
    readonly profileConfiguration?: outputs.databrew.JobProfileConfiguration;
    /**
     * Project name
     */
    readonly projectName?: string;
    readonly recipe?: outputs.databrew.JobRecipe;
    /**
     * Role arn
     */
    readonly roleArn?: string;
    /**
     * Timeout
     */
    readonly timeout?: number;
    /**
     * Data quality rules configuration
     */
    readonly validationConfigurations?: outputs.databrew.JobValidationConfiguration[];
}
/**
 * Resource schema for AWS::DataBrew::Job.
 */
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply((a: any) => getJob(a, opts))
}

export interface GetJobOutputArgs {
    /**
     * Job name
     */
    name: pulumi.Input<string>;
}
