import { useState } from "react";
import {
  Container,
  FormControl,
  InputLabel,
  Tabs,
  Tab,
  Typography,
  Button,
  TableContainer,
  TableHead,
  Table,
  TableRow,
  TableCell,
  TableBody,
  TablePagination,
  Select,
  MenuItem,
} from "@mui/material";
import AddIcon from "@mui/icons-material/Add";

function CustomTabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && children}
    </div>
  );
}

function App() {
  const [activeTab, setActiveTab] = useState(0);

  const count = 15;
  const rowsPerPage = 5;
  const page = 0;

  const handleChangePage = () => {};
  const handleChangeRowsPerPage = () => {};

  return (
    <Container>
      <Tabs value={activeTab} onChange={(_, newTab) => setActiveTab(newTab)}>
        <Tab value={0} label="weekly menu" />
        <Tab value={1} label="shopping list" />
        <Tab value={2} label="recipies" />
      </Tabs>

      <CustomTabPanel value={activeTab} index={0}>
        <Typography variant="h5">Plan your meals for the week</Typography>

        <TableContainer>
          <Table>
            <TableBody>
              <TableRow>
                <TableCell>Monday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Tuesday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      {" "}
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>{" "}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Wednesday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      {" "}
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>{" "}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Thursday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      {" "}
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>{" "}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Friday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      {" "}
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>{" "}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Saturday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      {" "}
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>{" "}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell>Sunday</TableCell>
                <TableCell>
                  <FormControl fullWidth>
                    <InputLabel id="demo-simple-select-label">
                      Select
                    </InputLabel>
                    <Select
                      variant="standard"
                      size="medium"
                      label="value"
                      value=""
                      onChange={() => {}}
                    >
                      {" "}
                      <MenuItem value="">Add some recipies</MenuItem>
                    </Select>
                  </FormControl>{" "}
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </CustomTabPanel>

      <CustomTabPanel value={activeTab} index={1}>
        <Typography variant="h5">Here's what to get</Typography>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Ingredient</TableCell>
                <TableCell>Quantity</TableCell>
                <TableCell></TableCell>
              </TableRow>
            </TableHead>
            <TableBody></TableBody>
          </Table>
        </TableContainer>
      </CustomTabPanel>

      <CustomTabPanel value={activeTab} index={2}>
        <Typography variant="h5">Add and edit recipies</Typography>
        <Button
          size="medium"
          color="primary"
          variant="contained"
          startIcon={<AddIcon />}
        >
          Add Recipie
        </Button>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Recipie</TableCell>
                <TableCell>Cell</TableCell>
                <TableCell></TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {Array.from({ length: 5 }, (_, i) => i).map((i) => (
                <TableRow>
                  <TableCell>Cell</TableCell>
                  <TableCell>Cell</TableCell>
                  <TableCell align="right">
                    <Button size="small" color="secondary" variant="outlined">
                      Edit
                    </Button>
                    <span style={{ marginLeft: "16px" }}> </span>
                    <Button size="small" color="error" variant="text">
                      Delete
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25]}
          component="div"
          count={count}
          rowsPerPage={rowsPerPage}
          page={page}
          onPageChange={handleChangePage}
          onRowsPerPageChange={handleChangeRowsPerPage}
        />
      </CustomTabPanel>
    </Container>
  );
}

export default App;
