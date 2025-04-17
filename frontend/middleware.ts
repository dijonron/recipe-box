import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers";
import { decrypt } from "@/lib/session";

const publicRoutes = ["/login", "/signup"];
const onboardingRoute = "/onboarding";

export default async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname;
  const isPublicRoute = publicRoutes.includes(path);
  const isOnboardingRoute = path === onboardingRoute;

  const cookie = (await cookies()).get("session")?.value;
  const session = await decrypt(cookie);

  // user is not logged in an accessing login or signup
  if (isPublicRoute && !session?.userId) {
    return NextResponse.next();
  }

  // user is logged in an accessing login or signup, redirect
  if (isPublicRoute && session?.userId) {
    console.log(session);
    if (!session?.tenantId) {
      return NextResponse.redirect(new URL("/onboarding", req.nextUrl));
    }
    return NextResponse.redirect(new URL("/", req.nextUrl));
  }

  // if user has not completed onboarding, redirect
  if (!isOnboardingRoute && !session?.tenantId) {
    return NextResponse.redirect(new URL("/onboarding", req.nextUrl));
  }

  // user is not logged in and accessing private route
  if (!isPublicRoute && !session?.userId) {
    return NextResponse.redirect(new URL("/login", req.nextUrl));
  }

  return NextResponse.next();
}

// Routes Middleware should not run on
export const config = {
  matcher: ["/((?!api|_next/static|_next/image|.*\\.png$).*)"],
};
