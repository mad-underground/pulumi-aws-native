// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AlertArgs } from "./alert";
export type Alert = import("./alert").Alert;
export const Alert: typeof import("./alert").Alert = null as any;
utilities.lazyLoad(exports, ["Alert"], () => require("./alert"));

export { AnomalyDetectorArgs } from "./anomalyDetector";
export type AnomalyDetector = import("./anomalyDetector").AnomalyDetector;
export const AnomalyDetector: typeof import("./anomalyDetector").AnomalyDetector = null as any;
utilities.lazyLoad(exports, ["AnomalyDetector"], () => require("./anomalyDetector"));

export { GetAlertArgs, GetAlertResult, GetAlertOutputArgs } from "./getAlert";
export const getAlert: typeof import("./getAlert").getAlert = null as any;
export const getAlertOutput: typeof import("./getAlert").getAlertOutput = null as any;
utilities.lazyLoad(exports, ["getAlert","getAlertOutput"], () => require("./getAlert"));

export { GetAnomalyDetectorArgs, GetAnomalyDetectorResult, GetAnomalyDetectorOutputArgs } from "./getAnomalyDetector";
export const getAnomalyDetector: typeof import("./getAnomalyDetector").getAnomalyDetector = null as any;
export const getAnomalyDetectorOutput: typeof import("./getAnomalyDetector").getAnomalyDetectorOutput = null as any;
utilities.lazyLoad(exports, ["getAnomalyDetector","getAnomalyDetectorOutput"], () => require("./getAnomalyDetector"));


// Export enums:
export * from "../types/enums/lookoutmetrics";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:lookoutmetrics:Alert":
                return new Alert(name, <any>undefined, { urn })
            case "aws-native:lookoutmetrics:AnomalyDetector":
                return new AnomalyDetector(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "lookoutmetrics", _module)
