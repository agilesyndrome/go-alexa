package skill

import (
  "context"
  "os"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-iopipe/iopipe"
  "github.com/agilesyndrome/go-alexa-dispatcher/dispatcher"

)

type LambdaHandler func (context.Context, alexa.Request) (alexa.Response, error)

type Skill struct {
    Name string
    UseIoPipe bool
    IntentMap map[string] dispatcher.AlexaHandler
    RequestMap map[string] dispatcher.AlexaHandler
}

var (
  My Skill = Skill {}
)

func init() {
  My.IntentMap = dispatcher.IntentMap
  My.RequestMap = dispatcher.RequestMap
  My.UseIoPipe = (os.Getenv("IOPIPE_TOKEN") != "")
}

func OnLambda() {
  
  //Wrap the LambdaHandler through the IOPipe Logging mechanism... 
  if (My.UseIoPipe == true) {
    lambda.Start(
      iopipe.WrapHandler(MyLambda),
    )
  } else {
    // or dont
    lambda.Start(MyLambda)
  }

}

func MyLambda(ctx context.Context, req alexa.Request) (alexa.Response, error) {
  if (My.UseIoPipe == true) {
    iopipe.Tag(ctx, req)
  }

  dispatcher.IntentMap = My.IntentMap
  dispatcher.RequestMap = My.RequestMap
  response, err := dispatcher.Dispatch(req)
  return response, err
}
