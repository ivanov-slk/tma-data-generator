# Changelog

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
