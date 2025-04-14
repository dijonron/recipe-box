"use client";

import { createContext, useContext, useState, useEffect } from "react";

type User = {
  email: string;
  name?: string;
  tenant_id?: string;
  role?: string;
};

const AuthContext = createContext<{ user: User | null; isLoading: boolean }>({
  user: null,
  isLoading: true,
});

export function AuthProvider({
  children,
  initialUser,
}: {
  children: React.ReactNode;
  initialUser?: any;
}) {
  const [user, setUser] = useState<User | null>(initialUser ?? null);
  const [isLoading, setIsLoading] = useState(!initialUser);

  useEffect(() => {
    if (!initialUser) {
      fetch("/api/me")
        .then((res) => res.json())
        .then((data) => {
          setUser(data.user);
          setIsLoading(false);
        })
        .catch(() => {
          setUser(null);
          setIsLoading(false);
        });
    }
  }, [initialUser]);

  return (
    <AuthContext.Provider value={{ user, isLoading }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within AuthProvider");
  return context;
};
