# Go-Alexa

An opinionated Alexa Skill written for Go and AWS (oh, and IOPipe too!)


## Main Features
 * Integration Testing Library (go-alexa-test)
 * Internatalization of Strings (go-alexa-i18n)
 * Static Intent (go-alexa-staticintent)
 * Say Float (go-alexa-sayfloat) (Coming Soon!)
 * IOPipe Logging Integration (go-alexa-iopipe)
 * Dynamic Dispatcher (go-alexa-dispatcher)
 
## Secondary Features
 * Testing framework helps minimize alexa skill certification time
 * CI Tool set (bash-alexa-ci) for rapid prototyping, versioning, and pubslihing



## Using this Library


Simply create a main.go file similar to this one.
There is an assumption (at the moment) that AWS Lambda is in use.

```
// main.go
package main

import (
	"github.com/agilesyndrome/go-alexa/alexaskill"
)

func init() {
  skill.My.Name = "myskill"
}

func main() {
	  skill.OnLambda()
}
```

Create a ./i18n.yml file. 

At a minimum you should always have
  * skill-name (The default card title)
  * welcome.text (Read after saying "Alexa open myskill")
  * fallback.text (When nothing else matches your skill intent routes)
  * help.text (To desscribe what your skill can do)
  * kaput.text (For when your skill returns an error)
```
en-US:
  skill-name: "My Alexa Skill"
  welcome:
    text: "Thanks for launching my skill"
  stop:
    text: "Goodbye!"
  fallback:
    text: "Hmmm, I'm not sure what you're trying to do."
  help:
    text: "Here are the things you can do."
  hello:
    text: "Hellllllooo World!"
  superpower:
    text: "These are my super duper super powers"
  kaput:
    text: "Something went very wrong"
    title: "I'm broken"
```

## StaticIntent Feature
 You do not need to write code for every intent you want to create.
 If your intent is to simply read text back to the user, without further processing, then you can create it as a section in the `i18n.yml` file (see superpower above).  Add an intent in the Alexa Skills console with `superpower` as the intent name.  This skill SDK will attempt to default to reading static text from the .yml file if an intent name is given to it that does not match other code routes.

## Custom Intent Code

So, you want to do more than just read plain text? No problem! We can do that too:

In your Skill's main.go file:
```
// main.go
package main

import (
	"github.com/agilesyndrome/go-alexa/alexaskill"
	"github.com/agilesyndrome/alexaskill-mycarpenter/src/actualsize"
)

func init() {
  skill.My.IntentMap["actual_size"] = ActualSizeIntent.Handler
  skill.My.Name = "mycarpenter"
}

func main() {
	  skill.OnLambda()
}
```

  Notice how the `actual_size` intent is recorded in the skill.My.IntentMap section of the `init` function.  This will cause the `actual_size` intent to be routed to your custom Handler instead.



## Create an Alexa Skill
1. Create your Intents


## Publish your Alexa Skill Go Files to Lambda

## Link Lambda to your Alexa Skill

