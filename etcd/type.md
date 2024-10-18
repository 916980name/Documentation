#### message KeyValue
Field
1. version
    - deletion set to 0.
    - modification increase.
1. lease
    - is ID of the lease attached to key.
    - lease expires, key deleted.
    - lease == 0, no lease attached.

#### V3 client function
1. (default disabled) Sync synchronizes client's endpoints with the known endpoints from the etcd membership.

    Will update `client.cfg.Endpoints`