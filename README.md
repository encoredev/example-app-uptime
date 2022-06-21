<img width="200px" src="https://encore.dev/assets/branding/logo/logo.svg" alt="Encore - The Backend Development Engine" />

# Uptime Monitoring System

This is an Encore application that continuously monitors the uptime of a list of websites.

When it detects a website going down it posts a Slack message notifying that the website is down, and another message when the website is back up again.

![Frontend](./images/frontend.png)
![Encore Flow](./images/encore-flow.png)

This took about 4 hours to build from scratch, including tests using [Encore](https://encore.dev). It took 2 minutes to deploy, including databases, Pub/Sub topics, and cron jobs.

## Tutorial

Check out the [written tutorial (30 min)](https://encore.dev/docs/tutorials/uptime) to learn how to build this application from scratch!

## Install

You will need the [Encore CLI](https://encore.dev/docs/install)
to run and deploy this application.

To install Encore, run:
```bash
# macOS 
brew install encoredev/tap/encore

# Windows
iwr https://encore.dev/install.ps1 | iex

# Linux
curl -L https://encore.dev/install.sh | bash
```

Clone and run the app locally:
```bash
git clone git@github.com:encoredev/example-app-uptime.git
cd example-app-uptime

# Log in to Encore
encore auth login

# Set the Slack webhook secret (see tutorial above)
encore secret set SlackWebhookURL

# Run the app
encore run
```

## Using the API

```bash
# Check if a given site is up (defaults to 'https://' if left out)
$ curl 'http://localhost:4000/ping/google.com'

# Add a site to be automatically pinged every 5 minutes
curl 'http://localhost:4000/site' -d '{"url":"google.com"}'

# Check all tracked sites immediately
curl -X POST 'http://localhost:4000/checkall'

# Get the current status of all tracked sites
curl 'http://localhost:4000/sites'
```

## Deployment

```bash
encore app create my-oncall-app-name
git push origin main
```

Then head over to <https://app.encore.dev> to find out your production URL, and off you go into the clouds!

## Testing

```bash
encore test ./...
```

## Contributing

All contributions are welcome! All we ask is that you adhere to the [Code of Conduct](https://github.com/encoredev/encore/blob/main/CODE_OF_CONDUCT.md)

- [Quick Start with Encore](https://encore.dev/docs/quick-start)
- [Create an Account with Encore](https://encore.dev/login)
- [Go Cheatsheet](https://encore.dev/guide/go.mod)