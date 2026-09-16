# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.4.1] - 2026-09-16

### Fixed

- Chart: Give Node Feature Discovery's `post-delete` prune hook a `CiliumNetworkPolicy` of its own (`allow-node-feature-discovery-prune-talk-to-apiserver`, a `post-delete` hook of lower weight, deleted with the Job once it succeeded). Helm runs the hook after the release's policies are gone, so on a Cilium cluster with a default-deny policy the `nfd-master -prune` Job could not reach the API server (`dial tcp 172.31.0.1:443: i/o timeout`): the uninstall was retried for about five minutes, every attempt left the Job and a pod in `Error` in `kube-system`, and the NFD labels were never pruned. `gpu-operator.node-feature-discovery.postDeleteCleanup` (upstream's default `true`) is an explicit, documented value now; `false` skips the prune and the policy. ([#167](https://github.com/giantswarm/gpu-operator-app/issues/167))

## [1.4.0] - 2026-09-16

### Changed

- Chart: Update `gpu-operator` to v26.7.0 (`appVersion` 26.7.0).

### Fixed

- Put Flatcar's `/opt/bin` on the toolkit validator's `PATH` (`gpu-operator.validator.toolkit.env`). The injected `nvidia-smi` lives there rather than on the default container `PATH`, so `nvidia-operator-validator` failed with `exec: "nvidia-smi": executable file not found in $PATH`, leaving the device plugin, GFD and DCGM stuck in `Init` and the node advertising no `nvidia.com/gpu`. ([#163](https://github.com/giantswarm/gpu-operator-app/pull/163))
- Add the missing `CiliumNetworkPolicy` for `gpu-feature-discovery`. The chart shipped policies for the operator, node-feature-discovery and the validator, but none selected `app: gpu-feature-discovery`, so GFD could not reach the API server (`dial tcp 172.31.0.1:443: i/o timeout`), crash-looped, and never published the `nvidia.com/gpu.*` node labels. ([#163](https://github.com/giantswarm/gpu-operator-app/pull/163))

### Added

- Chart: `gpu-operator.node-feature-discovery.worker.nodeSelector` and `.affinity` (documented, in the values schema, empty by default) scope Node Feature Discovery's worker — and with it `nvidia.com/gpu.present` and the operands — to the GPU node pool (`giantswarm.io/machine-pool=<cluster>-<pool>`). With the worker on every node, GPU-family nodes of a general Karpenter pool without a driver (AWS `g6f`, fractional-L4 spot) got the validator, device plugin and DCGM exporter stuck in `Init` and kept `ClusterPolicy` not ready. ([#164](https://github.com/giantswarm/gpu-operator-app/issues/164))

## [1.3.0] - 2026-05-14

### Changed

- Go: Update dependencies.
- Chart: Update `gpu-operator` to v26.3.1.

## [1.2.0] - 2026-03-03

### Changed

- Chart: Add `io.giantswarm.application.audience: all` annotation to `Chart.yaml`.
- Chart: Migrate team annotation from `application.giantswarm.io/team` to `io.giantswarm.application.team: tenet`.

## [1.1.1] - 2025-12-05

### Changed

- Chart: Update `gpu-operator` to v25.10.1. ([#29](https://github.com/giantswarm/gpu-operator-app/pull/29))

## [1.1.0] - 2025-10-27

### Changed

- Chart: Update `gpu-operator` to v25.10.0. ([#17](https://github.com/giantswarm/gpu-operator-app/pull/17))

## [1.0.1] - 2025-10-23

### Changed

- Updated E2E tests to use apptest-framework v1.14.0
- Repository: Some chores.
  - CircleCI: Update architect to v6.7.0.
  - Chart: Update GPU Operator to v25.3.4.
  - Tests: Update dependencies and configuration.
  - Repository: Update documentation.
  - Chart: Fix KubeLinter.

## [0.1.0] - 2025-03-11

### Added

- Initial `gpu-operator`.

[Unreleased]: https://github.com/giantswarm/gpu-operator-app/compare/v1.4.1...HEAD
[1.4.1]: https://github.com/giantswarm/gpu-operator-app/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/giantswarm/gpu-operator-app/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/giantswarm/gpu-operator-app/compare/v0.1.0...v1.0.1
[0.1.0]: https://github.com/giantswarm/gpu-operator-app/releases/tag/v0.1.0
