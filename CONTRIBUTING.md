Hello! Thank you for choosing to help contribute to the sendgrid-go project. There are many ways you can contribute and help is always welcome.

We use [Milestones](https://github.com/sendgrid/sendgrid-go/milestones) to help define current roadmaps, please feel free to grab an issue from the current milestone. Please indicate that you have begun work on it to avoid collisions. Once a PR is made, community review, comments, suggestions and additional PRs are welcomed and encouraged.

* [Feature Request](#feature_request)
* [Submit a Bug Report](#submit_a_bug_report)
* [Improvements to the Codebase](#improvements_to_the_codebase)
* [Understanding the Code Base](#understanding_the_codebase)
* [Testing](#testing)
* [Style Guidelines & Naming Conventions](#style_guidelines_and_naming_conventions)
* [Creating a Pull Request](#creating_a_pull_request)

There are a few ways to contribute, which we'll enumerate below:

<a name="feature_request"></a>
## Feature Request

If you'd like to make a feature request, please read this section.

The GitHub issue tracker is the preferred channel for library feature requests, but please respect the following restrictions:

- Please **search for existing issues** in order to ensure we don't have duplicate bugs/feature requests.
- Please be respectful and considerate of others when commenting on issues

<a name="submit_a_bug_report"></a>
## Submit a Bug Report

Note: DO NOT include your credentials in ANY code examples, descriptions, or media you make public.

A software bug is a demonstrable issue in the code base. In order for us to diagnose the issue and respond as quickly as possible, please add as much detail as possible into your bug report.

Before you decide to create a new issue, please try the following:

1. Check the Github issues tab if the identified issue has already been reported, if so, please add a +1 to the existing post.
2. Update to the latest version of this code and check if issue has already been fixed
3. Copy and fill in the Bug Report Template we have provided below

### Please use our Bug Report Template

In order to make the process easier, we've included a sample bug report template (borrowed from [Ghost](https://github.com/TryGhost/Ghost/)). The template uses [GitHub flavored markdown](https://help.github.com/articles/github-flavored-markdown/) for formatting.

```
Short and descriptive example bug report title

#### Issue Summary

A summary of the issue and the environment in which it occurs. If suitable, include the steps required to reproduce the bug. Please feel free to include screenshots, screencasts, code examples.

#### Steps to Reproduce

1. This is the first step
2. This is the second step
3. Further steps, etc.

Any other information you want to share that is relevant to the issue being reported. Especially, why do you consider this to be a bug? What do you expect to happen instead?

#### Technical details:

* sendgrid-go Version: master (latest commit: 2cb34372ef0f31352f7c90015a45e1200cb849da)
* Go Version: 1.5.1
```

<a name="improvements_to_the_codebase"></a>
## Improvements to the Codebase

We welcome direct contributions to the rest code base. Thank you!

### Development Environment ###

#### Install and run locally ####

##### Prerequisites #####

* Go 1.3 through 1.5
* github.com/sendgrid/rest

##### Initial setup: #####

```
git clone https://github.com/sendgrid/sendgrid-go.git
cd sendgrid-go
```

##### Execute: #####

See the [examples folder](https://github.com/sendgrid/rest/tree/v3beta/examples) to get started quickly.

You will need to setup the following environment to use the SendGrid example:

```
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > sendgrid.env
echo "sendgrid.env" >> .gitignore
source ./sendgrid.env
go run examples/example.go
```

<a name="understanding_the_codebase"></a>
## Understanding the Code Base

**/examples**

Working examples that demonstrate usage.

**sendgrid.go**

The main function that does the heavy lifting (and external entry point) is `API`.

<a name="testing"></a>
## Testing

All PRs require passing tests before the PR will be reviewed.

All test files are in `[sendgrid_test.go](https://github.com/sendgrid/sendgrid-go/tree/v3beta/sendgrid_test.go)`.

For the purposes of contributing to this repo, please update the [`sendgrid_test.go`](https://github.com/sendgrid/sendgrid-go/tree/v3beta/sendgrid_test.go) file with unit tests as you modify the code.

Run the test:

`go test -v`

<a name="style_guidelines_and_naming_conventions"></a>
## Style Guidelines & Naming Conventions

Generally, we follow the style guidelines as suggested by the official language. However, we ask that you conform to the styles that already exist in the library. If you wish to deviate, please explain your reasoning.

* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

Please run your code through [fmt](https://blog.golang.org/go-fmt-your-code)

### Directory Structure

* `examples` for example calls

## Creating a Pull Request<a name="creating_a_pull_request"></a>

1. [Fork](https://help.github.com/fork-a-repo/) the project, clone your fork,
   and configure the remotes:

   ```bash
   # Clone your fork of the repo into the current directory
   git clone https://github.com/sendgrid/sendgrid-go
   # Navigate to the newly cloned directory
   cd rest
   # Assign the original repo to a remote called "upstream"
   git remote add upstream https://github.com/sendgrid/sendgrid-go
   ```

2. If you cloned a while ago, get the latest changes from upstream:

   ```bash
   git checkout <dev-branch>
   git pull upstream <dev-branch>
   ```

3. Create a new topic branch (off the main project development branch) to
   contain your feature, change, or fix:

   ```bash
   git checkout -b <topic-branch-name>
   ```

4. Commit your changes in logical chunks. Please adhere to these [git commit
   message guidelines](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html)
   or your code is unlikely be merged into the main project. Use Git's
   [interactive rebase](https://help.github.com/articles/interactive-rebase)
   feature to tidy up your commits before making them public.

4a. Create tests.

4b. Create or update the example code that demonstrates the functionality of this change to the code.

5. Locally merge (or rebase) the upstream development branch into your topic branch:

   ```bash
   git pull [--rebase] upstream master
   ```

6. Push your topic branch up to your fork:

   ```bash
   git push origin <topic-branch-name>
   ```

7. [Open a Pull Request](https://help.github.com/articles/using-pull-requests/)
    with a clear title and description against the `master` branch. All tests must be passing before we will review the PR.

If you have any additional questions, please feel free to [email](mailto:dx@sendgrid.com) us or create an issue in this repo.
