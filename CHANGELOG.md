# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/gpu-operator-app/compare/v1.3.0...HEAD
[1.3.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/giantswarm/gpu-operator-app/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/gpu-operator-app/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/giantswarm/gpu-operator-app/compare/v0.1.0...v1.0.1
[0.1.0]: https://github.com/giantswarm/gpu-operator-app/releases/tag/v0.1.0
