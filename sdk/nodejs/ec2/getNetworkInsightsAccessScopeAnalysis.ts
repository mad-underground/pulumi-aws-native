// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis
 */
export function getNetworkInsightsAccessScopeAnalysis(args: GetNetworkInsightsAccessScopeAnalysisArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInsightsAccessScopeAnalysisResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getNetworkInsightsAccessScopeAnalysis", {
        "networkInsightsAccessScopeAnalysisId": args.networkInsightsAccessScopeAnalysisId,
    }, opts);
}

export interface GetNetworkInsightsAccessScopeAnalysisArgs {
    networkInsightsAccessScopeAnalysisId: string;
}

export interface GetNetworkInsightsAccessScopeAnalysisResult {
    readonly analyzedEniCount?: number;
    readonly endDate?: string;
    readonly findingsFound?: enums.ec2.NetworkInsightsAccessScopeAnalysisFindingsFound;
    readonly networkInsightsAccessScopeAnalysisArn?: string;
    readonly networkInsightsAccessScopeAnalysisId?: string;
    readonly startDate?: string;
    readonly status?: enums.ec2.NetworkInsightsAccessScopeAnalysisStatus;
    readonly statusMessage?: string;
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis
 */
export function getNetworkInsightsAccessScopeAnalysisOutput(args: GetNetworkInsightsAccessScopeAnalysisOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkInsightsAccessScopeAnalysisResult> {
    return pulumi.output(args).apply((a: any) => getNetworkInsightsAccessScopeAnalysis(a, opts))
}

export interface GetNetworkInsightsAccessScopeAnalysisOutputArgs {
    networkInsightsAccessScopeAnalysisId: pulumi.Input<string>;
}
