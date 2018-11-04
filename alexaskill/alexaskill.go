package alexaskill

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
  mySkill Skill = Skill {}
)

func init() {
  mySkill.IntentMap = dispatcher.IntentMap
  mySkill.RequestMap = dispatcher.RequestMap
}

func OnLambda(skill_name string) {
  
  
  mySkill.Name = skill_name
  
  //If you set IOPIPE_TOKEN, let's assume you want to use it
  mySkill.UseIoPipe = (os.Getenv("IOPIPE_TOKEN") != "")

  //Wrap the LambdaHandler through the IOPipe Logging mechanism... 
  if (mySkill.UseIoPipe == true) {
    lambda.Start(
      iopipe.WrapHandler(MyLambda),
    )
  } else {
    // or dont
    lambda.Start(MyLambda)
  }

}

func MyLambda(ctx context.Context, req alexa.Request) (alexa.Response, error) {
  if (mySkill.UseIoPipe == true) {
    iopipe.Tag(ctx, req)
  }

  dispatcher.IntentMap = mySkill.IntentMap
  dispatcher.RequestMap = mySkill.RequestMap
  response, err := dispatcher.Dispatch(req)
  return response, err
}
