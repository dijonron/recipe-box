import { AuthProvider } from "@/context/AuthContext";
// import { getCurrentUser } from "@/lib/auth";
import { redirect } from "next/navigation";

export default async function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  // const user = await getCurrentUser();
  // if (!user) {
  //   redirect("/login");
  // }

  // if (!user.tenant_id) {
  //   redirect("/onboarding");
  // }

  return <AuthProvider initialUser={{}}>{children}</AuthProvider>;
}
