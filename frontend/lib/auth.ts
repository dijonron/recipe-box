import { validate } from "@/actions/auth/validate";
import { getUser } from "@/actions/user/getUser";
import { cookies } from "next/headers";

export function decodeJWT(token: string) {
  try {
    const payload = JSON.parse(
      Buffer.from(token.split(".")[1], "base64").toString("utf-8")
    );
    return payload;
  } catch (e) {
    return null;
  }
}

export async function getCurrentUser() {
  const cookieStore = await cookies();
  const token = cookieStore.get("auth_token")?.value;
  if (!token) return null;
  try {
    const { valid } = await validate(token);
    if (!valid) return null;

    const payload = decodeJWT(token);
    if (!payload) return null;

    const { email } = payload;

    const { user } = await getUser(email);

    return user;
  } catch (err) {
    console.error("Error validating token:", err);
    return null;
  }
}

export async function logout() {
  const cookieStore = await cookies();
  cookieStore.delete("auth_token");
}
