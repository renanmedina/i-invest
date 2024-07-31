export default interface DiscordMessageBody {
  content?: string
  tts: boolean,
  embeds?: EmbedBody[]
}

interface EmbedBody {
  title: string
  description: string
}