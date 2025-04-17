"use client";

import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useState } from "react";

export default function Home() {
  const [orgName, setOrgName] = useState("");
  const [inviteCode, setInviteCode] = useState("");
  // const router = useRouter();

  const handleCreateOrg = async () => {
    // Call gRPC or API to create org
    // await createOrg({ name: orgName });
    // router.refresh(); // Trigger middleware re-eval
  };

  const handleJoinOrg = async () => {
    // Call gRPC or API to join org
    // await joinOrg({ inviteCode });
    // router.refresh();
  };

  return (
    <Tabs defaultValue="create" className="w-full max-w-md">
      <TabsList className="grid w-full grid-cols-2 mb-6">
        <TabsTrigger value="create">Create A New Recipe Box</TabsTrigger>
        <TabsTrigger value="join">Join An Existing Recipe Box</TabsTrigger>
      </TabsList>

      <TabsContent value="create">
        <Card>
          <CardContent className="space-y-4 pt-6">
            <Input
              placeholder="Recipe Box Name"
              value={orgName}
              onChange={(e) => setOrgName(e.target.value)}
            />
            <Button
              className="w-full"
              onClick={handleCreateOrg}
              disabled={!orgName}
            >
              Create
            </Button>
          </CardContent>
        </Card>
      </TabsContent>

      <TabsContent value="join">
        <Card>
          <CardContent className="space-y-4 pt-6">
            <Input
              placeholder="Invite code or box name"
              value={inviteCode}
              onChange={(e) => setInviteCode(e.target.value)}
            />
            <Button
              className="w-full"
              onClick={handleJoinOrg}
              disabled={!inviteCode}
            >
              Request to Join
            </Button>
          </CardContent>
          <Button>Logout</Button>
        </Card>
      </TabsContent>
    </Tabs>
  );
}
