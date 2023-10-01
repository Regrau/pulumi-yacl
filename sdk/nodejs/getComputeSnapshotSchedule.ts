// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about a Yandex Compute snapshot schedule. For more information, see
 * [the official documentation](https://cloud.yandex.ru/docs/compute/concepts/snapshot-schedule).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as yandex from "@pulumi/yandex";
 *
 * const mySnapshotSchedule = pulumi.output(yandex.getComputeSnapshotSchedule({
 *     snapshotScheduleId: "some_snapshot_schedule_id",
 * }));
 * ```
 */
export function getComputeSnapshotSchedule(args?: GetComputeSnapshotScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeSnapshotScheduleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("yandex:index/getComputeSnapshotSchedule:getComputeSnapshotSchedule", {
        "description": args.description,
        "diskIds": args.diskIds,
        "folderId": args.folderId,
        "labels": args.labels,
        "name": args.name,
        "retentionPeriod": args.retentionPeriod,
        "schedulePolicies": args.schedulePolicies,
        "snapshotCount": args.snapshotCount,
        "snapshotScheduleId": args.snapshotScheduleId,
        "snapshotSpecs": args.snapshotSpecs,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeSnapshotSchedule.
 */
export interface GetComputeSnapshotScheduleArgs {
    /**
     * An optional description of this snapshot schedule.
     */
    description?: string;
    /**
     * IDs of the disks of this snapshot schedule.
     */
    diskIds?: string[];
    /**
     * ID of the folder that the snapshot schedule belongs to.
     */
    folderId?: string;
    /**
     * A map of labels applied to this snapshot schedule.
     */
    labels?: {[key: string]: string};
    /**
     * The name of the snapshot schedule.
     */
    name?: string;
    /**
     * Retention period applied to snapshots created by this snapshot schedule.
     */
    retentionPeriod?: string;
    /**
     * Schedule policy of the snapshot schedule.
     */
    schedulePolicies?: inputs.GetComputeSnapshotScheduleSchedulePolicy[];
    /**
     * Maximum number of snapshots for every disk of the snapshot schedule.
     */
    snapshotCount?: number;
    /**
     * The ID of a specific snapshot schedule.
     */
    snapshotScheduleId?: string;
    /**
     * Additional attributes for snapshots created by this snapshot schedule.
     */
    snapshotSpecs?: inputs.GetComputeSnapshotScheduleSnapshotSpec[];
}

/**
 * A collection of values returned by getComputeSnapshotSchedule.
 */
export interface GetComputeSnapshotScheduleResult {
    /**
     * SnapshotSchedule creation timestamp.
     */
    readonly createdAt: string;
    /**
     * An optional description of this snapshot schedule.
     */
    readonly description: string;
    /**
     * IDs of the disks of this snapshot schedule.
     */
    readonly diskIds: string[];
    /**
     * ID of the folder that the snapshot schedule belongs to.
     */
    readonly folderId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A map of labels applied to this snapshot schedule.
     */
    readonly labels: {[key: string]: string};
    readonly name: string;
    /**
     * Retention period applied to snapshots created by this snapshot schedule.
     */
    readonly retentionPeriod: string;
    /**
     * Schedule policy of the snapshot schedule.
     */
    readonly schedulePolicies: outputs.GetComputeSnapshotScheduleSchedulePolicy[];
    /**
     * Maximum number of snapshots for every disk of the snapshot schedule.
     */
    readonly snapshotCount: number;
    readonly snapshotScheduleId: string;
    /**
     * Additional attributes for snapshots created by this snapshot schedule.
     */
    readonly snapshotSpecs: outputs.GetComputeSnapshotScheduleSnapshotSpec[];
    /**
     * The status of the snapshot schedule.
     */
    readonly status: string;
}

export function getComputeSnapshotScheduleOutput(args?: GetComputeSnapshotScheduleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeSnapshotScheduleResult> {
    return pulumi.output(args).apply(a => getComputeSnapshotSchedule(a, opts))
}

/**
 * A collection of arguments for invoking getComputeSnapshotSchedule.
 */
export interface GetComputeSnapshotScheduleOutputArgs {
    /**
     * An optional description of this snapshot schedule.
     */
    description?: pulumi.Input<string>;
    /**
     * IDs of the disks of this snapshot schedule.
     */
    diskIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the folder that the snapshot schedule belongs to.
     */
    folderId?: pulumi.Input<string>;
    /**
     * A map of labels applied to this snapshot schedule.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the snapshot schedule.
     */
    name?: pulumi.Input<string>;
    /**
     * Retention period applied to snapshots created by this snapshot schedule.
     */
    retentionPeriod?: pulumi.Input<string>;
    /**
     * Schedule policy of the snapshot schedule.
     */
    schedulePolicies?: pulumi.Input<pulumi.Input<inputs.GetComputeSnapshotScheduleSchedulePolicyArgs>[]>;
    /**
     * Maximum number of snapshots for every disk of the snapshot schedule.
     */
    snapshotCount?: pulumi.Input<number>;
    /**
     * The ID of a specific snapshot schedule.
     */
    snapshotScheduleId?: pulumi.Input<string>;
    /**
     * Additional attributes for snapshots created by this snapshot schedule.
     */
    snapshotSpecs?: pulumi.Input<pulumi.Input<inputs.GetComputeSnapshotScheduleSnapshotSpecArgs>[]>;
}