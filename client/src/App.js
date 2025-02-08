import {
  Box,
  Container,
  CssBaseline,
  Tab,
  Tabs,
  Typography,
} from "@mui/material";
import { Recipies } from "./components/Recipes";
import { TabPanel } from "./components/TabPanel";
import { WeeklyMenu } from "./components/WeeklyMenu";
import { ShoppingList } from "./components/ShoppingList";
import {
  ThemeProvider as MUIThemeProvider,
  createTheme,
} from "@mui/material/styles";
import DarkModeSwitch from "./components/DarkModeSwitch/DarkModeSwitch";
import { TabProvider, useTab } from "./contexts/TabContext";
import { ThemeProvider, useTheme } from "./contexts/ThemeContext";
import { SnackbarProvider } from "./contexts/SnackbarContext";

function App() {
  const { isDarkMode } = useTheme();
  const { selectedTab, handleTabChange } = useTab();

  const theme = createTheme({
    palette: {
      mode: isDarkMode ? "dark" : "light",
    },
  });

  return (
    <MUIThemeProvider theme={theme}>
      <CssBaseline>
        <Container>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            <Tabs
              value={selectedTab}
              onChange={(_, newTab) => handleTabChange(newTab)}
            >
              <Tab value={0} label="weekly menu" />
              <Tab value={1} label="shopping list" />
              <Tab value={2} label="recipies" />
            </Tabs>
            <DarkModeSwitch />
          </Box>

          <TabPanel value={selectedTab} index={0}>
            <Typography variant="h5">Plan your meals for the week</Typography>
            <WeeklyMenu />
          </TabPanel>

          <TabPanel value={selectedTab} index={1}>
            <Typography variant="h5">Here's what to get</Typography>
            <ShoppingList />
          </TabPanel>

          <TabPanel value={selectedTab} index={2}>
            <Typography variant="h5">Add and edit recipies</Typography>
            <Recipies />
          </TabPanel>
        </Container>
      </CssBaseline>
    </MUIThemeProvider>
  );
}

export default function AppWrapper() {
  return (
    <ThemeProvider>
      <SnackbarProvider>
        <TabProvider>
          <App />
        </TabProvider>
      </SnackbarProvider>
    </ThemeProvider>
  );
}
