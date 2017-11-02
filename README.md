# Typetalk Google Translate Bot

This repository contains a demo of how to use webhooks with Typetalk. 

If a Typetalk user posts a non-English message, this Bot will attempt to translate it. If successful, the bot will then post a reply with the English translation.

## Getting Started

### 1. Get API key from Google

Refer to the following Google documentation for instructions on how to retrieve your own API key:
https://support.google.com/cloud/answer/6158862

Save it - you'll need it later!

### 2. Clone this repository

Clone this repository on a publicly accessible server.

```
# via SSH:
$ git clone git@github.com:nulab/typetalk-google-translate-bot.git

# via HTTPS:
$ git clone https://github.com/nulab/typetalk-google-translate-bot.git
```

### 3. Run application

Fetch the package dependencies, set the Google API key you created earlier, then run the application.

```
$ go get all
$ export $GOOGLE_API_KEY=????????????????
$ go run app.go
```

### 4. Create Typetalk bot

Create a Typetalk bot in the **Topic** for which messages should be translated to English.

For more details, refer to our Typetalk webhook documentation:
https://developer.nulab-inc.com/docs/typetalk/#webhook

The Webhook URL should be your host name, and the port is `12345`.

e.g: `https://example.com:12345/`

If you'd like to use this demo app in a real environment, we strongly recommend that you configure a reverse proxy to use a different port. (We also recommend setting up HTTPS!)

### 5. Try it out! 

Post a non-English message, such as "こんにちは".
If everything went well, the bot should reply with the message, "Hello"!
