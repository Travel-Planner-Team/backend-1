package backend

import (
	"context"
	"fmt"
	"strings"

	//"regexp"

	"travel-planner/model"
	"travel-planner/util"

	"github.com/google/uuid"
	gogpt "github.com/sashabaranov/go-gpt3"
	//https://pkg.go.dev/github.com/aknuds1/go-gpt3
)

func SearchSitesInChatGPT(query string) ([]model.Site, error) {

  
   config, err := util.LoadApplicationConfig("conf", "chatGPT.yml")
   c := gogpt.NewClient(config.ChatGPTConfig.Key)
   ctx := context.Background()
 // get the config from config file
 // print client config
 fmt.Println(query)
//   fmt.Printf("key: %v\n", os.Getenv("GPT_KEY"))
//  c := gogpt.NewClient(os.Getenv("GPT_KEY"))
//  ctx := context.Background()

 req := gogpt.CompletionRequest{
  Model:     "text-davinci-003",
  MaxTokens: 200,
  Prompt:    query,
  Temperature: 0,
 }
 fmt.Printf("searchSites%v\n",req);
 resp, err := c.CreateCompletion(ctx, req)
 fmt.Printf("searchSites%v\n",resp);

 if err != nil {
   fmt.Printf("err: %v\n",err)
  return nil, err
 }
 
 reply :=resp.Choices[0].Text
 fmt.Println(reply)
 var sites []model.Site
 
 //a := regexp.MustCompile(`0-9\s`)
 //res := a.Split(reply, -1)
//  res := strings.Split(reply, ".")
//  fmt.Println(res)

// parse rsp_text return a list of string
 rsp_list := strings.Split(reply, "\n")
 // print each item in rsp_list
 for _, item := range rsp_list {
  // print item if not empty
  if item != "" {
   i := strings.Index(item, ".")
   item_clean := item[i+2:]
   var site model.Site
   site.Site_name = item_clean
   site.Id =uuid.New().ID()
   sites = append(sites,site)
  }
 }
 
  return sites, nil

}

func ReadSitesFromChatGPT(resp gogpt.CompletionResponse) ([]model.Site, error) {
fmt.Printf("ReadSitesFromCHatGPT%v\n",resp);
 var sites []model.Site

 var choices []gogpt.CompletionChoice
 choices = resp.Choices
 for _, item := range choices {
  var site model.Site
  site.Site_name = item.Text
  site.Id = uuid.New().ID()
  sites = append(sites, site)
 }

 return sites, nil
}