package main

import (
	"fmt"
	"io/ioutil"

	"github.com/assistant-ai/llmchat-client/gpt"
	"github.com/assistant-ai/prompt-tools/chainexecutor"
)

func main() {
	client, err := gpt.NewGptClientFromFile("/Users/vkovalevskyi/.jess/open-ai.key", 0, gpt.ModelGPT4, "", 8000, nil)
	if err != nil {
		panic(err)
	}
	// Read the content from the file
	data, err := ioutil.ReadFile("./transcript.txt")
	if err != nil {
		panic(err)
	}
	podcastTranscript := string(data)

	executor := chainexecutor.NewSimpleChainExecutor().
		LlmClient(client).
		Text(podcastTranscript).
		TaskPromt("Create podcast show notes from the podcast transcript. They should be short, but to the point with facts about what was discussed in the podcast. Will be published in iTunes. Good example of shownotes: podcast Radio-T, use it as example of how I want shownotes to look like.")

	result, err := executor.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("final result: ", result)
	// save to file
	err = ioutil.WriteFile("./shownotes.txt", []byte(result), 0644)
}
