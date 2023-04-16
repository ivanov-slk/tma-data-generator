# Changelog

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