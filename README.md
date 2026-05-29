GitHub actions
- github.com's automation framework
- usually used for CI/CD
- event driven

Workflow
- automation yaml file
- lives at `.github/workflows`
- contains one or more jobs

Job
- a collection of steps
- runs on an individual compute instance - runner
- by default jobs have no dependencies and run in parallely (if enough runners are available)

Steps
- tasks that run in the context of a job's runner
- consist of a shell command or an action
- share files
- changes to env vars don't carry over to the next step

Action
- actions are the superpower of GitHub actions
- a pre-defined, reusable set of jobs or code that performs specific tasks
- you can find actions in the GitHub Marketplace
- you can write your own actions (three types: container, JavaScript and composite)

Creating a release

```
$ git tag v0.0.1
$ git push origin v0.0.1 
```