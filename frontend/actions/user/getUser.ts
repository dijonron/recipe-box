import { userClient } from "@/lib/grpc/userClient";

export async function getUser(email: string): Promise<any> {
  return new Promise((resolve, reject) => {
    userClient.GetUserByEmail({ email }, (err: any, response: any) => {
      if (err) reject(err);
      else resolve(response);
    });
  });
}
