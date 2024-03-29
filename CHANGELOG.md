# Changelog

## [1.3.0](https://github.com/ivanov-slk/tma-data-generator/compare/v1.2.0...v1.3.0) (2024-01-27)


### Features

* The generator function now returns random temperature stats within pre-specified ranges. ([28dc46e](https://github.com/ivanov-slk/tma-data-generator/commit/28dc46ec78382bf55e126c3a7495d0bd662cecde))

## [1.2.0](https://github.com/ivanov-slk/tma-data-generator/compare/v1.1.2...v1.2.0) (2024-01-23)

### Maintenance

- Update go version in actions. ([9d29f47](https://github.com/ivanov-slk/tma-data-generator/commit/9d29f47c39a8eab8fc1d2a84a91c6469aa45f1ea))

### Features

- Producer now publishes the same message infinitely with 60-second interval. ([0da9ccb](https://github.com/ivanov-slk/tma-data-generator/commit/0da9ccbfadd34743f638a8a7fac7b8016ddb4907))

### Refactoring/Restructuring

- Use `slog` for logging throughout the service. ([c1a7f89](https://github.com/ivanov-slk/tma-data-generator/commit/c1a7f899f25307be7ee82cf060fbf55f620596d5))

## [1.1.2](https://github.com/ivanov-slk/tma-data-generator/compare/v1.1.1...v1.1.2) (2024-01-13)

### Maintenance

- Remove unneeded log messages. ([a6882eb](https://github.com/ivanov-slk/tma-data-generator/commit/a6882ebb943bf01bc9d569e4054553129ba05cd1))
- Use jetstream.MemoryStorage when creating the stream. This is sufficient for the current purposes of the service. ([3d89447](https://github.com/ivanov-slk/tma-data-generator/commit/3d89447f63f1d13455306161dd782e46f228267a))

## [1.1.1](https://github.com/ivanov-slk/tma-data-generator/compare/v1.1.0...v1.1.1) (2024-01-03)

### Documentation

- Remove references to Kafka, as the service does not use it right now and generally will not be constrained to a specific pub/sub provider. ([191e737](https://github.com/ivanov-slk/tma-data-generator/commit/191e737bd6f1e962421f7df951e09bada64df0c7))

### Refactoring/Restructuring

- Make the package public. ([97cb625](https://github.com/ivanov-slk/tma-data-generator/commit/97cb6251dce2d31d9917668d59006b1518ee6174))
- Make the data generator utilities public. ([9239145](https://github.com/ivanov-slk/tma-data-generator/commit/92391452dc6f473e7f52b999ba4aedd0b5ee217c))

### Testing

- Fix typos in the datetimes used for testing. ([f347299](https://github.com/ivanov-slk/tma-data-generator/commit/f34729912c61ce6167bfd2cd89ab8a8d7a30bdb5))

## [1.1.0](https://github.com/ivanov-slk/tma-data-generator/compare/v1.0.1...v1.1.0) (2024-01-03)

### Features

- Add object representing a temperature stats event and a static output as a working example. ([f1f40f0](https://github.com/ivanov-slk/tma-data-generator/commit/f1f40f0fbeb400e59185d8fb0668314bedcec37f))
- Output the hard-coded `TemperatureStats` object to NATS. ([7c86bb0](https://github.com/ivanov-slk/tma-data-generator/commit/7c86bb0a432e534e3719f8f9bf484b919fd0b2cd))

### Refactoring/Restructuring

- Remove unused code related to testcontainers logging. ([fa3cf09](https://github.com/ivanov-slk/tma-data-generator/commit/fa3cf090f34537e46be8c306fafcf095ce883564))

## [1.0.1](https://github.com/ivanov-slk/tma-data-generator/compare/v1.0.0...v1.0.1) (2023-09-22)

### Testing

- After the message is produced sleep the main process forever to make ArgoCD happy. ([6c5d30e](https://github.com/ivanov-slk/tma-data-generator/commit/6c5d30e040043d96939ec1657c9f18ed5885a961))

## [1.0.0](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.7...v1.0.0) (2023-09-22)

### Maintenance

- Leave nats.Connect() with default settings, as they suffice for the existing functionalities. ([83613e1](https://github.com/ivanov-slk/tma-data-generator/commit/83613e1fe8345605df97972840018869ab49d440))
- release 1.0.0 ([4569bae](https://github.com/ivanov-slk/tma-data-generator/commit/4569bae62877acb747217cf6b0cd6861c84e9559))

## [0.2.7](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.6...v0.2.7) (2023-09-10)

### Maintenance

- Increase NATS connection timeout to 20 seconds to take into account any Linkerd-related timeouts. ([1a740e2](https://github.com/ivanov-slk/tma-data-generator/commit/1a740e21fcd107b96a6dbc2ae50c51140d5db7bc))

## [0.2.6](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.5...v0.2.6) (2023-09-09)

### Maintenance

- Log nats options for debug purposes. ([de78122](https://github.com/ivanov-slk/tma-data-generator/commit/de781229bfdb3dd0d02e8d0d0402da785add7fa0))

## [0.2.5](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.4...v0.2.5) (2023-09-07)

### Maintenance

- Add log mode to error messages. ([07619b8](https://github.com/ivanov-slk/tma-data-generator/commit/07619b8568f2617db148565de14a37f2e65adf31))

### Bug Fixes

- Set connect timeout to nats.Connect(). ([5241b9a](https://github.com/ivanov-slk/tma-data-generator/commit/5241b9a65154bbc13d8da33f4f0afc46c69ce97a))

## [0.2.4](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.3...v0.2.4) (2023-09-02)

### Builds

- Dockerfile builds the main file in 'cmd'. ([ee74ac2](https://github.com/ivanov-slk/tma-data-generator/commit/ee74ac2939320d4b2ec61b3f3889f60ffbfee95f))

### Continuous Integration

- Do not run 'tests' action on PR changes, as PR-s are expected to be made only by release-please for documentation purposes. Track status of the main branch instead. ([58c9849](https://github.com/ivanov-slk/tma-data-generator/commit/58c98499e37d16ae18d485b86230f9cd3709ad4e))

### Bug Fixes

- Add correct executable path in the dockerfile. ([f8bd367](https://github.com/ivanov-slk/tma-data-generator/commit/f8bd3678e42beb78332494c6b1fa726c244de71e))

### Refactoring/Restructuring

- Acceptance test structure follows the DSL-driver model. ([ae56014](https://github.com/ivanov-slk/tma-data-generator/commit/ae56014650dd40a241989595701a175daaecdfdf))
- Delete directory 'mock', as not needed anymore. ([50870e8](https://github.com/ivanov-slk/tma-data-generator/commit/50870e883581d018706e9a1363579ad807cd8bf5))

### Testing

- Add a baseline acceptance test for the service. ([60b0ad7](https://github.com/ivanov-slk/tma-data-generator/commit/60b0ad7c4cd618e33c37c8c4bb26a24faa7fd2f4))

## [0.2.3](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.2...v0.2.3) (2023-04-16)

### Continuous Integration

- Add prefix 'v-' for image tags. ([274b59e](https://github.com/ivanov-slk/tma-data-generator/commit/274b59ebbc5b796f57f5bfa80163d0e102a8ffe1))
- Allow calling the container build workflow only when new tags are pushed. ([d7b8c5f](https://github.com/ivanov-slk/tma-data-generator/commit/d7b8c5fc61855285e4ca0d40fc408fd1fe1e75ae))
- Use the tag from the tag event directly; remove usage of external dependencies for tag fetching. ([a198d30](https://github.com/ivanov-slk/tma-data-generator/commit/a198d30d712a4fae5366417e4b1ae82eb67e6a10))

## [0.2.2](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.1...v0.2.2) (2023-04-16)

### Continuous Integration

- Add more fields for the automatic changelog generation as per the Conventional Changelog. ([b9d7723](https://github.com/ivanov-slk/tma-data-generator/commit/b9d772327daf8a2d7768df97a32015675510999f))

### Maintenance

- Delete CHANGES.md in favor of CHANGELOG.md. ([13283d3](https://github.com/ivanov-slk/tma-data-generator/commit/13283d35d8b4e03c2a91d00a37073a021f72d396))
- Move the unused github actions deprecated. They will be removed in a future version if release-please becomes the main tool for managing tags and releases. ([c909b14](https://github.com/ivanov-slk/tma-data-generator/commit/c909b1441defcf6f455d689638fad0ab1a1ed68d))

## [0.2.1](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.0...v0.2.1) (2023-04-16)

### Bug Fixes

- Run tests on pull request. ([3652f81](https://github.com/ivanov-slk/tma-data-generator/commit/3652f819fffcb7a3aab1a7284561e147f5ee822d))
- Turn off old workflows. ([1e4086d](https://github.com/ivanov-slk/tma-data-generator/commit/1e4086d84ed64681df2c50d405b505d4f7a89f63))

## [0.2.0](https://github.com/ivanov-slk/tma-data-generator/compare/v0.1.9...v0.2.0) (2023-04-16)

### Features

- Add release-please. ([0314357](https://github.com/ivanov-slk/tma-data-generator/commit/03143575d5f7822ca4beb12752583941debfd2d2))

### Bug Fixes

- Test if release-please updates PR. ([1b5256b](https://github.com/ivanov-slk/tma-data-generator/commit/1b5256b6381ca2d67ac291878081474ed97f6689))

---

# v0.1.9

### Bugfixes

- Wait after tag push to ensure tags have reached the repository before the latest is fetched for the container image build.

---

# v0.1.8

### Bugfixes

- Fetch the latest tag available during the container build steps.

---

# v0.1.7

### Maintenance

- Opt out for using compound and reusable workflows, as this is much easier to maintain.

---

# v0.1.6

### Bugfixes

- Configure the tag action to use the correct token

### Maintenance

- Remove some unneeded code.

---

# v0.1.5

### New

- Use `mathieudutour/github-tag-action` for tag.

### Bugfixes

- Correctly set the github token to be used.

---

# v0.1.4

### Bugfixes

- Use the correct secret name.

---

# v0.1.3

### Maintenance

- Use the correct action for fetching the Github application token.

---

# v0.1.2

### Maintenance

- Integrate a Github application to manage actions.

---

# v0.1.1

### Bugfixes

- Container image is now built when new tags are created;
- `Dockerfile` to copy all source code to the image.

---

# v0.1.0

### New

- Add main module returning static output;
- Add Github actions.
