import { FormControlLabel, Switch } from "@mui/material";
import { useTheme } from "../../contexts/ThemeContext";

function DarkModeSwitch() {
  const { isDarkMode, toggleDarkMode } = useTheme();

  return (
    <FormControlLabel
      control={
        <Switch sx={{ m: 1 }} checked={isDarkMode} onChange={toggleDarkMode} />
      }
    />
  );
}

export default DarkModeSwitch;
