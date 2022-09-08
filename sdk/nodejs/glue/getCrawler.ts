// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::Crawler
 */
export function getCrawler(args: GetCrawlerArgs, opts?: pulumi.InvokeOptions): Promise<GetCrawlerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:glue:getCrawler", {
        "id": args.id,
    }, opts);
}

export interface GetCrawlerArgs {
    id: string;
}

export interface GetCrawlerResult {
    readonly classifiers?: string[];
    readonly configuration?: string;
    readonly crawlerSecurityConfiguration?: string;
    readonly databaseName?: string;
    readonly description?: string;
    readonly id?: string;
    readonly recrawlPolicy?: outputs.glue.CrawlerRecrawlPolicy;
    readonly role?: string;
    readonly schedule?: outputs.glue.CrawlerSchedule;
    readonly schemaChangePolicy?: outputs.glue.CrawlerSchemaChangePolicy;
    readonly tablePrefix?: string;
    readonly tags?: any;
    readonly targets?: outputs.glue.CrawlerTargets;
}

export function getCrawlerOutput(args: GetCrawlerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCrawlerResult> {
    return pulumi.output(args).apply(a => getCrawler(a, opts))
}

export interface GetCrawlerOutputArgs {
    id: pulumi.Input<string>;
}
