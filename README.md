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
```

These values can be obtained by creating a free twitter developer account.

There are plans to make these values supplied via command line prompts, or ideally a self-hosted web interface but that day is not yet here.

## Facebook

Only way to manage this is to just delete your account.

## Instagram

Here is how you can manage your Instagram presence.

## Reddit

Here is how you can manage your Reddit presence.

### Kubernetes

https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/