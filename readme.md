# GHBan

To block multiple accounts across multiple organizations.

Useful during bots attacks.

```
NAME:
   ghban - Block multiple accounts on multiple GitHub organizations.

USAGE:
   ghban [global options] command [command options] 

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value                    GitHub personal access token [$GITHUB_TOKEN]
   --orgs value [ --orgs value ]    GitHub's organization names
   --users value [ --users value ]  GitHub usernames
   --help, -h                       show help
```

## Example

```bash

ghban --orgs myorg01,myorg02 --users spammer01,spammer02,spammer03
```
