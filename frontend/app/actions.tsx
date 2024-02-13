type PostMessageResponseType = { code: string; privateKey: string };

export async function postMessage(
  backendUrl: string,
  params: any,
): Promise<PostMessageResponseType> {
  const body = {
    message: params.message,
    email: params.email,
    shouldBurnInMinutes: params.burn ? 5 : 0,
  };

  const response = await fetch(backendUrl + "messages", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });
  const json = await response.json();
  if (!response.ok) {
    throw new Error(json.error);
  }
  return { code: json.code, privateKey: json.privateKey };
}

type ReadMessageResponseType = { message: string };

export async function readMessage(
  backendUrl: string,
  params: any,
): Promise<ReadMessageResponseType> {
  const body = {
    privateKey: params.privateKey,
  };

  const response = await fetch(backendUrl + "messages/" + params.code, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });
  const json = await response.json();

  if (!response.ok) {
    throw new Error(json.error);
  }
  return { message: json.message };
}
