"use server";

import { loginSchema } from "@/components/auth/types";
import * as grpc from "@grpc/grpc-js";
import { redirect } from "next/navigation";
import { userClient } from "../../lib/grpc/userClient";
import { createSession } from "@/lib/session";

function loginRequest(email: string, password: string): Promise<any> {
  return new Promise((resolve, reject) => {
    userClient.Login({ email, password }, (err: any, response: any) => {
      if (err) reject(err);
      else resolve(response);
    });
    // reject({ code: grpc.status.INTERNAL });
    // resolve({
    //   user: { id: "9152c9f4-9782-41c6-bd42-7bf7a58655c2", role: "user" },
    // });
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
    const resp = await loginRequest(email, password);
    const user = resp.user;
    console.log(user);

    if (!user) {
      throw Error("could not log user in");
    }

    await createSession(user.id, user.role);
  } catch (err: any) {
    if (err?.code === grpc.status.UNAUTHENTICATED) {
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
