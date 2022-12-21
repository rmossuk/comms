import { HandleRequest, HttpRequest, HttpResponse} from "@fermyon/spin-sdk"

const encoder = new TextEncoder()
const decoder = new TextDecoder()

const redisAddress = "redis://localhost:6379/"

interface InputPayload {
  from: string,
  to: string,
  message: string
}

export const handleRequest: HandleRequest = async function(request: HttpRequest): Promise<HttpResponse> {

    const input = request.json() as InputPayload

    const formatMessage = input.from+":"+input.message+":"+input.to
    spinSdk.redis.publish(redisAddress, "messages", encoder.encode(formatMessage).buffer)

    const outputMessage = "Added message '"+input.message+"' from '"+input.from+"' to '"+input.to+"'"

    return {
      status: 200,
        headers: { "foo": "bar" },
      body: encoder.encode(outputMessage).buffer
    }
}
