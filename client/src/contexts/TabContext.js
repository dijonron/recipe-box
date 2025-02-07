import { createContext, useState, useContext } from "react";

const TabContext = createContext();

export const TabProvider = ({ children }) => {
  const storedTab = localStorage.getItem("selectedTab");
  const [selectedTab, setSelectedTab] = useState(
    storedTab ? parseInt(storedTab, 10) : 0
  );

  const handleTabChange = (newTab) => {
    setSelectedTab(newTab);
    localStorage.setItem("selectedTab", newTab);
  };

  return (
    <TabContext.Provider value={{ selectedTab, handleTabChange }}>
      {children}
    </TabContext.Provider>
  );
};

export const useTab = () => useContext(TabContext);
