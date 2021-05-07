# psping
A GO implementation CLI tool, similar as Sysinterals' PSPing: https://docs.microsoft.com/en-us/sysinternals/downloads/psping

# Install as CLI

```shell
go install github.com/riveryc/psping/cmd/psping@latest
```

# Build it as docker image

```shell
docker build -t psping .
```

# Use it

```shell
# binanry
$ psping 8.8.8.8:53
$ psping google.com:443

# Docker
$ docker run -it --rm psping 8.8.8.8:53
```