# Contributing

PRs and issues are very welcome!
We use GitFlow, so new features should be developed on a `feature/myCoolFeature` branch.

## Releasing

This project uses [GoReleaser](https://goreleaser.com/) to automatically create builds and publish to homebrew, the snap store and the github releases page.

How to release:

- Make sure you have [installed GoReleaser](https://goreleaser.com/install/) and the tools required for releasing ([snapcraft](https://snapcraft.io/)).
- Set a tag, either manually of using a tool that helps with [Semantic Versioning](http://semver.org) & [GitFlow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) `git tag -a v0.4.0 -m "My new version!"`
- `goreleaser --rm-dist` Publish!
