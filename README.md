# forget-me-please

This project is all about the [Right to Be Forgotten](https://en.wikipedia.org/wiki/Right_to_be_forgotten):

- https://ico.org.uk/for-organisations/guide-to-data-protection/guide-to-the-general-data-protection-regulation-gdpr/individual-rights/right-to-erasure/
- https://www.mtsu.edu/first-amendment/article/1562/right-to-be-forgotten
- https://www.stanfordlawreview.org/online/privacy-paradox-the-right-to-be-forgotten/
- https://www.bbc.com/news/technology-49808208
- https://gdpr.eu/right-to-be-forgotten/

People make mistakes. People have a right to rehabilitate themselves. People can grow.

But the birth of the Internet along with cheaper and ever expanding storage capabilities created an "Eternal Memory" that not only preserves evidence of your accomplishments but also your misdeeds potentially forever. Okay, one could go esoteric and argue that this has always existed via God, Passage of Time, etc. having the same abilities to record our lives and for this argument I'll conceed that point (though I'm not sure I wholly agree). As the big change is now that other living people can access that information whenever they wish. 

## Twitter

At current this application runs locally and requires a few command line supplied arguments.

```
-twitterAuthBearer "" -twitterUsername "" -twitterAccessToken "" -twitterAccessTokenSecret "" -twitterApiKey "" -twitterApiKeySecret "" -twitterExemptUsers ""
```

```golang
var twitterAuthBearer = flag.String("twitterAuthBearer", "", "Twitter Authorization Bearer Token")
var twitterUsername = flag.String("twitterUsername", "", "Twitter User Name")
var twitterAccessToken = flag.String("twitterAccessToken", "", "Twitter Access Token")
var twitterAccessTokenSecret = flag.String("twitterAccessTokenSecret", "", "Twitter Access Token Secret")
var twitterApiKey = flag.String("twitterApiKey", "", "Twitter Consumer API Key")
var twitterApiKeySecret = flag.String("twitterApiKeySecret", "", "Twitter Consumer API Secret")
var twitterOAuthCallBackUrl = flag.String("twitterOAuthCallBackUrl", "oob", "OAuth Call Back URL")
var twitterExemptUsers = flag.String("twitterExemptUsers", "", "Twitter users whose (Re)Tweets you want to keep")
var twitterArchive = flag.String("twitterArchive", "", "Twitter users downloaded archive")
```

These values can be obtained by creating a free twitter developer account.

There are plans to make these values optionally supplied via command line prompts.

### Debugging 

You can check the values of any tweet using the following link format: https://twitter.com/anyuser/status/{tweet-id}

### Limitations

Currently the Twitter API only allows you to retrieve the latest 3200 tweets that a user has interacted with. Twitter has acknowledged this problem, made some empty comments about addressing it, and two years later there is no fix.

#### Archived Data

A way around this 3200 tweet limit is to manually download your twitter archive data here: https://twitter.com/settings/download_your_data. Usually these requests are processed within 24 - 72 hours. 

### Browsers Console Scripts

```javascript
setInterval(() => {
  for (const d of document.querySelectorAll('div[data-testid="unlike"]')) {
    d.click()
  }
  window.scrollTo(0, document.body.scrollHeight)
}, 3000)
```

## Facebook

Only way to manage this is to just delete your account.

## Instagram

Here is how you can manage your Instagram presence.

## Reddit

Here is how you can manage your Reddit presence.

### Kubernetes

https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/

https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#use-environment-variables-to-define-arguments
