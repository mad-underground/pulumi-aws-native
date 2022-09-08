// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DBClusterArgs } from "./dbcluster";
export type DBCluster = import("./dbcluster").DBCluster;
export const DBCluster: typeof import("./dbcluster").DBCluster = null as any;

export { DBClusterParameterGroupArgs } from "./dbclusterParameterGroup";
export type DBClusterParameterGroup = import("./dbclusterParameterGroup").DBClusterParameterGroup;
export const DBClusterParameterGroup: typeof import("./dbclusterParameterGroup").DBClusterParameterGroup = null as any;

export { DBInstanceArgs } from "./dbinstance";
export type DBInstance = import("./dbinstance").DBInstance;
export const DBInstance: typeof import("./dbinstance").DBInstance = null as any;

export { DBSubnetGroupArgs } from "./dbsubnetGroup";
export type DBSubnetGroup = import("./dbsubnetGroup").DBSubnetGroup;
export const DBSubnetGroup: typeof import("./dbsubnetGroup").DBSubnetGroup = null as any;

export { GetDBClusterArgs, GetDBClusterResult, GetDBClusterOutputArgs } from "./getDBCluster";
export const getDBCluster: typeof import("./getDBCluster").getDBCluster = null as any;
export const getDBClusterOutput: typeof import("./getDBCluster").getDBClusterOutput = null as any;

export { GetDBClusterParameterGroupArgs, GetDBClusterParameterGroupResult, GetDBClusterParameterGroupOutputArgs } from "./getDBClusterParameterGroup";
export const getDBClusterParameterGroup: typeof import("./getDBClusterParameterGroup").getDBClusterParameterGroup = null as any;
export const getDBClusterParameterGroupOutput: typeof import("./getDBClusterParameterGroup").getDBClusterParameterGroupOutput = null as any;

export { GetDBInstanceArgs, GetDBInstanceResult, GetDBInstanceOutputArgs } from "./getDBInstance";
export const getDBInstance: typeof import("./getDBInstance").getDBInstance = null as any;
export const getDBInstanceOutput: typeof import("./getDBInstance").getDBInstanceOutput = null as any;

export { GetDBSubnetGroupArgs, GetDBSubnetGroupResult, GetDBSubnetGroupOutputArgs } from "./getDBSubnetGroup";
export const getDBSubnetGroup: typeof import("./getDBSubnetGroup").getDBSubnetGroup = null as any;
export const getDBSubnetGroupOutput: typeof import("./getDBSubnetGroup").getDBSubnetGroupOutput = null as any;

utilities.lazyLoad(exports, ["DBCluster"], () => require("./dbcluster"));
utilities.lazyLoad(exports, ["DBClusterParameterGroup"], () => require("./dbclusterParameterGroup"));
utilities.lazyLoad(exports, ["DBInstance"], () => require("./dbinstance"));
utilities.lazyLoad(exports, ["DBSubnetGroup"], () => require("./dbsubnetGroup"));
utilities.lazyLoad(exports, ["getDBCluster","getDBClusterOutput"], () => require("./getDBCluster"));
utilities.lazyLoad(exports, ["getDBClusterParameterGroup","getDBClusterParameterGroupOutput"], () => require("./getDBClusterParameterGroup"));
utilities.lazyLoad(exports, ["getDBInstance","getDBInstanceOutput"], () => require("./getDBInstance"));
utilities.lazyLoad(exports, ["getDBSubnetGroup","getDBSubnetGroupOutput"], () => require("./getDBSubnetGroup"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:docdb:DBCluster":
                return new DBCluster(name, <any>undefined, { urn })
            case "aws-native:docdb:DBClusterParameterGroup":
                return new DBClusterParameterGroup(name, <any>undefined, { urn })
            case "aws-native:docdb:DBInstance":
                return new DBInstance(name, <any>undefined, { urn })
            case "aws-native:docdb:DBSubnetGroup":
                return new DBSubnetGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "docdb", _module)
