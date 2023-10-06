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
		TaskPromt("Generate podcast show notes from the provided transcript. The notes should be concise yet factual, summarizing key points discussed in the episode. It is ok to just list of main topics, do not try to describe the flow (* iPhone 15 - good example; * discussin coverd release of the iPhone 15 - bad example). Use list of books/TV shows/software programs that were used during the show, as many facts/titles from the show as possible. These show notes will be published on iTunes. Avoid excessive use of buzzwords; the aim is to inform listeners rather than attract attention. For formatting and style, refer to the show notes of the podcast Radio-T as a guideline. Do not explicitly mention that the style is similar to Radio-T. Additionally, here's the Discord link for further discussion: https://discord.gg/T38WpgkHGQ. Please format the output in Markdown. Also in the beggining of your output list 3-4 suggested podcast title names. Also include 2-3 suggestion for the prompt to Dall-e for the episode thumbnail.")

	result, err := executor.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("final result: ", result)
	err = ioutil.WriteFile("./shownotes.txt", []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}
