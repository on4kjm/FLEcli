# Developer Notes

* [Set up tooling on your mac](enableGo.md)
* [Build the project](buildNotes.md)
* [FLE file protocol notes](protocol.md)
* [FLE documentation](Fast%20Log%20Entry%20(FLE).pdf)

* `docker run -it --rm -v "$(PWD)":/usr/local/src/your-app githubchangeloggenerator/github-changelog-generator -u on4kjm -p FLEcli --token $GITHUB_TOKEN -o 1.tmp --unreleased-only`
