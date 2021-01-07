// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./clusterParameterGroup";
export * from "./clusterSecurityGroup";
export * from "./clusterSecurityGroupIngress";
export * from "./clusterSubnetGroup";

// Import resources to register:
import { Cluster } from "./cluster";
import { ClusterParameterGroup } from "./clusterParameterGroup";
import { ClusterSecurityGroup } from "./clusterSecurityGroup";
import { ClusterSecurityGroupIngress } from "./clusterSecurityGroupIngress";
import { ClusterSubnetGroup } from "./clusterSubnetGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:Redshift:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "cloudformation:Redshift:ClusterParameterGroup":
                return new ClusterParameterGroup(name, <any>undefined, { urn })
            case "cloudformation:Redshift:ClusterSecurityGroup":
                return new ClusterSecurityGroup(name, <any>undefined, { urn })
            case "cloudformation:Redshift:ClusterSecurityGroupIngress":
                return new ClusterSecurityGroupIngress(name, <any>undefined, { urn })
            case "cloudformation:Redshift:ClusterSubnetGroup":
                return new ClusterSubnetGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "Redshift", _module)
