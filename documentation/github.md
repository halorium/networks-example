## Github

The CI/CD pipeline begins with new commits arriving in Github.

The master branch does not allow direct pushes. The only way into master is through a PR, i.e. on a branch, and that branch must be green on Travis CI. Additionally, at least one other member of the team must review and approve.

Commits enter master via the Github merge PR function. It's limited to rebase merging only so that master is entirely free of merge commits. This choice was made to improve the engineering team's fluency in intermediate git usage, but is not necessary to the successful operation of the system.
