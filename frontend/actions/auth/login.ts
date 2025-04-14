"use server";

import { cookies } from "next/headers";
import { loginSchema } from "@/components/auth/types";
import * as grpc from "@grpc/grpc-js";
import { redirect } from "next/navigation";
import { authClient } from "../../lib/grpc/authClient";

function loginRequest(email: string, password: string): Promise<any> {
  return new Promise((resolve, reject) => {
    authClient.Login({ email, password }, (err: any, response: any) => {
      if (err) reject(err);
      else resolve(response);
    });
  });
}

export async function login(_: any, formData: FormData) {
  const { data, success: parsed, error } = loginSchema.safeParse(formData);

  if (!parsed) {
    return {
      success: false,
      errors: error.flatten().fieldErrors,
    };
  }

  try {
    const { email, password } = data;
    const res = await loginRequest(email, password);
    (await cookies()).set("auth_token", res.token, {
      httpOnly: true,
      path: "/",
      secure: process.env.NODE_ENV === "production",
      sameSite: "lax",
      maxAge: 60 * 60 * 24,
    });
  } catch (err: any) {
    if (err.code === grpc.status.UNAUTHENTICATED) {
      return {
        success: false,
        errors: { _form: ["Invalid email or password."] },
      };
    }
    return {
      success: false,
      errors: { _form: ["An error occured. Please try again."] },
    };
  }
  redirect("/");
}
