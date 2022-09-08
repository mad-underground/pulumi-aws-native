// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Theme Resource Type.
 */
export function getTheme(args: GetThemeArgs, opts?: pulumi.InvokeOptions): Promise<GetThemeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:quicksight:getTheme", {
        "awsAccountId": args.awsAccountId,
        "themeId": args.themeId,
    }, opts);
}

export interface GetThemeArgs {
    awsAccountId: string;
    themeId: string;
}

export interface GetThemeResult {
    /**
     * <p>The Amazon Resource Name (ARN) of the theme.</p>
     */
    readonly arn?: string;
    /**
     * <p>The date and time that the theme was created.</p>
     */
    readonly createdTime?: string;
    /**
     * <p>The date and time that the theme was last updated.</p>
     */
    readonly lastUpdatedTime?: string;
    /**
     * <p>A display name for the theme.</p>
     */
    readonly name?: string;
    /**
     * <p>A valid grouping of resource permissions to apply to the new theme.
     * 			</p>
     */
    readonly permissions?: outputs.quicksight.ThemeResourcePermission[];
    /**
     * <p>A map of the key-value pairs for the resource tag or tags that you want to add to the
     * 			resource.</p>
     */
    readonly tags?: outputs.quicksight.ThemeTag[];
    readonly type?: enums.quicksight.ThemeType;
    readonly version?: outputs.quicksight.ThemeVersion;
}

export function getThemeOutput(args: GetThemeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetThemeResult> {
    return pulumi.output(args).apply(a => getTheme(a, opts))
}

export interface GetThemeOutputArgs {
    awsAccountId: pulumi.Input<string>;
    themeId: pulumi.Input<string>;
}
