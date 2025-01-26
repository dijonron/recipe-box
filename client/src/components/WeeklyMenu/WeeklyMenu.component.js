import React from "react";
import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
} from "@mui/material";

function WeeklyMenu(props) {
  return (
    <TableContainer>
      <Table>
        <TableBody>
          <TableRow>
            <TableCell>Monday</TableCell>
            <TableCell>
              <FormControl fullWidth>
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
                <InputLabel id="demo-simple-select-label">Select</InputLabel>
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
  );
}

export default WeeklyMenu;
