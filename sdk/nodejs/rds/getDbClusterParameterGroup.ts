// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::RDS::DBClusterParameterGroup resource creates a new Amazon RDS DB cluster parameter group. For more information, see Managing an Amazon Aurora DB Cluster in the Amazon Aurora User Guide.
 */
export function getDbClusterParameterGroup(args: GetDbClusterParameterGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDbClusterParameterGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:rds:getDbClusterParameterGroup", {
        "dbClusterParameterGroupName": args.dbClusterParameterGroupName,
    }, opts);
}

export interface GetDbClusterParameterGroupArgs {
    dbClusterParameterGroupName: string;
}

export interface GetDbClusterParameterGroupResult {
    /**
     * An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
     */
    readonly parameters?: any;
    /**
     * The list of tags for the cluster parameter group.
     */
    readonly tags?: outputs.rds.DbClusterParameterGroupTag[];
}
/**
 * The AWS::RDS::DBClusterParameterGroup resource creates a new Amazon RDS DB cluster parameter group. For more information, see Managing an Amazon Aurora DB Cluster in the Amazon Aurora User Guide.
 */
export function getDbClusterParameterGroupOutput(args: GetDbClusterParameterGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDbClusterParameterGroupResult> {
    return pulumi.output(args).apply((a: any) => getDbClusterParameterGroup(a, opts))
}

export interface GetDbClusterParameterGroupOutputArgs {
    dbClusterParameterGroupName: pulumi.Input<string>;
}