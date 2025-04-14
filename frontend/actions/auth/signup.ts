"use server";

import { cookies } from "next/headers";
import { signupSchema } from "@/components/auth/types";
import * as grpc from "@grpc/grpc-js";
import { redirect } from "next/navigation";
import { userClient } from "../../lib/grpc/userClient";

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
    const res = await signupRequest(name, email, password);
    (await cookies()).set("auth_token", res.token, {
      httpOnly: true,
      path: "/",
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      maxAge: 60 * 60 * 24,
    });
  } catch (err: any) {
    // TODO: already exists!
    return {
      success: false,
      errors: { _form: ["An error occured. Please try again."] },
    };
  }
  redirect("/");
}
