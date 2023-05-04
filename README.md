**The project is under development. His idea is to synchronize two task managers: asana and todoist.**

Before working with this project, create a file with the extension .env in the root directory and fill in the following parameters:
```
TOKENTODOIST=
ASANA_TOKEN=
FOLDERNAME=
USER_NAME=
WORKSPACE_NAME=
REDIS_ADDR=
REDIS_PASS=
```
You also need to install redis on your device. Installation Guide: https://redis.io/docs/getting-started/installation/.

If you are using windows, you may encounter a problem, since redis is not officially supported on this operating system. To solve this problem, you can use docker or wsl 2 versions, these options are discussed in the official redis documentation.

As an alternative solution, I can suggest using a ported version redis for windows from microsoft https://github.com/microsoftarchive/redis/releases.
This is a simpler solution, in my opinion