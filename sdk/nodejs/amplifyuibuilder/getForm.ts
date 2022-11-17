// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::AmplifyUIBuilder::Form Resource Type
 */
export function getForm(args: GetFormArgs, opts?: pulumi.InvokeOptions): Promise<GetFormResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:amplifyuibuilder:getForm", {
        "appId": args.appId,
        "environmentName": args.environmentName,
        "id": args.id,
    }, opts);
}

export interface GetFormArgs {
    appId: string;
    environmentName: string;
    id: string;
}

export interface GetFormResult {
    readonly appId?: string;
    readonly cta?: outputs.amplifyuibuilder.FormCTA;
    readonly dataType?: outputs.amplifyuibuilder.FormDataTypeConfig;
    readonly environmentName?: string;
    readonly fields?: outputs.amplifyuibuilder.FormFieldsMap;
    readonly formActionType?: enums.amplifyuibuilder.FormActionType;
    readonly id?: string;
    readonly name?: string;
    readonly schemaVersion?: string;
    readonly sectionalElements?: outputs.amplifyuibuilder.FormSectionalElementMap;
    readonly style?: outputs.amplifyuibuilder.FormStyle;
}

export function getFormOutput(args: GetFormOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFormResult> {
    return pulumi.output(args).apply(a => getForm(a, opts))
}

export interface GetFormOutputArgs {
    appId: pulumi.Input<string>;
    environmentName: pulumi.Input<string>;
    id: pulumi.Input<string>;
}