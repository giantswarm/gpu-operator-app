# gpu-operator

![Version: 1.0.1](https://img.shields.io/badge/Version-1.0.1-informational?style=flat-square) ![AppVersion: 25.3.4](https://img.shields.io/badge/AppVersion-25.3.4-informational?style=flat-square)

A Helm chart to deploy NVIDIA GPU Operator with custom Cilium Network Policies.

**Homepage:** <https://github.com/giantswarm/gpu-operator-app>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://helm.ngc.nvidia.com/nvidia | gpu-operator | 25.3.4 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.imageRegistry | string | `"gsoci.azurecr.io"` |  |
| gpu-operator.ccManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.dcgm.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.dcgmExporter.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.devicePlugin.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.driver.enabled | bool | `false` |  |
| gpu-operator.driver.manager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.gdrCopy.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.gds.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.gfd.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.kataManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.manager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.migManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.node-feature-discovery.image.repository | string | `"gsoci.azurecr.io/giantswarm/node-feature-discovery"` |  |
| gpu-operator.nodeStatusExporter.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.operator.initContainer.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.operator.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.sandboxDevicePlugin.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.toolkit.enabled | bool | `false` |  |
| gpu-operator.validator.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vfioManager.driverManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vfioManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vgpuDeviceManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
| gpu-operator.vgpuManager.repository | string | `"gsoci.azurecr.io/giantswarm"` |  |
