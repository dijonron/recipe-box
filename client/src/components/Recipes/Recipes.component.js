import React, { useState } from "react";
import AddIcon from "@mui/icons-material/Add";
import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TablePagination,
  TableRow,
} from "@mui/material";
import RecipeModal from "../RecipeModal/RecipeModal.component";

function Recipies() {
  const [open, setOpen] = useState(false);

  const count = 15;
  const rowsPerPage = 5;
  const page = 0;

  const handleChangePage = () => {};
  const handleChangeRowsPerPage = () => {};

  return (
    <>
      <Button
        size="medium"
        color="primary"
        variant="contained"
        startIcon={<AddIcon />}
        onClick={() => {
          setOpen(true);
        }}
      >
        Add Recipie
      </Button>
      {open && (
        <RecipeModal
          open={open}
          handleClose={() => {
            setOpen(false);
          }}
        />
      )}
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Recipie</TableCell>
              <TableCell>Chef</TableCell>
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
    </>
  );
}

export default Recipies;
