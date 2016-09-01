# git-mr

Do merge request from terminal to GitLab

## Installation

Download binaries or install with go get and build. Then move to executable path. 

```
go get github.com/frozzare/git-mr
```

## Usage

Required environment variables:

```
GITLAB_URL=http://git.example.com
GITLAB_TOKEN=xxx
```

```
git mr -s=source -t=develop -m='Merge request title' -p=example/web
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
