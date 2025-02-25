// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::B2BI::Transformer Resource Type
 */
export function getTransformer(args: GetTransformerArgs, opts?: pulumi.InvokeOptions): Promise<GetTransformerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:b2bi:getTransformer", {
        "transformerId": args.transformerId,
    }, opts);
}

export interface GetTransformerArgs {
    transformerId: string;
}

export interface GetTransformerResult {
    readonly createdAt?: string;
    readonly ediType?: outputs.b2bi.TransformerEdiTypeProperties;
    readonly fileFormat?: enums.b2bi.TransformerFileFormat;
    readonly mappingTemplate?: string;
    readonly modifiedAt?: string;
    readonly name?: string;
    readonly sampleDocument?: string;
    readonly status?: enums.b2bi.TransformerStatus;
    readonly tags?: outputs.Tag[];
    readonly transformerArn?: string;
    readonly transformerId?: string;
}
/**
 * Definition of AWS::B2BI::Transformer Resource Type
 */
export function getTransformerOutput(args: GetTransformerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransformerResult> {
    return pulumi.output(args).apply((a: any) => getTransformer(a, opts))
}

export interface GetTransformerOutputArgs {
    transformerId: pulumi.Input<string>;
}
