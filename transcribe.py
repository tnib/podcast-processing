# Note: you need to be using OpenAI Python v0.27.0 for the code below to work
import openai

openai.api_key_path="/Users/vkovalevskyi/.jess/open-ai.key"

audio_file= open("./podcast.mp3", "rb")
transcript = openai.Audio.transcribe("whisper-1", audio_file, response_format="text")

# save to file transript.txt
with open('transcript.txt', 'w') as f:
    f.write(transcript)
