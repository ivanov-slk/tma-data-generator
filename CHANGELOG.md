# Changelog

## [0.2.2](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.1...v0.2.2) (2023-04-16)


### Continuous Integration

* Add more fields for the automatic changelog generation as per the Conventional Changelog. ([b9d7723](https://github.com/ivanov-slk/tma-data-generator/commit/b9d772327daf8a2d7768df97a32015675510999f))


### Maintenance

* Delete CHANGES.md in favor of CHANGELOG.md. ([13283d3](https://github.com/ivanov-slk/tma-data-generator/commit/13283d35d8b4e03c2a91d00a37073a021f72d396))
* Move the unused github actions deprecated. They will be removed in a future version if release-please becomes the main tool for managing tags and releases. ([c909b14](https://github.com/ivanov-slk/tma-data-generator/commit/c909b1441defcf6f455d689638fad0ab1a1ed68d))

## [0.2.1](https://github.com/ivanov-slk/tma-data-generator/compare/v0.2.0...v0.2.1) (2023-04-16)


### Bug Fixes

* Run tests on pull request. ([3652f81](https://github.com/ivanov-slk/tma-data-generator/commit/3652f819fffcb7a3aab1a7284561e147f5ee822d))
* Turn off old workflows. ([1e4086d](https://github.com/ivanov-slk/tma-data-generator/commit/1e4086d84ed64681df2c50d405b505d4f7a89f63))

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
