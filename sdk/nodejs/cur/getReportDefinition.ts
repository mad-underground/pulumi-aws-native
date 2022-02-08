// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::CUR::ReportDefinition resource creates a Cost & Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.
 */
export function getReportDefinition(args: GetReportDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetReportDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cur:getReportDefinition", {
        "reportName": args.reportName,
    }, opts);
}

export interface GetReportDefinitionArgs {
    /**
     * The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.
     */
    reportName: string;
}

export interface GetReportDefinitionResult {
    /**
     * A list of manifests that you want Amazon Web Services to create for this report.
     */
    readonly additionalArtifacts?: enums.cur.ReportDefinitionAdditionalArtifactsItem[];
    /**
     * The compression format that AWS uses for the report.
     */
    readonly compression?: enums.cur.ReportDefinitionCompression;
    /**
     * The format that AWS saves the report in.
     */
    readonly format?: enums.cur.ReportDefinitionFormat;
    /**
     * Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees.
     */
    readonly refreshClosedReports?: boolean;
    /**
     * The S3 bucket where AWS delivers the report.
     */
    readonly s3Bucket?: string;
    /**
     * The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces.
     */
    readonly s3Prefix?: string;
    /**
     * The region of the S3 bucket that AWS delivers the report into.
     */
    readonly s3Region?: string;
}

export function getReportDefinitionOutput(args: GetReportDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReportDefinitionResult> {
    return pulumi.output(args).apply(a => getReportDefinition(a, opts))
}

export interface GetReportDefinitionOutputArgs {
    /**
     * The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.
     */
    reportName: pulumi.Input<string>;
}
