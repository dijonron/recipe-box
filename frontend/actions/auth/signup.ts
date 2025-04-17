"use server";

import { signupSchema } from "@/components/auth/types";
import { redirect } from "next/navigation";
import { userClient } from "../../lib/grpc/userClient";
import { createSession } from "@/lib/session";

function signupRequest(
  name: string,
  email: string,
  password: string
): Promise<any> {
  return new Promise((resolve, reject) => {
    userClient.CreateUser(
      { name, email, password },
      (err: any, response: any) => {
        if (err) reject(err);
        else resolve(response);
      }
    );
  });
}

export async function signup(_: any, formData: FormData) {
  const { data, success: parsed, error } = signupSchema.safeParse(formData);

  if (!parsed) {
    return {
      success: false,
      errors: error.flatten().fieldErrors,
    };
  }

  try {
    const { name, email, password } = data;
    const resp = await signupRequest(name, email, password);
    const user = resp.user;

    if (!user) {
      throw Error("could not sign user up");
    }

    await createSession(user.id, user.role);
  } catch (err) {
    return {
      success: false,
      message:
        "An error occurred while creating your account. Please try again.",
    };
  }
  redirect("/onboarding");
}
