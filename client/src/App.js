import { useState } from "react";
import { Container, Tab, Tabs, Typography } from "@mui/material";
import { Recipies } from "./components/Recipes";
import { TabPanel } from "./components/TabPanel";
import { WeeklyMenu } from "./components/WeeklyMenu";
import { ShoppingList } from "./components/ShoppingList";

function App() {
  const [activeTab, setActiveTab] = useState(0);

  return (
    <Container>
      <Tabs value={activeTab} onChange={(_, newTab) => setActiveTab(newTab)}>
        <Tab value={0} label="weekly menu" />
        <Tab value={1} label="shopping list" />
        <Tab value={2} label="recipies" />
      </Tabs>

      <TabPanel value={activeTab} index={0}>
        <Typography variant="h5">Plan your meals for the week</Typography>
        <WeeklyMenu />
      </TabPanel>

      <TabPanel value={activeTab} index={1}>
        <Typography variant="h5">Here's what to get</Typography>
        <ShoppingList />
      </TabPanel>

      <TabPanel value={activeTab} index={2}>
        <Typography variant="h5">Add and edit recipies</Typography>
        <Recipies />
      </TabPanel>
    </Container>
  );
}

export default App;
