# gpu-operator

![Version: 1.3.0](https://img.shields.io/badge/Version-1.3.0-informational?style=flat-square) ![AppVersion: 26.7.0](https://img.shields.io/badge/AppVersion-26.7.0-informational?style=flat-square)

A Helm chart to deploy NVIDIA GPU Operator with custom Cilium Network Policies.

**Homepage:** <https://github.com/giantswarm/gpu-operator-app>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://helm.ngc.nvidia.com/nvidia | gpu-operator | v26.7.0 |

## Scoping the operands to a GPU node pool

The operator finds its nodes through Node Feature Discovery (NFD). The NFD worker runs on every node and labels each one that carries an NVIDIA PCI device `feature.node.kubernetes.io/pci-10de.present=true`; the operator turns that into `nvidia.com/gpu.present=true` and the `nvidia.com/gpu.deploy.*=true` state labels its operand DaemonSets (validator, device plugin, GPU feature discovery, DCGM exporter) select on. A node with an NVIDIA device but no driver gets the operands too, and they hang in `Init` at the driver validation while `ClusterPolicy` stays not ready. That is what happens when a general Karpenter pool picks a GPU-family instance as an ordinary spot node (AWS `g6f`, the fractional-L4 family) on a cluster whose driver lives in the GPU pool's machine image.

On a cluster whose GPU nodes come from a dedicated pool, scope the worker to that pool — `giantswarm.io/machine-pool=<cluster>-<pool>` is the label every node of a Giant Swarm node pool carries:

```yaml
gpu-operator:
  node-feature-discovery:
    worker:
      nodeSelector:
        giantswarm.io/machine-pool: mycluster-gpu-l4
```

Several pools take the set form:

```yaml
gpu-operator:
  node-feature-discovery:
    worker:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: giantswarm.io/machine-pool
                    operator: In
                    values: [mycluster-gpu-l4, mycluster-gpu-a10g]
```

A node without a worker gets no NFD labels, so neither `nvidia.com/gpu.present` nor operands; the NFD master and garbage collector and the operator itself are Deployments and stay where they are. The worker's `NodeFeature` object is owned by its pod, so a node the worker leaves loses its NFD labels, the operator withdraws `nvidia.com/gpu.present` and the state labels, and operands already stuck there end — scoping a running installation needs no manual clean-up. Both values are empty by default: the chart then behaves as upstream and runs the worker everywhere. cluster-manager's `<cluster>-gpu-operator` release sets the set form for the cluster's GPU pools. The other way to keep the operands off such nodes is to keep GPU instance families out of the general pool's requirements.

## Uninstalling: the Node Feature Discovery prune hook

Node Feature Discovery (NFD) labels the nodes its worker runs on (`feature.node.kubernetes.io/*`, and the `nvidia.com/gpu.*` labels GPU feature discovery publishes through it). Uninstalling the release removes the workloads, not the labels, so NFD's subchart ships a Helm `post-delete` hook — `gpu-operator.node-feature-discovery.postDeleteCleanup`, `true` by default as upstream — a Job `<release>-node-feature-discovery-prune` running `nfd-master -prune` that takes the NFD-managed labels, annotations and extended resources off every node.

Helm runs the hooks of an uninstall after the release's objects are gone, the chart's `allow-node-feature-discovery-talk-to-apiserver` policy included. On a cluster with a default-deny network policy the prune pod cannot reach the API server then (`dial tcp <apiserver>:443: i/o timeout`): the hook fails, the uninstall is retried for minutes, every attempt leaves the Job and a pod in `Error` in the release namespace, and the labels stay. The chart therefore gives the hook a policy of its own, `allow-node-feature-discovery-prune-talk-to-apiserver`: a `post-delete` hook too, with a lower `helm.sh/hook-weight` so it exists before the Job, selecting the prune pod (`role: prune`) and allowing egress to the `kube-apiserver` entity. Both carry `helm.sh/hook-delete-policy: before-hook-creation,hook-succeeded`: they are deleted once the prune succeeded, and a failed attempt is cleaned up before the next one.

`postDeleteCleanup: false` renders neither. The NFD labels then outlive the uninstall until NFD runs on the node again or the node is replaced; with the worker scoped to a GPU node pool that is confined to the pool's nodes.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.imageRegistry | string | `"gsoci.azurecr.io"` |  |
| gpu-operator.validator.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.validator.toolkit.env[0].name | string | `"PATH"` |  |
| gpu-operator.validator.toolkit.env[0].value | string | `"/opt/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"` |  |
| gpu-operator.operator.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.driver.enabled | bool | `false` |  |
| gpu-operator.driver.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.driver.manager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.toolkit.enabled | bool | `false` |  |
| gpu-operator.toolkit.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.devicePlugin.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.dcgm.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.dcgmExporter.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.gfd.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.migManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.nodeStatusExporter.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.gds.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.gdrcopy.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vgpuManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vgpuManager.driverManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vgpuDeviceManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vfioManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vfioManager.driverManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.sandboxDevicePlugin.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.kataSandboxDevicePlugin.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.ccManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.node-feature-discovery.postDeleteCleanup | bool | `true` | Run Node Feature Discovery's `post-delete` hook, upstream's default: a Job `nfd-master -prune` that takes the NFD-managed labels, annotations and extended resources off the nodes when the release is uninstalled. The chart gives the hook a `CiliumNetworkPolicy` of its own (a `post-delete` hook of lower weight, deleted with the Job), so the prune reaches the API server on a cluster with a default-deny network policy — see "Uninstalling: the Node Feature Discovery prune hook". `false` skips the prune and its policy; the labels then stay on the nodes until Node Feature Discovery runs there again or the nodes are replaced. |
| gpu-operator.node-feature-discovery.image.repository | string | `"gsoci.azurecr.io/giantswarm/node-feature-discovery"` |  |
| gpu-operator.node-feature-discovery.worker.nodeSelector | object | `{}` | Node selector of the Node Feature Discovery worker. Empty runs the worker on every node, as upstream does: every node with an NVIDIA PCI device is then labelled `nvidia.com/gpu.present=true` and gets the operands, driver or not. On a cluster whose general node pool may pick GPU instance families without a driver (AWS `g6f`, fractional-L4 spot), set it to the GPU node pool's label, `giantswarm.io/machine-pool: <cluster>-<pool>` — see "Scoping the operands to a GPU node pool". |
| gpu-operator.node-feature-discovery.worker.affinity | object | `{}` | Node affinity of the Node Feature Discovery worker, the set form of `nodeSelector`: a `giantswarm.io/machine-pool In [<cluster>-<pool>, ...]` expression covers several GPU node pools. |
