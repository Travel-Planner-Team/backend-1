## github configuration

- Pull from team main branch

```shell
git remote add upstream https://github.com/Travel-Planner-Team/backend.git
git pull upstream main
```

- Push to local main branch

```shell
git remote set-url origin https://github.com/YOUR_GITHUB_ID/backend.git
git push origin main
```

- Do Pull Request
- Check your remote status with `git remote -v`, it will show something like below:

```shell
origin  https://github.com/YOUR_GITHUB_ID/backend.git (fetch)
origin  https://github.com/YOUR_GITHUB_ID/backend.git (push)
upstream        https://github.com/Travel-Planner-Team/backend.git (fetch)
upstream        https://github.com/Travel-Planner-Team/backend.git (push)
```

## Resources
1. [gorm API](https://gorm.io/docs/index.html)
