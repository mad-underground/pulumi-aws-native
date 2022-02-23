// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceCatalog::Portfolio
 */
export function getPortfolio(args: GetPortfolioArgs, opts?: pulumi.InvokeOptions): Promise<GetPortfolioResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:servicecatalog:getPortfolio", {
        "id": args.id,
    }, opts);
}

export interface GetPortfolioArgs {
    id: string;
}

export interface GetPortfolioResult {
    readonly acceptLanguage?: string;
    readonly description?: string;
    readonly displayName?: string;
    readonly id?: string;
    readonly portfolioName?: string;
    readonly providerName?: string;
    readonly tags?: outputs.servicecatalog.PortfolioTag[];
}

export function getPortfolioOutput(args: GetPortfolioOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortfolioResult> {
    return pulumi.output(args).apply(a => getPortfolio(a, opts))
}

export interface GetPortfolioOutputArgs {
    id: pulumi.Input<string>;
}